package main

import (
	"byte"
	"github.com/bmizerany/aws4"
	"io/ioutil"
	"log"
	"net/http"
	//"github.com/davecgh/go-spew/spew"
)

func main() {
	url := "your call url"

	k := &aws4.Keys{
		AccessKey: "your ak",
		SecretKey: "your sk",
	}

	ser := &aws4.Service{
		Name:   "",
		Region: "",
	}

	bodystring := ""
	body, _ := json.Marshal(bodystring)

	reqbody := bytes.NewReader(body)
	r, _ := http.NewRequest("POST", url, bodys)
	//GET: r, _ := http.NewRequest("GET", url, nil)

	//add you need header here
	r.Header.Set("Content-Type", "application/json")

	ser.Sign(k, r)
	resp, err := aws4.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//print the response
	fmt.Println(string(responseData))

}
