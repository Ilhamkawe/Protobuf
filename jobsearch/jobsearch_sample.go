package jobsearch

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"protobuf/protogen/basic"
	"protobuf/protogen/dummy"
	"protobuf/protogen/jobsearch"
)

func BasicJobSearchSoftware() {
	s := jobsearch.JobSoftware{
		JobSoftwareId: 1,
		Application:   &basic.Application{ApplicationId: 1, ApplicationName: "Kalibrrr", ApplicationVersion: "1.0.0"},
	}

	jsonByte, _ := protojson.Marshal(&s)
	log.Println(string(jsonByte))
}

func BasicJobSearchCandidate() {
	c := jobsearch.JobCandidate{
		JobCandidateId: 1,
		Application: &dummy.Application{
			ApplicationId:     1,
			Email:             "Kawekaweha123@gmail.com",
			Phone:             "087876897004",
			ApplicantFullName: "Muhammad Ilham Kusumawardhana",
		},
	}

	jsonByte, _ := protojson.Marshal(&c)
	log.Println(string(jsonByte))
}
