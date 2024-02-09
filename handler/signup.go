package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/webtaken/go-templ/view/pages"
)

func SignUp(c echo.Context) error {
	return render(c, pages.SignUp())
}
