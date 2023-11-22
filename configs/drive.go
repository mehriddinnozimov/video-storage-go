package configs

import (
	"context"
	"log"
	"os"
	"path"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func NewDrive() *drive.Service {
	cd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	credentials := path.Join(cd, "secret", "drive.json")

	d, err := drive.NewService(context.Background(), option.WithCredentialsFile(credentials))
	if err != nil {
		log.Fatal(err)
	}

	return d
}
