package main

import(
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"sync"
	"time"
	//"fmt"
	//"github.com/aws/aws-sdk-go-v2/aws"
)

func main() {
	//variables
	region := "eu-west-2"
	byte_display_option := "MB"

	//create a wait group
	var wait_group sync.WaitGroup

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

	//for every bucket
	for _, bucket := range(bucket_list.Buckets){
		//wait for the bucket to be created and printed
		wait_group.Add(1)
		//Create the rest of the bucket from the name.
		go Create_and_display_bucket(bucket.Name, bucket.CreationDate, region, client, byte_display_option, &wait_group)
	}

	wait_group.Wait()

}

func Create_and_display_bucket(name *string, creation_date *time.Time, region string, client *s3.Client, byte_display_option string, wait_group *sync.WaitGroup){

	//when this fuction exits main fuction can stop waiting 
	defer wait_group.Done()

	//TODO filter list by name
	//TODO filter by view type
	//TODO get more information on bucket
	Complete_Bucket := Create_bucket(name, creation_date, region, client)
	Print_bucket(Complete_Bucket, byte_display_option)
}