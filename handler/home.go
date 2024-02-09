package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/webtaken/go-templ/view/pages"
)

func Home(c echo.Context) error {
	return render(c, pages.Home())
}
