package handler

import (
	"net/http"
	"net/mail"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/webtaken/go-templ/view/pages"
)

func SignUp(c echo.Context) error {
	return render(c, pages.SignUp())
}

type RegisterFormData struct {
	Email     string `form:"email" json:"email"`
	Username  string `form:"username" json:"username"`
	Firstname string `form:"firstname" json:"firstname"`
	Lastname  string `form:"lastname" json:"lastname"`
	Phone     string `form:"phone" json:"phone"`
	Password1 string `form:"password1" json:"password1"`
	Password2 string `form:"password2" json:"password2"`
}

func SignUpValidation(c echo.Context) error {
	alreadyRegisteredUsers := []string{"koki", "popi", "unknown"}
	field := c.QueryParam("field")
	var registerData RegisterFormData
	err := c.Bind(&registerData)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	if field == "email" {
		_, err := mail.ParseAddress(registerData.Email)
		if err != nil {
			return c.String(http.StatusOK, "Your e-mail is invalid")
		}
	} else if field == "username" {
		if registerData.Username == "" {
			return c.String(http.StatusOK, "Username cannot be empty")
		}
		for _, username := range alreadyRegisteredUsers {
			if username == registerData.Username {
				return c.String(http.StatusOK, "This username is already registered")
			}
		}
	} else if field == "firstname" {
		if registerData.Firstname == "" {
			return c.String(http.StatusOK, "Firstname cannot be empty")
		}
	} else if field == "lastname" {
		if registerData.Lastname == "" {
			return c.String(http.StatusOK, "Lastname cannot be empty")
		}
	} else if field == "phone" {
		if registerData.Phone == "" {
			return c.String(http.StatusOK, "Phone cannot be empty")
		}
		if _, err := strconv.Atoi(registerData.Phone); err != nil {
			return c.String(http.StatusOK, "Doesn't look like a Phone")
		}
	} else if field == "password1" {
		eightOrMore, number, upper, special := verifyPassword(registerData.Password1)
		if !eightOrMore || !number || !upper || !special {
			return c.String(
				http.StatusOK,
				"Password rules:<br/>at least 8 characters long<br/>at least 1 number<br/>at least 1 uppercase<br/>at least 1 special character",
			)
		}
	} else if field == "password2" {
		if registerData.Password2 != registerData.Password1 {
			return c.String(http.StatusOK, "Passwords don't match")
		}
	}
	return c.String(http.StatusOK, "")
}
