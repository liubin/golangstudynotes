package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    res, err := http.Get("http://t.nanshapo.com/ip.php")
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    result, err := ioutil.ReadAll(res.Body)

    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n",result)
}