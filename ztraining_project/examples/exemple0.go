package testing

import "github.com/gin-gonic/gin"

// import "net/http" // optional

// func main() {
// r := gin.Default()
// r.GET("/ping", func(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})
// })
// 	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

func example0() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}