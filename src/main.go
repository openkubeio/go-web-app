package main


import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    H "./handlers" 
    P "./handlers/pods"
)


func defaultHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, r.URL.Path )
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "Hello!")
}

func remoteAPIHandler(w http.ResponseWriter, r *http.Request) {

    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
    if err != nil {
       log.Fatalln(err)
    }

    // Read the response body on the line below.
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
       log.Fatalln(err)
    }
    
    // Convert the body to type string
    sb := string(body)
    fmt.Fprintf(w, sb)
}



func main() {

    http.HandleFunc("/",     defaultHandler)
    http.HandleFunc("/hello",helloHandler)
    http.HandleFunc("/remoteAPI", remoteAPIHandler)
    
    H.FirstBasicHandler()
    H.SecondBasicHandler()
    P.PodHandler()
    
    log.Printf("Starting server at port 8080\n")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}