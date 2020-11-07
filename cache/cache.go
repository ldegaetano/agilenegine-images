package cache

import (
	"github.com/ldegaetano/agilenegine-images/external"
	"github.com/ldegaetano/agilenegine-images/images"
)

type cache struct {
	external external.External
	pages    map[int][]images.Image
}

//New return a cache implemntation
func New() (images.Cache, error) {
	c := cache{}
	err := c.Load()
	return c, err
}

func (c cache) GetImagesPage(page int) ([]images.Image, int) {
	return c.pages[page], len(c.pages)
}

func (c cache) GetPages() map[int][]images.Image {
	return c.pages
}

func (c cache) Load() error {
	pages, err := c.external.GetPages()
	c.pages = pages
	return err
}
