package main

import (
	"fmt"

	"strconv"

	"log"

	"html/template"

	"net/http"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title": "Personal Web",
}

type Project struct {
	Title       string
	Start_date  string
	End_date    string
	Description string
}

var Projects = []Project{
	{
		Title:       "Dumbways Mobile App",
		Start_date:  "20 Agustus 2022",
		End_date:    "20 Oktober 2022",
		Description: "Test",
	},
}

func main() {
	route := mux.NewRouter()

	// route path folder untuk public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/formProject", formProject).Methods("GET")
	route.HandleFunc("/projectPage", projectPage).Methods("GET")
	route.HandleFunc("/projectDetail/{index}", projectDetail).Methods("GET")
	route.HandleFunc("/addProject", addProject).Methods("POST")
	route.HandleFunc("/deleteProject/{index}", deleteProject).Methods("GET")
	route.HandleFunc("/updateForm", updateForm).Methods("GET")
	route.HandleFunc("/updateProject", updateProject).Methods("POST")

	fmt.Println("Server Running on port 8000")
	http.ListenAndServe("localhost:8000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	respData := map[string]interface{}{
		"Projects": Projects,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contact-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func formProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contact-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/form-project.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func projectPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contact-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/project-page.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	respData := map[string]interface{}{
		"Projects": Projects,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contact-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/project-page.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	var ProjectDetail = Project{}

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	for i, data := range Projects {
		if index == i {
			ProjectDetail = Project{
				Title:       data.Title,
				Start_date:  data.Start_date,
				End_date:    data.End_date,
				Description: data.Description,
			}
		}
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	fmt.Println(data)

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	var title = r.PostForm.Get("input-title")
	var startDate = r.PostForm.Get("start-date")
	var endDate = r.PostForm.Get("end-date")
	var description = r.PostForm.Get("project-description")

	var newProject = Project{
		Title:       title,
		Start_date:  startDate,
		End_date:    endDate,
		Description: description,
	}

	Projects = append(Projects, newProject)

	// fmt.Println(Projects)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	fmt.Println(index)

	Projects = append(Projects[:index], Projects[index+1:]...)

	http.Redirect(w, r, "/", http.StatusFound)
}

func updateForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contact-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/form-update.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func updateProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Masih Belum Mengerti Cara Update")
}
