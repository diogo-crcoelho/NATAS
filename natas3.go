package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	username = "natas3"
	password = "G6ctbMJ5Nb4cbFwhpMPSvxGHhQ7I6W8Q"
	url = "http://natas3.natas.labs.overthewire.org" 
)

func main (){
	request, _ := http.NewRequest("GET", url + "/s3cr3t/users.txt", nil)

	request.SetBasicAuth(username, password)
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()	

	content ,_ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))
	regex := regexp.MustCompile(`natas(.*)`)
	almost := strings.Split(regex.FindAllString(string(content), 1)[0], ":")
	answer := almost[len(almost) - 1]
	fmt.Println(answer)
}