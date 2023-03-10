package main

import (
	"fmt"
	iframeparse "main/iframeParse"
	loadfiles "main/loadFiles"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: iframeParser <root directory> <filename>")
		os.Exit(1)
	}

	root := os.Args[1]
	fileName := os.Args[2]

	doc1 := loadfiles.LoadHtml(root, fileName)

	iframeparse.IframeParse(doc1, root)

	outputPath := filepath.Join(root, "output.html")
	htmlFile2, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer htmlFile2.Close()

	err = html.Render(htmlFile2, doc1)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
