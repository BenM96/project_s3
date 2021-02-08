package main

import(
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"sync"
	"time"
	"flag"
)

func main() {

	//get the run options from the user via flags. In of type map[string]string
	run_options := Get_run_options()

	//create a wait group
	var wait_group sync.WaitGroup

	//setup s3 client
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(config)
	

	//get list of all buckets in the specified region
	//The Bucket list will contain the name and creation date of every bucket
	bucket_list := Get_all_buckets(run_options["region"], client)

	//print the bucket headder
	Print_bucket_headder(run_options["byte_display_option"])

	//for every bucket:
	for _, bucket := range(bucket_list.Buckets){
		//wait for this bucket to be created and printed before finishing the main function
		wait_group.Add(1)
		//Create the rest of the bucket from the name and display to output
		go Create_and_display_bucket(bucket.Name, bucket.CreationDate, client, &wait_group, run_options)
	}


	//waiting for all buckets to be processed
	wait_group.Wait()

}

func Create_and_display_bucket(name *string, creation_date *time.Time, client *s3.Client, wait_group *sync.WaitGroup, run_options map[string]string){
	//This function takes the name of a bucket, generates the rest of the information about the bucket and then prints that informataion
	//There are a number of filters that will premeturly end this function if it does not pass the users filter requirments

	//when this fuction exits let the main function know to stop it waiting for it
	defer wait_group.Done()
	
	//if the name doesn't contain the name_filter_string end the function here
	if !Name_filter(*name, run_options["name_filter_string"]){return}
	
	//Gather the rest of the bucket's information
	complete_bucket := Create_bucket(name, creation_date, run_options["region"], client)
	
	//Filters
	if !Storage_type_filter(complete_bucket.Storage_types, run_options["storage_type_filter_string"]){return}
	if !Region_filter(complete_bucket.Region, run_options["region_filter_string"]){return}

	//If passed all filters print the bucket
	Print_bucket(complete_bucket, run_options["byte_display_option"])
}

func Get_run_options() map[string]string{
		//This fuction gathers infromation from flags on the command line and returns them as a map
		
		//List of flags. Name, deafult value, information
		region_flag := flag.String("region", "eu-west-2", "The region used to get the initial list of buckets")
		byte_display_option_flag := flag.String("byte-unit", "MB", "How the number of bytes will be displayed. Allowed values are KB, MB or GB")
		name_filter_string_flag := flag.String("name-filter", "", "Will only return buckets with names that contain the 'name-filter' value")
		storage_type_filter_string_flag := flag.String("storage-filter", "", "Will only return buckets that have an object with the 'storage-filter' value. Examples of accepted values are STANDARD, DEEP_ARCHIVE or REDUCED_REDUNDANCY")
		region_filter_string_flag := flag.String("region-filter", "", "Will only return buckets from the given region")
	
		
		//get values from the flags given by the user
		flag.Parse()
		
		//make a map of all of the run options the use has given
		run_options := make(map[string]string)

		run_options["region"] = *region_flag
		run_options["byte_display_option"] = *byte_display_option_flag
		run_options["name_filter_string"] = *name_filter_string_flag
		run_options["storage_type_filter_string"] = *storage_type_filter_string_flag
		run_options["region_filter_string"] = *region_filter_string_flag

		return run_options
}