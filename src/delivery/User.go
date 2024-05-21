package delivery

import (
	"net/http"

	"github.com/fajrulaulia/minder/src/usecase"
	User "github.com/fajrulaulia/minder/src/usecase/user"
	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	User usecase.UserUsecaseIface
}

func NewUserDelivery(c usecase.UserUsecaseIface) UserDelivery {
	return UserDelivery{
		User: c,
	}
}

func (d *UserDelivery) Apply(e *echo.Echo) {
	e.POST("/v1/signup", d.Signup)
	e.POST("/v1/login", d.Login)
}

func (d *UserDelivery) Signup(c echo.Context) error {
	var (
		err   error
		token = ""
		form  = new(User.UserRequest)
	)

	if err := c.Bind(form); err != nil {
		return err
	}

	if token, err = d.User.Signup(c.Request().Context(), form); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	c.Response().Header().Set("Authorization", "Bearer "+token)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
	})
}

func (d *UserDelivery) Login(c echo.Context) error {
	var (
		err   error
		token = ""
		form  = new(User.LoginRequest)
	)

	if err := c.Bind(form); err != nil {
		return err
	}

	if token, err = d.User.Login(c.Request().Context(), form); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	c.Response().Header().Set("Authorization", "Bearer "+token)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
	})
}
