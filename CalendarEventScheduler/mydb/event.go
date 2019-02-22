package mydb

import (
	"log"
	/*"database/sql"
	"fmt"
	"encoding/json"*/
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
/*
func ModifyEventById(id,name,date,time,duration string) {
	_,err := db.Exec(`
		UPDATE public."events" SET "ID"=$1, "NAME"=$2, "DATE"=$3, "TIME"=$4, "DURATION"=$5
		WHERE "ID"=$6`,id,name,date,time,duration,id)
	
	if err != nil {
		log.Printf("Updation Error : %v",err)
	}else{
		log.Printf("Updated successfully")
	}
}

func DeleteEventById(id string) {
	_,err := db.Exec(`
		DELETE from public."events"
		WHERE "ID"=$1`,id)
	
	if err != nil {
		log.Printf("Deletion Error : %v",err)
	}else{
		log.Printf("Deleted successfully")
	}
}

func getJSON(sqlString string) (string, error) {
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil 
}

func GetAllEvents() (string, error){
	
	query :=`SELECT "ID","NAME","DATE","TIME","DURATION" FROM public."events"`
	eventslist,err := getJSON(query)
	if err!=nil{
		log.Printf("Error")
	}else{
		log.Printf("%v",eventslist)
	}
	return eventslist,err
}

func GetEventByDate(date string) (*Event, error) {
	result := &Event{}
	row := db.QueryRow(`
		SELECT "ID","NAME","DATE","TIME","DURATION"
		FROM public."events"
		WHERE "DATE" = $1`, date)
	err := row.Scan(&result.id, &result.name, &result.date,&result.time,&result.duration)
	if err!=nil{
		log.Printf("Error:%v",err)
	}
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("No event found")
	case err != nil:
		
		return nil, err
	}
	return result, nil
}
*/