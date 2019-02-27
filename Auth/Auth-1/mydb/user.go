package mydb

import (
	"log"
	"database/sql"
	"fmt"
	"crypto/sha512"
	"encoding/base64"
	
)


type User struct {
	email     string 
	password  string 
	username string  
}

func Signup(username, email, password string) {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	_,err := db.Exec(`
		INSERT INTO public.user ("USERNAME", "EMAIL", "PASSWORD")
		VALUES ($1,$2,$3)`,username, email, pwd)
	
	if err != nil {
		log.Printf("Insertion Error : %v",err)
	}else{
		log.Printf("Registered successfully")
	}
}

func Login(email, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	row := db.QueryRow(`
		SELECT "USERNAME", "EMAIL", "PASSWORD"
		FROM public."user"
		WHERE "EMAIL" = $1 
		  AND "PASSWORD" = $2`, email, pwd)
	err := row.Scan(&result.username, &result.email, &result.password)
	if err != nil {
		log.Printf("Error:%v",err)
	}
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}
	return result, nil
}
