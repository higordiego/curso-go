package controllers

import (
	"net/http"
	"strconv"

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

// listar
func Listar(c echo.Context) error {
	var user []models.User
	if err := models.UserModel.Find().All(&user); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{
			"message": "Error em retornar os dados!",
		})
	}
	if user != nil {
		return c.JSON(http.StatusAccepted, user)
	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": "Não conseguimos achar o que procura!",
	})
}

// deletar
func Deletar(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))

	resultado := models.UserModel.Find("id = ?", usuarioID)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Usuário não encontrado!",
		})
	}

	if err := resultado.Delete(); err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"message": "Não foi possivel deletar o usuário!",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"message": "usuário deletado com sucesso!",
	})
}

// Alterar
func Alterar(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	email := c.FormValue("email")
	resultado := models.UserModel.Find("id=?", usuarioID)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusBadGateway, map[string]string{
			"mensagem": "Usuário não encontrado!",
		})
	}

	if err := resultado.Update(models.User{
		ID:    usuarioID,
		Name:  name,
		Email: email,
	}); err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{
			"messagem": "Error em alterar o registro",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"menssagem": "Usuário modificado com sucesso!",
	})

}
