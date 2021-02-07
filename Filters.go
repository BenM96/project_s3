//The file contains the various filter than can be applied to the tool
//All of the functions will take a bucket or part of a bucket and return a bool. return value will be true if the bucket should be printed
//If a filter string is ever empty functions will always return true

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



func Storage_type_filter (storage_types []string, filter_string string) bool {
	//if the filter string is in the list of storage types then return true

	//if filter string is empty always return true
	if filter_string==""{
		return true
	}

	//for every stype in stroage types
	for _, stype := range(storage_types){
		if stype == filter_string {
			return true
		}
	}

	return false
}

func Region_filter (reigon string, filter_string string) bool {
	//if filter string is the same as the reigion return true

	//if filter string empty return true
	if filter_string == "" {
		return true
	}

	if filter_string==reigon{
		return true
	}else{
		return false
	}
}