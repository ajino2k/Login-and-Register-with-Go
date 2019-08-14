package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	

	helper "./helpers"
)

func main () {
	
	uName, email, pwd, pwdConfirm := "", "", "", ""

	mux := http.NewServeMux()

	//Signup
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()
		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")

		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmcheck := helper.IsEmpty(pwdConfirm)


		if 	uNameCheck || emailCheck || pwdCheck || pwdConfirmcheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
			return 
		}
		
		if pwd == pwdConfirm {
			//Save to database (username, email and password)
			fmt.Fprintf(w, "Registration successful.")
		} else {
			fmt.Fprintf(w, "Password infomation must be the same.")
		}

	})

	// Login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()
		email = r.FormValue("email")
		pwd = r.FormValue("password")

		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
			return
		}

		dbPwd := "cuongnp2!*."
		dbEmail := "cuongnp2@vng.com.vn"


		if email == dbEmail && pwd == dbPwd {
			fmt.Println("Login successful!")
			fmt.Fprintf(w, "Login successful!")
		} else {
			fmt.Println("Login failed!")
			fmt.Fprintf(w, "Login failed!")
		}
	})

	// Write log to file 
	f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
			log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "prefix", log.LstdFlags)
	logger.Println("text to append")
	logger.Println("more text to append")

	http.ListenAndServe(":8080", mux)
}