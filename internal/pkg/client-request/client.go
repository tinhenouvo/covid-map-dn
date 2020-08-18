package client_request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"crawlerdatacovid/internal/app/model"
)

type CustomJar struct {
	jar map[string][]*http.Cookie
}

func (p *CustomJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	p.jar[u.Host] = cookies
}

func (p *CustomJar) Cookies(u *url.URL) []*http.Cookie {
	return p.jar[u.Host]
}

func noRedirect(req *http.Request, via []*http.Request) error {
	return errors.New("Don't redirect!")
}

func ClientRequest() (*http.Response, *http.Client, error) {
	client := &http.Client{
		CheckRedirect: noRedirect,
	}
	jar := &CustomJar{}
	jar.jar = make(map[string][]*http.Cookie)
	client.Jar = jar

	req, err := http.NewRequest("GET", "https://covidmaps.danang.gov.vn/danang?locale=vn", nil)
	if err != nil {
		return nil, nil, err
	}
	transport := http.Transport{}
	resp, err := transport.RoundTrip(req)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 302 {
		return nil, nil, errors.New(fmt.Sprintf("Failed with status %v", resp.Status))
	}
	return resp, client, nil
}

func GetData() ([]model.DataCovid, error) {
	link := "https://api-open.busmap.vn/open_api/sc/get_pois?lat=16.073393070451104&lng=108.21673488616943&regionCode=dn&type=10&filter="
	//reader := bytes.NewBuffer()
	req, err := http.NewRequest("GET", link, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "vi-VN,vi;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("key", "0e8eea40d9638a9175b692ed01ae6add4de0ed439333dd052f2803360f2bcee6")
	req.Header.Set("key0", "aea38cdcb4d780538ea36b55e1e9b83d4617324951a0364595e1e9b472ff9bc5")
	req.Header.Set("origin", "https://covidmaps.danang.gov.vn")
	req.Header.Set("timestamp", "1597743495790")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var v = make([]model.DataCovid, 0)
	err = json.Unmarshal(body, &v)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(v)
	return v, err
}
