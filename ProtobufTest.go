package main

import (
	"GoTest/example"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func ProtobufTest() {
	fmt.Println("protobuf测试")
	test := &example.Test{
		Label: proto.String("this is protobuf test"),
		Type:  proto.Int32(17),
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	}
	if test.GetLabel() != newTest.GetLabel() {
		fmt.Println("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
	fmt.Println(newTest.GetLabel())
	fmt.Println(newTest.GetType())
	fmt.Println(newTest.GetOptionalgroup())
}
