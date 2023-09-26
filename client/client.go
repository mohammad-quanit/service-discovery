package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

var url string

func main() {
	serviceDiscoveryWithConsul()
	fmt.Println("Starting Client.")
	var client = &http.Client{
		Timeout: time.Second * 30,
	}
	callServerEvery(time.Second*10, client)
}

func serviceDiscoveryWithConsul() {
	consulClient, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		fmt.Println(err)
	}

	services, err := consulClient.Agent().Services()
	if err != nil {
		fmt.Println(err)
	}

	service := services["helloworld-server"]
	address := service.Address
	port := service.Port
	url = fmt.Sprintf("http://%s:%v/hw", address, port)

}

func hello(t time.Time, client *http.Client) {
	response, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, _ := io.ReadAll(response.Body)
	fmt.Printf("%s. Time is %v\n", body, t)
}

func callServerEvery(d time.Duration, client *http.Client) {
	for x := range time.Tick(d) {
		hello(x, client)
	}
}
