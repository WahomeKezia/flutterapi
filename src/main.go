package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
) ;
type task struct {
	ID         string `json:"id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}

var tasks []task

func allTasks() {
	Task1 := task{
		ID:         "1",
		TaskName:   "New Projects",
		TaskDetail: "Must DO's for new projects",
		Date:       "2023-05-07",
	}

	tasks = append(tasks, Task1)
	Task2 := task{
		ID:         "2",
		TaskName:   "Flutter Projects",
		TaskDetail: "Finish Flutter Projects",
		Date:       "2023-05-07",
	}

	tasks = append(tasks, Task2)
	fmt.Println("Your tasks are:", tasks)
}

func homePage(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getTask(w http.ResponseWriter, r *http.Request) {
     taskId := mux.Vars(r)["id"]
	 flag := false
	 for i:=0; i<len(tasks); i++ {
		 if tasks[i].ID == taskId {
			 flag = true 
			 json.NewEncoder(w).Encode(tasks[i])
			 break
		 }
	 }
	 if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"message": "Task not found"})

	 }
	 
	}
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application")
	json.NewEncoder(w).Encode(tasks)
		}
func createTask(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Welcome to the HomePage!")
			}
func deleteTask(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Welcome to the HomePage!")
				}
func updateTask(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintf(w, "Welcome to the HomePage!")
					}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettask", getTask).Methods("GET") 
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")



	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRoutes()
	fmt.Println("Hello Flutter Girls")
	allTasks()
}
