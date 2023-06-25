package renderTools

import (
	"regexp"
	"strings"
)

func TextRenderByRegexp(body string) (render string, links [][2]string, images []string) {
	render = textTag(body)
	links = linksTag(body)
	images = imagesTag(body)
	return
}

func textTag(body string) (renderString string) {
	regex := `<[^h].*?>` // поиск всех тегов кроме заголовком
	re := regexp.MustCompile(regex)
	renderString = re.ReplaceAllString(body, " ")
	renderString = strings.ReplaceAll(renderString, `\n`, "\n") // редактирование для консоли
	return renderString
}

func imagesTag(body string) (renderImagesLink []string) {
	regex := `<img[^>]*src=["'](http?://[^"']{4,})["'][^>].*?>` // выборка изображений

	re := regexp.MustCompile(regex)
	matches := re.FindAllStringSubmatch(body, -1)

	if len(matches) > 0 {
		for _, match := range matches {
			if len(match) >= 2 {
				url := match[1]
				renderImagesLink = append(renderImagesLink, url)
			}
		}
	}
	return renderImagesLink
}

func linksTag(body string) (renderLinks [][2]string) {
	regex := `<a[^>]*href=["'](https?://[^"']{4,})["'][^>]*>(.*?)</a>` // выбирается всё

	re := regexp.MustCompile(regex)
	matches := re.FindAllStringSubmatch(body, -1)

	if len(matches) > 0 {
		for _, match := range matches {
			if len(match) >= 3 {
				url := match[1]
				text := match[2]
				if regexp.MustCompile("<[^>]*>").MatchString(text) {
					continue
				}
				renderLinks = append(renderLinks, [2]string{text, url})
			}
		}
	}
	return renderLinks
}
