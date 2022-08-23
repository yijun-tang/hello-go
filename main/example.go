package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yijun-tang/hello-go/ginexamples"
)

var upgrader = websocket.Upgrader{
	//check origin will check the cross region source (note : please not using in production)
	CheckOrigin: func(r *http.Request) bool {
		//Here we just allow the chrome extension client accessable (you should check this verify accourding your client source)
		return r.Header.Get("Origin") == "chrome-extension://cbcbkhdmedgianpaifchdaddpnmgnknn"
	},
}

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:8899", nil)
	}()

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

	r.GET("/getb", ginexamples.GetDataB)
	r.GET("/getc", ginexamples.GetDataC)
	r.GET("/getd", ginexamples.GetDataD)

	r.GET("/:name/:id", ginexamples.BindUri)

	// TODO:
	// Current: Custome Validators

	r.GET("/", func(c *gin.Context) {
		//upgrade get request to websocket protocol
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ws.Close()
		for {
			//Read Message from client
			mt, message, err := ws.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}
			//If client message is ping will return pong
			if string(message) == "ping" {
				message = []byte("pong")
			}
			//Response message to client
			err = ws.WriteMessage(mt, message)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("inner websocket")
		}
		fmt.Println("stop connection...")
	})

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

/* func main() {

	done := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 2)
		close(done)
	}()

	<-done
	fmt.Println("channel closed...")
} */

// close chan
/* func main() {

	done := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 2)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("channel closed...")
	}

} */
