package delivery

import (
	"net/http"

	"github.com/fajrulaulia/minder/helper"
	"github.com/fajrulaulia/minder/src/usecase"
	"github.com/fajrulaulia/minder/src/usecase/matcher"
	"github.com/labstack/echo/v4"
)

type MatcherDelivery struct {
	Matcher usecase.MatcherUsecaseIface
}

func NewMatcherDelivery(c usecase.MatcherUsecaseIface) MatcherDelivery {
	return MatcherDelivery{
		Matcher: c,
	}
}

func (d *MatcherDelivery) Apply(e *echo.Echo) {
	e.POST("/v1/matcher/action", d.Action, helper.ValidateAuth)
	e.POST("/v1/matcher/list", d.ListRecomendation, helper.ValidateAuth)
}

func (d *MatcherDelivery) Action(c echo.Context) error {
	var (
		err  error
		form = new(matcher.LikePass)
	)

	form.UserEmail = c.Get("email").(string)
	if err := c.Bind(form); err != nil {
		return err
	}

	if err = d.Matcher.LikePassAction(c.Request().Context(), form); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
	})
}

func (d *MatcherDelivery) ListRecomendation(c echo.Context) error {
	var (
		err  error
		form = new(matcher.LikePass)
		res  = []string{}
	)

	if err := c.Bind(form); err != nil {
		return err
	}

	if res, err = d.Matcher.ListPartnerCoupleRecommedation(c.Request().Context(), c.Get("email").(string)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success":  true,
		"partners": res,
	})
}
