package main

import(
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	//"fmt"
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
	bucket_list := Get_buckets_in_region(region, client)


	for _, bucket := range(bucket_list.Buckets){
		//TODO filter list by name
		//TODO filter by view type
		//TODO get more information on bucket
		Print_bucket(Create_bucket(bucket.Name, bucket.CreationDate, region, client))
	}

	

}