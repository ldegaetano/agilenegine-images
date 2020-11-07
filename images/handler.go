package images

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type handler struct {
	srv Service
}

func NewHandler(cache Cache) handler {
	return handler{NewService(cache)}
}

func (h handler) GetImages(c echo.Context) error {
	r := imagesRequest{}
	page := c.QueryParam("page")
	if page != "" {
		pageNumber, _ := strconv.ParseInt(page, 10, 64)
		r.Page = int(pageNumber)
	}
	res := h.srv.GetImages(r)
	if len(res.Pictures) == 0 {
		return c.JSON(http.StatusNotFound, "Not found")
	}
	return c.JSON(http.StatusOK, res)
}
