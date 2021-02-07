//This file comtains miscellaneous functions that are usefull
package main

import (
	"log"
)

func byte_conversion (number_of_bytes int64, output_unit string) int64 {
	//takes in a number of bytes and returns the value in the desired unit
	//Accepts "MB", "KB", "GB"
	if output_unit == "KB"{
		return number_of_bytes/1000
	}
	
	if output_unit == "MB"{
		return number_of_bytes/1000000
	}

	if output_unit == "GB"{
		return number_of_bytes/1000000000
	}

	log.Fatal("Invalid output unit for byte_conversion. Wants KB, MG, GB. received", output_unit)
	//will never reach this return
	return -1
	

}