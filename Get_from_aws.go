//This file contains functions relating to getting information from aws via apis

package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"context"
	"log"
	"time"
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

func Get_file_data (region string, bucket_name string, client *s3.Client) (int, int64, time.Time) {
	//Function will get information of the files in a bucket and return: number of files, total size of files, last modified date of most recent file

	//Gets list of files
	ObjectList, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{Bucket: &bucket_name}, func(o *s3.Options) {
		o.Region = region
	})
	if err != nil {
		log.Fatal(err)
	}
	
	FileList := &ObjectList.Contents

	//calculate number of files in a bucket
	number_of_files := len(*FileList)

	//calculate total size of files in bucket and when a file it was last modified
	var total_file_size int64 = 0
	var most_recent_modification_time time.Time
	for _ , file := range(*FileList){
		//add file size to total bucket size
		total_file_size = total_file_size + file.Size

		//If the current file was eddited after the current most recent modified time then it becomes the most recently modified file
		if most_recent_modification_time.Before(*(file.LastModified)){
			most_recent_modification_time=*(file.LastModified)
		}
	}


	return number_of_files, total_file_size, most_recent_modification_time
}