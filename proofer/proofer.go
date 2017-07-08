package proofer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func New(key string) (*Proofer, error) {
	if len(key) < 10 {
		return nil, fmt.Errorf("Invalid api key")
	}
	return &Proofer{
		MeaningcloudKey: key,
	}, nil
}

func (proofer *Proofer) Check(text string) (Result, error) {
	apiUrl := URL + "key=" + proofer.MeaningcloudKey + "&lang=es&ilang=en&of=" + OF_JSON
	data := url.Values{}
	data.Set("txt", text)
	res, err := http.Post(apiUrl, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return Result{}, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	if res.StatusCode == 200 {
		result := Result{}
		err := json.Unmarshal(body, &result)
		if err != nil {
			return Result{}, err
		}
		return result, nil
	}

	return Result{}, nil
}
