package iframePackage

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func findAllIframes2(n *html.Node) []*html.Node {
	var iframes []*html.Node

	if n == nil {
		return iframes
	}

	if n.Type == html.ElementNode && n.Data == "iframe" {
		iframes = append(iframes, n)
		str := getAttrValue(n, "src")
		doc := loadFile("./html/", str)
		iframeChildren := findAllIframes2(doc)
		iframes = append(iframes, iframeChildren...)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		iframeChildren := findAllIframes2(c)

		iframes = append(iframes, iframeChildren...)
	}

	return iframes
}

func getAttrValue(n *html.Node, attrName string) string {
	print(n.Data)
	for _, a := range n.Attr {
		if a.Key == attrName {
			return a.Val
		}
	}
	fmt.Println("node with no attribut ", attrName, " found")
	os.Exit(1)
	return ""
}

func findAllIframes(n *html.Node) []*html.Node {
	var iframes []*html.Node

	if n == nil {
		return iframes
	}

	if n.Type == html.ElementNode && n.Data == "iframe" {
		iframes = append(iframes, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		iframeChildren := findAllIframes(c)
		iframes = append(iframes, iframeChildren...)
	}

	return iframes
}

func FindAllIframes(n *html.Node) []*html.Node {
	temp := findAllIframes2(n)

	if temp == nil {
		fmt.Println("Error: 99")
		os.Exit(1)
	}

	return temp
}

func loadFile(dest string, src string) *html.Node {
	htmlFile, err := os.Open(dest + src)
	if err != nil {
		fmt.Println("Error:", err)
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

func LoadFile(dest string, src string) *html.Node {
	return loadFile(dest, src)
}

func GetAttrValue(n *html.Node, attrName string) string {
	return getAttrValue(n, attrName)
}
