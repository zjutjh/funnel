package zfCaptchaCracker

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"
)

// ! NOT READY FOR PRODUCTION USE !

// AnalyzeCracker is a ZfCaptchaCracker
type AnalyzeCracker struct {
	labConverter RgbTOLabConverter
	interest     int
}

func (c *AnalyzeCracker) Init(config string) error {
	c.labConverter.Init()
	return nil
}

func (c *AnalyzeCracker) Crack(img image.Image) (string, error) {

	if file, err := os.Create("debug_bg.png"); err == nil {
		png.Encode(file, img)
		file.Close()
	}
	fmt.Print("input interest: ")
	fmt.Scanln(&c.interest)

	pos, err := c.findGapLeftEdge(img)
	if err != nil {
		return "", err
	}
	return c.genMouseTrack(pos), nil
}

func (c *AnalyzeCracker) judgeColorLightenOnly(c1 color.Color, c2 color.Color, log bool) bool {
	L1, a1, b1 := c.labConverter.RgbToLab(c1)
	L2, a2, b2 := c.labConverter.RgbToLab(c2)

	if log {
		fmt.Printf("(%04d,%04d,%04d)(%04d,%04d,%04d)(%04d,%04d,%04d)\n ", L1, a1, b1, L2, a2, b2, L2-L1, a1-a2, b1-b2)
	}

	return (L2 >= L1+(30-L1/4))
}

func (c *AnalyzeCracker) findGapLeftEdge(img image.Image) (int, error) {
	// 扫描Y跳过的步长 此值根据拼图高度确定 对效率成倍影响
	const scanStepY int = 60
	// 每行前后跳过的像素数 此值根据拼图宽度确定 对效率影响较小
	const scanIgnoreSideX int = 60
	// 一列向上/下最大扫描的像素数 此值根据图块高度决定
	const scanMaxHeight int = 60
	// 一列同时多少像素符合规则判定为图块边缘 此值根据图块高度决定
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
				//blendColor := c.getBlendColor(img.At(curX, curStartY-offsetY))
				if curX < c.interest+3 && curX > c.interest-3 {
					fmt.Printf("(%03d,%03d)", curX, curStartY-offsetY)
				}
				if c.judgeColorLightenOnly(img.At(curX, curStartY-offsetY), img.At(curX+1, curStartY-offsetY), (curX < c.interest+1 && curX > c.interest-1)) {
					checkPassed++
				}
			}
			// 向下扫描符合规则的像素
			for offsetY := 0; offsetY < scanMaxHeight; offsetY++ {
				if curStartY+offsetY > maxStartY { // 超过边缘则退出
					break
				}
				if c.judgeColorLightenOnly(img.At(curX, curStartY-offsetY), img.At(curX+1, curStartY-offsetY), false) {
					checkPassed++
				}
			}
			if checkPassed >= pixelThreshold {
				return curX, nil
			}
		}
	}
	return 0, fmt.Errorf("targeted left edge not found") // ERR_HANDLE_8
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

// AI-Generated Code Below
// Nobody can read this code except AI

type RgbTOLabConverter struct {
	rToX        [256]float64
	rToY        [256]float64
	rToZ        [256]float64
	gToX        [256]float64
	gToY        [256]float64
	gToZ        [256]float64
	bToX        [256]float64
	bToY        [256]float64
	bToZ        [256]float64
	initialized bool
}

func (r *RgbTOLabConverter) RgbToLab(c color.Color) (L, a, b int) {
	if !r.initialized {
		r.Init()
	}

	rr, gg, bb, _ := color.RGBAModel.Convert(c).RGBA()
	r8 := uint8(rr >> 8)
	g8 := uint8(gg >> 8)
	b8 := uint8(bb >> 8)

	X := r.rToX[r8] + r.gToX[g8] + r.bToX[b8]
	Y := r.rToY[r8] + r.gToY[g8] + r.bToY[b8]
	Z := r.rToZ[r8] + r.gToZ[g8] + r.bToZ[b8]

	const (
		refX    = 0.95047
		refY    = 1.0
		refZ    = 1.08883
		epsilon = 0.008856
		kappa   = 903.3
	)

	x := X / refX
	y := Y / refY
	z := Z / refZ

	fx := labPivot(x, epsilon, kappa)
	fy := labPivot(y, epsilon, kappa)
	fz := labPivot(z, epsilon, kappa)

	Lf := 116*fy - 16
	af := 500 * (fx - fy)
	bf := 200 * (fy - fz)

	L = int(Lf)
	a = int(af)
	b = int(bf)

	return
}

func (r *RgbTOLabConverter) Init() {
	const (
		rX = 0.4124564
		rY = 0.2126729
		rZ = 0.0193339
		gX = 0.3575761
		gY = 0.7151522
		gZ = 0.1191920
		bX = 0.1804375
		bY = 0.0721750
		bZ = 0.9503041
	)

	for i := 0; i < 256; i++ {
		linear := srgbToLinear(float64(i) / 255.0)
		r.rToX[i] = linear * rX
		r.rToY[i] = linear * rY
		r.rToZ[i] = linear * rZ
		r.gToX[i] = linear * gX
		r.gToY[i] = linear * gY
		r.gToZ[i] = linear * gZ
		r.bToX[i] = linear * bX
		r.bToY[i] = linear * bY
		r.bToZ[i] = linear * bZ
	}

	r.initialized = true
}

func srgbToLinear(v float64) float64 {
	if v <= 0.04045 {
		return v / 12.92
	}
	return math.Pow((v+0.055)/1.055, 2.4)
}

func labPivot(t, epsilon, kappa float64) float64 {
	if t > epsilon {
		return math.Cbrt(t)
	}
	return (kappa*t + 16) / 116
}
