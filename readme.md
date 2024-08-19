introduction about the echo
...
   go
...
type ContentRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`

	// demo only!
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello world",
		})

	})

	e.POST("/contents", func(c echo.Context) error {
		var contentRequest ContentRequest
		// read the request from the client
		if err := c.Bind(&contentRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "invalid request",
			})
		}

		if contentRequest.Title == "" || contentRequest.Content == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "input cannot be empty",
			})
		}

		// return it to the client
		return c.JSON(http.StatusCreated, map[string]string{
			"message": "success",
			"title":   contentRequest.Title,
			"content": contentRequest.Content,
		})
	})
