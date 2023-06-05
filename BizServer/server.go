package main

import (
	"context"
	_ "fmt"
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
	UnimplementedBizServer
}

func (s *server) GetUsers(ctx context.Context, input *GetUserInput) (*GetUserOutput, error) {
// 	fmt.Println(input.Nonce)
	return &GetUserOutput{
		User: User{
			name:      "Kiana",
			family:    "Masoudzadeh"
			id:        2
			age:       25
			sex:       male
			createdAt: 103,
		},
	}, nil
}

func (s *server) GetUsersWithSqlInjection (ctx context.Context, input *GetUserInput) (*GetUserOutput, error) {
// 	fmt.Println(input.MessageId)
	return &GetUserOutput{
		user:       User{
                    	name : "Kiana",
                    	family: "Masoudzadeh"
                    	id : 2
                    	age : 25
                    	sex : male
                    	createdAt : 103
                    }
		message_id:   3,
	}, nil
}