package model

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"io"
	"net/http"

	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/image/draw"
)

const (
	thumbnailSize = 256 // px
)

type Image struct {
	img image.Image
}

func NewImage(reader io.Reader) (Image, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return Image{}, fmt.Errorf("failed decode image. err: %w", err)
	}
	return Image{
		img: img,
	}, nil
}

func (i *Image) Thumbnail() Image {
	return i.SquareTrim(thumbnailSize)
}

func (i *Image) EncodeJpeg(w io.Writer) error {
	if err := jpeg.Encode(w, i.img, nil); err != nil {
		return fmt.Errorf("failed jpeg.Encode. err: %w", err)
	}
	return nil
}

// 正方形のトリミング
func (i *Image) SquareTrim(size int) Image {
	width := i.img.Bounds().Max.X
	height := i.img.Bounds().Max.Y

	// 短辺を基に正方形を作成
	shorter := min(width, height)

	// トリミングする正方形の左上座標を計算
	top := (height - shorter) / 2
	left := (width - shorter) / 2

	// 新しい画像を作成
	newImg := image.NewRGBA(image.Rect(0, 0, size, size))

	// 元の画像から正方形部分を切り出し、新しいサイズにリサイズ
	srcSub := image.NewRGBA(image.Rect(0, 0, shorter, shorter))
	draw.Draw(srcSub, srcSub.Bounds(), i.img, image.Point{X: left, Y: top}, draw.Src)
	draw.BiLinear.Scale(newImg, newImg.Bounds(), srcSub, srcSub.Bounds(), draw.Over, nil)

	return Image{
		img: newImg,
	}
}

// サイズを特定のサイズに変更
func (i *Image) Resize(width, height int) Image {
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.BiLinear.Scale(newImage, newImage.Bounds(), i.img, i.img.Bounds(), draw.Over, nil)
	return Image{
		img: newImage,
	}
}

// アスペクト比を保ったままサイズを変える
func (i *Image) ResizeKeepAspect(size int) Image {
	width := i.img.Bounds().Max.X
	height := i.img.Bounds().Max.Y

	if width > height {
		height = height * size / width
		width = size
	} else {
		width = width * size / height
		height = size
	}
	return i.Resize(width, height)
}

func (i *Image) ToContent(fileName string) (*Content, error) {
	b := &bytes.Buffer{}
	buf := bufio.NewWriter(b)
	if err := jpeg.Encode(buf, i.img, nil); err != nil {
		return nil, fmt.Errorf("failed encode. err: %w", err)
	}
	readCloser := io.NopCloser(b)

	contentType := http.DetectContentType(b.Bytes())

	return NewContent(readCloser, contentType, int64(b.Len()), fileName), nil
}
