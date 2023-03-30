package cld

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

func NewCloudinary() *cloudinary.Cloudinary {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}

	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_SECRET"))
	// fmt.Println(res)
	if err != nil {
		return nil
	}
	return cld
}

func UploadImage(ctx context.Context, folder, name string, file interface{}) (*uploader.UploadResult, error) {
	cld := NewCloudinary()
	return cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:   folder,
		PublicID: name,
	})
}
