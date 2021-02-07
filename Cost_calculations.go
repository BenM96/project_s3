package main

import (
	//"fmt"
)


func Get_bucket_cost (total_size_of_files int64) float64{
	//retuns the estimated cost of a bucket.

	//convert total size of files to GB as it is what aws uses for pricing
	total_size_of_files_GB := byte_conversion(total_size_of_files, "GB")	
	cost := float64(total_size_of_files_GB) * 0.024

	return cost
}