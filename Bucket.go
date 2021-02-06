package main

import (
    "time"
)

type Bucket struct {
    Name *string
    Creation_date *time.Time
    Number_of_files *int
    Total_size_of_files *int
    Last_modified_date_of_most_recent_files *time.Time
    Cost *int

}

func CreateBucket(name *string, creation_date *time.Time) Bucket {

    NewBucket := Bucket{ Name: name, Creation_date: creation_date}

    return NewBucket
}