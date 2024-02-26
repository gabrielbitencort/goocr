package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabrielbitencort/goocr/modules"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	docExtensions := []string{".pdf", ".epub", ".mobi", ".odt", ".docx", ".doc", ".txt", ".html"}
	imageExtensions := []string{".png", ".jpeg", ".jp2", ".tiff", ".gif", ".webp", ".bmp", ".pnm"}

	client := gosseract.NewClient()
	defer client.Close()

	client.SetLanguage("por")

	fmt.Print("File Path: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		file := scanner.Text()

		fmt.Println(file)

		if file != "" {
			if hasExtension(file, docExtensions) {
				imgList, err := modules.DocToPng(file)
				if err != nil {
					fmt.Println("Erro ao transformar documento em imagens: ", err)
					return
				}

				for _, img := range imgList {
					client.SetImage(img)

					text, err := client.Text()
					if err != nil {
						fmt.Println("Erro ao executar OCR na imagem: ", err)
						return
					}

					fmt.Println(text)

					if err := modules.SaveText(text); err != nil {
						fmt.Println("Erro ao salvar o texto: ", err)
						return
					}
				}
			} else if hasExtension(file, imageExtensions) {
				client.SetImage(file)

				text, err := client.Text()
				if err != nil {
					fmt.Println("Erro ao executar OCR na imagem: ", err)
					return
				}

				fmt.Println(text)

				if err := modules.SaveText(text); err != nil {
					fmt.Println("Erro ao salvar o texto: ", err)
					return
				}
			} else {
				fmt.Println("Unsupported file format!")
				fmt.Println("Supported document formats: ", docExtensions)
				fmt.Println("Supported image formats: ", imageExtensions)
				return
			}
		} else {
			fmt.Println("Please enter a file path")
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}
}

// hasExtension checks if the file extension is in the list of valid extensions
func hasExtension(file string, extensions []string) bool {
	fileExt := strings.ToLower(filepath.Ext(file))
	for _, ext := range extensions {
		if fileExt == ext {
			return true
		}
	}

	return false
}
