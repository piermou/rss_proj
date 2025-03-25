package archives

import (
	"net/http"
	"text/template"
)

var testTemplate *template.Template

type ViewData struct {
	User User
}

type User struct {
	ID    int
	Email string
	// HasPermission func(string) bool
}

// func (u User) NeverMind(feature string) bool {
// 	if feature == "feature-a" {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	vd := ViewData{User: User{ID: 1, Email: "jon@calhoun.io"}}

	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// var err error
// // testTemplate, err = template.ParseFiles("hello.html")
// testTemplate, err = template.New("hello.html").Funcs(template.FuncMap{
// 	"hasPermission": func(user User, feature string) bool {
// 		if user.ID == 1 && feature == "feature-a" {
// 			return true
// 		}
// 		return false
// 	},
// }).ParseFiles("hello.html")
// if err != nil {
// 	panic(err)
// }

// http.HandleFunc("/", handler)
// http.ListenAndServe(":3000", nil)
