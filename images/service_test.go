package images

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

type cacheMock struct {
	mock.Mock
}

func (m *cacheMock) GetImagesPage(page int) ([]Image, int) {
	res := m.Called(page)
	return res.Get(0).([]Image), res.Int(1)
}
func (m *cacheMock) GetPages() map[int][]Image {
	res := m.Called()
	return res.Get(0).(map[int][]Image)
}

func TestGetImages(t *testing.T) {
	cache := &cacheMock{}
	srv := NewService(cache)
	imageMock := []Image{
		Image{"id": "432432",
			"author": "Author"},
	}
	page := 2
	totalPages := 4
	cache.On("GetImagesPage", page).Return(imageMock, totalPages)
	res := srv.GetImages(imagesRequest{Page: page})

	assert.Equal(t, totalPages, res.PageCount)
	assert.Equal(t, 2, res.Page)
	assert.Equal(t, true, res.HasMore)
	assert.Equal(t, "432432", res.Pictures[0]["id"])
}

func TestGetImagesByID(t *testing.T) {
	cache := &cacheMock{}
	srv := NewService(cache)
	imageID := "432432"
	imageMock := []Image{
		Image{"id": "432432",
			"author": "Author"},
	}
	pagesMock := map[int][]Image{
		1: imageMock,
	}
	cache.On("GetPages").Return(pagesMock)
	res := srv.GetImagesByID(imagesRequest{ID: imageID})

	assert.Equal(t, imageID, res["id"])
	assert.Equal(t, "Author", res["author"])
}
