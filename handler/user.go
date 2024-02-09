package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/webtaken/go-templ/model"
	"github.com/webtaken/go-templ/view/user"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	userData := model.User{Email: "luckly083@gmail.com"}
	return render(c, user.Show(userData))
}
