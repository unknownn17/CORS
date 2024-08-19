package adjust

import (
	email1 "conn/internal/email"
	jwttoken "conn/internal/jwt"
	"conn/internal/models"
	redis17 "conn/internal/redis/methods"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

type Adjust struct {
	R         *redis17.Redis
	Generated map[int]bool
	Rng       *rand.Rand
}

func (u *Adjust) Register(ctx context.Context, req *models.Register) error {
	byted, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := u.R.Register(req.Email, byted); err != nil {
		log.Println(err)
		return err
	}
	code, err := email1.Sent(req.Email)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := u.R.Verify(code, req.Email); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Adjust) Verify(ctx context.Context, req *models.Verify) error {
	val, err := u.R.VerifyUser(req.SecretCode, req.Email)
	if err != nil {
		log.Println(err)
		return err
	}
	if val != req.SecretCode {
		return errors.New("secretcode isn't match please double check it")
	}
	return nil
}

func (u *Adjust) LogIn(ctx context.Context, req *models.LogIn) (string, error) {
	val, err := u.R.LogIn(req.Email)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var user models.Register
	if err := json.Unmarshal(val, &user); err != nil {
		log.Println(err)
		return "", err
	}
	if user.Password != req.Password {
		return "", errors.New("password isn't match")
	}
	token, err := jwttoken.CreateToken(&user)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return token, nil
}

func (u *Adjust) OriginAdd(ctx context.Context, req *models.OriginCreate) (string,error) {
	id := u.GenerateUniqueRandomNumber()
	fmt.Printf("The origin is %v",req.Origin)
	byted, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return "",err
	}
	if err := u.R.OriginaddAndUpdate(id, byted); err != nil {
		log.Println("there is an error",err)
		return "",err
	}
	return id,nil
}

func (u *Adjust) OriginGetbyId(ctx context.Context, req string) (*models.OriginCreate, error) {
	val, err := u.R.OriginGet(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var origin models.OriginCreate
	fmt.Println(origin)
	if err := json.Unmarshal(val, &origin); err != nil {
		log.Println("error occured here",err)
		return nil, err
	}
	return &origin, nil
}

func (u *Adjust) OriginGetAll(ctx context.Context) ([]*models.OriginCreate, error) {
	val, err := u.R.OriginsGet()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var origins []*models.OriginCreate

	for _, v := range val {
		var all models.OriginCreate
		if err := json.Unmarshal([]byte(v), &all); err != nil {
			log.Println(err)
			return nil, err
		}
		origins = append(origins, &all)
	}
	return origins, nil
}

func (u *Adjust) OriginPut(ctx context.Context, req *models.OriginGet) error {
	new:=models.OriginCreate{Origin: req.Origin}
	byted, err := json.Marshal(new)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := u.R.OriginaddAndUpdate(req.Id, byted); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Adjust) OriginDelete(ctx context.Context, req string) error {
	if err := u.R.OriginDelete(req); err != nil {
		log.Println(err)
	}
	return nil
}

func (u *Adjust) GenerateUniqueRandomNumber() string {
	if len(u.Generated) >= (10000 - 1 + 1) {
		log.Printf("no more unique numbers available in the range")
	}

	for {
		num := u.Rng.Intn(10000-1+1) + 1
		if !u.Generated[num] {
			u.Generated[num] = true
			return strconv.Itoa(num)
		}
	}
}