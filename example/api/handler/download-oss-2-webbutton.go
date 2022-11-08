package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	aws_s3 "github.com/org-lib/bus/aws/aws-s3"
	"github.com/org-lib/bus/config"
)

var (
	Info *aws_s3.Info
)

func init() {
	var viper = config.Config.V
	Info = &aws_s3.Info{
		Endpoint:  viper.GetString("s3.s3-endpoint"),
		AccessKey: viper.GetString("s3.s3-access-key"),
		SecretKey: viper.GetString("s3.s3-secret-key"),
		Region:    viper.GetString("s3.region"),
		Currency:  viper.GetString("s3.currency"),
		Bucket:    viper.GetString("s3.bucket"),
		TmpDir:    viper.GetString("s3.tmpdir"),
	}
	if Info.Region == "" {
		//ap-southeast-1
		Info.Region = endpoints.ApSoutheast1RegionID
	}
}

func Do2wb(ctx *gin.Context) {

	sses, _ := aws_s3.NewSession(Info)
	svc := s3.New(sses)
	response, _ := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(Info.Bucket),
		Key:    aws.String(fmt.Sprintf("dbhouse/yuandeqiao/999911/db_lulu_test/testa.zip")),
	})
	ctx.DataFromReader(200, *response.ContentLength, *response.ContentType, response.Body, nil)
}
