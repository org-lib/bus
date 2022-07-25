package awsoss

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/s3"
	aws_s3 "github.com/org-lib/bus/aws/aws-s3"
	"github.com/org-lib/bus/config"
	"os"
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
func lists() {
	sses, err := aws_s3.NewSession(Info)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}
	svc := s3.New(sses)
	result, err := svc.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	for _, b := range result.Buckets {
		fmt.Printf("%s\n", aws.StringValue(b.Name))
	}

}
func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	lists()
}
