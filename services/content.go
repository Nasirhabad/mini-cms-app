package services

import (
	"mini-cms-api/models"
	"mini-cms-api/repositories"
)

type Contentservice struct {
	repository repositories.ContentRepository
}

func Initcontentservice() Contentservice {
	return Contentservice{
		repository: &repositories.ContentRepositoryImpl{},
	}
}

func (cs *Contentservice) Getall() ([]models.Content, error) {
	return cs.repository.Getall()
}

func (cs *Contentservice) GetByID(id string) (models.Content, error) {
	return cs.repository.GetByID(id)
}

func (cs *Contentservice) Create(contentReq models.ContentRequest) (models.Content, error) {
	return cs.repository.Create(contentReq)
}

func (cs *Contentservice) Update(contentReq models.ContentRequest, id string) (models.content, error) {
	return cs.repository.Update(contentReq)
}

func (cs *Contentservice) Delete(id string) error {
	return &cs.repository.Delete(id)
}
