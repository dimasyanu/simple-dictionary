package main

import (
	"net/http"

	"github.com/dimasyanu/simple-dictionary/models"
	"github.com/dimasyanu/simple-dictionary/server"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	d, err := server.NewDictionary()
	if err != nil {
		panic(err)
	}

	r.GET("/list", func(c *gin.Context) {
		items, err := d.GetItems()
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, models.ResponseDto[[]models.DictionaryItem]{
			StatusCode: http.StatusOK,
			Message:    "Success",
			Data:       items,
		})
	})

	r.POST("/create", func(c *gin.Context) {
		var model models.DictionaryItem
		if c.Bind(&model) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
			})
			return
		}

		if model.Id == "" || model.Term == "" || model.Description == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
			})
			return
		}

		if !d.Insert(model) {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Success",
		})
	})

	r.PUT("/update", func(c *gin.Context) {
		var model models.DictionaryItem
		if c.Bind(&model) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
			})
			return
		}

		if model.Id == "" || model.Term == "" || model.Description == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
			})
			return
		}

		if !d.Update(model) {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Success",
		})
	})

	return r
}

func main() {
	engine := setupRouter()
	engine.Run(":8000")
}
