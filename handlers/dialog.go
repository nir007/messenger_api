package handlers

import (
	"messenger/drepository"
	"messenger/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindAllDialogs handles http request
func FindAllDialogs(c *gin.Context) {
	find := &dto.FindDialogs{}

	err := c.ShouldBind(find)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if uid := c.Param("uid"); len(uid) > 0 {
		find.UID1 = uid
		find.ApplicationID = c.Request.Header["Application-Id"][0]
	}

	dialog := &drepository.Dialog{ApplicationID: find.ApplicationID}
	dialogs, total, err := dialog.Find(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": dialogs, "total": total})
}

// FindDialog handles http request
func FindDialog(c *gin.Context) {
	find := &dto.FindDialogs{}
	err := c.ShouldBind(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	dialog := &drepository.Dialog{ApplicationID: find.ApplicationID}

	err = dialog.FindOne(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": dialog})
}
