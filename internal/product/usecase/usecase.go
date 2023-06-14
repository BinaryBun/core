package usecase

import (
	"Core/config"
	"Core/internal/product"
	"Core/pkg/renderTools"
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
	// TODO: Артём делай тут
	render, err = "Final point!!", nil // example
	return
}

func (p *productUC) renderImage(url string) (render string, err error) {
	return renderTools.ImageRenderByURL(url, 20)
}
