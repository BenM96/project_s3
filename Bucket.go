//Contains the Bucket struct and associated functions

package main

import (
    "time"
    "fmt"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)


type Bucket struct {
    Name string
    Creation_date *time.Time
    Region string
    Number_of_files int
    Total_size_of_files int64
    Object_last_modified time.Time
    Storage_types []string
    //cost in in USD per month
    Cost float64
}

func Create_bucket(name *string, creation_date *time.Time, region string, client *s3.Client) *Bucket {
    //creats a bucket from the buckets name and region

    //Make a new item bucket
    NewBucket := Bucket{ Name: *name, Creation_date: creation_date}

    //populate it with vales
    Complete_Bucket(&NewBucket, region, client)

    //using pointers to save on processing time
    return &NewBucket
}

func Print_bucket_headder(byte_display_option string){
    //Prints the headder for the list of buckets
    fmt.Printf("Name:Creation date:Number of files:Total size of files(%v):Object last modified:Storage types:Cost(US$/Month)\n")
}

func Print_bucket(bucket *Bucket, byte_display_option string) {
    //Prints all information on the bucket to the commandline
    fmt.Printf("%v:%v:%v:%v:%v:%v:%v\n", bucket.Name , bucket.Creation_date , bucket.Number_of_files, byte_conversion(bucket.Total_size_of_files, byte_display_option), bucket.Object_last_modified, bucket.Storage_types, bucket.Cost)
}

func Complete_Bucket (bucket *Bucket, region string, client *s3.Client){
    //This will collect all of the data other than the name and creation time.

    bucket.Region = Get_bucket_region(bucket.Name, client, region)

    //get all of the data relating to the files and edit values
    bucket.Number_of_files, bucket.Total_size_of_files, bucket.Object_last_modified, bucket.Storage_types = Get_file_data(bucket.Region, bucket.Name, client)

    //get the cost of the bucket
    bucket.Cost = Get_bucket_cost()
}