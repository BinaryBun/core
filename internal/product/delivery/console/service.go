package console

import (
	"Core/config"
	"Core/internal/product"
)

type UserInterface struct {
	cfg      *config.Config
	uc       product.UseCase
	httpRepo product.Repository
	startUrl string
}

func NewUI(
	cfg *config.Config,
	uc product.UseCase,
	httpRepo product.Repository,
	startUrl string,
) *UserInterface {
	return &UserInterface{
		cfg:      cfg,
		uc:       uc,
		httpRepo: httpRepo,
		startUrl: startUrl,
	}
}

func (u *UserInterface) CreateWorker() {
	u.createConsoleWorker()
}
