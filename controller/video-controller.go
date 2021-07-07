package controller

import (
	"net/http"

	"gin-test/entity"
	"gin-test/service"
	"gin-test/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

)

type VideoController interface {
	// Save(ctx *gin.Context) entity.Video
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService //interface
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	// return &controller{
	// 	service: service,
	// }
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error{
	var video entity.Video
	// ctx.BindJSON(&video)  //json=>struct
	err := ctx.BindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title" : "Video Page",
		"videos" : videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
