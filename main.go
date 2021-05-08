package main

import (
	"fmt"
	"io/ioutil"
	"log"

	enum_example "github.com/rprajapati0067/protocol-buffer-go/src/enum_example"
	example_simple "github.com/rprajapati0067/protocol-buffer-go/src/simple"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	// sm := doSimple()

	//	readAndWriteDemo(sm)
	// jsonDemo(sm)
	doEnum()

}

func doEnum() {
	em := enum_example.EnumMessage{
		Id:           42,
		DayOfTheWeek: enum_example.DayOfTheWeek_MONDAY,
	}
	fmt.Println(em)
}
func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)
	sm2 := &example_simple.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct ", sm2)
}
func readAndWriteDemo(sm proto.Message) {
	sm2 := &example_simple.SimpleMessage{}
	writeToFile("simple.bin", sm)
	readFromFile("simple.bin", sm2)

	fmt.Println(sm2)
}

func toJSON(pb proto.Message) string {
	ouput, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON")
	}
	return string(ouput)
}

func fromJSON(in string, pb proto.Message) {
	unmarshal := protojson.UnmarshalOptions{}
	err := unmarshal.Unmarshal([]byte(in), pb)

	if err != nil {
		log.Fatalln("Couldn'tunmarshal the JSON into pb struct", err)
	}

}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}
	fmt.Println("Data has been written")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("something went wrong while reading the file", err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffer struct", err)
		return err2
	}
	return nil
}

func doSimple() *example_simple.SimpleMessage {
	sm := example_simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	return &sm
}
