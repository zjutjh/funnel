package zfCaptchaCracker

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"os"
	"strings"
	"time"
)

// CompareCracker is a ZfCaptchaCracker
// 通过和本地模板库比对的方式破解验证码
type CompareCracker struct {
	templatePath         string
	templateFingerprints map[string]string
}

// PUBLIC METHODS

func (c *CompareCracker) Init(config string) error {
	c.templatePath = config // 该破解器将参数视为模板路径
	c.templateFingerprints = make(map[string]string)

	entries, err := os.ReadDir(c.templatePath) // 读入模板目录
	if err != nil {
		return fmt.Errorf("cannot open dir: %s", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue // 跳过子目录（不递归）
		}
		name := entry.Name()
		if strings.HasSuffix(strings.ToLower(name), ".png") {
			// 打开并解码文件
			img, err := c.openTemplateImage(c.templatePath + "/" + entry.Name())
			if err != nil {
				return err
			}
			// 生成模板图像指纹并存储
			c.templateFingerprints[c.genImgFingerprint(img)] = entry.Name()
		}
	}
	return nil
}

func (c *CompareCracker) Crack(img image.Image) (string, error) {
	templateName, found := c.templateFingerprints[c.genImgFingerprint(img)]
	if !found {
		return "", fmt.Errorf("no matching template found")
	}
	templateImg, err := c.openTemplateImage(c.templatePath + "/" + templateName)
	if err != nil {
		return "", err
	}
	gapPos, err := findGapLeftEdge(img, templateImg)
	if err != nil {
		return "", err
	}
	return c.genMouseTrack(gapPos), nil
}

// PRIVATE METHODS

func (c *CompareCracker) openTemplateImage(fileName string) (image.Image, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open template file: %s", fileName)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("cannot decode template image: %s", fileName)
	}
	return img, nil
}

func (c *CompareCracker) genImgFingerprint(img image.Image) string {
	bounds := img.Bounds()
	points := []image.Point{
		{X: bounds.Min.X, Y: bounds.Min.Y},
		{X: bounds.Max.X - 1, Y: bounds.Min.Y},
		{X: bounds.Min.X, Y: bounds.Max.Y - 1},
		{X: bounds.Max.X - 1, Y: bounds.Max.Y - 1},
	}
	var parts []string
	for _, p := range points {
		r, g, b, _ := img.At(p.X, p.Y).RGBA()
		parts = append(parts, fmt.Sprintf("%d%d%d", r, g, b))
	}
	return strings.Join(parts, "")
}

// subtractColorSum 计算像素色值减去(最多到0)的差值之和
func subtractColorSum(fg, bg color.Color) int {
	r1, g1, b1, _ := fg.RGBA()
	r2, g2, b2, _ := bg.RGBA()
	rSub := max(int(r1)-int(r2), 0)
	gSub := max(int(g1)-int(g2), 0)
	bSub := max(int(b1)-int(b2), 0)
	sum := rSub + gSub + bSub
	return sum
}

func findGapLeftEdge(bgImg, templateImg image.Image) (int, error) {

	// 每次扫描中要向上校验的纵像素列数量 此值根据拼图突出部分纵向最大长度确定 对效率线性影响
	// 这个值不能大于scanStepY，否则会漏边界
	const chkPixelNum int = 15
	// 扫描Y跳过的步长 此值根据拼图高度确定 对效率成倍影响
	const scanStepY int = 60 - chkPixelNum
	// 初次扫描X跳过的步长 此值根据拼图宽度确定 对效率成倍影响
	const scanStepX int = 60

	dbg_repeat := 0

	// 图像大小校验
	imgBounds := bgImg.Bounds()
	if imgBounds.Dx() != templateImg.Bounds().Dx() || imgBounds.Dy() != templateImg.Bounds().Dy() {
		return 0, fmt.Errorf("bgImg(downloaded) and templateImg size mismatch")
	}

	maxStartY := imgBounds.Dy() - 1 // 可能作为指纹的右侧不检测
	maxStartX := imgBounds.Dx() - 1 // 可能作为指纹的底边不检测
	// 从上到下，逐步移动起始Y位置，初始Y的位置应该确保向上查询不越界
	for curStartY := scanStepY; curStartY <= maxStartY; curStartY += scanStepY {
		// 在当前起始Y位置，逐列扫描X轴
		for curX := scanStepX; curX < maxStartX; curX += scanStepX {
			// 当步内逐像素对比当前列
			inChunk := true
			// 从开始Y开始，逐像素对比当前列，比对checkPerStep个像素
			for curYOffset := 0; curYOffset < chkPixelNum; curYOffset++ {
				dbg_repeat++
				// 从起始位置向上判断，因为拼图突出的部分向上
				if subtractColorSum(bgImg.At(curX, curStartY-curYOffset), templateImg.At(curX, curStartY-curYOffset)) == 0 {
					inChunk = false // 发现像素相同，判定为拼图突出部分
					break           // 当步从下到上检测，任意像素相同则跳出
				}
			}
			if inChunk {
				// 开始向左滑动寻找边缘
				for subtractColorSum(bgImg.At(curX, curStartY), templateImg.At(curX, curStartY)) != 0 {
					dbg_repeat++
					curX--
				}
				curX++ // 回退一步到边缘位置
				return curX, nil
			}
		}
	}
	return 0, fmt.Errorf("targeted left edge not found")
}

func (c *CompareCracker) genMouseTrack(distance int) string {
	// @author: CodeBoy2006
	type move struct {
		X int `json:"x"`
		Y int `json:"y"`
		T int `json:"t"`
	}
	var track []move
	startX, startY := 630, rand.Intn(10)+480
	startTime := time.Now().UnixMilli()
	track = append(track, move{X: startX, Y: startY, T: int(startTime)})

	totalDuration := int64(rand.Intn(400) + 300) // 总耗时在300-700ms之间
	for i := 1; i <= distance; i += rand.Intn(5) + 5 {
		if i > distance {
			break
		}
		currentTime := startTime + (int64(i) * totalDuration / int64(distance))
		track = append(track, move{X: startX + i, Y: startY + rand.Intn(4) - 2, T: int(currentTime)})
	}
	track = append(track, move{X: startX + distance, Y: startY, T: int(startTime + totalDuration)})

	trackBytes, _ := json.Marshal(track)
	return string(trackBytes)
}
