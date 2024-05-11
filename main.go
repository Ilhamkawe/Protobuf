package main

import (
	"fmt"
	"log"
	"protobuf/basic"
	"time"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05" + " " + string(bytes)))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	//basic.BasicHello()
	//basic.BasicUser()
	//basic.ProtoToJsonUser()
	//basic.JsonToProtoUser()
	//basic.BasicUserGroup()
	//basic.BasicUserOneOf()
	//basic.WriteToFileSample()
	//basic.ReadFromFileSample()
	//basic.WriteJsonToFileSample()
	//basic.ReadProtoFromJson()
	//basic.DoWriteUserContentToFile()
	basic.DoReadUserContentFromFile()
}
