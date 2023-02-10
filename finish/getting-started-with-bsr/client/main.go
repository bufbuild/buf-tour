package main

import (
	"context"
	"log"
	"net/http"

	"buf.build/gen/go/<USER>/petapis/bufbuild/connect-go/pet/v1/petv1connect"
	petv1 "buf.build/gen/go/<USER>/petapis/protocolbuffers/go/pet/v1"
	"github.com/bufbuild/connect-go"
)

func main() {
	client := petv1connect.NewPetStoreServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	_, err := client.PutPet(
		context.Background(),
		connect.NewRequest(&petv1.PutPetRequest{
			PetType: petv1.PetType_PET_TYPE_SNAKE,
			Name:    "Ekans",
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully PutPet")
}
