package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"messenger/application"
	"messenger/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Blacklist check access to send a message
func Blacklist(c *gin.Context) {
	message := &application.Message{}

	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	err := json.Unmarshal(bodyBytes, message)

	fmt.Println(message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if len(message.UID2) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "undefined user id"})
		c.Abort()
		return
	}

	findUser := &dto.FindUsers{
		UID:       message.UID2,
		BlackList: []string{message.UID1},
	}

	user := application.User{ApplicationID: message.ApplicationID}
	if err := user.FindOne(findUser); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you are in blacklist"})
		c.Abort()
		return
	}

	c.Next()
}
