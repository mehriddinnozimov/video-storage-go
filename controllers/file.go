package controllers

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"video-storage/configs"
	"video-storage/services"
	"video-storage/types"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/drive/v3"
)

type FileController struct {
	storageService *services.StorageService
	videoService   *services.VideoService
}

func NewFileController(db mongo.Database, storage *drive.Service) *FileController {
	return &FileController{
		storageService: services.NewStorageService(storage),
		videoService:   services.NewVideoService(db),
	}
}

func (fc *FileController) Create(c *gin.Context) {
	fileHeader, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	file, err := fileHeader.Open()

	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	contentTypeArray := fileHeader.Header["Content-Type"]

	if contentTypeArray == nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: "File not allowed"})
		return
	}

	contentType := contentTypeArray[0]
	var ALL_ALLOW_MIME_TYPES = []string{}

	ALL_ALLOW_MIME_TYPES = append(ALL_ALLOW_MIME_TYPES, configs.ALLOW_VIDEO_MIME_TYPES...)
	ALL_ALLOW_MIME_TYPES = append(ALL_ALLOW_MIME_TYPES, configs.ALLOW_IMAGE_MIME_TYPES...)

	if !slices.Contains(ALL_ALLOW_MIME_TYPES, contentType) {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: "File not allowed"})
		return
	}

	fileId, err := fc.storageService.Create(file, services.Info{
		FileName: fileHeader.Filename,
		MimeType: contentType,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.File{FileId: fileId})
}

func (fc *FileController) GetOne(c *gin.Context) {
	var filter types.File
	var headers = map[string]string{}

	err := c.ShouldBindUri(&filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	stat, err := fc.storageService.Stat(filter.FileId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	if slices.Contains(configs.ALLOW_VIDEO_MIME_TYPES, stat.MimeType) {
		var rangeHeader = c.GetHeader("range")

		var start int64 = 0
		var end = stat.Size

		if len(rangeHeader) > 0 {
			var parts = strings.Split(strings.Replace(rangeHeader, "bytes=", "", 1), "-")
			startConv, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
				return
			}

			start = startConv

			if len(parts[1]) > 0 {
				endConv, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
					return
				}

				end = endConv
			} else {
				end = stat.Size - 1
			}

			headers["Range"] = rangeHeader
		}

		var chunkSize = end - start + 1
		c.Status(http.StatusPartialContent)
		c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, stat.Size))
		c.Header("Accept-Ranges", "bytes")
		c.Header("Content-Length", fmt.Sprint(chunkSize))

	} else {
		c.Header("Content-Length", fmt.Sprint(stat.Size))
	}

	c.Header("Content-Type", stat.MimeType)

	file, err := fc.storageService.Get(filter.FileId, headers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.CustomResponse{Ok: false, Message: err.Error()})
		return
	}

	p := make([]byte, 1024)

	for {
		n, err := file.Read(p)
		if err != nil {
			break
		}
		c.Writer.Write(p[:n])
	}
}
