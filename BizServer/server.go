package main

import (
	"context"
	_ "fmt"
	gen "main/gen/go"
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
	// 	fmt.Println(input.Nonce)
	return &gen.GetUserOutput{
		User: User{
			Name:      "Kiana",
			Family:    "Masoudzadeh",
			Id:        2,
			Age:       25,
			Sex:       "male",
			CreatedAt: 103,
		},
		MessageId: 3,
	}, nil
}

func (s *server) GetUsersWithSqlInjection(ctx context.Context, input *gen.GetUserInput2) (*gen.GetUserOutput, error) {
	// 	fmt.Println(input.MessageId)
	return &gen.GetUserOutput{
		User: User{
			Name:      "Kiana",
			Family:    "Masoudzadeh",
			Id:        2,
			Age:       25,
			Sex:       "male",
			CreatedAt: 103,
		},
		MessageId: 3,
	}, nil
}

type User struct {
	Name      string `json:"name"`
	Family    string `json:"family"`
	Id        int32  `json:"id"`
	Age       int32  `json:"age"`
	Sex       string `json:"sex"`
	CreatedAt int64  `json:"createdAt"`
}
