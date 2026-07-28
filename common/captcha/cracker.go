package captcha

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/anthonynsimon/bild/blend"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/segment"
	"github.com/zeromicro/go-zero/core/logx"
)

//go:embed static/*.png
var templateFS embed.FS

var (
	templates map[string]image.Image
	once      sync.Once
	initErr   error
)

// loadTemplates 懒加载图片模板并生成指纹。
func loadTemplates() error {
	logger := logx.WithContext(context.Background())
	once.Do(func() {
		templates = make(map[string]image.Image)
		entries, err := templateFS.ReadDir("static")
		if err != nil {
			initErr = fmt.Errorf("验证码初始化失败, 无法读取嵌入目录: %w", err)
			logger.Errorf("%v", initErr)
			return
		}
		for _, entry := range entries {
			name := entry.Name()
			if !strings.EqualFold(filepath.Ext(entry.Name()), ".png") {
				logger.Infof("跳过非PNG文件: %s", name)
				continue
			}
			data, err := templateFS.ReadFile("static/" + name)
			if err != nil {
				logger.Errorf("无法读取嵌入文件 %s: %v", name, err)
				continue
			}
			img, _, err := image.Decode(bytes.NewReader(data))
			if err != nil {
				logger.Errorf("无法解码嵌入图像 %s: %v", name, err)
				continue
			}
			templates[generateFingerprint(img)] = img
		}
		logger.Infof("验证码模板初始化完成, 共%d个模板", len(templates))
	})

	return initErr
}

func Crack(img image.Image) (string, error) {
	if img == nil {
		return "", fmt.Errorf("输入图片不能为空")
	}

	if err := loadTemplates(); err != nil {
		return "", err
	}

	templateImg, ok := templates[generateFingerprint(img)]
	if !ok {
		return "", fmt.Errorf("验证码匹配模板失败")
	}
	gapPos, err := match(img, templateImg)
	if err != nil {
		return "", err
	}
	return generateMouseTrack(gapPos), nil
}

// generateFingerprint 生成图片指纹
func generateFingerprint(img image.Image) string {
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

func match(bgImg, templateImg image.Image) (int, error) {
	// diff
	diff := blend.Difference(bgImg, templateImg)

	// gray
	gray := effect.Grayscale(diff)

	// binary
	binary := segment.Threshold(gray, 30)

	bounds := binary.Bounds()
	minWhitePixels := bounds.Dy() / 10

	// scan
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		whiteCount := 0
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			if binary.GrayAt(x, y).Y == 0xFF {
				whiteCount++
			}
		}
		if whiteCount >= minWhitePixels {
			return x, nil
		}
	}

	return 0, fmt.Errorf("识别失败")
}

func generateMouseTrack(distance int) string {
	type move struct {
		X int `json:"x"`
		Y int `json:"y"`
		T int `json:"t"`
	}
	var track []move
	startX, startY := 630, rand.Intn(10)+480
	startTime := time.Now().UnixMilli()
	track = append(track, move{X: startX, Y: startY, T: int(startTime)})

	totalDuration := int64(rand.Intn(400) + 300) // 模拟总耗时在300-700ms之间
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
