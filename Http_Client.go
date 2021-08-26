// Go - HTTP Client

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	buf := new(bytes.Buffer)
	//json.NewEncoder(buf).Encode(body)
	req, _ := http.NewRequest("GET", "http://localhost:8080/view/A", buf)

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Fatal(e)
	}

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)

	io.Copy(os.Stdout, res.Body)
}
