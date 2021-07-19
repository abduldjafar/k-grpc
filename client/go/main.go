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
		Name: "abdul",
	}

	requestById := &entitypb.ID{
		Id: "60f4e08897bef70725119399",
	}

	resp, _ := client.AddTalent(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Name)

	resp2, _ := client.GetTalent(context.Background(), requestById)
	fmt.Printf("Receive response => [%v]", resp2.Name)

}
