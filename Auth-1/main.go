package main
 
import (
    "fmt"
    "net/http"
    "encoding/json"
	"database/sql"
    "log"
    "os"
    "io/ioutil"
    "crypto/sha512"
	"encoding/base64"
	mydb "programs/AUTH-1/mydb"
    _ "github.com/lib/pq"
    helper "programs/AUTH-1/helpers"
    "github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func main() {
    
    uName, email, pwd, pwdConfirm := "", "", "", ""
    id,subject,StartDateTime,EndDateTime := "", "", "", ""
    mux := http.NewServeMux()
    db := connectToDatabase()
    
    defer db.Close()
    
    // Signup
    mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
 
        uName = r.FormValue("username")     // Data from the form
        email = r.FormValue("email")        // Data from the form
        pwd = r.FormValue("password")       // Data from the form
        pwdConfirm = r.FormValue("confirm") // Data from the form
 
        // Empty data checking
        uNameCheck := helper.IsEmpty(uName)
        emailCheck := helper.IsEmpty(email)
        pwdCheck := helper.IsEmpty(pwd)
        pwdConfirmCheck := helper.IsEmpty(pwdConfirm)
 
        if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
            fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
            return
        }
 
        if pwd == pwdConfirm {
            mydb.Signup(uName,email,pwd)
        } else {
            fmt.Fprintln(w, "Password information must be the same.")
		}
    })

    // Login
    mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
 
        email = r.FormValue("email")  // Data from the form
        pwd = r.FormValue("password") // Data from the form
 
        // Empty data checking
        emailCheck := helper.IsEmpty(email)
        pwdCheck := helper.IsEmpty(pwd)
 
        if emailCheck || pwdCheck {
            fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
            return
        }

		if user, err := mydb.Login(email, pwd); err == nil {
            hasher := sha512.New()
	        hasher.Write([]byte(pwd))
	        pwd1 := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
            token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
                "username": email,
                "password": pwd1,
            })
            tokenString, error := token.SignedString([]byte("secret"))
            if error != nil {
                fmt.Println(error)
            }
            json.NewEncoder(w).Encode(JwtToken{Token: tokenString})

            var file, err = os.Create(`path`)
            if err != nil {
                
            }  
            fmt.Fprintf(file,tokenString) 
            defer file.Close()
            
            log.Printf("User has logged in: %v\n", user)
			return
		} else {
			log.Printf("Failed to log user in with email: %v %v, error was: %v\n", email,pwd, err)
		}
    })
    
    mux.HandleFunc("/protected",func(w http.ResponseWriter, req *http.Request) {
        b, err := ioutil.ReadFile("creds.txt")
        if err != nil {
            fmt.Print(err)
        }
        fmt.Println(string(b))
        token, _ := jwt.Parse(string(b), func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("There was an error")
            }
            return []byte("secret"), nil
        })
        
        if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            
            fmt.Println("Hi Authenticated")
        } else {
            json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
        }
    })

    mux.HandleFunc("/AddEvent", func(w http.ResponseWriter, r *http.Request) {
        b, err := ioutil.ReadFile("creds.txt")
        if err != nil {
            fmt.Print(err)
        }
        fmt.Println(string(b))
        token, _ := jwt.Parse(string(b), func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("There was an error")
            }
            return []byte("secret"), nil
        })
        
        if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            
            fmt.Println("Hi Authenticated")
        } else {
            json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
        }
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
    http.ListenAndServe(":8000", mux)
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	mydb.SetDatabase(db)
	return db
}