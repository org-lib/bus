package main

import (
	"fmt"
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

func main() {
	lists()
}
