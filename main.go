package main

import (
	"io/ioutil"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"

	"github.com/drhanlondon/protobuf-example-go/src/simple"
	"github.com/drhanlondon/protobuf-example-go/src/enum_simple"
	
	
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)	
	jsonDemo(sm)
	
}

func jsonDemo(sm proto.Message) {

	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct: ", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}

	//why passing a pointer ??
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}


func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("can't write to file", err)
		return err
	}

	fmt.Println("Data has been written")

	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("something went wrong when reading the file", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("couldn't put the bytes into the protocol buffer struct", err)
		return err
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id: 12345,
		IsSimple: true,
		Name: "Seokhyun you did now",
		SampleList: []int32{1, 3, 5},
	}

	fmt.Println(sm.GetId())

	return &sm
}