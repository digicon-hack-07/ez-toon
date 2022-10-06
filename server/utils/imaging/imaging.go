package imaging

import (
	"image"

	"github.com/digicon-hack-07/ez-toon/server/repository"
)

func GenPageThumbnail(data repository.Page) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 128, 128))
	return img
}
