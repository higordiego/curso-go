package controllers

import (
	"net/http"

	"curso/models"

	"github.com/labstack/echo"
)

// home
func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"menssage": "testando",
	})
}

// Inserir
func Insert(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	if name != "" && email != "" {
		if _, err := models.UserModel.Insert(models.User{
			Name:  name,
			Email: email,
		}); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusCreated, map[string]string{
			"mensagem": "O registro foi armazenado com sucesso!",
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"messagem": "Os campos precisam ser passados!",
	})
}
