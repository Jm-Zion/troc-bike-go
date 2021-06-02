package media

import (
	"github.com/aws/aws-sdk-go/aws" 
	"github.com/aws/aws-sdk-go/aws/session" 
	"github.com/aws/aws-sdk-go/aws/credentials" 
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/Jm-Zion/troc-bike-go/app"
	"fmt"
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/labstack/echo/v4"
	"net/http"
)

var sess = connectAWS()

func connectAWS() *session.Session{

	sess := session.New(&aws.Config{
		Credentials:      credentials.NewStaticCredentials("id", "secret", "token"),
		Endpoint:         aws.String("http://localhost:4567"),
		Region:           aws.String("us-west-2"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})

	return sess
}

const (
	AWS_S3_REGION = "eu-east1"
	AWS_S3_BUCKET = "Bike"
)

func CreateMedia(c echo.Context) error{

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(file.Filename),
		Body:   src,
	})

	thumbnailDecoded, err := imaging.Decode(src)
	thumbnailBytes := imaging.Resize(thumbnailDecoded, 300, 0, imaging.Box)
    var thumbBuffer bytes.Buffer
	err = imaging.Encode(&thumbBuffer, thumbnailBytes, imaging.JPEG)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Error encoding thumbnail, %v", err),
		})
	}

	fileType := http.DetectContentType(thumbBuffer.Bytes())
	
	thumbnail, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(file.Filename),
        Body:   bytes.NewReader(thumbBuffer.Bytes()),
        ContentType:   aws.String(fileType),
	})

	media := new(Media)
	media.Thumbnail = &thumbnail.Location
	media.Raw = &result.Location

	_, err = app.PGMain().Model(media).Insert()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Error uploading media, %v.", media),
		})
	}

	return c.JSON(http.StatusOK, media)
}