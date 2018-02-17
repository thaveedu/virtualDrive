package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// our main function
func main() {
    router := mux.NewRouter()

    router.HandleFunc("/file", CreateFile).Methods("POST")
    router.HandleFunc("/file", GetFileList).Methods("GET")
    router.HandleFunc("/file/{id}", GetFile).Methods("GET")
    router.HandleFunc("/file/{id}", UpdateFile).Methods("PUT")
    router.HandleFunc("/file/{id}/content", UploadFile).Methods("PUT")

    files = append(files, File{ID: "1", Name: "The Matrix", CreatedDate: "2018-02-01",UpdatedDate: "2018-02-01"})
    files = append(files, File{ID: "2", Name: "The Matrix Reloaded", CreatedDate: "2018-02-01",UpdatedDate: "2018-02-01"})
    files = append(files, File{ID: "3", Name: "The Matrix Revolutions", CreatedDate: "2018-02-01",UpdatedDate: "2018-02-01"})



    log.Fatal(http.ListenAndServe(":8000", router))
}

var files []File
func GetFile(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var file []File
    for _, item := range files {
        if item.ID == params["id"]{
            file = append(file,item)
            break
        }
    }
        json.NewEncoder(w).Encode(file)
}
func GetFileList(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(files)
}
func CreateFile(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    var file File
    _ = json.NewDecoder(r.Body).Decode(&file)
    file.ID = params["id"]
    files = append(files, file)
    json.NewEncoder(w).Encode(files)

}
func UpdateFile(w http.ResponseWriter, r *http.Request) {}
func UploadFile(w http.ResponseWriter, r *http.Request) {}
func DelteFile(w http.ResponseWriter, r *http.Request) {}


type File struct {
    ID        string   `json:"Id,omitempty"`
    Name string   `json:"Name,omitempty"`
    CreatedDate  string   `json:"CreatedDate,omitempty"`
    UpdatedDate  string   `json:"UpdatedDate,omitempty"`
}


