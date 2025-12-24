package storage

import (
	"context"
	"log/slog"
	"maneko/media-proxy/internal/config"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage struct {
	Client *s3.Client
	Bucket string
}

func New(cfg *config.Storage, log *slog.Logger) (*Storage, error) {
	client := s3.New(s3.Options{
		BaseEndpoint: aws.String(cfg.Endpoint),
		Region:       cfg.Region,
		Credentials: aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKey,
				cfg.SecretKey,
				"",
			),
		),
		UsePathStyle: true,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.HeadBucket(ctx, &s3.HeadBucketInput{Bucket: &cfg.Bucket})
	if err != nil {
		return nil, err
	}

	log.Info("storage connected",
		slog.String("bucket", cfg.Bucket),
		slog.String("endpoint", cfg.Endpoint),
		slog.String("region", cfg.Region),
	)

	return &Storage{
		Client: client,
		Bucket: cfg.Bucket,
	}, nil
}
