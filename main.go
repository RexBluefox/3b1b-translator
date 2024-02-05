package main

import (
	//"html/template"
	"fmt"
	"log"

	//"net/http"
	"github.com/go-git/go-git/v5"
	//"github.com/philippta/go-template/html/template"
)

type Todo struct {
	Title string
	Done  bool
}
type Item struct {
	Content string
}
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
	Items     []Item
}

func main() {
	repoURL := "https://github.com/maragudk/gomponents.git"
	repo, err := git.PlainOpen(repoURL)
	if err != nil {
		log.Fatal(err)
	}
	/* repoInfo, err := repo.Remotes()
	if err != nil {
		log.Fatal(err)
	} */
	fmt.Println(repo)
	/* tmpl := template.Must(template.ParseFiles("layout.html", "dropdown.html"))
	dropdown := template.Must(template.ParseFiles("dropdown.html"))
	dropdown.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
			Items: []Item{
				{Content: "Lets"},
				{Content: "fucking"},
				{Content: "go"},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":1337", nil) */
}
func test() {
	fmt.Println("Test")
}
