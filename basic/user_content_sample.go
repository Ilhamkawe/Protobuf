package basic

import (
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"protobuf/protogen/basic"
)

func basicUserContent() basic.UserContent {
	return basic.UserContent{
		UserContentId: 1,
		Slug:          "cara-masak-dimsum",
		Title:         "Cara Masak Dimsum",
		HtmlContent:   "<p>ini cara masak dimsum</p>",
		AuthorId:      3,
		//Category:      "test",
		//Subcategory: "test",
	}
}

func WriteUserContentToFile(msg proto.Message, fname string) {
	uc, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalln("Can not marshal message", err)
	}

	if err := ioutil.WriteFile(fname, uc, 0644); err != nil {
		log.Fatalln("Can not write to file", err)
	}

}

/*
Membaca biner , lalu melakukan unmarshalling menjadi message proto
*/
func ReadUserContentFromFile(fname string, dest proto.Message) {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can not read file", err)
	}

	if err := proto.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can not unmarshal", err)
	}
}

func DoWriteUserContentToFile() {
	uc := basicUserContent()

	WriteUserContentToFile(&uc, "user_content_v3.bin")
}

func DoReadUserContentFromFile() {
	dest := basic.UserContent{}

	ReadProtoFromFile("user_content_v3.bin", &dest)

	log.Println(&dest)
}
