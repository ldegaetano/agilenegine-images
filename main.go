package main

import (
	"fmt"

	"github.com/ldegaetano/agilenegine-images/cache"
	"github.com/ldegaetano/agilenegine-images/images"

	"github.com/labstack/echo"
)

func main() {
	c, err := cache.New()
	if err != nil {
		fmt.Printf("Could not initialize the cache: %s", err.Error())
		return
	}
	h := images.NewHandler(c)
	e := echo.New()
	e.GET("/images", h.GetImages)
	e.Logger.Fatal(e.Start(":8080"))
}
