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
	Firstname string `form:"firstname" json:"firstname"`
	Lastname  string `form:"lastname" json:"lastname"`
	Phone     string `form:"phone" json:"phone"`
	Password1 string `form:"password1" json:"password1"`
	Password2 string `form:"password2" json:"password2"`
}

func SignUpValidation(c echo.Context) error {
	alreadyRegisteredEmails := []string{
		"koki@gmail.com", "popi@hotmail.com", "babyBoomer123@outlook.com",
	}
	field := c.QueryParam("field")
	var registerData RegisterFormData
	err := c.Bind(&registerData)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	if field == "email" {
		for _, email := range alreadyRegisteredEmails {
			if email == registerData.Email {
				return c.String(http.StatusOK, "<span id=\"email-errors\" class=\"text-sm text-error\">This username is already registered</span>")
			}
		}

		_, err := mail.ParseAddress(registerData.Email)
		if err != nil {
			return c.String(http.StatusOK, "<span id=\"email-errors\" class=\"text-sm text-error\">Your e-mail is invalid</span>")
		}

		return c.String(http.StatusOK, "<span id=\"email-errors\" class=\"text-sm text-success\">Correct email</span>")
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
