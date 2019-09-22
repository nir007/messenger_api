package handlers

import (
	"bytes"
	"io"
	"messenger/drepository"
	"messenger/dto"
	"messenger/s3"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateManager creates new manager
func CreateManager(c *gin.Context) {
	manager := &drepository.Manager{}
	err := c.Bind(manager)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "binding params: " + err.Error(),
		})
		return
	}

	_, err = manager.Insert()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	manager.Password = ""
	c.JSON(http.StatusOK, gin.H{"result": manager})
}

// FindOneManager search manager
func FindOneManager(c *gin.Context) {
	id, _ := c.Get("managerID")
	objectID, _ := primitive.ObjectIDFromHex(id.(string))

	find := &dto.FindManagers{ID: objectID}
	err := c.ShouldBind(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	manager := &drepository.Manager{}
	err = manager.FindOne(find)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": manager})
}

// FindAllManagers search managers
func FindAllManagers(c *gin.Context) {}

// UpdateManager changes manager
func UpdateManager(c *gin.Context) {
	updateManager := &dto.UpdateManager{}
	err := c.ShouldBind(updateManager)

	id, _ := c.Get("managerID")
	objectID, _ := primitive.ObjectIDFromHex(id.(string))

	findMenager := &dto.FindManagers{ID: objectID}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	manager := &drepository.Manager{}
	_, err = manager.Update(findMenager, updateManager)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	manager.Password = ""

	c.JSON(http.StatusOK, gin.H{"result": manager})

}

//UpdateManagerAvatar updates manager avatar
func UpdateManagerAvatar(c *gin.Context) {
	locationFull := ""
	previewImage := ""
	pathName := "avatars"

	formFile, _ := c.FormFile("file")

	file, _ := formFile.Open()
	defer file.Close()

	if isImage(formFile.Header["Content-Type"][0]) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "require jpeg or png format"})
		c.Abort()
		return
	}

	var copiedFile io.Reader
	var buf bytes.Buffer

	copiedFile = io.TeeReader(file, &buf)
	preview, err := makePreview(copiedFile)

	if err == nil {
		newFileName, _ := newFileName(pathName, formFile.Header["Content-Type"][0])
		previewImage, _ = s3.Upload(preview, newFileName)
	}

	newFileName, err := newFileName(pathName, formFile.Header["Content-Type"][0])

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var errUpload error
	if buf.Len() == 0 {
		locationFull, errUpload = s3.Upload(file, newFileName)
	} else {
		r := bytes.NewReader(buf.Bytes())
		locationFull, errUpload = s3.Upload(r, newFileName)
	}

	if errUpload != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errUpload.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": map[string]string{
			"url":        locationFull,
			"previewUrl": previewImage,
			"type":       formFile.Header["Content-Type"][0],
		}})
}

// DeleteManager removes manager
func DeleteManager(c *gin.Context) {}
