package http

import (
	"Core/internal/product"
	"github.com/go-resty/resty/v2"
)

type Service struct {
	restyCl *resty.Client
}

func NewHttpService(
	restyCl *resty.Client,
) product.Repository {
	return &Service{
		restyCl: restyCl,
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
		resp, err = s.restyCl.R().Get(url)
	case "POST":
		resp, err = s.restyCl.R().Post(url)
	}

	return
}

func (s *Service) parseBody(resp *resty.Response) (body []byte) {
	return resp.Body()
}
