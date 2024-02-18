package orchestrator

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func SendHTTPRequest(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func SendHTTPRequestWithContext(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func SendPostRequestToDemon(req Request) (*http.Response, error) {

	//data, err := ioutil.ReadFile("config.yaml")
	//if err != nil {
	//	return nil, err
	//}
	//
	//config := config2.Config{}
	//err = yaml.Unmarshal(data, &config)
	//if err != nil {
	//	return nil, err
	//}
	//url := config.Demon.Host + ":" + strconv.Itoa(config.Demon.Port)

	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := "http://localhost:8081"
	r := bytes.NewReader(reqData)
	return http.Post(url, "application/json", r)
}

func SendGetRequestToDemon() (Response, error) {

	//data, err := ioutil.ReadFile("config.yaml")
	//if err != nil {
	//	return nil, err
	//}
	//
	//config := config2.Config{}
	//err = yaml.Unmarshal(data, &config)
	//if err != nil {
	//	return nil, err
	//}
	//url := config.Demon.Host + ":" + strconv.Itoa(config.Demon.Port)
	client := http.Client{}
	var r io.Reader
	url := "http://localhost:8081/"
	get, err := http.NewRequest("GET", url, r)

	if err != nil {
		return Response{
			Status: "",
			Result: 0,
		}, err
	}
	resp, err := client.Do(get)
	if err != nil {
		panic(err)
	}
	get.Header.Set("Content-Type", "application/json")
	// Прочитайте ответ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var respns Response
	err = json.Unmarshal(body, &respns)
	return respns, err
}
