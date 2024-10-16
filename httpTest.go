package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func httpTest() {
	httpRequest()
	//httpPostUrl := "https://video.belet.tm/ru/watch"
	//fmt.Println("HTTP POST JSON URL:", httpPostUrl)
	//
	//var jsonData = []byte(`
	//      {
	//       "name":"John",
	//       "job":"leader"
	//      }`)
	//req, err := http.NewRequest("POST", httpPostUrl, bytes.NewBuffer(jsonData))
	//if err != nil {
	//	panic(err)
	//}
	//req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	//
	//client := &http.Client{}
	//res, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer res.Body.Close()
	//
	//fmt.Println("Response header: ", res.Header)
	//fmt.Println("Response status: ", res.Status)
	//
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println("Response body: ", string(body))
}

func httpRequest() {
	// GET request
	//httpUrl := "https://httpbin.org/get"
	//resp, err := http.Get(httpUrl)
	//if err != nil {
	//	panic(err)
	//}
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	// POST request

	//httpUrl := "https://httpbin.org/post"
	//message := map[string]interface{}{
	//	"1": "First",
	//	"2": 2,
	//	"embedded": map[string]string{
	//		"3": "third",
	//	},
	//}
	//
	//messageBytes, err := json.Marshal(message)
	//if err != nil {
	//	panic(err)
	//}
	//resp, err := http.Post(httpUrl, "application/json", bytes.NewBuffer(messageBytes))
	//if err != nil {
	//	panic(err)
	//}
	//var result map[string]string
	//json.NewDecoder(resp.Body).Decode(&result)
	//fmt.Println(result)
	//fmt.Println(result["data"])
	//
	//formData := url.Values{ // map[string][]string
	//	"name": {"John"},
	//}
	//resp, err := http.PostForm("https://httpbin.org/post", formData)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//var result map[string]interface{}
	//json.NewDecoder(resp.Body).Decode(&result)
	//
	//fmt.Println(result)
	//fmt.Println(result["form"])

	// Кастомный Clients/Requests
	client := http.Client{}
	request, _ := http.NewRequest("GET", "https://httpbin.org/get", nil)

	resp, _ := client.Do(request)
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)
}
