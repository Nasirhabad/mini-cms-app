package repositories

import "mini-cms-api/models"

type ContentRepository interface {
	Getall() ([]models.Content, error)
	GetByID(id string) (models.Content, error)
	Create(contentReq models.ContentRequest) (models.Content, error)
	Update(contentReq models.ContentRequest, id string) (models.Content, error)
	Delete(id string) error
}

type UserRepository interface {
	Register(registerReq models.RegisterRequest) (models.User, error)
	GetByEmail(loginReq models.LoginRequest) (models.User, error)
	GetUserInfo(id string) (models.User, error)
}

type CategoryRepository interface {
	Getall() ([]models.Category, error)
	GetByID(id string) (models.Category, error)
	Create(contentReq models.CategoryRequest) (models.Category, error)
	Update(contentReq models.CategoryRequest, id string) (models.Category, error)
	Delete(id string) error
}
