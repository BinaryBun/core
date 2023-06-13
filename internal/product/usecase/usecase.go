package usecase

import (
	"Core/config"
	"Core/internal/product"
	"Core/internal/product/delivery/http"
	log "github.com/sirupsen/logrus"
)

type productUC struct {
	cfg         *config.Config
	HTTPService *http.Service
}

func NewProductUC(
	cfg *config.Config,
	httpServ *http.Service,
) product.UseCase {
	return &productUC{
		cfg:         cfg,
		HTTPService: httpServ,
	}
}

func (p *productUC) RenderBodyToAscii(body []byte) (render string) {
	return
}

func (p *productUC) Print() {
	log.Info("Oll good")
}
