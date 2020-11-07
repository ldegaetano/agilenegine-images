package images

type Service interface {
	GetImages(imagesRequest) imagesResponse
	GetImageByID(imagesRequest) Image
	SearchImages(imagesRequest) imagesResponse
}

type Cache interface {
	GetImagesPage(int) ([]Image, int)
	GetPages() map[int][]Image
}

type service struct {
	cache Cache
}

func NewService(c Cache) service {
	return service{c}
}

func (s service) GetImages(ir imagesRequest) imagesResponse {
	if ir.Page == 0 {
		ir.Page = 1
	}
	images, totalPages := s.cache.GetImagesPage(ir.Page)

	return imagesResponse{
		Pictures:  images,
		Page:      ir.Page,
		PageCount: totalPages,
		HasMore:   ir.Page < totalPages,
	}
}

func (s service) GetImageByID(ir imagesRequest) Image {
	pages := s.cache.GetPages()

	var i Image
	for _, images := range pages {
		i = find(images, "id", ir.ID)
	}

	return i
}

func (s service) SearchImages(ir imagesRequest) imagesResponse {
	pages := s.cache.GetPages()

	var images []Image
	for _, images := range pages {
		images = append(images, findAll(images, ir.Key, ir.ID)...)
	}

	paginated := paginate(images)
	return imagesResponse{
		Pictures:  paginated[ir.Page],
		Page:      ir.Page,
		PageCount: len(paginated),
		HasMore:   ir.Page < len(paginated),
	}
}

func find(il []Image, key, value string) Image {
	for k, i := range il {
		v, ok := i[key]
		if ok && v == value {
			return il[k]
		}
	}
	return Image{}
}

func findAll(il []Image, key, value string) (res []Image) {
	for k, i := range il {
		v, ok := i[key]
		if ok && v == value {
			res = append(res, il[k])
		}
	}
	return
}

func paginate(il []Image) (res map[int][]Image) {
	page := []Image{}
	imageCount := 0
	pageCount := 0
	for k := range il {
		imageCount++
		if imageCount > 10 {
			pageCount++
			res[pageCount] = page
			page = []Image{}
		}
		page = append(page, il[k])
	}
	return
}
