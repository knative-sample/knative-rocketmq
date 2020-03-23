package orderservice

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type OrderInfo struct {
	OrderId     string `json:"orderId"`
	OrderStatus string `json:"orderStatus"`
	UserPhoneNo string `json:"userPhoneNo"`
	ProdId      string `json:"prodId"`
	ProdName    string `json:"prodName"`
	ChargeMoney string `json:"chargeMoney"`
	ChargeTime  string `json:"chargeTime"`
	FinishTime  string `json:"finishTime"`
}

type CityInfo struct {
	Name     string `json:"name"`
	Adcode   string `json:"adcode"`
	Citycode string `json:"citycode"`
	IsCity   string `json:"iscity"`
}

func QueryWeather(url, requestBody string) (body []byte, err error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", url, strings.NewReader(requestBody))
	if err != nil {
		fmt.Printf("http send request url %s fails -- %v ", url, err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http send request url %s fails -- %v ", url, err)
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	//status code not in [200, 300) fail
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Printf("response status code %d, error messge: %s", resp.StatusCode, string(body))
		return
	}
	if err != nil {
		fmt.Printf("read the result of get url %s fails, response status code %d -- %v", url, resp.StatusCode, err)
	}
	return
}
