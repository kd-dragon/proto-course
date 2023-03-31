package main

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
)

func toJSON(pb proto.Message) string {

	opt := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := opt.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return string(out)
}

func fromJSON(in string, pb proto.Message) {

	opt := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	if err := opt.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Couldn't unmarshal from JSON", err)
	}
}
