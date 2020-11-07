package images

type (
	imagesRequest struct {
		Page  int
		ID    string
		Key   string
		Value string
	}
	imagesResponse struct {
		Pictures  []Image `json:"pictures"`
		Page      int     `json:"page"`
		PageCount int     `json:"page_count"`
		HasMore   bool    `json:"has_more"`
	}
	imageResponse struct{}
	Image         map[string]string
)
