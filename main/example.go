package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ge "github.com/yijun-tang/hello-go/gin_examples"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/getb", ge.GetDataB)
	r.GET("/getc", ge.GetDataC)
	r.GET("/getd", ge.GetDataD)

	r.Run() // listen and serve on 0.0.0.0:8080
}
