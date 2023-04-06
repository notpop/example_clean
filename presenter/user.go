package presenter

import (
	"example_clean/entity"
)

type UserPresenter interface {
	Response(user *entity.User) map[string]interface{}
}

type UserPresenterImpl struct{}

func NewUserPresenter() UserPresenter {
	return &UserPresenterImpl{}
}

func (up *UserPresenterImpl) Response(user *entity.User) map[string]interface{} {
	return map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}
}
