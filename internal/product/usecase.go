package product

type UseCase interface {
	RenderBodyToAscii(body []byte) (render string, err error)
}
