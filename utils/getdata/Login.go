package getdata

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	// third party dependencies
	"github.com/PuerkitoBio/goquery"
)

// inccu登入
const (
	loginUrl      = "https://i.nccu.edu.tw/Login.aspx?ReturnUrl=%2f"
	getSSOBaseUrl = "https://i.nccu.edu.tw/sso_app/PersonalInfoSSO.aspx?p=1&sid="
	ssoLoginUrl   = "https://moltke.nccu.edu.tw/SSO/login.sso"
	// 全人網
	selfDevelopMenuUrl = "https://moltke.nccu.edu.tw/selfDevelop_SSO/login.selfDevelop"
)

func getSSOPassword(client *http.Client, studentId string, password string) string {
	// inccu單一入口登入
	getloginHtmlRequest, err := http.NewRequest("GET", loginUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	getloginHtmlResponse, err := client.Do(getloginHtmlRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer getloginHtmlResponse.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(getloginHtmlResponse)
	if err != nil {
		log.Fatal(err)
	}
	viewstate := doc.Find("#__VIEWSTATE").First().AttrOr("value", "")
	viewstategenerator := doc.Find("#__VIEWSTATEGENERATOR").First().AttrOr("value", "")
	data := url.Values{
		"__VIEWSTATE":             {viewstate},
		"__VIEWSTATEGENERATOR":    {viewstategenerator},
		"__EVENTTARGET":           {"captcha$Login1$LoginButton"},
		"__EVENTARGUMENT":         {""},
		"captcha$Login1$UserName": {studentId},
		"captcha$Login1$Password": {password},
	}
	if err != nil {
		log.Fatal(err)
	}
	loginResponse, err := client.Post(loginUrl,
		"application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	defer loginResponse.Body.Close()
	getSSOResponse, err := client.Get(getSSOBaseUrl + studentId)
	if err != nil {
		log.Fatal(err)
	}
	getSSOdoc, err := goquery.NewDocumentFromResponse(getSSOResponse)
	if err != nil {
		log.Fatal(err)
	}
	ssoPassword := getSSOdoc.Find("#password").First().AttrOr("value", "")
	return ssoPassword
}

func loginSelfDevelopment(client *http.Client, studentId string, ssoPassword string) {
	// 全人網登入
	data := url.Values{
		"id":       {studentId},
		"password": {ssoPassword},
		"p":        {"1"},
		"url":      {selfDevelopMenuUrl},
	}
	loginSelfDevelopmentRequest, err := http.NewRequest("GET", ssoLoginUrl+"?"+data.Encode(), nil)
	if err != nil {
		log.Fatal(err)
	}
	loginSelfDevelopmentResponse, err := client.Do(loginSelfDevelopmentRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer loginSelfDevelopmentResponse.Body.Close()
}
