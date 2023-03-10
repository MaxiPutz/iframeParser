package iframeparse

import (
	"fmt"
	"os"
	"strings"

	"main/fp"
	loadfiles "main/loadFiles"

	"golang.org/x/net/html"
)

func IframeParse(n *html.Node, root string) {

	iframeParse(n, root)

	deleteTag(n)

	arr := make([]html.Node, 0)
	getLink(n, &arr)

	var maxiVal = []string{}

	for i, ele := range arr {

		for j, atr := range ele.Attr {
			if atr.Key == "href" {
				str := ele.Attr[j].Val
				str = strings.TrimPrefix(str, "../")
				ele.Attr[j].Val = str

				if !isStingInArr(str, maxiVal) {
					maxiVal = append(maxiVal, str)

				}
				// fmt.Println(len(maxiVal))
			}
		}
		arr[i] = ele
	}

	maxiVal = fp.Map(maxiVal, func(ele string) string {
		return loadfiles.LoadCss(root, ele)
	})

	styleTag := &html.Node{
		Type: html.ElementNode,
		Data: "style",
		FirstChild: &html.Node{
			Type: html.TextNode,
			Data: strings.Join(maxiVal, "\n"),
		},
	}

	head := getFirstHeader(n)

	n.LastChild.AppendChild(creatDiv())

	divHeadStyle := creatDiv()
	divHeadStyle.AppendChild(styleTag)
	head.AppendChild(divHeadStyle)

	fmt.Println(head.Data)

	removeLink(n)
}

func removeLink(n *html.Node) {
	if n == nil {
		return
	}
	var prev *html.Node
	for ele := n.FirstChild; ele != nil; ele = ele.NextSibling {
		if ele.Data == "link" {
			if prev != nil {
				prev.NextSibling = ele.NextSibling
			} else {
				n.FirstChild = ele.NextSibling
			}
			continue
		}
		prev = ele
		removeLink(ele)
	}
}

type foundHeader struct {
	isFound bool
	n       *html.Node
}

func getFirstHeader(n *html.Node) *html.Node {
	tmp := &html.Node{
		Data: "div",
	}
	fh := &foundHeader{false, tmp}

	_getFirstHeader(n, fh)

	if fh.isFound == false {
		fmt.Println("no header is found")
		os.Exit(1)
	}
	fmt.Println(fh.n.Data)
	return fh.n
}

func _getFirstHeader(n *html.Node, fh *foundHeader) {
	if fh.isFound {
		return
	}
	for ele := n.FirstChild; ele != nil; ele = ele.NextSibling {
		if ele.Data == "head" {
			fh.isFound = true
			fh.n = ele
			return
		}
		_getFirstHeader(ele, fh)
	}
}

func isStingInArr(str string, arr []string) bool {
	for _, ele := range arr {

		if ele == str {
			return true
		}
	}

	return false
}

func getLink(n *html.Node, arr *[]html.Node) {
	if n == nil {
		return
	}
	for ele := n.FirstChild; ele != nil; ele = ele.NextSibling {
		if ele.Data == "link" {

			*arr = append(*arr, *ele)

		}
		getLink(ele, arr)
	}
}

func deleteTag(n *html.Node) {
	if n == nil {
		return
	}
	for ele := n.FirstChild; ele != nil; ele = ele.NextSibling {
		if ele.Data == "iframe" {
			ele.Data = "div"
			ele.Attr = []html.Attribute{{Key: "class", Val: "iframe"}}
		}
		// if ele.Data == "html" {
		// 	ele.Data = "div"
		// 	ele.Attr = []html.Attribute{{Key: "class", Val: "embedded"}}
		// }
		deleteTag(ele)
	}

}

func iframeParse(n *html.Node, root string) *html.Node {

	if n == nil {
		return nil
	}

	if n.Type == html.ElementNode && n.Data == "iframe" {
		str := getAttrValue(n, "src")

		doc := loadfiles.LoadHtml(root+"/", str)
		// div := creatDiv()

		iframeRoot := creatDiv()
		iframeRoot.Attr = []html.Attribute{{Key: "class", Val: "iframe"}}

		n.AppendChild(doc)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		iframeParse(c, root)
	}
	return n
}

func getAttrValue(n *html.Node, attrName string) string {
	for _, a := range n.Attr {
		if a.Key == attrName {
			return a.Val
		}
	}
	fmt.Println("node with no attribut ", attrName, " found")
	os.Exit(1)
	return ""
}

func creatDiv() *html.Node {
	return &html.Node{
		Data: "div",
		Type: html.ElementNode,
		Attr: []html.Attribute{
			{
				Key: "class",
				Val: "iframe",
			},
		},
	}
}

func PrintHtml(n *html.Node) {
	if n == nil {
		return
	}
	fmt.Println(n.DataAtom)

	for ele := n.FirstChild; ele != nil; ele = ele.NextSibling {
		PrintHtml(ele)
	}
}
