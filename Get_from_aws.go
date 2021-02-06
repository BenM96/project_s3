//This file contains functions relating to getting information from aws via apis

package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"context"
	"log"
	//"fmt"
)

func Get_buckets_in_region (region string, client *s3.Client ) (*s3.ListBucketsOutput){
	//function send api request to aws to get all buckets in a region
	//return a pointer to the bucket_list
	//api request to aws
	bucket_list, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{}, func(o *s3.Options) {
		o.Region = region
	})
	if err != nil {
		log.Fatal(err)
	}

	return bucket_list
}

func Get_number_of_files (region string, bucket_name string, client *s3.Client) int{
	//Function that will count the number of files in 

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{Bucket: &bucket_name}, func(o *s3.Options) {
		o.Region = region
	})
	if err != nil {
		log.Fatal(err)
	}

	number_of_files := len(output.Contents)

	return number_of_files
}