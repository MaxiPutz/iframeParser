package loadfiles

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/net/html"
)

func loadHtml(dest string, src string) *html.Node {
	htmlFile, err := os.Open(dest + src)
	if err != nil {
		fmt.Println("Error: html not found", err)
		os.Exit(1)
	}
	defer htmlFile.Close()
	doc, err := html.Parse(htmlFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return doc
}

func LoadHtml(dest string, src string) *html.Node {
	return loadHtml(dest, src)
}

func loadCss(dest string, src string) string {
	cssFile, err := os.Open(dest + src)
	if err != nil {
		fmt.Println("Error: css not found", err)
		os.Exit(1)
	}
	defer cssFile.Close()

	cssBytes, err := ioutil.ReadAll(cssFile)
	if err != nil {
		fmt.Println("Error: css not bytes are parsed", err)
	}

	cssString := string(cssBytes)

	return cssString
}

func LoadCss(dest string, src string) string {
	return loadCss(dest, src)
}
