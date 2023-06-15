package usecase

import (
	"Core/config"
	"Core/internal/product"
	"bytes"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"strings"
)

type productUC struct {
	cfg *config.Config
}

func NewProductUC(
	cfg *config.Config,
) product.UseCase {
	return &productUC{
		cfg: cfg,
	}
}

func (p *productUC) RenderBodyToAscii(body []byte) (render string, err error) {
	mainNode, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Errorf("productUC.RenderBodyToAscii()->html.Parse(): %v", err)
	}

	render = parseLinks(mainNode)
	return
}

func parseLinks(node *html.Node) string {
	var text string
	if node.Type == html.ElementNode {
		buf := &bytes.Buffer{}
		collectText(node, buf)
		text += strings.TrimSpace(buf.String())

	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		text += parseLinks(c)
	}

	return text
}

func collectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, buf)
	}
}

// TODO: Максим сделает потом, пока что не смотри
/*func (p *productUC) renderImage() {

}*/
