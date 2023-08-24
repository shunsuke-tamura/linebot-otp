package handler

import (
	"line-bot-otp-back/logic"
	"line-bot-otp-back/model"
	"line-bot-otp-back/util"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

type UserCreateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (*UserHandler) Create(ctx *gin.Context) {
	var req UserCreateRequest
	err := ctx.Bind(&req)
	if err != nil {
		s := err.Error()
		r := util.BadRequest(&s)
		log.Errorln("[Error]request parse error: ", s)
		ctx.JSON(r.StatusCode, r.Message)
		return
	}
	ul := logic.UserLigic{User: &model.User{Name: req.Name, Password: req.Password}}
	err = ul.Create()
	if err != nil {
		log.Errorln("[Error]exec error: ", err.Error())
		r := util.InternalServerError(nil)
		ctx.JSON(r.StatusCode, r.Message)
		return
	}
	r := util.Ok(nil)
	ctx.JSON(r.StatusCode, r.Message)
}