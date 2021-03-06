package activity

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"shuZhiNet/model/student"
)

func OptOutActivity(student student.Student, id string) {
	jar, _ := cookiejar.New(nil)
	cancelURL, _ := url.Parse("http://www.sz.shu.edu.cn")
	jar.SetCookies(cancelURL, student.Cookies)
	client := http.Client{Jar: jar}
	getURL := "http://www.sz.shu.edu.cn/api/HuoDong/HuoDBMXX/DeleteHuoDBM?hdbmid=" + id
	_, _ = client.Get(getURL)
}
