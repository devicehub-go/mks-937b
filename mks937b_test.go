package mks937b_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	mks937b "github.com/devicehub-go/mks-937b"
	"github.com/devicehub-go/unicomm"
	"github.com/devicehub-go/unicomm/protocol/unicommtcp"
)

func TestReadPressure(t *testing.T) {
	fmt.Println("Initialing read pressure example...")

	options := unicomm.UnicommOptions{
		Protocol: unicomm.TCP,
		TCP: unicommtcp.TCPOptions{
			Host: "10.0.4.135",
			Port: 4001,
			ReadTimeout: 500 * time.Millisecond,
			WriteTimeout: 500 * time.Millisecond,
		},
		Delimiter: "\r",
	}

	mks := mks937b.New(48, options)
	if err := mks.Connect(); err != nil {
		log.Fatal(err)
	}
	defer mks.Disconnect()

	if err := mks.SetPowerStatus(1, false); err != nil {
		log.Fatal(err)
	}

	response, err := mks.GetPressure(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
	fmt.Println(mks.GetSensorStatus(1))
}

func TestSystemCommands(t *testing.T) {
	fmt.Println("Initializing system commands test...")

	options := unicomm.UnicommOptions{
		Protocol: unicomm.TCP,
		TCP: unicommtcp.TCPOptions{
			Host: "10.0.4.135",
			Port: 4001,
			ReadTimeout: 500 * time.Millisecond,
			WriteTimeout: 500 * time.Millisecond,
		},
		Delimiter: "\r",
	}

	mks := mks937b.New(48, options)
	if err := mks.Connect(); err != nil {
		log.Fatal(err)
	}
	defer mks.Disconnect()

	// Testing address functions
	fmt.Println("Address Test")
	address, err := mks.GetAddress()
	if err != nil {
		t.Error(err)
	}
	if err := mks.SetAddress(int(address)); err != nil {
		t.Error(err)
	}

	// Testing baudrate function
	fmt.Println("Baud Rate Test")
	baudrate, err := mks.GetBaudRate()
	if err != nil {
		t.Error(err)
	}
	if err := mks.SetBaudRate(int(baudrate)); err != nil {
		t.Error(err)
	}
}