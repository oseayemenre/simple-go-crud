package types

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/oseayemenre/go_crud_scratch/internal/sql/database"
)

type ApiConfig struct{
	DB *database.Queries
}

type Response [T interface{}] struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data T `json:"data"`
}

type Claims struct {
	*jwt.RegisteredClaims
	User string
}