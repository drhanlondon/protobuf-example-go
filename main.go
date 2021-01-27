package main

import (
	"fmt"
	"github.com/drhanlondon/protobuf-example-go/src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := example_simple.SimpleMessage{
		Id: 12345,
		IsSimple: true,
		Name: "Seokhyun",
		SampleList: []int32{1,4,7,8},
	}

	//fmt.Println(sm)
	fmt.Println(sm.GetId())
	fmt.Println(sm.Id)
}