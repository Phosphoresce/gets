package main

import (
	"flag"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	bucketPtr := flag.String("bucket", "", "bucket: bucket name")
	keyPtr := flag.String("file", "", "file: file path")
	regionPtr := flag.String("region", "us-east-1", "region: default is fine for global")
	flag.Parse()

	//parse 'key' and grab last bit, as it is the filename
	fileAr := strings.Split(*keyPtr, "/")
	fileName := fileAr[len(fileAr)-1]

	//have to make an ec2 call to get instance role creds due to amz bug
	ec2.New(session.New(&aws.Config{Region: aws.String(*regionPtr), DisableSSL: aws.Bool(true)}))
	client := s3.New(session.New(&aws.Config{Region: aws.String(*regionPtr), DisableSSL: aws.Bool(true)}))

	params := &s3.GetObjectInput{
		Bucket: aws.String(*bucketPtr),
		Key:    aws.String(*keyPtr),
	}

	result, err := client.GetObject(params)
	if err != nil {
		println(err.Error())
		return
	}
	file, err := os.Create(fileName)
	if err != nil {
		println(err.Error())
	}
	if _, err := io.Copy(file, result.Body); err != nil {
		println(err.Error())
	}
	result.Body.Close()
	file.Close()
}
