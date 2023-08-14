package user

import (
	"context"
	"log"

	"github.com/MuhAndriJP/dealls.git/action/user"
	"github.com/MuhAndriJP/dealls.git/entity"
	"github.com/MuhAndriJP/dealls.git/helper"
	"github.com/labstack/echo/v4"
)

type UserLogin struct {
}

func (u *UserLogin) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	req := new(entity.Users)
	if err = c.Bind(req); err != nil {
		return
	}

	log.Println("Register User Request", req)
	res, err := user.NewUserLogin().Handle(ctx, req)
	if err != nil {
		log.Println("[ERROR] User Login:", err)
		return c.JSON(helper.HTTPStatusFromCode(helper.InvalidArgument), &helper.Response{
			Code:    helper.InvalidArgument,
			Message: helper.StatusMessage[helper.InvalidArgument],
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		})
	}

	return c.JSON(helper.HTTPStatusFromCode(helper.Success), &helper.Response{
		Code:    helper.SuccessCreated,
		Message: helper.StatusMessage[helper.SuccessCreated],
		Data: map[string]interface{}{
			"data": res,
		},
	})
}

func NewUserLogin() *UserLogin {
	return &UserLogin{}
}
