package ali_oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Info struct {
	Endpoint   string `json:"-"`
	AccessKey  string `json:"-"`
	SecretKey  string `json:"-"`
	Region     string `json:"-"`
	Currency   string `json:"-"`
	Bucket     string `json:"-"`
	TmpDir     string `json:"-"`
	ExpireTime int64  `json:"-"`
}

func NewSession(info Info) (*oss.Client, error) {
	// 创建OSSClient实例。
	client, err := oss.New(info.Endpoint, info.AccessKey, info.SecretKey)
	if err != nil {
		return nil, err
	}
	return client, nil
}
