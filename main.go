package main

import(
	"fmt"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	//"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	//variables
	region := "eu-west-2"

	//setup client
	config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(config)


	//get list of all buckets in the specified region
	bucket_list, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	//TODO filter list by name
	//TODO convert

	fmt.Println("buck: ", *bucket_list.Buckets[0].Name, *bucket_list.Buckets[1].Name)

	bucky := CreateBucket(bucket_list.Buckets[0].Name , bucket_list.Buckets[0].CreationDate)

	fmt.Println(bucky)
}