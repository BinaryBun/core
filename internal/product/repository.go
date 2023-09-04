package product

type Repository interface {
	GetRequestBody(url string) (body []byte, err error)
}
