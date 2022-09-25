package repository

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"wnexam/app"
)

//Data struct create
// type Data struct{
// 	Data []JSONData `json:"Data"`
// }

//JSONData struct create for url
// type JSONData struct{
// 	ConfirmDate string `json:"ConfirmDate"`
// 	No string `json:"No"`
// 	Age int `json:"Age"`
// 	Gender string `json:"Gender"`
// 	GenderEn string `json:"GenderEn"`
// 	Nation string `json:"Nation"`
// 	NationEn string `json:"NationEn"`
// 	Province string `json:"Province"`
// 	ProvinceID int `json:"ProvinceId"`
// 	District string `json:"District"`
// 	ProvinceEn string `json:"ProvinceEn"`
// 	StatQuarantine int `json:"StatQuarantine"`
// }

//DataRepository interface create
// type DataRepository interface {
// 	DecodeJSONUrl() (Data, error)
// 	GetJSONUrl() ([]byte,error)
// 	QueryInput() error
// 	GetInput() Data
// }

type dataRepository struct{
	url string
	input app.Data
}

//NewDataRepository create
func NewDataRepository(url string) *dataRepository{
	return &dataRepository{url:url}
}

//decode json data
func (r *dataRepository) DecodeJSONUrl() (app.Data, error) {
	body, err := r.GetJSONUrl()
	if err != nil {
		log.Panicln(err)
	}
	decode := app.Data{}
	err = json.Unmarshal(body, &decode)
	if err != nil {
		log.Println(err)
	}
	return decode, nil
}

//Get JSON data from url
func (r dataRepository) GetJSONUrl() ([]byte,error) {
	spaceClient := http.Client{
		Timeout: time.Second *2,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, //Skip to verify https cert
				},
				},
			}
	
	//send request
	req, err := http.NewRequest(http.MethodGet, r.url, nil)
	if err != nil {
		log.Println(err)
	}

	//get response
	res, err := spaceClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	//get body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	
	return body,nil
}

func (r dataRepository) QueryInput() error {
	spaceClient := http.Client{
		Timeout: time.Second *2,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, //Skip to verify https cert
				},
				},
			}
	
	//send request
	req, err := http.NewRequest(http.MethodGet, r.url, nil)
	if err != nil {
		return err
	}

	//get response
	res, err := spaceClient.Do(req)
	if err != nil {
		return err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	//get body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &r.input)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (r dataRepository) GetInput() app.Data {
	return r.input
}