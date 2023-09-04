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
func (p *productUC) RenderBodyToAscii(body string) (render string, links [][2]string, images []string) {
	// TODO: Возвращать из packege renderString, ArrImagesLinks.
	/* TODO: For _, imageLink := arrImagesLinks{

	_, _ :=  p.RenderImages(imageLink)

	}*/
	return renderTools.TextRenderByRegexp(body)
}

func (p *productUC) RenderImage(url string) (render string, err error) {
	return renderTools.ImageRenderByURL(url, 20)
}
