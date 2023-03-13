package db

import (
	"github.com/blastertwist/antex-dash/config"
	"github.com/blastertwist/antex-dash/pkg/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

func ConnectMinio(cfg *config.Config, logger logger.Logger) *minio.Client {
	// Initialize Minio Client Object
	minioClient, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.SecretID, cfg.Minio.SecretKey, ""),
		Secure: cfg.Minio.UseSSL,
	})

	if err != nil {
		logger.Panic("[MinioClient]Failed to connect to minio server...", zap.String("[Error]: ", err.Error()))
	}

	return minioClient
}
