package storage

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type gcs struct {
	client     *storage.Client
	bucketName string
}

// ref: https://developers.google.com/accounts/docs/application-default-credentials
func NewGCSWithCredential(bucketName string, credentialADCPath string) (*gcs, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialADCPath))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &gcs{
		client:     client,
		bucketName: bucketName,
	}, nil
}

func NewGCS(bucketName string) (*gcs, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &gcs{
		client:     client,
		bucketName: bucketName,
	}, nil
}

func (s *gcs) Upload(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer file.Close()
	// TODO: ハッシュはスクショの生成時間+ユーザー名とかで生成したい。
	b := []byte(filepath)
	hashed := fmt.Sprintf("%x", md5.Sum(b))
	bkt := s.client.Bucket(s.bucketName)
	obj := bkt.Object(hashed)
	ctx := context.Background()
	w := obj.NewWriter(ctx)

	if _, err := io.Copy(w, file); err != nil {
		return "", errors.WithStack(err)
	}
	if err := w.Close(); err != nil {
		return "", errors.WithStack(err)
	}

	return hashed, nil
}

func (s *gcs) GetURL(hashedFileName string) (string, error) {
	const base = "https://storage.cloud.google.com/%s/%s"
	u, err := url.Parse(fmt.Sprintf(base, s.bucketName, hashedFileName))
	if err != nil {
		return "", errors.WithStack(err)
	}
	return u.String(), nil
}

func (s *gcs) GetAttrs() ([][4]string, error) {
	bkt := s.client.Bucket(s.bucketName)
	ctx := context.Background()
	query := &storage.Query{Prefix: ""}

	var ret [][4]string
	// ref: https://pkg.go.dev/cloud.google.com/go/storage#hdr-Listing_objects
	it := bkt.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		url, err := s.GetURL(attrs.Name)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		// name/url/owner/created
		var attr = [4]string{}
		attr[0] = attrs.Name
		attr[1] = url
		attr[2] = attrs.Owner
		attr[3] = attrs.Created.String()
		ret = append(ret, attr)
	}
	return ret, nil
}

func (s *gcs) Close() error {
	return s.client.Close()
}
