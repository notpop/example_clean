package controller

import (
	"example_clean/interactor"
	"example_clean/presenter"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userInteractor interactor.UserInteractor
	userPresenter  presenter.UserPresenter
}

func NewUserController(ui interactor.UserInteractor, up presenter.UserPresenter) *UserController {
	return &UserController{userInteractor: ui, userPresenter: up}
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, _ := uc.userInteractor.GetUser(id)
	// if err != nil {
	// 	return
	// }

	c.JSON(200, uc.userPresenter.Response(user))
}
