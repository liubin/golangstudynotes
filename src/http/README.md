http in go lang


#1. basic GET method
```go
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
```

**Attention: we must close the file handler before exit**

#2. basic POST method

#3. change HTTP Headers

#4. POST with binary data

#5. http server

##5.1. basic server without TLS(HTTPS)

The simple example likes below:


```go
// step 1. prepare a function to handle request
// func( http.ResponseWriter, *http.Request)
func HelloServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("Hello World.\n"))
}

func main() {

    // somthing like url mapping(route) in Rails
    http.HandleFunc("/hello", HelloServer)

    // start listening on port 8000
    err := http.ListenAndServe(":8000", nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

##5.2. secure http with TLS(HTTPS)

Just use **ListentAndServerTLS** insteads of **ListenAndServe** to ensure a secure http transfer.

```go
err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
```

##5.3. read and write cookie

write cookie:

```go
    expire := time.Now().AddDate(0, 0, 1)
    cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", Expires: expire}
    http.SetCookie(w, &cookie)
```


read cookie:

```go
    var cookie,err = req.Cookie("testcookiename")
    if err == nil {
        var cookievalue = cookie.Value
        io.WriteString(w, "get cookie value is " + cookievalue + "\n")
    }

```