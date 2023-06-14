package renderTools

import (
	"fmt"
	"image"
	"io"
	"net/http"
)

func ImageRenderByURL(url string, maxSize int) (asciiImage string, err error) {
	image, err := downloadImage(url)
	if err != nil {
		fmt.Println("Не удалось загрузить изображение:", err)
		return
	}

	// Сжимаем изображение до 10 пикселей в длину
	compressedImage := resizeImage(image, maxSize, 0)

	// Преобразуем изображение в ASCII-графику
	asciiImage = convertToASCII(compressedImage)

	// Выводим ASCII-графику
	return asciiImage, nil

}

// Загружает изображение по ссылке и возвращает объект image.Image
func downloadImage(url string) (img image.Image, err error) {
	// Отправляем GET-запрос для загрузки изображения
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	// Декодируем изображение
	img, _, err = image.Decode(response.Body)
	if err != nil {
		return
	}

	return
}

// Изменяет размер изображения
func resizeImage(img image.Image, width, height int) image.Image {
	bounds := img.Bounds()
	imgWidth := bounds.Max.X
	imgHeight := bounds.Max.Y

	// Рассчитываем пропорциональную высоту, если она не указана
	if height == 0 {
		ratio := float64(imgHeight) / float64(imgWidth)
		height = int(float64(width) * ratio)
	}

	// Изменяем размер изображения
	resizedImg := image.NewRGBA(image.Rect(0, 0, width, height))
	resizedBounds := resizedImg.Bounds()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			xOrig := x * imgWidth / width
			yOrig := y * imgHeight / height
			resizedImg.Set(x, y, img.At(resizedBounds.Min.X+xOrig, resizedBounds.Min.Y+yOrig))
		}
	}

	return resizedImg
}

// Преобразует изображение в ASCII-графику с окрашиванием ANSI-кодами
func convertToASCII(img image.Image) string {
	// Определяем размеры изображения
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	// Палитра символов
	asciiChars := []string{"██", "██", "██", "██", "██", "██", "██", "██", "██", "██", "██"}
	charCount := len(asciiChars)

	// Результат ASCII-графики
	asciiImage := ""

	// Преобразуем каждый пиксель в символ ASCII с окрашиванием
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Получаем цвет пикселя
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()

			// Вычисляем индекс символа в палитре ASCII
			grayValue := (r + g + b) / 3
			charIndex := int(grayValue) * charCount / 65535

			// Проверяем, что индекс находится в допустимых границах
			if charIndex >= charCount {
				charIndex = charCount - 1
			}

			// Добавляем символ с окрашиванием ANSI-кодом
			char := asciiChars[charIndex]
			asciiImage += fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r>>8, g>>8, b>>8, char)
		}
		asciiImage += "\n"
	}

	return asciiImage
}
