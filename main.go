package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var rand_ *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const Char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func isRandomChar(len_ int, RandChar string) string {
	index := make([]byte, len_) 

	for i := range index { 
		index[i] = RandChar[rand_.Intn(len(RandChar))]
	}
	return string(index) 
}

func main() {
	var name string

	var domain string = "WhiteHole_@gmail.com"

	fmt.Print("ENTER KOREA NAME >> ")
	fmt.Scanln(&name)

	var FullEmail string = isRandomChar(5, Char) + domain

	var ID string = isRandomChar(5, Char)

	var PW string = isRandomChar(5, Char)

	data := url.Values{
		"reg_name":  {name},
		"reg_email": {FullEmail},
		"reg_ids":   {ID},
		"reg_pw":    {PW}}

	var BaseURL string = "https://asked.kr/sing_ups.php"

	res, er := http.Get(BaseURL)

	if er != nil {
		log.Fatal(er)
		fmt.Println(er)
	}

	fmt.Println(res.StatusCode)

	req, err := http.PostForm("https://asked.kr/sing_ups.php", data)

	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	defer req.Body.Close()

	reqBody, erro := ioutil.ReadAll(req.Body) //리스폰 체크

	if erro == nil {
		check := string(reqBody)
		fmt.Println(check)
	}

	if req.StatusCode == 200 {

		fmt.Println("success")

		fmt.Println("ID : " + ID + "\n" + "PW : " + PW)
	}
}
