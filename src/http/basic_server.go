package main

import (
    "io"
    "net/http"
    "log"
    "fmt"
    "time"
)


func HelloServer(w http.ResponseWriter, req *http.Request) {
    // set cookies.
    expire := time.Now().AddDate(0, 0, 1)

    cookie2 := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", Expires: expire, MaxAge: 86400}

    http.SetCookie(w, &cookie2)



    w.Header().Set("Content-Type", "text/plain")

    // 
    // we can not set cookie after writing something to ResponseWriter
    // if so ,we cannot set cookie succefully.


    // print out request info
    w.Write([]byte("Hello World.\n"))
    io.WriteString(w, req.Method + "\n")
    io.WriteString(w, req.RemoteAddr + "\n")
    io.WriteString(w, req.RequestURI + "\n")


    // read cookie
    // the cookievalue will be avalibale after the second time you access the page
    // before setting it 
    var cookie,err = req.Cookie("testcookiename")
    if err == nil {
        var cookievalue = cookie.Value
        io.WriteString(w, "get cookie value is " + cookievalue + "\n")
    }


}


func main() {

    http.HandleFunc("/", HelloServer)

    fmt.Println("listen on 3000")
    err := http.ListenAndServe(":3000", nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}