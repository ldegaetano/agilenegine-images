package external

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ldegaetano/agilenegine-images/images"
	"github.com/ldegaetano/go-http-client/rest"
)

type External interface {
	GetPages() (map[int][]images.Image, error)
}

type external struct {
	client rest.HttpClient
	token  string
}

func New() External {
	cfg := rest.ClientCfg{
		BasePath: "http://interview.agileengine.com",
	}
	ex := external{client: rest.NewClient(cfg)}
	token := ex.getToken()
	return external{client: rest.NewClient(cfg), token: token}
}

func (ex external) GetPages() (map[int][]images.Image, error) {
	pages := map[int][]images.Image{}
	n := 1
	next := true
	images := []images.Image{}
	var err error
	for next && err != nil {
		images, next, err = ex.getPage(n)
		pages[n] = images
		n++
	}
	return pages, err
}

func (ex external) getPage(page int) ([]images.Image, bool, error) {
	header := &http.Header{}
	header.Add("Authorization", ex.token)

	uri := fmt.Sprintf("/images?page=%d", page)
	_, body, err := ex.client.GetWithHeader(uri, header)

	response := &imagesResponse{}
	json.Unmarshal(body, response)
	return response.Pictures, response.HasMore, err
}

func (ex external) getToken() string {
	tokenRequest := struct {
		APIKey string `json:"apiKey"`
	}{APIKey: "23567b218376f79d9415"}

	tokenResponse := struct {
		Token string `json:"token"`
		Auth  bool   `json:"auth"`
	}{}

	b, _ := json.Marshal(tokenRequest)
	_, bytes, _ := ex.client.Post("/auth", b)

	json.Unmarshal(bytes, &tokenResponse)

	if tokenResponse.Auth {
		return tokenResponse.Token
	}
	return ""
}

type imagesResponse struct {
	Pictures  []images.Image `json:"pictures"`
	Page      int            `json:"page"`
	PageCount int            `json:"page_count"`
	HasMore   bool           `json:"has_more"`
}
