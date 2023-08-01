package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	aws_s3 "github.com/org-lib/bus/aws/aws-s3"
	"github.com/org-lib/bus/config"
	"net/http"
)

var (
	Info aws_s3.Info
)

func init() {
	var viper = config.Config.V
	Info = aws_s3.Info{
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
	//桶名称
	bkname := ctx.Query("bucket")
	//路径名称
	path := ctx.Query("path")

	// 初始化桶信息
	ifo := Info
	ifo.Bucket = bkname
	err, existsBk := existBucket(ifo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":   err.Error(),
			"statu": -1,
		})
		return
	}
	if !existsBk {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":   fmt.Sprintf("无法在OSS上找到桶名称：%v", bkname),
			"statu": -1,
		})
		return
	}
	// 初始化 OSS 连接
	sses, _ := aws_s3.NewSession(ifo)
	svc := s3.New(sses)
	response, _ := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(ifo.Bucket), //example: bz-dba-backup
		Key:    aws.String(path),       //example : dbhouse/yuandeqiao/999911/db_lulu_test/testa.zip
	})
	// 响应 OSS 对象
	ctx.DataFromReader(200, *response.ContentLength, *response.ContentType, response.Body, nil)
}
func existBucket(info aws_s3.Info) (error, bool) {
	var exi bool
	ses, _ := aws_s3.NewSession(info)
	svc := s3.New(ses)
	result, err := svc.ListBuckets(nil)
	if err != nil {
		return err, false
	}
	for _, b := range result.Buckets {
		if info.Bucket == aws.StringValue(b.Name) {
			exi = true
			break
		}
	}
	return nil, exi
}
