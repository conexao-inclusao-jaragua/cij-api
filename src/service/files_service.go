package service

import (
	"context"
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type filesService struct {
	cloudinaryIntegration *cloudinary.Cloudinary
}

func NewFilesService(cloudinaryIntegration *cloudinary.Cloudinary) *filesService {
	return &filesService{
		cloudinaryIntegration: cloudinaryIntegration,
	}
}

func (f *filesService) UploadFile(file io.Reader) (string, error) {
	uploadResult, err := f.cloudinaryIntegration.Upload.Upload(
		context.TODO(),
		file,
		uploader.UploadParams{},
	)
	
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}