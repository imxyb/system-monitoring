package main

import (
    "html/template"
    "net/http"
    "log"
    "github.com/gorilla/websocket"
)

func watch(w http.ResponseWriter, r *http.Request) {
    t := template.New("watch.html")
    t, _ = t.ParseFiles("src/template/watch.html")
    t.Execute(w, "abc")
}

var upgrader = websocket.Upgrader{} // use default options

func serveWs(w http.ResponseWriter, r *http.Request) {
    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()

}

func main() {
    http.HandleFunc("/", watch)
    http.HandleFunc("/ws", serveWs)
    log.Fatal(http.ListenAndServe(":8081", nil))
}
