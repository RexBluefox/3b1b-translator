package main

import (
	//"html/template"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	//"net/http"
	/* "github.com/go-git/go-git/v5" */ //"github.com/philippta/go-template/html/template"
)

func main() {
	var jBody []DirItem
	jBody = getDirContent("https://api.github.com/repos/3b1b/captions/contents/")
	for i := 0; i < len(jBody); i++ {
		if jBody[i].Type == "dir" && !strings.HasPrefix(jBody[i].Name, ".") {
			body := getDirContent(jBody[i].Url)
			
		}
	}
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

func getDirContent(url string) []DirItem {
	resp, _ := http.Get(url)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var jBody []DirItem
	json.Unmarshal(body, &jBody)
	return jBody
}
func getFileContent(url string) FileItem {
	resp, _ := http.Get(url)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var jBody FileItem
	json.Unmarshal(body, &jBody)
	return jBody
}
