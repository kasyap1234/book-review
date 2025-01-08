package user

import (
	"context"
	"fmt"
	"net/http"
	"strings"

)
type UserKey string 

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("authorization")
        if tokenString == "" {
            fmt.Println("empty token")
            http.Error(w, "missing authorization header", http.StatusUnauthorized)
        }
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
       claims, err := VerifyToken(tokenString)
	   if err !=nil {
		fmt.Println("error authenticating middleware")
		http.Error(w,"invalid token",http.StatusUnauthorized)
	   }
        ctx := r.Context()
        ctx = context.WithValue(ctx, UserKey("userID"), claims.userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}