# project_s3

## Installation:

To install the tool download the binary s3.exe. This can be done by clicking this link [here](https://raw.githubusercontent.com/BenM96/project_s3/master/s3.exe)

## Prerequisites 

The script needs access to aws credentials in order to run. The most common way to do this is by using the ~/.aws/credentials (C:\Users\<username>\.aws\credentials for windows) folder. For more information on this click [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)

## Running the tool

To run the tool you can simply run s3.exe from windows command prompt or ./s3.exe on linux

To see the options avalable run s3.exe --help

# Understanding the repo

## Main.go

Contains the main function and is where the program starts.

## Bucket.go

Contains the struct for the "Bucket" object and functions relating to the creation usage of Bucket objects.

## Get_from_aws.go

Contains functions that send api requests to aws, extracts the data wanted and returns the value

## Filters.go

Contains a series of filters that can be applied by the user to reduce the output and potentaly speed up the program

## Filters_test.go

Contains go unit tests for functions in Filters.gg

## Cost_calculations.go

Contains functions relating to calculating to the cost of an s3 bucket

## Misc_functions

Contains untility functions that dont belong anywhere else

## s3.exe

The compiled binary of the go files


# Please note

This is not a finnished product, there are still many improvments that can be made. However there is enough content here to show my coding style and prove my compitance (hopefully).
