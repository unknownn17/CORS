package redis17

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	R *redis.Client
	C context.Context
}

func (u *Redis) Register(email string, user []byte) error {
	if err := u.R.Set(u.C, fmt.Sprintf("%s+register", email), user,0).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Redis) Verify(code, email string) error {
	if err := u.R.Set(u.C, fmt.Sprintf("%s+verify", email), code, time.Minute*5).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Redis) VerifyUser(code, email string) (string, error) {
	code1, err := u.R.Get(u.C, fmt.Sprintf("%s+verify", email)).Result()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return code1, nil
}

func (u *Redis) LogIn(email string) ([]byte, error) {
	user, err := u.R.Get(u.C, fmt.Sprintf("%s+register", email)).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return []byte(user), nil
}

func (u *Redis) OriginaddAndUpdate(id string, origin []byte) error {
	byted, err := u.OriginGet(id)
	if err != nil {
		if err1 := u.R.Set(u.C, id, origin, 0).Err(); err1 != nil {
			log.Println(err1)
			return err
		}
		if err := u.StoreOrigins(origin); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	if err := u.R.Set(u.C, id, origin, 0).Err(); err != nil {
		log.Println(err)
		return err
	}
	if err := u.RemoveOrigin(byted); err != nil {
		log.Println(err)
		return err
	}
	if err := u.StoreOrigins(origin); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Redis) OriginGet(id string) ([]byte, error) {
	user, err := u.R.Get(u.C, id).Bytes()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return []byte(user), nil
}

func (u *Redis) OriginsGet() ([]string, error) {
	origins, err := u.R.SMembers(u.C, "origins").Result()
	fmt.Println("origins",origins)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return origins, nil
}

func (u *Redis) OriginDelete(id string) error {
	origin, err := u.OriginGet(id)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := u.R.Del(u.C, id).Err(); err != nil {
		log.Println(err)
		return err
	}
	if err := u.RemoveOrigin(origin); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Redis) StoreOrigins(origin []byte) error {
	if err := u.R.SAdd(u.C, "origins", origin).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Redis) RemoveOrigin(origin []byte) error {
	if err := u.R.SRem(u.C, "origins", origin).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
