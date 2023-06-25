package product

type UseCase interface {
	RenderImage(url string) (render string, err error)
	RenderBodyToAscii(body string) (render string, links [][2]string, images []string)
}
