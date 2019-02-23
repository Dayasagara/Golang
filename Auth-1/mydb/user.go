package mydb

import (
	"log"
	"database/sql"
	"fmt"
)


type User struct {
	email     string 
	password  string 
	username string 
	confirm  string 
}

func Signup(username, email, password, confirm string) {
	_,err := db.Exec(`
		INSERT INTO public.user ("USERNAME", "EMAIL", "PASSWORD", "C0NFIRM")
		VALUES ($1,$2,$3,$4)`,username, email, password, confirm)
	
	if err != nil {
		log.Printf("Insertion Error : %v",err)
	}else{
		log.Printf("Registered successfully")
	}
}

func Login(email, password string) (*User, error) {
	result := &User{}
	row := db.QueryRow(`
		SELECT "USERNAME", "EMAIL", "PASSWORD", "C0NFIRM"
		FROM public."user"
		WHERE "EMAIL" = $1 
		  AND "PASSWORD" = $2`, email, password)
	err := row.Scan(&result.username, &result.email, &result.password,&result.confirm)
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
