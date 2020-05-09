package action

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var body []byte

func Test(c *gin.Context) {
	var err error
	body, err = ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	checkResult := "success"
	traceid := c.Request.Header.Get("traceid")
	if traceid != string(body) {
		checkResult = "mismatch"
	}
	c.JSON(200, gin.H{
		"message": "pong",
		"traceid": traceid,
		"body":    string(body),
		"result":  checkResult,
	})
}
