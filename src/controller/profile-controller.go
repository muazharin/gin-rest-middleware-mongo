package controller

import (
	"ginrest/gin-rest-middleware-mongo/src/entity"
	"ginrest/gin-rest-middleware-mongo/src/service"

	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	FindAll() []entity.Profile
	Save(c *gin.Context) entity.Profile
}

type profileController struct {
	service service.ProfileService
}

func New(service service.ProfileService) ProfileController {
	return &profileController{
		service: service,
	}
}
func (pc *profileController) Save(c *gin.Context) entity.Profile {
	var profile entity.Profile
	c.ShouldBind(&profile)
	pc.service.Save(profile)
	return profile
}
func (pc *profileController) FindAll() []entity.Profile {
	return pc.service.FindAll()
}
