package services

import (
	"io"
	"mime/multipart"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

type StorageService struct {
	storage *drive.Service
}

type Info struct {
	FileName string
	MimeType string
}

func NewStorageService(storage *drive.Service) *StorageService {
	return &StorageService{
		storage: storage,
	}
}

func (ss *StorageService) Create(file multipart.File, info Info) (string, error) {
	f := &drive.File{Name: info.FileName, MimeType: info.MimeType}

	uploadType := googleapi.QueryParameter("uploadType", "resumable")

	f, err := ss.storage.Files.Create(f).Media(file).Do(uploadType)
	return f.Id, err
}

func (ss *StorageService) Stat(file_id string) (*drive.File, error) {
	f, err := ss.storage.Files.Get(file_id).Fields("size", "mimeType", "videoMediaMetadata").Do()
	return f, err
}

func (ss *StorageService) Get(file_id string, headers map[string]string) (io.ReadCloser, error) {
	prepare := ss.storage.Files.Get(file_id)
	header := prepare.Header()

	for k, v := range headers {
		header.Set(k, v)
	}

	response, err := prepare.Download()
	if err != nil {
		return nil, err
	}

	return response.Body, err
}

func (ss *StorageService) Files() ([]*drive.File, error) {
	list, err := ss.storage.Files.List().Do()
	if err != nil {
		return nil, err
	}

	return list.Files, err
}

func (ss *StorageService) Remove(file_id string) error {
	err := ss.storage.Files.Delete(file_id).Do()
	return err
}
