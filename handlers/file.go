package handlers

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"messenger/s3"
	"net/http"
	"strings"
	"time"

	uuid "github.com/google/uuid"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

// UploadFiles uploads files for message
func UploadFiles(c *gin.Context) {
	locationFull := ""
	previewImage := ""
	pathName := "files_" + c.Request.Header["Application-Id"][0]

	formFile, _ := c.FormFile("file")

	file, _ := formFile.Open()
	defer file.Close()

	var copiedFile io.Reader
	var buf bytes.Buffer

	if isImage(formFile.Header["Content-Type"][0]) {
		copiedFile = io.TeeReader(file, &buf)
		preview, err := makePreview(copiedFile)

		if err == nil {
			newFileName, _ := newFileName(pathName, formFile.Header["Content-Type"][0])
			previewImage, _ = s3.Upload(preview, newFileName)
		}
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

func newFileName(path, mimeType string) (string, error) {
	result := path
	mimeTypes := map[string]string{
		"image/jpeg":      "jpeg",
		"image/jpg":       "jpeg",
		"image/JPG":       "jpeg",
		"image/png":       "png",
		"application/pdf": "pdf",
	}

	uuid := uuid.New().String()

	time := time.Now()
	result = result + "/" + string(uuid) + "_" + strings.Replace(time.String(), " ", "_", -1)

	if ext, ok := mimeTypes[mimeType]; ok {
		return result + "." + ext, nil
	}

	return "", errors.New("unknown mimeType: " + mimeType)
}

func isImage(mimeType string) bool {
	mimeTypes := map[string]string{
		"image/jpeg": "jpeg",
		"image/jpg":  "jpeg",
		"image/JPG":  "jpeg",
		"image/png":  "png",
	}

	_, ok := mimeTypes[mimeType]

	return ok
}

func makePreview(file io.Reader) (io.Reader, error) {
	image, imageType, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	creppedImage := imaging.Resize(image, 100, 0, imaging.Lanczos)

	buf := new(bytes.Buffer)

	if imageType == "jpeg" {
		err = jpeg.Encode(buf, creppedImage, nil)
	}

	if imageType == "png" {
		err = png.Encode(buf, creppedImage)
	}

	if err != nil {
		return nil, err
	}

	return buf, nil
}
