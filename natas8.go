package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"regexp"
	"strings"
)

var (
	username = "natas8"
	password = "a6bZCNYwdKqN5cGP11ZdtPg0iImQQhAB"
	urls = fmt.Sprintf("http://%s.natas.labs.overthewire.org",username) 
)

func main (){
	// request, _ := http.NewRequest("GET", urls, nil)
	data := url.Values{}
	data.Set("secret", "oubWYf2kBq")
	data.Set("submit", "submit")
	request, _ := http.NewRequest("POST", urls, strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(username, password)
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()

	content ,_ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))
	regex := regexp.MustCompile(`The password for natas(.*)`)
	almost := strings.Split(regex.FindAllString(string(content), 1)[0], " ")
	answer := almost[len(almost) -1]
	fmt.Println(answer)
}
