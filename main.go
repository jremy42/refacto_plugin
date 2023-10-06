package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

//api startcoder ai

func main() {

	apiKey := ""

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "post",
		})
	})

	r.GET("/chatgpt", func(c *gin.Context) {
		client := &http.Client{}
		req, err := http.NewRequest("POST", apiEndpoint, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
		}
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		fmt.Print(resp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
		}
		defer resp.Body.Close()
		c.JSON(http.StatusOK, resp.Body)
	})

	r.Run()
}
