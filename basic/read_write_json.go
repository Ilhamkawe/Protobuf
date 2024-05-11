package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"protobuf/protogen/basic"
)

/*
Proto -> JSON -> Biner,
Merubah Proto Menjadi JSON lalu convert manjadi Biner
*/
func WriteProtoToJson(msg proto.Message, fname string) {
	b, err := protojson.Marshal(msg)
	if err != nil {
		log.Fatalln("Can not marshal message", err)
	}

	log.Println(b)

	if err := ioutil.WriteFile(fname, b, 0644); err != nil {
		log.Fatalln("Can not write to file", err)
	}
}

func WriteJsonToFileSample() {
	u := dummyUser()
	WriteProtoToJson(&u, "superman_file.json")
}

/*
Membaca JSON , lalu melakukan unmarshalling menjadi message proto
*/
func ReadProtoFromJSON(fname string, dest proto.Message) {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can not read file", err)
	}

	if err := protojson.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can not unmarshal", err)
	}
}

func ReadProtoFromJson() {
	dest := basic.User{}

	ReadProtoFromJSON("superman_file.json", &dest)

	log.Println(&dest)
}
