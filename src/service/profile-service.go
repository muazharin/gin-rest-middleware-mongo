package service

import "ginrest/gin-rest-middleware-mongo/src/entity"

type ProfileService interface {
	Save(entity.Profile) entity.Profile
	FindAll() []entity.Profile
}

type profileService struct {
	profile []entity.Profile
}

func New() ProfileService {
	return &profileService{}
}

func (service *profileService) Save(profile entity.Profile) entity.Profile {
	service.profile = append(service.profile, profile)
	return profile
}
func (service *profileService) FindAll() []entity.Profile {
	return service.profile
}
