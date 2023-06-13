package usecase

import (
	"Core/config"
	"Core/internal/product"
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

// TODO: Максим сделает потом, пока что не смотри
/*func (p *productUC) renderImage() {

}*/
