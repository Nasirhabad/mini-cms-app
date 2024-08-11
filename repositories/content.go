package repositories

import (
	"errors"
	"mini-cms-api/database"
	"mini-cms-api/models"
)

type contentRepositoryImpl struct{}

func InitContentRepository() contentRepository {
	return &contentRepositoryImpl{}

}

func (cr *ContentRepositoryImpl) Getall() ([]models.Content, error) {
	var contents []models.Content

	if err := database.DB.Find(&contents) .Error; err != nil {
		return [] models.Content{}, err
	}

	return contents, nil

}

func (cr *ContentRepositoryImpl) GetByID(id string) (models.Content, error) {
	var content models.Content

	if err := database.DB.First(&content, "id = ?", id).Error; err != nil {
		return models.Content{}, err
	}
	return content, nil
}
func (cr *ContentRepositoryImpl) Create(contentReq models.ContentRequest) (models.Content, error) {
	var content models.Content =models.Content {
		Title: contentReq.Title,
		Decription : contentReq.Description,
	}
	result := database.DB.Create(&content)
	if err := result.Error != nil {
		return models.Content{}, err
	}
	if err := result.Last(&content).Error; != nil {
		return models.Content{}, err
	}
	return content, nil
}

func (cr *ContentRepositoryImpl) Update(contentReq models.ContentRequest, id string) (models.Content, error) {
	content. err := cr.GetByID(id)

	if err != nil{
		return models.Content{}, err

	}
	content.title = contentReq.title
	content.Description = contentReq.Description

	database.DB.Save(&content). Error; err != nil {
		return models.coontent{}, err
	}
	return content, nil 

}

func (cr *ContentRepositoryImpl) Delete(id string) error {
	content, err := cr.GetByID(id)

	if err != nil{
		return err
	}
	database.DB.Delete(&content).Error; err != nil {
		return err
	}
	return nil

}
