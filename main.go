package main
//import ("fmt"; "net/http"; "html/template")
/*import ("fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)*/
import ("fmt"
		"net/http"
		"html/template"
)


func index(w http.ResponseWriter, r *http.Request) { 
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	http.HandleFunc("/", index) 
	http.ListenAndServe(":8080", nil)
}
/*type User struct {
	Name string `json:"name"`
	Age uint16 `json:"age"`
}*/

//func main() {
	//handleFunc()
	//fmt.Println("Hello")
	/*db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	 insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Alex', 35)")
	if err != nil {
		panic(err)
	}
	defer insert.Close() 


	res, err := db.Query("SELECT `name`, `age` FROM `users`")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	} */
/* 

	fmt.Println("Подключено") */
//}
/*
type User struct {
	Name string
	Age uint16
	Money int16
	Avg_grades, Happiness float64
	Hobbies []string
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name %s. He is %d years old, has money %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string)  {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) { 
	bob := User{"Bob", 25, -50, -4.2, 0.8, []string{"Football", "Skate", "Dance"}}
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
	
	//bob.setNewName("Alex")
	//fmt.Fprintf(w, bob.getAllInfo())
	//fmt.Fprintf(w, "Go is super easy!")
}
func contacts_page(w http. ResponseWriter, r *http.Request) { 
	fmt.Fprintf(w, "Contacts page!")
}
func about_page(w http. ResponseWriter, r *http.Request) { 
	fmt.Fprintf(w, "about page!")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page) 
	http.HandleFunc("/about/", about_page) 
	http.ListenAndServe(":8080", nil)
}
func main() {
	handleRequest()
}*/

