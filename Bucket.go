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
    Cost int

}

func Create_bucket(name *string, creation_date *time.Time, region string, client *s3.Client) *Bucket {

    NewBucket := Bucket{ Name: *name, Creation_date: creation_date}

    Complete_Bucket(&NewBucket, region, client)

    return &NewBucket
}

func Print_bucket(bucket *Bucket) {
    //TODO imporve this when other information is harvesed
    fmt.Println( "Name: ", bucket.Name ," Creation date: ", bucket.Creation_date ," number of files: ", bucket.Number_of_files, " Total size of files: ", bucket.Total_size_of_files, " object last modified: ", bucket.Object_last_modified, " cost: ", bucket.Cost)
}

func Complete_Bucket (bucket *Bucket, region string, client *s3.Client){
    //This will collect all of the data other than the name and creation time.

    bucket.Region = Get_bucket_region(bucket.Name, client)

    //get all of the data relating to the files and edit values
    bucket.Number_of_files, bucket.Total_size_of_files, bucket.Object_last_modified = Get_file_data(bucket.Region, bucket.Name, client)

    //get the cost of the bucket
    bucket.Cost = Get_bucket_cost(bucket.Total_size_of_files)


}