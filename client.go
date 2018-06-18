package main

import (
	"compress/flate"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://localhost:1323"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "2f744697-1f6e-4b61-ad30-52b4ee8bd0dd")
	req.Header.Add("Accept-Encoding", "gzip, flate")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	// body, _ := httputil.DumpResponse(res, true)
	rd := flate.NewReader(res.Body)
	body, _ := ioutil.ReadAll(rd)

	fmt.Println(res)
	fmt.Println(string(body))
}
