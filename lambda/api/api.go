package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty parameters")
	}

	// Does username already exists ?
	userExist, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("error on checking if user exists %w", err)
	}

	if userExist {
		return fmt.Errorf("a user with that username already exists")
	}

	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("error on registering an user %w", err)
	}

	return nil
}
