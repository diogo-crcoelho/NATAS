package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	username = "natas1"
	password = "g9D9cREhslqBKtcA2uocGHPfMZVzeFK6"
	url = "http://natas1.natas.labs.overthewire.org" 
)

func main (){
	request, _ := http.NewRequest("GET", url, nil)

	request.SetBasicAuth(username, password)
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()	
	content ,_ := ioutil.ReadAll(response.Body)
	regex := regexp.MustCompile(`<!--The password for natas(.*)>`)
	almost := strings.Split(regex.FindAllString(string(content), 1)[0], " ")
	answer := almost[len(almost) - 2]
	fmt.Println(answer)
}