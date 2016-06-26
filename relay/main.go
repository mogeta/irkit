//Package relay
//irkit doc http://getirkit.com/
package relay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/koron/go-dproxy"
	"github.com/miekg/dns"
	"github.com/soh335/go-mdns"
)

var address string

func init() {
	var irkitName = "_irkit._tcp.local."
	client := new(mdns.Client)
	client.Discover(irkitName, func(msg *dns.Msg) {
		for _, rr := range msg.Extra {
			switch rr := rr.(type) {
			case *dns.A:
				fmt.Println(rr.Header().Name, "=>", rr.A)
				address = "http://" + rr.A.String()
			default:
			}
		}
	})
}

//SendMessage is ...
func SendMessage(data string) error {
	url := address + "/messages"
	req, _ := http.NewRequest("POST", url, bytes.NewBufferString(data))
	req.Header.Set("X-Requested-With", "curl")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

//GetMessages get a Infrared data from irkit.
func GetMessages() (string, error) {
	url := address + "/messages"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Requested-With", "curl")

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return "can not read" + url, err
	}

	defer resp.Body.Close()

	//get response
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "can not read", err
	}

	return string(result), nil
}

//RequestClientToken request irkit client token
func RequestClientToken() (string, error) {
	url := address + "/keys"
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("X-Requested-With", "curl")

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return "can not read" + url, err
	}

	defer resp.Body.Close()

	//get response
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return "can not read", err
	}

	//json
	var data interface{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		fmt.Println(result)
		return "error", err
	}

	//read data
	s, err := dproxy.New(data).M("clienttoken").String()
	if err != nil {
		return "err", err
	}

	return s, nil
}

//GetIPAddress get irkit server address
func GetIPAddress() string {
	return address
}
