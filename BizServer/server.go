package main

import (
	"context"
	"fmt"
	_ "fmt"
	gen "main/gen/go"
	"main/service"
)

//type User struct {
// name string
// family string
// id int
// age int
// sex string
// createdAt int
//}

type server struct {
	gen.UnimplementedBizServer
}

func (s *server) GetUsers(ctx context.Context, input *gen.GetUserInput1) (*gen.GetUserOutput, error) {
	//fmt.Println("Received getUsers request", input)
	//pg := service.GetUsers(input.UserId, int(input.MessageId))
	//return &ReqPqResponse{
	//	Nonce:       pg.Nonce,
	//	ServerNonce: pg.ServerNonce,
	//	MessageId:   int32(pg.MessageId),
	//	P:           int32(pg.P),
	//	G:           int32(pg.G),
	//}, nil	return &gen.GetUserOutput{
	//	User: User{
	//		Name:      "Kiana",
	//		Family:    "Masoudzadeh",
	//		Id:        2,
	//		Age:       25,
	//		Sex:       "male",
	//		CreatedAt: 103,
	//	},
	//	MessageId: 3,
	//}, nil
	fmt.Println("Received GetUsers request", input)
	res := service.GetUsers(fmt.Sprint(input.UserId), int(input.MessageId))
	return &gen.GetUserOutput{
		User:      res,
		MessageId: 3,
	}, nil
}

func (s *server) GetUsersWithSqlInjection(ctx context.Context, input *gen.GetUserInput2) (*gen.GetUserOutput, error) {
	fmt.Println("Received GetUsersWithSqlInjection request", input)
	res := service.GetUsers(input.UserId, int(input.MessageId))
	return &gen.GetUserOutput{
		User:      res,
		MessageId: 3,
	}, nil
}
