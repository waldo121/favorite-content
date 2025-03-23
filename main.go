package main

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/waldo121/brainrot-index/internal/service"
)

func main() {
	godotenv.Load()
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"safeEmbedURL": func(url url.URL) template.URL {
			return template.URL(url.String())
		},
	})
	router.LoadHTMLGlob("views/*")
	router.Static("/assets", "./assets")
	router.GET("/", func(c *gin.Context) {
		content := service.GetRandomContent()
		embedCode := content.Embed()

		c.HTML(
			http.StatusOK,
			"index",
			gin.H{
				"Title":     "My Favorite Content",
				"EmbedCode": template.HTML(embedCode),
			},
		)

	})
	router.GET("/next", func(c *gin.Context) {
		newContent := service.GetRandomContent()
		embedCode := newContent.Embed()

		c.HTML(
			http.StatusOK,
			"frame",
			gin.H{
				"EmbedCode": template.HTML(embedCode),
			},
		)

	})
	router.POST("/content", func(c *gin.Context) {
		urlForm := c.PostForm("url")
		parsed, err := url.Parse(urlForm)
		if err == nil {
			err = service.AddContent(*parsed)
		}
		c.HTML(
			http.StatusOK,
			"add-notification",
			gin.H{
				"AddContentError": err != nil,
			},
		)
	})
	router.Run()
}
