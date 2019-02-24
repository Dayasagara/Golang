package mydb

import (
	"log"
)

type Event struct {
	id     string   
	subject  string	
	StartDateTime string		
	EndDateTime  string	
}

func AddEvent(id,name,date,time string) (int){
	_,err := db.Exec(`
		INSERT INTO public."events" ("ID","SUBJECT","STARTDATETIME","ENDDATETIME")
		VALUES ($1,$2,$3,$4)`,id,name,date,time)
	
	if err != nil {
		log.Printf("Insertion Error : %v",err)
		return 1
	}else{
		log.Printf("Added successfully")
		return 0
	}
	
}

