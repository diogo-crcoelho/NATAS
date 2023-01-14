package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	// "regexp"
	// "strings"
)

var (
	username = "natas2"
	password = "h4ubbcXrWqsTo7GGnnUMLppXbOogfBZ7"
	url = "http://natas2.natas.labs.overthewire.org" 
)

func main (){
	request, _ := http.NewRequest("GET", url, nil)

	request.SetBasicAuth(username, password)
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()	
	content ,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	// regex := regexp.MustCompile(`<!--The password for natas(.*)>`)
	// almost := strings.Split(regex.FindAllString(string(content), 1)[0], " ")
	// answer := almost[len(almost) - 2]
	// fmt.Println(answer)
}