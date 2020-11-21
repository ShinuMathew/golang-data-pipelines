package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shinz9474/InsightAps/DataProcessor/Processors"
)

type Arguments struct {
	cmdArgs []string
}

var args Arguments

func main() {

	for i := 0; i < len(os.Args); i++ {
		args.cmdArgs = append(args.cmdArgs, os.Args[i])
	}

	Init_Processor(args.cmdArgs[1])
}

func Init_Processor(processor_name string) {

	log.Printf("===========================================\n")
	log.Printf("Initializing DataProcessor...\n")
	log.Printf("===========================================\n")
	log.Printf("Processor name: '%s'", processor_name)

	switch processor_name {

	case "TestCase_CSV_To_MYSQL_Writter":

		fmt.Printf("\n\nExecuting ''TestCase_CSV_To_MYSQL_Writter'' processor for TC: '%s'\n\n", args.cmdArgs[2])

		TestCase_CSV_To_MYSQL_Writter.Start_Processor(args.cmdArgs[2])
		break
	}

}

// fmt.Printf("\n\n Running build : %s\n\n", cmdArgs[0])
// fmt.Printf("\n\n Executing processor: %s\n\n", cmdArgs[1])
// fmt.Printf("\n\n Test Case to be processed: %s\n\n", cmdArgs[2])
// "github.com/Shinz9474/InsightAps/DataProcessor/Processors"

// test_string := "ROU_QA_Samsung_Smoke_CODTest"

// test_strings := strings.Split(test_string, "_")

// fmt.Printf("Test_strings: %v", test_strings)

// fmt.Printf("Total no. of args: %v\n\n", len(os.Args))
