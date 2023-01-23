package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "bytes"
)

var Form map[string]string

func get(w http.ResponseWriter, r *http.Request) {
    fmt.Println("our server")
    if len(Form) == 0 {
        fmt.Println("there is no data yet :(")
    } else {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(Form)
    }
}

func post(w http.ResponseWriter, r *http.Request) {
    fmt.Println("inserting value... ", r.Method)
    Form = map[string]string{
        "ID": "1",
        "Name": "moein",
    }

    jsonData, err := json.Marshal(Form)
    errRes := "error occur!"
    if err != nil {
        http.Error(w, errRes, http.StatusMethodNotAllowed)
        return
    }

    resp, err := http.Post("https://localhost:8080/insert", "application/json; charset=utf-8", bytes.NewBuffer(jsonData))
    if err != nil {
        http.Error(w, errRes, http.StatusMethodNotAllowed)
        return
    }
    defer resp.Body.Close()

    bodyBytes, _ := ioutil.ReadAll(resp.Body)
    bodyString := string(bodyBytes)
    fmt.Println(bodyString)
    json.Unmarshal(bodyBytes, &Form)
    fmt.Printf("%+v\n", Form)
}

func del(w http.ResponseWriter, r *http.Request) {
    fmt.Println("for delete value")
    // for example delete item id: 1
    // also we can get input from user to specefic item
    delete(Form, "ID")
    fmt.Printf("%+v\n", Form)
}

func main() {
    http.HandleFunc("/", get)
    http.HandleFunc("/insert", post)
    http.HandleFunc("/delete", del)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// I know these problems:
// status code somewhere is not right
// some bugs when got errors
