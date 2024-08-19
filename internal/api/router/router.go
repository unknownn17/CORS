package router

import (
	"conn/internal/connections"
	embaded17 "conn/internal/embaded"
	jwttoken "conn/internal/jwt"
	"fmt"
	"log"
	"net/http"
)

func Router() {
	r := http.NewServeMux()
	handler := connections.NewHandler()

	r.HandleFunc("POST /users/register", handler.Register)
	r.HandleFunc("POST /users/verify", handler.Verify)
	r.HandleFunc("POST /users/login", handler.LogIn)

	r.HandleFunc("POST /origins", jwttoken.JWTMiddleware(handler.Create))
	r.HandleFunc("GET /origins/{id}", jwttoken.JWTMiddleware(handler.Get))
	r.HandleFunc("GET /origins", jwttoken.JWTMiddleware(handler.Getall))
	r.HandleFunc("PUT /origins/{id}", jwttoken.JWTMiddleware(handler.Update))
	r.HandleFunc("DELETE /origins/{id}", jwttoken.JWTMiddleware(handler.Delete))
	r.HandleFunc("GET /cors", handler.EnableCORS(handler.CorsMessage))
	go func() {
		fmt.Println("server startedn on port 8081")
		http.HandleFunc("GET /cors", func(w http.ResponseWriter, r *http.Request) {
			t := embaded17.GetTemplates()
			w.WriteHeader(http.StatusOK)
			w.Write(t)
		})
		http.ListenAndServe(":8081",nil)
	}()
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
