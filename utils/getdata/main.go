package getdata

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	// third party dependencies
	// "github.com/PuerkitoBio/goquery"
)

const (
	// 通識
	comUrl = "https://moltke.nccu.edu.tw/selfDevelop_SSO/tscomsub.selfDevelop"
	// 共同必修
	other0Url = "https://moltke.nccu.edu.tw/selfDevelop_SSO/tscomsubothers.selfDevelop"
	// 系必修
	reqUrl = "https://moltke.nccu.edu.tw/selfDevelop_SSO/tsreqsub.selfDevelop"
	// 選修
	other1Url = "https://moltke.nccu.edu.tw/selfDevelop_SSO/tsothersub.selfDevelop?type=1"
	// 外系
	otehr2Url = "https://moltke.nccu.edu.tw/selfDevelop_SSO/tsothersub.selfDevelop?type=2"
)

// 取得全人網資料
func GetData(studentId string, password string) {
	//
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	var client = &http.Client{
		Jar: jar,
	}
	ssoPassword := getSSOPassword(client, studentId, password)
	fmt.Println(ssoPassword)
	if len(ssoPassword) == 0 {
		fmt.Println("invalid login information")
	}
	loginSelfDevelopment(client, studentId, ssoPassword)
}
