package oss

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/os/gfile"
	"github.com/minio/minio-go/v6"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var defaultBucketName = "worker"
var defaultLocation = "us-east-1"

var minioServerPool *sync.Pool

func init() {
	minioServerPool = &sync.Pool{
		New: func() interface{} {
			return &MinioServe{}
		},
	}
}

type MinioBucket struct {
	BucketName string
	Location   string
}

type MinioFileObject struct {
	ObjectName  string
	FilePath    string
	ContentType string
	Bucket      MinioBucket
	Config      MinioConfig
}

type MinioConfig struct {
	EndPoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
	UseSSL          bool   `json:"useSSL"`
}

func (c *MinioConfig) Validate() error {
	if c.EndPoint == "" {
		return errors.New("endpoint is not specified")
	}
	if c.AccessKeyID == "" {
		return errors.New("accessKeyId is not specified")
	}
	if c.SecretAccessKey == "" {
		return errors.New("secretAccessKey is not specified")
	}
	return nil
}

func (c *MinioConfig) GetClient() (*minio.Client, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	client, err := minio.New(c.EndPoint, c.AccessKeyID, c.SecretAccessKey, c.UseSSL)
	if err != nil {
		return nil, err
	}
	return client, nil
}

type MinioServerOption func(*MinioServe)

func WithMinioServerClient(c *minio.Client) MinioServerOption {
	return func(s *MinioServe) {
		s.Client = c
	}
}

type MinioServe struct {
	Client *minio.Client
}

func NewMinioServer(opts ...MinioServerOption) (*MinioServe, error) {

	m := &MinioServe{}
	for _, opt := range opts {
		opt(m)
	}
	return m, nil
}

func (m *MinioServe) Reset() {
	m.Client = nil
}

func (m *MinioServe) createBucket(name string, loc string) (bool, error) {
	// 判断桶是否存在
	exists, err := m.Client.BucketExists(name)
	if err == nil && exists {
		return true, nil
	}
	// 创建新桶
	err = m.Client.MakeBucket(name, loc)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *MinioServe) FileUpload(obj *MinioFileObject) (int64, error) {
	_, err := m.createBucket(obj.Bucket.BucketName, obj.Bucket.Location)
	if err != nil {
		return 0, err
	}
	n, err := m.Client.FPutObject(obj.Bucket.BucketName, obj.ObjectName, obj.FilePath, minio.PutObjectOptions{ContentType: obj.ContentType})
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (m *MinioServe) FileDownloadToList(obj *MinioFileObject) ([]string, error) {
	res := make([]string, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()
	file, err := m.Client.GetObjectWithContext(ctx, obj.Bucket.BucketName, obj.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		return res, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return res, err
			}
		}
		res = append(res, string(data))
	}
	return res, nil
}

func (m *MinioServe) FileDownloadToDic(obj *MinioFileObject) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()
	file, err := m.Client.GetObjectWithContext(ctx, obj.Bucket.BucketName, obj.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		return "", err
	}
	defer file.Close()
	name := filepath.Base(obj.ObjectName)
	if gfile.Exists(name) {
		log.Printf("[minio] %s already exists", name)
		return name, nil
	}

	locatFile, err := os.Create(fmt.Sprintf("%s", name))

	if err != nil {
		return "", err
	}
	defer locatFile.Close()
	_, err = io.Copy(locatFile, file)
	if err != nil {
		return "", err
	}

	return name, nil
}

func acquireMinioServer() *MinioServe {
	return minioServerPool.Get().(*MinioServe)
}

func releaseMinioServer(m *MinioServe) {
	m.Reset()
	minioServerPool.Put(m)

}

func FileUpload(obj *MinioFileObject) (int64, error) {

	m := acquireMinioServer()
	client, err := obj.Config.GetClient()
	if err != nil {
		return 0, err
	}
	m.Client = client
	n, err := m.FileUpload(obj)
	if err != nil {
		return 0, err
	}
	releaseMinioServer(m)
	log.Printf("[minio] file upload successed")
	return n, nil
}

func FileDownloadToList(obj *MinioFileObject) ([]string, error) {

	m := acquireMinioServer()
	client, err := obj.Config.GetClient()
	if err != nil {
		return nil, err
	}
	defer func() {
		releaseMinioServer(m)
	}()
	m.Client = client
	res, err := m.FileDownloadToList(obj)
	if err != nil {
		return res, err
	}
	log.Printf("[minio] file download successed")
	return res, nil
}

func FileDownloadToDic(obj *MinioFileObject) (string, error) {

	m := acquireMinioServer()
	client, err := obj.Config.GetClient()
	if err != nil {
		return "", err
	}
	defer func() {
		releaseMinioServer(m)
	}()

	m.Client = client
	name, err := m.FileDownloadToDic(obj)
	if err != nil {
		return name, err
	}
	log.Printf("[minio] file download successed")
	return name, nil
}


func GetFile(obj *MinioFileObject) ([]byte, error) {

	m := acquireMinioServer()
	client, err := obj.Config.GetClient()
	if err != nil {
		return nil, err
	}
	m.Client = client
	n, err := m.GetFile(obj)
	if err != nil {
		return nil, err
	}
	releaseMinioServer(m)
	log.Printf("[minio] file get successed")
	return n, nil
}

func (m *MinioServe) GetFile(obj *MinioFileObject) ([]byte, error) {
	object, err := m.Client.GetObject(obj.Bucket.BucketName, obj.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()
	reader := bufio.NewReader(object)
	btyes, errr := ioutil.ReadAll(reader)
	if errr  != nil {
		return nil, errr
	}
	return btyes, nil
}