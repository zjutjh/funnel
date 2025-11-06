package zfCaptchaCracker

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"
)

// AnalyzeCracker is a ZfCaptchaCracker
type AnalyzeCracker struct {
}

func (c *AnalyzeCracker) Init(config string) error {
	return nil
}

func (c *AnalyzeCracker) Crack(img image.Image) (string, error) {
	pos, err := c.findGapLeftEdge(img)
	if err != nil {
		return "", err
	}
	return c.genMouseTrack(pos), nil
}

func (c *AnalyzeCracker) judgeColorLightenOnly(c1 color.Color, c2 color.Color) bool {
	L1 := c.calcGrayscale(c1)
	L2 := c.calcGrayscale(c2)

	// MAGIC NUMBER 65: 经验值 颜色变亮的阈值
	// MAGIC NUMBER 4: 经验值 像素基础亮度修正
	return (L2 >= L1+(65-L1/4))
}

func (c *AnalyzeCracker) calcGrayscale(raw color.Color) int32 {
	rr2, gg2, bb2, _ := color.RGBAModel.Convert(raw).RGBA()
	r2f := float64(uint8(rr2 >> 8))
	g2f := float64(uint8(gg2 >> 8))
	b2f := float64(uint8(bb2 >> 8))
	return int32(0.299*r2f + 0.587*g2f + 0.114*b2f)
}

func (c *AnalyzeCracker) findGapLeftEdge(img image.Image) (int, error) {
	// 扫描Y跳过的步长 此值根据拼图高度确定 对效率成倍影响
	const scanStepY int = 60
	// 每行前后跳过的像素数 此值根据拼图宽度确定 对效率影响较小
	const scanIgnoreSideX int = 60
	// 一列向上/下最大扫描的像素数 此值根据图块高度决定
	const scanMaxHeight int = 60
	// 一列同时多少像素符合规则判定为图块边缘 此值根据图块高度决定 且是经验值
	const pixelThreshold int = 40

	imgBounds := img.Bounds()
	maxStartX := imgBounds.Dx() - scanIgnoreSideX // 右边缘小于图块的区域不检测
	maxStartY := imgBounds.Dy() - 1               // 底边不检测防止溢出
	// 从上到下，逐步移动起始Y位置，初始Y的位置应该确保向上查询不越界
	for curStartY := scanStepY; curStartY <= maxStartY; curStartY += scanStepY {
		// 在当前起始Y位置，扫描X轴
		for curX := scanIgnoreSideX; curX < maxStartX; curX += 1 {
			checkPassed := 0 // 向上扫描并通过了的像素数
			// 向上扫描符合规则的像素
			for offsetY := 0; offsetY < scanMaxHeight; offsetY++ {
				if curStartY-offsetY < 0 { // 超过边缘则退出
					break
				}
				if c.judgeColorLightenOnly(img.At(curX, curStartY-offsetY), img.At(curX+1, curStartY-offsetY)) {
					checkPassed++
				}
			}
			// 向下扫描符合规则的像素
			for offsetY := 0; offsetY < scanMaxHeight; offsetY++ {
				if curStartY+offsetY > maxStartY { // 超过边缘则退出
					break
				}
				if c.judgeColorLightenOnly(img.At(curX, curStartY-offsetY), img.At(curX+1, curStartY-offsetY)) {
					checkPassed++
				}
			}
			if checkPassed >= pixelThreshold {
				return (curX + 1), nil
			}
		}
	}
	return 0, fmt.Errorf("targeted left edge not found")
}

func (c *AnalyzeCracker) genMouseTrack(distance int) string {
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
