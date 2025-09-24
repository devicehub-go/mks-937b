/*
Author: Leonardo Rossi Leao
Created at: September 23rd, 2025
Last update: September 23rd, 2025
*/

package protocol

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/devicehub-go/unicomm"
)

type MKS937B struct {
	Communication unicomm.Unicomm
	Address int

	mutex sync.Mutex
}

/*
Establishes a connection with the device
*/
func (m *MKS937B) Connect() error {
	if m.Address < 1 || 254 < m.Address {
		return NewErrInvalidAddress(m.Address)
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Communication.Connect()
}

/*
Closes the connection with the device
*/
func (m *MKS937B) Disconnect() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Communication.Disconnect()
}

/*
Returns true if the device is connected
*/
func (m *MKS937B) IsConnected() bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.Communication.IsConnected()
}

/*
Queries a value from the device
*/
func (m *MKS937B) Query(command string) (string, error) {
	if !m.IsConnected() {
		return "", ErrNotConnected
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	addressStr := fmt.Sprintf("%03d", m.Address)
	message := fmt.Sprintf("@%s%s?;FF", addressStr, command)
	m.Communication.Write([]byte(message))

	response, err := m.Communication.ReadUntil(";FF")
	if err != nil {
		return "", err
	}
	responseStr := string(response)
	regex := regexp.MustCompile(`@([0-9]+)(?:ACK|NAK)(.*?);FF`)
	matches := regex.FindStringSubmatch(responseStr)

	if len(matches) < 3 {
		return "", NewErrUnexpectedReply(message, responseStr)
	}
	if matches[1] != addressStr {
		return "", NewErrUnexpectedAddress(addressStr, matches[1])
	}
	return matches[2], nil
}

/*
Sets a value to the device
*/
func (m *MKS937B) Set(command string, parameter string) error {
	if !m.IsConnected() {
		return fmt.Errorf("no MKS937B is connected")
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	addressStr := fmt.Sprintf("%03d", m.Address)
	message := fmt.Sprintf("@%s%s!%s;FF", addressStr, command, parameter)
	m.Communication.Write([]byte(message))

	response, err := m.Communication.ReadUntil(";FF")
	if err != nil {
		return err
	}
	responseStr := string(response)
	regex := regexp.MustCompile(`@([0-9]+)(?:ACK|NAK)(.*?);FF`)
	matches := regex.FindStringSubmatch(responseStr)

	if len(matches) < 3 {
		return NewErrUnexpectedReply(message, responseStr)
	}
	if matches[1] != addressStr {
		return NewErrUnexpectedAddress(addressStr, matches[1])
	}
	if matches[2] != parameter {
		return NewErrUnexpectedParamater(parameter, matches[2])
	}
	return nil
}