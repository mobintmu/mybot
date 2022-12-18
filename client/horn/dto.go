package horn

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewClient(url string) *HornStruct {
	return &HornStruct{
		Url: url,
	}
}

func (h HornStruct) Call(message string) {

	fmt.Println("calling")

	url := h.Url
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"text":"` + message + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
