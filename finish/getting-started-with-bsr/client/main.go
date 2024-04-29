package main

import (
	"context"
	"log"
	"net/http"

	// Replace <USERNAME> with your BSR username if username isn't present
	"buf.build/gen/go/:username/petapis/connectrpc/go/pet/v1/petv1connect"
	petv1 "buf.build/gen/go/:username/petapis/protocolbuffers/go/pet/v1"
	connect "connectrpc.com/connect"
)

func main() {
	client := petv1connect.NewPetStoreServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.PutPet(
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
	log.Println(res.Msg, "Successfully PutPet")
}
