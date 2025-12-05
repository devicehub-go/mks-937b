/*
Author: Leonardo Rossi Leao
Created at: September 23rd, 2025
Last update: September 23rd, 2025
*/

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/devicehub-go/mks-937b/protocol"
	"github.com/devicehub-go/unicomm"
	"github.com/devicehub-go/unicomm/protocol/unicommtcp"
)

/*
Creates a new MKS 937B instance that allow to communicate
with the device through the selected protocol.

For MKS 937B some usual character format are: 8 data bits,
1 stop bit, and no parity. Baudrate by default is 9600
*/
func New(address int, options unicomm.UnicommOptions) *protocol.MKS937B {
	return &protocol.MKS937B{
		Communication: unicomm.New(options),
		Address:       address,
	}
}

func main() {
	options := unicomm.UnicommOptions{
		Protocol: unicomm.TCP,
		TCP: unicommtcp.TCPOptions{
			Host:         "10.0.28.38",
			Port:         4001,
			ReadTimeout:  500 * time.Millisecond,
			WriteTimeout: 500 * time.Millisecond,
		},
		Delimiter: "\r",
	}

	mks := New(3, options)
	if err := mks.Connect(); err != nil {
		log.Fatal(err)
	}
	defer mks.Disconnect()

	if err := mks.SetPowerStatus(3, true); err != nil {
		log.Fatal(err)
	}

	for {
		response, err := mks.GetPressure(3)
		if err != nil {
			log.Fatal(err)
		}
		f, _ := os.OpenFile("measure.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		fmt.Fprintf(f, "%s, %.3e\n", time.Now().Format(time.RFC3339), response.Value)
		f.Close()

		time.Sleep(5 * time.Second)
	}
}
