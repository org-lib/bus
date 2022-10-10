package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	ali_oss "github.com/org-lib/bus/aliyun/ali-oss"
	"github.com/org-lib/bus/config"
	"os"
)

var (
	Info *ali_oss.Info
)

func init() {
	var viper = config.Config.V
	Info = &ali_oss.Info{
		Endpoint:   viper.GetString("oss.endpoint"),
		AccessKey:  viper.GetString("oss.access-key"),
		SecretKey:  viper.GetString("oss.secret-key"),
		Region:     viper.GetString("oss.region"),
		Currency:   viper.GetString("oss.currency"),
		Bucket:     viper.GetString("oss.bucket"),
		TmpDir:     viper.GetString("oss.tmpdir"),
		ExpireTime: int64(30),
	}
	if Info.Region == "" {
		//ap-southeast-1
		Info.Region = endpoints.ApSoutheast1RegionID
	}
}
func lists() {

	client, err := ali_oss.NewSession(Info)
	if err != nil {
		fmt.Println(err.Error())
	}
	bucket, err := client.Bucket(Info.Bucket)
	if err != nil {
		// HandleError(err)
	}
	lsRes, err := bucket.ListObjects()
	if err != nil {
		// HandleError(err)
	}

	for _, object := range lsRes.Objects {
		fmt.Println("Objects:", object.Key)
	}

}
func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func Upload() {
	client, err := ali_oss.NewSession(Info)
	if err != nil {
		fmt.Println(err.Error())
	}
	bucket, err := client.Bucket(Info.Bucket)
	if err != nil {
		// HandleError(err)
	}
	err = bucket.PutObjectFromFile("my-object/11.jpg", "/Users/yuandeqiao/Downloads/da.sql")
	if err != nil {
		// HandleError(err)
		fmt.Println(err.Error())
	}

}
func DownLoadUrl() {
	/**
	  这里以png图片为例，故此设置为 image/png
	*/
	//options := []oss.Option{
	//	oss.ContentType("image/png"),
	//}

	client, err := ali_oss.NewSession(Info)
	if err != nil {
		fmt.Println(err.Error())
	}
	bucket, err := client.Bucket(Info.Bucket)
	if err != nil {
		fmt.Println(err.Error())
	}
	signedURL, err := bucket.SignURL("dbhouse/yuandeqiao/699/db_lulu_test/testa.zip", oss.HTTPGet, 6000)
	fmt.Println(signedURL)
}
func main() {
	lists()
	DownLoadUrl()
}
