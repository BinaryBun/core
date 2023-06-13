package http

import (
	"Core/config"
	"github.com/go-resty/resty/v2"
)

type Service struct {
	cfg   *config.Config
	resty *resty.Client
}

func NewHttpService(
	cfg *config.Config,
	resty *resty.Client,
) *Service {
	return &Service{
		cfg:   cfg,
		resty: resty,
	}
}

func (s *Service) GetRequestBody(url string) (body []byte, err error) {
	resp, err := s.makeRequest(url, "GET")
	if err != nil {
		return nil, err
	}
	body = s.parseBody(resp)
	return
}

func (s *Service) makeRequest(url string, method string) (resp *resty.Response, err error) {
	switch method {
	case "GET":
		resp, err = s.resty.R().Get(url)
	case "POST":
		resp, err = s.resty.R().Post(url)
	}

	return
}

func (s *Service) parseBody(resp *resty.Response) (body []byte) {
	return resp.Body()
}
