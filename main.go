package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type ServiceConfig struct {
	ServicePath string
	ServiceDomain string
}

type Service struct {
	Name string
	Configs []ServiceConfig
}

type Stack struct {
	Name string
	Id int
	Services []Service
}

type TodoPageData struct {
	PageTitle string
	Stacks []Stack
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	c1 := ServiceConfig{ServicePath: "C1"}
	c2 := ServiceConfig{ServicePath: "C2"}
	c3 := ServiceConfig{ServicePath: "C3", ServiceDomain: "c3.com"}
	c4 := ServiceConfig{ServicePath: "C4"}
	c5 := ServiceConfig{ServicePath: "C5"}
	c6 := ServiceConfig{ServicePath: "C6"}

	s1 := Service{Name: "S1", Configs: []ServiceConfig{c1}}
	s2 := Service{Name: "S2", Configs: []ServiceConfig{c2, c3}}
	s3 := Service{Name: "S3", Configs: []ServiceConfig{c4}}
	s4 := Service{Name: "S4", Configs: []ServiceConfig{c5}}
	s5 := Service{Name: "S5", Configs: []ServiceConfig{c6}}

	stack1 := Stack{Name: "Stack A", Id: 1, Services: []Service{s1, s2}}
	stack2 := Stack{Name: "Stack B", Id: 2, Services: []Service{s3}}
	stack3 := Stack{Name: "Stack C", Id: 3, Services: []Service{s4, s5}}

	stacks := make([]Stack, 3)
	stacks[0] = stack1
	stacks[1] = stack2
	stacks[2] = stack3


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Stacks: stacks,
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":9777", nil)
}