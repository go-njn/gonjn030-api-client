package main

import (
	"context"
	"encoding/json"
	api "github.com/go-njn/gonjn030-api-client/pkg"
	"github.com/go-njn/gonjn030-api-client/pkg/domain"
	"github.com/go-njn/gonjn030-api-client/shared"
	"time"
)

func main() {
	client := api.NewUserApiClient(shared.NewConfig()) //use default config with default token
	ctx := context.Background()

	println("GET ALL ")
	users, err := client.GetAll(ctx)
	if err != nil {
		println("ERROR ", err.Error())
	}

	prettyPrint(users)

	user := users[0]
	trace := user

	println("GET BY ID : ", user.Id)
	if trace, err = client.GetById(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}

	prettyPrint(trace)

	println("UPDATE NAME\\EMAIL for ID : ", user.Id, " name : ", user.Name)
	if err := client.Update(ctx, user.Id,
		domain.User{
			Name:  user.Name + time.Now().Format(time.RFC3339),
			Email: "NEW-EMAIL-" + user.Email,
		}); err != nil {
		println("ERROR ", err.Error())
	}

	if trace, err = client.GetById(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}
	prettyPrint(trace)

	println("UPDATE STATUS ID : ", user.Id, " status : ", user.Status, "  email : ", user.Email)
	if err := client.UpdateStatus(ctx, user.Id, getNextStatus(user.Status)); err != nil {
		println("ERROR ", err.Error())
	}

	if trace, err = client.GetById(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}
	prettyPrint(trace)

	println("UPDATE GENDER ID : ", user.Id, " gender : ", user.Gender)
	if err := client.UpdateGender(ctx, user.Id, getNextGender(user.Gender)); err != nil {
		println("ERROR ", err.Error())
	}

	if trace, err = client.GetById(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}
	prettyPrint(trace)

	println("DELETE USER ID : ", user.Id)
	if err := client.Delete(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}

	if trace, err = client.GetById(ctx, user.Id); err != nil {
		println("ERROR ", err.Error())
	}
	prettyPrint(trace)
}

func getNextStatus(status domain.UserStatus) domain.UserStatus {
	if status == domain.ActiveStatus {
		return domain.InactiveStatus
	} else {
		return domain.ActiveStatus
	}
}

func getNextGender(gender domain.UserGender) domain.UserGender {
	if gender == domain.MaleGender {
		return domain.FemaleGender
	} else {
		return domain.MaleGender
	}
}

func prettyPrint(an any) {
	bytes, _ := json.MarshalIndent(an, "", "\t")
	println(string(bytes))
}
