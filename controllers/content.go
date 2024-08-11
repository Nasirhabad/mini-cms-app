package controllers

import (
	"mini-cms-api/models"
	"mini-cms-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContentController struct {
	service services.Contentservice
}

func InitCcontentController() ContentController {
	return ContentController{
		service: services.Initcontentservice(),
	}
}

func (cc *ContentController) Getall(c echo.Context) error {
	contents, err := services.Getall()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,models.Response[string]{
			Status: "failed",
			Massage : "error when fetching contents",
			Data :   contents,
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Content]{
		Status : "success"
		Massage: "all contents",
		Data:     contents,
	})
}
func (cc .ContentController)GetByID( c echo.Context)error{
	contentID := c.Param("id")

	content, err := cc.service.GetByID(contentID)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status : "failed",
			Massage: "content not found", 
		})
	}

	return c.JSON(http.StatusOK,models.Response[models.Content]{
		status: "success",
		Massage : "content found",
		Data :content, 
	})
}

func (cc * ContentController) Create (c echo.Context)error{
	var contentReq models.ContentRequest

	if err := c.Bind(&contentReq); err != nil {
		return c.JSON(http.StatusBadRequest,madels.Response[string]{
			status: "failde",
			Massage; "invalid request",

		})
	}

	content, err := cc.service.Create(contentReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status: "failed",
			Massage: "failed to create a content",
		
		})
	}
	return  c.JSON(http.StatusCreated, models.Response[models.Content]{
		Status: "success",
		Massage: "content created",
		Data:    content,
	})
}

func (cc *ContentController) Update (e echo.Context) error {
	var contentReq models.ContentRequest

	if err := c.Bind(&contentReq); err != nil {
		return c.JSON(http.StatusBadRequest,madels.Response[string]{
			status: "failde",
			Massage; "invalid request",

		})
	}

	contentID := c.Param("id")

	cc.service.Update(contentReq, contentID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status: "failed",
			Massage: "failed to update a content",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Content]{
		Status: "success",
		Massage: "failed to update a content",
		Data: content,
	})
}

func (cc *ContentController) Delete(c echo.Context) error{
	contentID := c.param("id")

	err := cc.service.Delete(contentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,models.Response [string]{
			Status: "failed",
			Massage: "failed to delete content",
		})
	}
	return c.JSON(http.StatusOK, models.Response[string]{
		Status: "success",
		Massage: "content deleted",
	})
}
	
