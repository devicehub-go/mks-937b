/*
Author: Leonardo Rossi Leao
Created at: September 24rd, 2025
Last update: September 24rd, 2025
*/

package protocol

import (
	"fmt"
	"slices"
	"strconv"
)

/*
Gets the controller address (1 to 254)
*/
func (m *MKS937B) GetAddress() (int, error) {
	response, err := m.Query("AD")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(response)
}

/*
Sets the controller address
*/
func (m *MKS937B) SetAddress(address int) error {
	if address < 1 || 254 < address {
		return NewErrInvalidAddress(address)
	}
	return m.Set("AD", fmt.Sprintf("%03d", address))
}

/*
Gets the controller baud rate
*/
func (m *MKS937B) GetBaudRate() (int, error) {
	response, err := m.Query("BR")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(response)
}

/*
Sets the controller baud rate (valid values include 9600, 19200,
38400, 57600, 115200)
*/
func (m *MKS937B) SetBaudRate(baudrate int) error {
	valid := []int{9600, 19200, 38400, 57600, 115200}
	if !slices.Contains(valid, baudrate) {
		return NewErrInvalidBaudRate(baudrate)
	}
	return m.Set("BR", fmt.Sprint(baudrate))
}

/*
Gets the controller parity
*/
func (m *MKS937B) SetParity(parity string) error {
	valid := []string{"NONE", "EVEN", "ODD"}
	if !slices.Contains(valid, parity) {
		return NewErrInvalidParity(parity)
	}
	return m.Set("PAR", parity)
}

/*
Gets delay time of RS485 communication in milliseconds
*/
func (m *MKS937B) GetDelayTime() (int, error) {
	response, err := m.Query("DLY")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(response)
}

/*
Sets the delay time of RS485 communication in milliseconds.
For a reliable communication the time must be greater than 1 ms.
Default is 8 ms.
*/
func (m *MKS937B) SetDelayTime(delay int) error {
	return m.Set("DLY", fmt.Sprint(delay))
}

/*
Gets the pressure unit
*/
func (m *MKS937B) GetPressureUnit() (string, error) {
	return m.Query("U")
}

/*
Sets the pressure unit (Torr, MBAR, PASCAL, Micron)
*/
func (m *MKS937B) SetPressureUnit(unit string) error {
	valid := []string{"Torr", "MBAR", "PASCAL", "Micron"}
	if !slices.Contains(valid, unit) {
		return NewErrInvalidUnit(unit)
	}
	return m.Set("U", unit)
}