package imaging

import (
	"bytes"
	"image"
	"image/jpeg"
	"os"

	"github.com/devlights/gomy/output"
)

// Grayscale は、Go の標準パッケージ image を利用して 8-bit Grayscale JPEG を作成するサンプルです.
//
// REFERENCES:
//   - https://riptutorial.com/go/example/31693/convert-color-image-to-grayscale
//   - https://pkg.go.dev/image@go1.17#example-package
//   - https://unsplash.com/photos/hgUcyDIWPfM
func Grayscale() error {
	const (
		imgUrl = "https://unsplash.com/photos/hgUcyDIWPfM/download?force=true"
	)

	//////////////////////////////////////////////////////
	// Download
	//////////////////////////////////////////////////////
	var (
		dl  = downloder(imgUrl)
		buf *bytes.Buffer
		err error
	)

	if buf, err = dl.download(); err != nil {
		return err
	}
	output.Stdoutl("[download]", buf.Len())

	//////////////////////////////////////////////////////
	// Decode
	//////////////////////////////////////////////////////
	var (
		img    image.Image
		imgFmt string
	)

	if img, imgFmt, err = image.Decode(buf); err != nil {
		return err
	}
	output.Stdoutl("[image.Decode]", imgFmt, img.Bounds().Size())

	//////////////////////////////////////////////////////
	// Convert
	//////////////////////////////////////////////////////
	var (
		bounds  = img.Bounds()
		grayImg *image.Gray
	)

	grayImg = image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	//////////////////////////////////////////////////////
	// Output
	//////////////////////////////////////////////////////
	var (
		f *os.File
		o = &jpeg.Options{
			Quality: jpeg.DefaultQuality,
		}
	)

	if f, err = os.CreateTemp("", "*.jpg"); err != nil {
		return err
	}
	defer f.Close()

	if err = jpeg.Encode(f, grayImg, o); err != nil {
		return err
	}
	output.Stdoutl("[output]", f.Name())

	return nil
}
