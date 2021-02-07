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

func Get_bucket_region (name string, client *s3.Client) string {
	location, err := client.GetBucketLocation(context.TODO(), &s3.GetBucketLocationInput{Bucket: &name}, func(o *s3.Options) {
		o.Region = "eu-west-2"
	})
	if err != nil {
		log.Fatal(err)
	}

	region := location.LocationConstraint

	return string(region)
}

func Get_file_data (region string, bucket_name string, client *s3.Client) (int, int64, time.Time, []string) {
	//Function will get information of the files in a bucket and return: number of files, total size of files, last modified date of most recent file

	//Gets list of objects in bucket
	ObjectList, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{Bucket: &bucket_name}, func(o *s3.Options) {
		o.Region = region
	})
	if err != nil {
		log.Fatal(err)
	}
	
	//Get files from the list of objects
	FileList := &ObjectList.Contents

	//calculate number of files in a bucket
	number_of_files := len(*FileList)

	//declare the variables that will be built up as the script runs through every file in the bucket
	var total_file_size int64 = 0
	var most_recent_modification_time time.Time
	var storage_types []string

	//For every file in the bucket
	for _ , file := range(*FileList){
		//add file size to total bucket size
		total_file_size = total_file_size + file.Size



		//If the current file was eddited after the current most recent modified time then it becomes the most recently modified file
		if most_recent_modification_time.Before(*(file.LastModified)){
			most_recent_modification_time=*(file.LastModified)
		}

		

		//checks if storage type is in the current list of storage_type, if not add it
		current_file_storage_type := string(file.StorageClass)
		type_already_exists := false
		for _, existing_type := range(storage_types){
			if current_file_storage_type == existing_type{
				//if the current storage type has already been recorded, note it and break for loop
				type_already_exists = true
				break
			}
		}
		//if storage type doesn't already exist add it to the list of storage types
		if !type_already_exists{
			storage_types = append(storage_types, current_file_storage_type)
		}
	}


	return number_of_files, total_file_size, most_recent_modification_time, storage_types
}