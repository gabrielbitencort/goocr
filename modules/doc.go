package modules

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

// DocToPng recebe um documento e transforma em uma imagem PNG.
func DocToPng(docPath string) ([]string, error) {
	doc, err := fitz.New(docPath)
	if err != nil {
		return nil, err
	}
	defer doc.Close()

	tmpDir, err := os.MkdirTemp(os.TempDir(), "fitz")
	if err != nil {
		return nil, err
	}

	var imgList []string

	for page := 0; page < doc.NumPage(); page++ {
		img, err := doc.Image(page)
		if err != nil {
			return nil, err
		}

		outputFile := filepath.Join(tmpDir, fmt.Sprintf("page_%d", page))

		file, err := os.Create(outputFile)
		if err != nil {
			return nil, err
		}

		if err = png.Encode(file, img); err != nil {
			return nil, err
		}
		file.Close()

		imgList = append(imgList, outputFile)
	}

	return imgList, nil
}
