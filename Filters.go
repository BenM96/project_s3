//The file contains the various filter than can be applied to the tool
//All of the functions will take a bucket or part of a bucket and return a bool. return value will be true if the bucket should be printed

package main

import (
	"strings"
)

func Name_filter (name string, filter_string string) bool{
	//if the filter string is in the name then return true
	if strings.Contains(name, filter_string){
		return true
	}else{
		return false
	}
}