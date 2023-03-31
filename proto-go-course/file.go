package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func writeToFile(fileName string, pb proto.Message) {
	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return
	}

	if err = os.WriteFile(fileName, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return
	}

	fmt.Println("Data has been written!")

}

func readFromFile(fileName string, pb proto.Message) error {
	in, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatalln("Can't read file", err)
		return err
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn't unmarshal", err)
		return err
	}
	return nil
}
