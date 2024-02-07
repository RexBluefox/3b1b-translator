package main

import (
	//"html/template"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	//"net/http"
	/* "github.com/go-git/go-git/v5" */ //"github.com/philippta/go-template/html/template"
)

func main() {
	godotenv.Load()
	var jBody []DirItem
	jBody = getDirContent("https://api.github.com/repos/3b1b/captions/contents/")
	fmt.Println(len(jBody))
	body := getDirContent(jBody[2].Url)
	body1 := getDirContent(body[0].Url)
	for i := 0; i < len(body1); i++ {
		if body1[i].Type == "dir" && !strings.HasPrefix(body1[i].Name, ".") {
			// body1 := getDirContent(body[i].Url)
			// fmt.Println(body[0].Url)
			if body1[i].Name == "german" {
				body2 := getDirContent(body1[i].Url)
				for j := 0; j < len(body2); j++ {
					//fmt.Println(body2[j].Name)
					r := getFileContent(body2[j].Url)
					esc := strings.Replace(r.Content, "\n", "", -1)
					decoded, _ := base64.StdEncoding.DecodeString(esc)
					fmt.Println(string(decoded))
				}
			}

		}
	}
	// for i := 0; i < len(jBody); i++ {
	// 	if jBody[i].Type == "dir" && !strings.HasPrefix(jBody[i].Name, ".") {
	// 		body := getDirContent(jBody[i].Url)
	// 		fmt.Println(body[0].Url)
	// 	}
	// }
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
	client := http.DefaultClient
	req, _ := http.NewRequest("GET",url,nil)
	req.Header.Set("Authorization",os.Getenv("API_TOKEN"))
	resp, _ := client.Do(req)
	//resp, _ := http.Get(url)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var jBody []DirItem
	json.Unmarshal(body, &jBody)
	return jBody
}
func getFileContent(url string) FileItem {
	client := http.DefaultClient
	req, _ := http.NewRequest("GET",url,nil)
	req.Header.Set("Authorization",os.Getenv("API_TOKEN"))
	resp, _ := client.Do(req)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var jBody FileItem
	json.Unmarshal(body, &jBody)
	return jBody
}
