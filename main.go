package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/nafisfaysal/ProtoBufReadWrite/proto"
	"io/ioutil"
	"log"
)

func main() {
	var message = PbInitValue()
	err := WriteToFile("test.txt", message)
	CatchErr(err)
	err = ReadFromFile("test.txt", message)
	CatchErr(err)
	pbToJson := PbToJSON(message)
	fmt.Println("Proto Document to JSON: \t", pbToJson)
	FromJSONToPb(pbToJson, message)
	fmt.Println("JSON to Proto Document: \t ", message)
}

func WriteToFile(fname string, pb proto.Message) error {
	b, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Error Serialise to Bytes", err)
		return err
	}

	err = ioutil.WriteFile(fname, b, 0644)
	if err != nil {
		log.Fatalln("Can't Write to File", err)
		return err
	}

	return nil
}

func ReadFromFile(fname string, pb proto.Message) error {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can't Read From the File", err)
		return err
	}

	err = proto.Unmarshal(b, pb)
	if err != nil {
		log.Fatalln("Can't Unmarshal the Bytes", err)
		return err
	}

	return nil
}

func PbToJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	str, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Error to Converting JSON")
		return ""
	}

	return str
}

func FromJSONToPb(s string, pb proto.Message) {
	err := jsonpb.UnmarshalString(s, pb)
	if err != nil {
		log.Fatalln("Error JSON Can't Format to Proto Document")
	}
}

func PbInitValue() *messagepb.SimpleMessage {
	mp := messagepb.SimpleMessage{
		Id:   1234,
		Name: "Nafis Faysal",
		List: []int32{100, 200, 300, 900},
	}
	fmt.Println(mp)

	return &mp
}

func CatchErr(err error) {
	if err != nil {
		log.Fatalln("Catch Error: ", err)
	}
}
