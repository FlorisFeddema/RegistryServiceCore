package registry

import (
	"CoreService/src/util"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Test()  {
	registry := util.GetConfig().Registry

	host := fmt.Sprintf("https://%s:%d/", registry.Host, registry.Port )
	uri := fmt.Sprintf("%s%s", host, "v2/")

	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Docker-Distribution-API-Version", "registry/2.0")
	req.SetBasicAuth(registry.Username, registry.Password)

	if err != nil {
		log.Fatal(err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	log.Fatal()
}
