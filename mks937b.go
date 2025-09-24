/*
Author: Leonardo Rossi Leao
Created at: September 23rd, 2025
Last update: September 23rd, 2025
*/

package mks937b

import (
	"github.com/devicehub-go/mks-937b/protocol"
	"github.com/devicehub-go/unicomm"
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
		Address: address,
	}
}