package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"

	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/oseayemenre/go_crud_scratch/internal/response"
	"github.com/oseayemenre/go_crud_scratch/internal/types"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origins", "https://*, http://*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Link")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "300")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		godotenv.Load()

		access_key := os.Getenv("ACCESS_KEY")

		if access_key == "" {
			log.Fatalf("Access token not provided")
		}

		token := r.Header.Get("Authorization")

		if token == "" {
			response.WriteToJSON(w, 401, &types.Response[interface{}]{
				Status: "failed",
				Message: "No token in header",
			})	
			return
		}

		value := strings.Split(token, " ")

		if len(value) != 2 && value[0] != "Bearer" {
			response.WriteToJSON(w, 401, &types.Response[interface{}]{
				Status: "failed",
				Message: "Malformed token",
			})	
			return
		}

		claims := &types.Claims{}
	
		decode, err := jwt.ParseWithClaims(value[1], claims, func(token *jwt.Token) (interface{}, error){
			return access_key, nil
		})

		if err != nil {
			log.Fatalf("error: %v", err)
		}

		if !decode.Valid {
			response.WriteToJSON(w, 401, &types.Response[interface{}]{
				Status: "failed",
				Message: "Token could not be validated",
			})	
			return
		}

		var user = "user"

		ctx := context.WithValue(r.Context(), user, claims.User)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}