package main

import (
	"context"
	"fmt"
	"k-grpc/entitypb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := entitypb.NewTalentServiceClient(cc)
	request := &entitypb.TalentRequest{
		Name:  "abdul",
		Email: "abdul@gmail.com",
		BodySize: &entitypb.BodySizeDataType{
			Height:    12,
			Bust:      120,
			Waist:     150,
			Hips:      100,
			BodyType:  100,
			HairType:  "straight",
			HairColor: "black",
			EyesColor: "black",
			ShoesSize: 35,
		},
	}

	requestById := &entitypb.ID{
		Id: "60fb7524009132101038ef0c",
	}

	resp, _ := client.AddTalent(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp)

	resp2, _ := client.GetTalent(context.Background(), requestById)
	fmt.Printf("Receive Get Talent response => [%v]", resp2)

}
