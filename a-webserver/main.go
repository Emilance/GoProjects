package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	fmt.Println("Hello world")
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmp1 := template.Must(template.ParseFiles("index.html"))
		users := map[string][]User{
			"Users": {
				{Name: "The GodFather", Email: "FrancisFordf@gami.com ", Age: 20},
				{Name: "Blade Runner", Email: "Ridleyscott@yahoo", Age: 23},
				{Name: "The Thing", Email: "JohnCarpenter@muu.com", Age: 40},
			},
		}
		tmp1.Execute(w, users)
	}

	addUser := func(w http.ResponseWriter, r *http.Request) {
		Name := r.PostFormValue("name")
		Email := r.PostFormValue("email")
		// Parse Age as an integer
		Age, err := strconv.Atoi(r.PostFormValue("age"))
		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}
		tmp1 := template.Must(template.ParseFiles("index.html"))
		tmp1.ExecuteTemplate(w, "user-list-element", User{Name: Name, Email: Email, Age: Age})

		// htmlStr := fmt.Sprintf("  <li> <span>I %v </span><span> %v </span><span> %v </span></li>", Name, Email, Age)
		// tmp1, _ := template.New("t").Parse(htmlStr)
		// tmp1.Execute(w, nil)
		// fmt.Println("Name: ", Name)
		// fmt.Println("Email: ", Email)
		// fmt.Println("Age: ", Age)

	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-user/", addUser)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
