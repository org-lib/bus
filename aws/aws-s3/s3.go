package aws_s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Info struct {
	Endpoint  string `json:"-"`
	AccessKey string `json:"-"`
	SecretKey string `json:"-"`
	Region    string `json:"-"`
	Currency  string `json:"-"`
	Dir       string `json:"-"`
	TmpDir    string `json:"-"`
}

func NewSession(info *Info) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(info.AccessKey, info.SecretKey, ""),
		Endpoint:         aws.String(info.Endpoint),
		Region:           aws.String(info.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
}
