
"net/http""

type contentRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"massage": "hello world!",
		})
	})
	e.POST("/contents", func(c echo.Context) error {
		var contentRequest contentRequest
		// read the request body from the client
		if err := c.Bind(&contentRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"massage": "invalid",
			})
		}
		// return in to the clie nt
		return c.JSON(http.StatusCreated, map[string]string{
			"massage": "success",
			"Title":   contentRequest.Title,
			"content": contentRequest.Content,
		})

	})
