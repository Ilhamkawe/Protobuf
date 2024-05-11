package basic

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"protobuf/protogen/basic"
)

func BasicUserGroup() {

	users := []*basic.User{
		&basic.User{
			Id:       12,
			Username: "Kawekaweha123",
			IsActive: true,
			Password: []byte("passwordnyakawe"),
			Gender:   basic.Gender_GENDER_MALE,
			Emails:   []string{"Kawekaweha123@gmail.com", "Muhammad.ilham.kusumawardhana@gmail.com"},
			Address:  &basic.Address{Street: "Karang Satria", Country: "Indonesia", City: "Bekasi", Coordinate: &basic.Address_Coordinate{Long: 19.12321, Lat: 40.2888}},
		},
		&basic.User{
			Id:       11,
			Username: "Kawekaweha321",
			IsActive: true,
			Password: []byte("passwordnyakawe"),
			Gender:   basic.Gender_GENDER_MALE,
			Emails:   []string{"Kawekaweha321@gmail.com", "Muhammad.ilham.kusumawardhana@gmail.com"},
			Address:  &basic.Address{Street: "Karang Satria", Country: "Indonesia", City: "Bekasi", Coordinate: &basic.Address_Coordinate{Long: 19.12321, Lat: 40.2888}},
		},
	}

	uG := basic.UserGroup{
		GroupId:   11,
		GroupName: "Avengers",
		Roles: []string{
			"Captain",
			"Scientist",
			"Warrior",
		},
		Users:       users,
		Description: "User Group Betawi Gang",
	}

	jsonBytes, _ := protojson.Marshal(&uG)
	log.Println(string(jsonBytes))
}
