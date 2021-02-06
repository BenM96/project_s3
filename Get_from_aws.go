//This file contains functions relating to getting information from aws via apis

package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"context"
	"log"
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