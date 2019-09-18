package main

import (
	"context"
	"../db"
	"../model"
	"github.com/aws/aws-lambda-go/lambda"
)

type GetRoleRequest struct {
	ID       uint
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, request GetUserRequest) ([]model.Role, error) {
	postgresConnector := db.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	defer db2.Close()
	if err != nil {
		return []model.Role{}, err
	}
	db2.AutoMigrate(&model.Role{})
	account := &model.Role{}
	if request.Name != "" {
		account.Name = request.Name
	}
	var roles []model.Role
	db2.Where(account).Find(&roles)
	return roles, nil
}
func main() {
	lambda.Start(HandleRequest)
}
