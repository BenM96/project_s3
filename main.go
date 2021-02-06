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
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(config)
	
	//TODO allow getting multiple regions
	//get list of all buckets in the specified region
	//The Bucket list will contain the name and creation date of every bucket
	bucket_list, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{}, func(o *s3.Options) {
		o.Region = region
	})
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("buck: ", *bucket_list.Buckets[0].Name, *bucket_list.Buckets[1].Name)

	for _, bucket := range(bucket_list.Buckets){
		//TODO filter list by name
		//TODO filter by view type
		//TODO get more information on bucket
		fmt.Println(bucket.Name)
	}

}