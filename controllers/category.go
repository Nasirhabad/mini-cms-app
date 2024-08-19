package controllers

import (
	"mini-cms-api/models"
	"mini-cms-api/services"
	"net/http"
	"github.com/go-delve/delve/service"
	"github.com/labstack/echo/v4"
)



type CategoryController struct {
	service services.CategoryService

}

func InitCategoryController() CategoryController {
	return CategoryController{
		service: services.InitCategoryService(),
	}
	
}

func (cc *CategoryController) Getall(c echo.Context) error {
	categories, err := cc.service.Getall()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[]{
			status : "failed",
			Massage :"failed to retrieve categories",
		})
	}

	return c.JSON(http.StatusOK, models.Response[]models.Category{
		status: "success",
		Massage: "all categories",
		Data: categories,
	})
}

func (cc *CategoryController) GetByID (c echo.Context) error {
	CategoryID := C.Param("id")

	category, err := cc.service.GetByID(categoryID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			status : "failed",
			Massage: "category not found",

		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Category]{
		status: "success",
		Massage: "category found",
		Data : category,
	})

}
	
func (cc *CategoryController) Create (c echo.Contex) error {
	var categoryRequest models.CategoryRequest

	if err := c.Bind(&categoryRequest); err != nil{
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			status : "failed",
			Massage: "invalid request",
		})
	}

	err := categoryRequest.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status: "failed",
			Massage: "please fill the all required fields",
		})
	}

	category, err := cc.service.Create(categoryRequest)
	
	if err := nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			status: "failed",
			Massage: "failed to create a categry",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Category]{
		status : " success",
		Massage : "category created",
		Data : category,
	})

}

func (cc *CategoryController)update(c echo.Contex) error {
	categoryID := c.Param("id")

	var categoryRequest models.CategoryRequest

	if err := c.Bind(&categoryrRequest); err != nil{
	return c.JSON(http.StatusBadRequest, models.Response[string]{
		status: "failed",
		Massage: "invalid request",

	})
}

err := categoryRequest.validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status: "failed",
			Massage: "please fill the required fields",
		})
	}


category, err := cc.service.update(categoryRequest,cotegoryID)

if err != nil {
	return c.JSON(http.StatusInternalServerError, models.Response[string]{
		Status: "failed",
		Massage: "failed to update a category",

	})
}

return c.JSON(htt.StatusOK,models.Response[models.Category]{
	status: "success",
	Massage:"category updated",
	Data: category,

})

}

func (cc *categoryController) Delete (c echo.Context) error {
	categoryID := c.Param ("id")

	err := cc.service.Delete(categoryID)

	if err != nil{
		return c.JSON(http.StatusInternalServerError,models.Response[string]{
			status: "failed",
			masMassage: "failed to delete category",
		})
	}
	return c.JSON(http.StatusOK, models.Response[string]{
		status: "success",
		Massage: "category deleted",
		
	})
}

