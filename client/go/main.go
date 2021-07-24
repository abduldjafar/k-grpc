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

	resp, _ := client.AddTalent(context.Background(), request)

	talentData, _ := client.GetTalent(context.Background(), resp)

	talentData.BodySize.BodyType = 190
	updatedRequest := &entitypb.TalentUpdaterequest{
		Id: resp,
		RequestsData: &entitypb.TalentRequest{
			Email:            talentData.Email,
			Name:             "roberto",
			Address:          talentData.Address,
			Age:              talentData.Age,
			BirthDate:        talentData.BirthDate,
			Gender:           talentData.Gender,
			Verified:         talentData.Verified,
			BodySize:         talentData.BodySize,
			ProductionType:   talentData.ProductionType,
			Languages:        talentData.Languages,
			Skills:           []string{"walking beauty"},
			Keywords:         talentData.Keywords,
			DataType:         talentData.DataType,
			PhotoProfile:     talentData.PhotoProfile,
			Password:         talentData.Password,
			CodeVerification: talentData.CodeVerification,
		},
	}

	updateData, _ := client.UpdateTalent(context.Background(), updatedRequest)

	deleteTalent, _ := client.DeleteTalent(context.Background(), &entitypb.ID{
		Id: resp.Id,
	})

	allTalents, _ := client.GetListTalents(context.Background(), &entitypb.Pagination{
		Limit: 20,
		Page:  0,
	})
	log.Printf("Receive Get Talent response => [%v]", talentData)
	log.Printf("Receive response => [%v]", resp)
	log.Println(updateData)
	log.Println(deleteTalent)
	log.Println(allTalents)

}
