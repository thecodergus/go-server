package main

import (
  "html/template"
  "log"
  "net/http"
  "time"
)

type PageVariables struct {
	Date         string
	Time         string
}

type User struct{
	Name string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/Ola", Ola)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request){

    now := time.Now() // find the time right now
    HomePageVars := PageVariables{ //store the date and time in a struct
      Date: now.Format("02-01-2006"),
      Time: now.Format("15:04:05"),
    }

    t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
    if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
  	}
    err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
}

func Ola(w http.ResponseWriter, r *http.Request){
	HomePageVars := User{Name: "Gustavo"}

	t, err := template.ParseFiles("ola.html")
	if err != nil{
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, HomePageVars)
	if err != nil{
		log.Print("template executin error: ", err)
	}
}
