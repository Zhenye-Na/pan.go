package ceph

import (
    "gopkg.in/amz.v1/aws"
    "gopkg.in/amz.v1/s3"

    cfg "pan.go/config"
)

var cephConn *s3.S3

// GetCephConnection : 获取 Ceph 连接
func GetCephConnection() *s3.S3 {
    if cephConn != nil {
        return cephConn
    }

    // 1. 初始化 ceph 的一些信息
    auth := aws.Auth{
        AccessKey: cfg.CephAccessKey,
        SecretKey: cfg.CephSecretKey,
    }

    curRegion := aws.Region{
        Name:                 "default",
        EC2Endpoint:          cfg.CephGWEndpoint,
        S3Endpoint:           cfg.CephGWEndpoint,
        S3BucketEndpoint:     "",
        S3LocationConstraint: false,
        S3LowercaseBucket:    false,
        Sign:                 aws.SignV2,
    }

    // 2. 创建 S3 类型的连接
    return s3.New(auth, curRegion)
}

// GetCephBucket : 获取指定的 bucket 对象
func GetCephBucket(bucket string) *s3.Bucket {
    conn := GetCephConnection()
    return conn.Bucket(bucket)
}

// PutObject : 上传文件到 ceph 集群
func PutObject(bucket string, path string, data []byte) error {
    return GetCephBucket(bucket).Put(path, data, "octet-stream", s3.PublicRead)
}
