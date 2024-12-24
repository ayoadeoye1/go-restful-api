package helper

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
)

func ProcessUploadedImages(images []*multipart.FileHeader, ctx *gin.Context) ([]string, error) {
	var uploadedImageURLs []string

	for _, image := range images {
		fileName := "img" + time.Now().Format("01-02-03 15-04-.995") + image.Filename
		filePath := fmt.Sprintf("public/uploads/%s", fileName)

		err := ctx.SaveUploadedFile(image, filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to save image %s: %w", fileName, err)
		}

		imageURL := fmt.Sprintf("%s/uploads/%s", ctx.Request.Host, fileName)
		uploadedImageURLs = append(uploadedImageURLs, imageURL)
	}

	return uploadedImageURLs, nil
}
