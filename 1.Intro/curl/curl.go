package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	curlClient := http.Client{}
	if len(os.Args) != 3 {
		log.Fatalln("Need exact 2 argument <Method> and <URL>")
	}
	req2, err := http.NewRequest(os.Args[1], os.Args[2], nil)
	if err != nil {
		log.Fatalln("could not create a request")
	}

	res2, err := curlClient.Do(req2)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		log.Fatalln("Could not read response data.")
	}

	log.Println(string(data))

	err = res2.Body.Close()
	if err != nil {
		panic(err)
	}
}
