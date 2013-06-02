http in go lang


#1. basic GET method

resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)

**Attention: we must close the file handler before exit**

#2. basic POST method

#3. change HTTP Headers

#4. POST with binary data

#5.