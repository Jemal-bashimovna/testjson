package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//jsonStruct, err := UnmarshalJson("jsonFile/data.json")
	//fmt.Println(jsonStruct, err)
	//UnmarshalList()
	//test()
	//httpTest()
	if len(os.Args) < 2 {
		fmt.Println("provide arguments")
	}

	for i, arg := range os.Args {
		fmt.Printf("arguments %d => %s\n", i, arg)
	}

	city := os.Args[1]
	fmt.Println(city)

	name := flag.String("name", "guest", "Your name")
	flag.Parse()
	fmt.Printf("hello, %s\n", *name)
}
