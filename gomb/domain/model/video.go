package model

import (
	"fmt"
	"io"

	"os"

	"github.com/gofrs/uuid/v5"
	"gocv.io/x/gocv"
)

type Video struct {
	video *gocv.VideoCapture
}

func NewVideo(reader io.Reader) (Video, error) {
	// opencvはファイルでやる
	tmpF, err := os.CreateTemp("/tmp", "temp_video_*.mp4")
	if err != nil {
		return Video{}, fmt.Errorf("failed create temp. err: %w", err)
	}
	defer os.Remove(tmpF.Name())
	defer tmpF.Close()

	if _, err := io.Copy(tmpF, reader); err != nil {
		return Video{}, fmt.Errorf("failed copy to temp. err: %w", err)
	}

	video, err := gocv.VideoCaptureFile(tmpF.Name())
	if err != nil {
		return Video{}, fmt.Errorf("failed gocv.VideoCaptureFile. fileName: %s, err: %w", tmpF.Name(), err)
	}
	return Video{
		video: video,
	}, nil
}

func (v *Video) Thumbnail() (Image, error) {
	totalFrames := v.video.Get(gocv.VideoCaptureFrameCount)
	if totalFrames <= 0 {
		return Image{}, fmt.Errorf("failed get video frames")
	}

	// 中間のフレームを取得
	// ここは調整可能
	frameNumber := totalFrames / 2
	v.video.Set(gocv.VideoCapturePosFrames, frameNumber)

	img := gocv.NewMat()
	defer img.Close()

	if ok := v.video.Read(&img); !ok {
		return Image{}, fmt.Errorf("failed read frame. frameNumber: %f", frameNumber)
	}

	if img.Empty() {
		return Image{}, fmt.Errorf("failed image is empty")
	}

	u, err := uuid.NewV4()
	if err != nil {
		return Image{}, fmt.Errorf("failed uuid.NewV4. err: %w", err)
	}
	tmpFilePath := fmt.Sprintf("/tmp/temp_video_thumbnail_%s.jpg", u.String())
	if ok := gocv.IMWrite(tmpFilePath, img); !ok {
		return Image{}, fmt.Errorf("failed gocv.IMWrite")
	}
	defer os.Remove(tmpFilePath)

	f, err := os.Open(tmpFilePath)
	if err != nil {
		return Image{}, fmt.Errorf("failed open. file: %s, err: %w", tmpFilePath, err)
	}
	defer f.Close()

	thumbnail, err := NewImage(f)
	if err != nil {
		return Image{}, fmt.Errorf("failed NewImage. err: %w", err)
	}
	return thumbnail.Thumbnail(), nil
}

func (v *Video) Close() error {
	return v.video.Close()
}
