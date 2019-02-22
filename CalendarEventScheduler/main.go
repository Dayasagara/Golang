package main
 
import (
    "fmt"
    "net/http"
    "database/sql"
	"log"
	mydb "./mydb"
    _ "github.com/lib/pq"
    helper "./helpers"
)
 
func main() {
 
    id,subject,StartDateTime,EndDateTime := "", "", "", ""
    mux := http.NewServeMux()
	db := connectToDatabase()
	defer db.Close()
    
    //For adding an event
    mux.HandleFunc("/AddEvent", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
 
        id = r.FormValue("id")     // Data from the form
        subject = r.FormValue("subject")   // Data from the form
        StartDateTime = r.FormValue("StartDateTime")   // Data from the form
        EndDateTime = r.FormValue("EndDateTime") // Data from the form

        idCheck := helper.IsEmpty(id)  //Check if the data is empty to prevent inserting them
        subjectCheck := helper.IsEmpty(subject)
        StartDateTimeCheck := helper.IsEmpty(StartDateTime)
        EndDateTimeCheck := helper.IsEmpty(EndDateTime)
 
        if idCheck || subjectCheck || StartDateTimeCheck || EndDateTimeCheck {
            fmt.Fprintf(w, "There is empty data.")
            return
        }
 
        status:=mydb.AddEvent(id,subject,StartDateTime,EndDateTime)
        if status==0{
            fmt.Fprintf(w,"Added Successfully")
        }
    })
 
    //To get event by date
    /*mux.HandleFunc("/GetEventByDate", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()

        date = r.FormValue("date") // Data from the form
		
		if event, err := mydb.GetEventByDate(date); err == nil {
            log.Printf("%v\n", event)
			return
		} else {
			log.Printf("error was: %v\n",err)
		}
    })*/
        /*
    //List all events
    mux.HandleFunc("/GetAllEvents", func(w http.ResponseWriter, r *http.Request) {
		
        if event, err := mydb.GetAllEvents()
        err == nil {
            log.Printf("%v\n", event)
            fmt.Fprintf(w,""+event)
			return
		} else {
			log.Printf("error was: %v\n",err)
		}
    })

    //Modify an event based on unique identifier
    mux.HandleFunc("/ModifyEventById", func(w http.ResponseWriter,r *http.Request){
        r.ParseForm()
        id = r.FormValue("id")     // Data from the form
        name = r.FormValue("name")  // Data from the form
        date = r.FormValue("date")  // Data from the form
        time = r.FormValue("time") // Data from the form
        duration = r.FormValue("duration")

        mydb.ModifyEventById(id,name,date,time,duration)

    })*/

    //Delete Event by Id
    /*mux.HandleFunc("/DeleteEventById", func(w http.ResponseWriter,r *http.Request){
        r.ParseForm()
        id = r.FormValue("id")     // Data from the form
        mydb.DeleteEventById(id)

    })*/
 
    http.ListenAndServe(":8080", mux)
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	mydb.SetDatabase(db)
	return db
}