package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"math/rand"
	"protobuf/protogen/basic"
	"protobuf/protogen/communication"
	"time"
)

func BasicUser() {
	comm := randomCommunicationChannel()
	sr := map[string]uint32{"Power": 5, "Speed": 10, "Intelegent": 10}
	u := basic.User{
		Id:       1,
		Username: "Kawekaweha123",
		IsActive: true,
		Password: []byte("passwordnyakawe"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"Kawekaweha123@gmail.com", "Muhammad.ilham.kusumawardhana@gmail.com"},
		Address: &basic.Address{
			City:    "Bekasi",
			Country: "Indonesia",
			Street:  "Karang Satria",
		},
		ComunincationChannel: &comm,
		SkillRate:            sr,
	}
	log.Println(&u)
}

// contoh penggunaan oneof
func BasicUserOneOf() {
	instantMessaging := communication.InstantMessaging{
		Id:     1,
		Note:   "XL",
		Number: "087876897004",
	}

	//jika ingin menggunakan oneof , haurs membuat instance ini dulu
	ecomm := basic.User_InstantMessaging{
		InstantMessaging: &instantMessaging,
	}

	comm := randomCommunicationChannel()

	u := basic.User{
		Id:       1,
		Username: "Kawekaweha123",
		IsActive: true,
		Password: []byte("passwordnyakawe"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"Kawekaweha123@gmail.com", "Muhammad.ilham.kusumawardhana@gmail.com"},
		Address: &basic.Address{
			City:    "Bekasi",
			Country: "Indonesia",
			Street:  "Karang Satria",
		},
		ComunincationChannel: &comm,
		ElectonicCommChannel: &ecomm,
	}
	log.Println(&u)

}

// marshalling mengubah proto menjadi json
func ProtoToJsonUser() {

	u := basic.User{
		Id:       1,
		Username: "Kawekaweha123",
		IsActive: true,
		Password: []byte("passwordnyakawe"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"Kawekaweha123@gmail.com", "Muhammad.ilham.kusumawardhana@gmail.com"},
		Address: &basic.Address{
			City:    "Bekasi",
			Country: "Indonesia",
			Street:  "Karang Satria",
		},
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

// unmarshalling mengubah json jadi proto
func JsonToProtoUser() {
	json := `
		{
			"ID":4,
			"UserName":"Kawekaweha4136",
			"is_active":true,
			"Password":"cGFzc6dvcmRueWFrYXdl",
			"Emails":["Kawekaweha4136@gmail.com", "Muhammad.ilham.kusumawardhana@gmail.com"],
			"Gender":"GENDER_MALE"
		}
	`

	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		log.Fatalln("Got Error : ", err)
	}

	log.Println(&p)
}

func randomCommunicationChannel() anypb.Any {
	rand.Seed(time.Now().UnixNano())

	paperMail := communication.PaperMail{
		Id:      1,
		Address: "Alamat Surat",
		Note:    "Dekat rumah kosong",
	}

	socialMedia := communication.SocialMedia{
		Id:       1,
		Address:  "Kawe123_",
		Platform: communication.SocialMediaPlatform_PLATFORM_INSTAGRAM,
	}

	instantMessaging := communication.InstantMessaging{
		Id:     1,
		Note:   "XL",
		Number: "087876897004",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 2:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	case 3:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

// unmarshal tipe yang diketahui
func BasicUnmarshalAnyKnown() {
	socialMedia := communication.SocialMedia{
		Platform: communication.SocialMediaPlatform_PLATFORM_INSTAGRAM,
		Address:  "kawe123_",
	}

	var a anypb.Any

	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	sm := communication.SocialMedia{}

	if err := a.UnmarshalTo(&sm); err != nil {
		return
	}

	jsonBytes, _ := protojson.Marshal(&sm)
	log.Println(string(jsonBytes))
}

// unmarshal tipe yang tidak diketahui
func BasicUnmarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := a.UnmarshalNew()

	if err != nil {
		return
	}

	log.Println("Unmarshal As", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshaled)
	log.Println(string(jsonBytes))
}

// mengetahui apakah tipe tertentu
func BasicUnmasrshalAnyIs() {
	a := randomCommunicationChannel()
	pm := communication.PaperMail{}

	if a.MessageIs(&pm) {
		if err := a.UnmarshalTo(&pm); err != nil {
			log.Fatalln(err)
		}

		jsonBytes, _ := protojson.Marshal(&pm)
		log.Println(string(jsonBytes))
	} else {
		log.Println("Not PaperMail, But : ", a.TypeUrl)
	}
}
