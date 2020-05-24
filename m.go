package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	cok, err := r.Cookie(checkUser().Username)
	if err != nil || cok.Value != checkUser().Password {
		tpl.ExecuteTemplate(w, "index2.html", nil)
	} else {
		tpl.ExecuteTemplate(w, "upload2.html", nil)
		fmt.Println(*cok)
	}
	//w.Write([]byte("hello112323"))
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		username := r.FormValue("username")
		password := r.FormValue("password")
		myuser := checkUser()
		if username != myuser.Username || password != myuser.Password {
			cokie := &http.Cookie{Name: username, MaxAge: -1, Path: "/"}
			http.SetCookie(w, cokie)
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("forbidden"))
		} else {
			cokie := &http.Cookie{Name: username, Value: password, Path: "/"}
			http.SetCookie(w, cokie)
			w.WriteHeader(http.StatusOK)
			tpl.ExecuteTemplate(w, "upload2.html", nil)
			//	tpl.ExecuteTemplate(w, "upload.html", nil)
		}
		//}
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie(checkUser().Username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
	} else {
		cokie := &http.Cookie{Name: checkUser().Username, MaxAge: -1, Path: "/"}
		http.SetCookie(w, cokie)
		tpl.ExecuteTemplate(w, "index2.html", nil)
	}
}
func handleForgot(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "forget2.html", nil)
	} else if r.Method == "POST" {
		if r.FormValue("username") == checkUser().Username {
			if r.FormValue("Password Now") == checkUser().Password {
				changePassword(user{r.FormValue("username"), r.FormValue("New Password")})
				http.Redirect(w, r, "/index", 307)
			} else {
				//wrong password
				w.WriteHeader(http.StatusForbidden)
			}
		} else {
			//wrong account
			w.WriteHeader(http.StatusForbidden)

		}
	}

}

func main() {
	fmt.Println(checkUser())
	m := http.NewServeMux()
	m.HandleFunc("/", handleHome)
	m.HandleFunc("/form", handleForm)
	m.HandleFunc("/forgetPassword", handleForgot)
	m.HandleFunc("/logout", handleLogout)
	m.HandleFunc("/upload", handleHome)
	http.ListenAndServe(":8000", m)
}
