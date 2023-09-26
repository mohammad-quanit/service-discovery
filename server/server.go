package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	serviceRegistryWithConsul()
	log.Println("Starting Server.")
	http.HandleFunc("/hw", helloWorld)
	http.HandleFunc("/check", check)
	http.ListenAndServe(getPort(), nil)
}

func serviceRegistryWithConsul() {
	// Get a new client
	consulClient, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}

	serviceID := "helloworld-server"
	port, _ := strconv.Atoi(getPort()[1:len(getPort())])
	address := "127.0.0.1"

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "helloworld-server",
		Port:    port,
		Address: address,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/check", address, port),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	regErr := consulClient.Agent().ServiceRegister(registration)

	if regErr != nil {
		log.Printf("Failed to register service: %s:%v ", address, port)
	} else {
		log.Printf("successfully register service: %s:%v", address, port)
	}

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("helloworld service is called.")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world.")
}

func check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Consul check")
}

func getPort() (port string) {
	port = os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	port = ":" + port
	return port
}

// func getHostname() (hostname string) {
// 	hostname, _ = os.Hostname()
// 	return hostname
// }
