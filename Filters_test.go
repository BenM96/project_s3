//This file contains go tests for the functions found in Filters.go
package main

import ( 
	"testing"
)

func Test_Name_filter( t *testing.T){

	//get the list of value that be used in the test
	test_value_table := []struct{
		Name string
		Filter_string string
		Expected bool
	}{
		{"NameOfABucket", "", true},
		{"NameOfABucket", "NameOfABucket", true},
		{"NameOfABucket", "Name", true},
		{"NameOfABucket", "Bucket", true},
		{"NameOfABucket", "OfA", true},

		{"NameOfABucket", "nameofabucket", false},
		{"NameOfABucket", "notASubString", false},
		{"NameOfABucket", "NameOfABuckett", false},
		{"NameOfABucket", "NNameOfABucket", false},
		{"NameOfABucket", "NameBucket", false},
	}

	//for every set of values listed above:
	for _, test_value := range(test_value_table){
		//Generate what the function returns
		result := Name_filter(test_value.Name, test_value.Filter_string)

		//if the result isn't what was expected then the test fails
		if result != test_value.Expected{
			t.Errorf("Name_filter failed for values %v, %v. Expected: %v Got: %v", test_value.Name, test_value.Filter_string, test_value.Expected, result)
		}
	}
}
