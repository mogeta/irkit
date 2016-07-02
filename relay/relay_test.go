package relay

import (
	"fmt"
	"testing"
)

func TestGetIPAddress(t *testing.T) {
	GetIPAddress()
}

func TestRequestClientToken(t *testing.T) {
	token, err := RequestClientToken()
	if err != nil {
		t.Fail()
	}
	fmt.Println(token)
}

func TestGetMessage(t *testing.T) {
	message, err := GetMessages()
	if err != nil {
		t.Fail()
	}
	fmt.Println(message)
}
func TestSendMessaget(t *testing.T) {
	var data = `{"format":"raw","freq":38,"data":[17421,8755,1190,1037,1190,1037,1190,1037,1190,1037,1190,1037,1190,1037,1190,3228,1190,1037,1150,3228,1190,3228,1190,3228,1190,3228,1190,3228,1190,3228,1190,1037,1150,3228,1150,1037,1190,3228,1190,1037,1190,1037,1190,3228,1150,1037,1190,1037,1190,1037,1190,3228,1190,1037,1190,3228,1190,3228,1190,1037,1190,3228,1150,3228,1150,3228,1150,65535,0,13693,17421,4400,1150]}`
	err := SendMessage(data)
	if err != nil {
		t.Fail()
	}

}
