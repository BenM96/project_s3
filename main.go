package main

import(
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"sync"
	"time"
	"flag"
	//"fmt"
)

func main() {
	//variables
	region_flag := flag.String("region", "eu-west-2", "The region used to get the initial list of buckets")
	byte_display_option_flag := flag.String("byte-unit", "MB", "How the number of bytes will be displade. Allowed values are KB, MB or GB")
	name_filter_string_flag := flag.String("name-filter", "", "Will only return buckets with names that contain the 'name-filter' value")
	storage_type_filter_string_flag := flag.String("storage-filter", "", "Will only return buckets that have an object with the 'storage-filter' value")
	region_filter_string_flag := flag.String("reigion-filter", "", "Will only return buckets from the given region")

	
	//get values for all flags
	flag.Parse()

	region := *region_flag
	byte_display_option := *byte_display_option_flag
	name_filter_string := *name_filter_string_flag
	storage_type_filter_string := *storage_type_filter_string_flag
	region_filter_string := *region_filter_string_flag


	//create a wait group
	var wait_group sync.WaitGroup

	//setup client
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(config)
	

	//get list of all buckets in the specified region
	//The Bucket list will contain the name and creation date of every bucket
	bucket_list := Get_buckets_in_region(region, client)

	//for every bucket
	for _, bucket := range(bucket_list.Buckets){
		//wait for the bucket to be created and printed
		wait_group.Add(1)
		//Create the rest of the bucket from the name.
		go Create_and_display_bucket(bucket.Name, bucket.CreationDate, region, client, byte_display_option, &wait_group, name_filter_string, storage_type_filter_string, region_filter_string)
	}

	wait_group.Wait()

}

func Create_and_display_bucket(name *string, creation_date *time.Time, region string, client *s3.Client, byte_display_option string, wait_group *sync.WaitGroup, name_filter_string string, storage_type_filter_string string, region_filter_string string){

	//when this fuction exits main fuction can stop waiting 
	defer wait_group.Done()


	//if the name doesn't contain the name_filter_string end the function here
	if !Name_filter(*name, name_filter_string){return}


	//TODO filter list by name
	//TODO filter by view type
	//TODO get more information on bucket
	complete_bucket := Create_bucket(name, creation_date, region, client)

	if !Storage_type_filter(complete_bucket.Storage_types, storage_type_filter_string){return}

	if !Region_filter(complete_bucket.Region, region_filter_string){return}

	

	Print_bucket(complete_bucket, byte_display_option)
}