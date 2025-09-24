/*
Author: Leonardo Rossi Leao
Created at: September 24rd, 2025
Last update: September 24rd, 2025
*/

package protocol

import (
	"fmt"
	"strconv"
	"strings"
)

type PressureReading struct {
	value float64
	status string
}

var stringResponse = map[string]string{
	"LO<": "Pressure lower than minimum",
	"ATM": "PR when pressure is lower than 450 Torr",
	"OFF": "Cold cathode HV if OFF, or HC/PR/CP power if OFF",
	"WAIT": "CC or HC startup delay",
	"LowEmis": "HC OFF due to lowe emission",
	"CTRL_OFF": "CC or HC if OFF in controlled state",
	"PROT_OFF": "CC or HC if OFF in protected state",
	"MISCONN": "Sensor improperly connected, or broken filament (PR, CP only)",
	"NOGAUGE": "Controller unable to determine sensor connection",
	"NO_GAUGE": "Controller unable to determine sensor connection",
	"COMB_DISABLED": "Combination disabled",
}

/*
Parses a pressure reading from device
*/
func parsePressure(reading string) (PressureReading, error) {
	var pressure PressureReading

	for key, value := range stringResponse {
		if strings.Contains(reading, key) {
			pressure.status = value
			return pressure, nil
		}
	}
	value, err := strconv.ParseFloat(reading, 64)
	if err != nil {
		return pressure, err
	}

	pressure.value = value
	pressure.status = "OK"
	return pressure, nil
}

/*
Reads the pressure of a target channel
*/
func (m *MKS937B) GetPressure(channel int) (PressureReading, error) {
	var pressure PressureReading

	if 1 < channel || channel > 6 {
		return pressure, NewErrInvalidChannel(1, 6, channel)
	}
	command := fmt.Sprintf("PR%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return pressure, err
	}
	return parsePressure(response)
}

/*
Reads the pressures from all device channels
*/
func (m *MKS937B) GetPressures() ([]PressureReading, error) {
	response, err := m.Query("PRZ")
	if err != nil {
		return nil, err
	}

	pressures := make([]PressureReading, 6)
	for idx, value := range strings.Split(response, " ") {
		pressure, err := parsePressure(value)
		if err != nil {
			return nil, err
		}
		pressures[idx] = pressure
	}

	return pressures, nil
}

/*
Reads pressure on target channel (1 or 2) and its combination
sensor
*/
func (m *MKS937B) GetPressureCombination(channel int) (PressureReading, error) {
	var pressure PressureReading

	if channel < 1 || 2 < channel {
		return pressure, NewErrInvalidChannel(1, 2, channel)
	}
	command := fmt.Sprintf("PC%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return pressure, err
	}
	return parsePressure(response)
}