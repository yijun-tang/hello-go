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

	r.GET("/:name/:id", ge.BindUri)

	// TODO:
	// Current: Custome Validators

	r.Run() // listen and serve on 0.0.0.0:8080
}

/* var (
	jobMap  map[int]int = make(map[int]int)
	idx     int         = 0
	jobChan chan int    = make(chan int)
)

func main() {
	fmt.Println("started....")
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			jobChan <- idx
			idx++
		}
	}()

	go func() {
		for {
			select {
			case idx := <-jobChan:
				if idx == 1 {
					fmt.Println("continue...")
					continue
				}
				Job(idx, "insert")
			case <-time.After(2 * time.Second):
				for idx := range jobMap {
					Job(idx, "update")
				}
			}
		}
	}()

	<-make(chan int)
}

func Job(id int, jobType string) {
	time.Sleep(time.Second)
	jobMap[id] = id
	fmt.Printf("[%v]: Job %v has finished...\n", jobType, id)
} */
