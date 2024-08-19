package handler

import (
	"conn/internal/models"
	"conn/internal/redis/services"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	S *services.Services
	C context.Context
}

func (u *Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.Register

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Decoding error %v", err), http.StatusNotAcceptable)
	}
	if err := u.S.Register(u.C, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": "Confirmation has been sent to your email"})
}

func (u *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.Verify
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Decoding error %v", err), http.StatusNotAcceptable)
	}

	if err := u.S.Verify(u.C, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": "You have successfully confirmed your account and you can LOGIN! "})
}

func (u *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.LogIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Decoding error %v", err), http.StatusNotAcceptable)
	}
	res, err := u.S.LogIn(u.C, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]string{"Your Token": res})
}

func (u *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.OriginCreate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Decoding error %v", err), http.StatusNotAcceptable)
	}
	res, err := u.S.OriginAdd(u.C, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]string{"Your Origin has been added with this id": res})
}

func (u *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	res, err := u.S.OriginGetbyId(u.C, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(res)
}

func (u *Handler) Getall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := u.S.OriginGetAll(u.C)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(res)
}

func (u *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.OriginGet
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Decoding error %v", err), http.StatusNotAcceptable)
	}
	req.Id = r.PathValue("id")
	fmt.Printf("request id %v",req.Id)
	if err := u.S.OriginPut(u.C, &req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]string{"Origin has Updated with this id":req.Id})
}

func (u *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	if err := u.S.OriginDelete(u.C, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]string{"Message": "Origin has beem deleted with this id"})
}

func (u *Handler) CorsMessage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(map[string]string{"Message":"Congrats buddy you are now one of the trusted person"})
}

func (u *Handler) EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		res, err := u.S.OriginGetAll(u.C)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		origin := r.Header.Get("Origin")
		if origin != "" {
			for _, v := range res {
				if origin == v.Origin {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
		}
		next.ServeHTTP(w, r)
	}
}
