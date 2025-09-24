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
Gets protection set point value for sensor on a
target channel that must be 1, 3 or 5
*/
func (m *MKS937B) GetProtectionTarget(channel int) (float64, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return 0, NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("PRO%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(response, 64)
}

/*
Sets a protection set point value for sensor on a
target channel that must be 1, 3 or 5.

The valid PRO range is 1e-5 to 1e-2 Torr. Use 0 for disable
and the default value is 5e-3 Torr
*/
func (m *MKS937B) SetProtectionTarget(channel int, target float64) error {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	if target != 0 && target < 1e-5 && 1e-2 < target {
		return NewErrInvalidPRO(target)
	}
	command := fmt.Sprintf("PRO%d", channel)
	return m.Set(command, fmt.Sprintf("%.2E", target))
}

/*
Gets the set point value for a sensor on a target channel
*/
func (m *MKS937B) GetTarget(channel int) (float64, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return 0, NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("CSP%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(response, 64)
}

/*
Sets a target for a sensor on a desired channel.

Valid CSP range is 5e-4 to 1e-2 Torr for Pirani, 
2e-3 to 1e-2 Torr for Convention Pirani, and 0.2% of 
full scale to 0.02 Torr for Capacitance Manometer
*/
func (m *MKS937B) SetTarget(channel int, target float64) error {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	if target < 5e-4 && 1e-2 < target {
		return NewErrInvalidRangeExp(5e-4, 1e-2, target)
	}
	command := fmt.Sprintf("CSP%d", channel)
	return m.Set(command, fmt.Sprintf("%.2E", target))
}

/*
Get upper control set point status
*/
func (m *MKS937B) GetUpperControlStatus(channel int) (bool, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return false, NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("XCS%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return false, err
	}
	return response == "ON", nil
}

/*
Sets the upper control set point. If enabled the
range is extended from 1e-2 Torr to 9.5e-1 Torr
*/
func (m *MKS937B) SetUpperControlStatus(channel int, status bool) error {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("XCS%d", channel)
	if status {
		return m.Set(command, "ON")
	}
	return m.Set(command, "OFF")
}

/*
Gets control set point hysterises value for a 
target channel
*/
func (m *MKS937B) GetHysterisesTarget(channel int) (float64, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return 0, NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("CHP%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(response, 64)
}

/*
Sets a target hysterises value for a target channel

Valid CHP range in 1.2*CSP to 1.1e-2 Torr for convention
pirani and pirani, and 1.2*CSP to 0.03 Torr for capacitance
manometer. Default value is 1.5*CSP
*/
func (m *MKS937B) SetHysterisesTarget(channel int, target float64) error {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	CSP, err := m.GetTarget(channel)
	if err != nil {
		return err
	}
	if target < 1.2*CSP || 0.03 < target {
		return NewErrInvalidRangeExp(1.2*CSP, 0.03, target)
	}
	command := fmt.Sprintf("CHP%d", channel)
	return m.Set(command, fmt.Sprintf("%.2E", target))
}

/*
Gets the control channel for a sensor on a desired channel
*/
func (m *MKS937B) GetControlChannelStatus(channel int) (string, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return "", NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("CSE%d", channel)
	return m.Query(command)
}

/*
Sets the control channel status for a sensor on
a desired channel.

Valid target options are A1, A2, B1, B2, C1, C2 or OFF
*/
func (m *MKS937B) SetControlChannelStatus(channel int, target string) error {
	validChannels := []int{1, 3, 5}
	validTargets := []string{"A1", "B1", "A2", "B2", "C1", "C2", "OFF"}

	if !slices.Contains(validChannels, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	if !slices.Contains(validTargets, target) {
		return NewErrInvalidCSE(target)
	}
	command := fmt.Sprintf("CSE%d", channel)
	return m.Set(command, target)
}

/*
Gets the control mode for a desired channel
*/
func (m *MKS937B) GetControlMode(channel int) (string, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return "", NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("CTL%d", channel)
	return m.Query(command)
}

/*
Sets the control mode for a desired channel

Valid mode are:
	- AUTO: HC/CC can be turned ON or OFF by controlling sensor
	- SAFE: Sensor can be turned OFF, but not be turned ON by controlling
	- OFF: disable control
*/
func (m *MKS937B) SetControlMode(channel int, mode string) error {
	validChannels := []int{1, 3, 5}
	validMode := []string{"AUTO", "SAFE", "OFF"}

	if !slices.Contains(validChannels, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	if !slices.Contains(validMode, mode) {
		return NewErrInvalidControlMode(mode)
	}

	command := fmt.Sprintf("CTL%d", channel)
	return m.Set(command, mode)
}

/*
Gets active filament for Hot Cathode
*/
func (m *MKS937B) GetActiveFilament(channel int) (int, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return 0, NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("AF%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(response)
}

/*
Sets active filament for Hot Cathode
*/
func (m *MKS937B) SetActiveFilament(channel int, filament int) error {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	if filament < 1 && 2 < filament {
		return NewErrInvalidFilament(filament)
	}
	
	command := fmt.Sprintf("AF%d", channel)
	return m.Set(command, fmt.Sprint(filament))
}

/*
Gets the emission current
*/
func (m *MKS937B) GetEmissionCurrent(channel int) (string, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return "", NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("EC%d", channel)
	return m.Query(command)
}

/*
Sets the emission current

Valid value for emission are 20UA, 100UA, AUTO20 and AUTO100
*/
func (m *MKS937B) SetEmissionCurrent(channel int, current string) error {
	validChannels := []int{1, 3, 5}
	validCurrent := []string{"20UA", "100UA", "AUTO20", "AUTO100"}

	if !slices.Contains(validChannels, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	if !slices.Contains(validCurrent, current) {
		return NewErrInvalidControlMode(current)
	}

	command := fmt.Sprintf("EC%d", channel)
	return m.Set(command, current)
}

/*
Gets the gas correction factor for an HC sensor on 
a desired channel
*/
func (m *MKS937B) GetGasCorrection(channel int) (float64, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return 0, NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("GC%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(response, 64)
}

/*
Sets the gas correction factor for an HC sensor on
a desired channel

Valid range for factor is from 0.1 to 50.0
*/
func (m *MKS937B) SetGasCorrection(channel int, factor float64) error {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	if factor < 0.1 || 50.0 < factor {
		return NewErrInvalidRangeExp(0.1, 50, factor)
	}
	command := fmt.Sprintf("GC%d", channel)
	return m.Set(command, fmt.Sprintf("%.1f", factor))
}

/*
Gets the channel power status for PR, CP, HC or high
voltage status for CC
*/
func (m *MKS937B) GetPowerStatus(channel int) (bool, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return false, NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("CP%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return false, err
	}
	return response == "ON", nil
}

/*
Sets the channel power status for PR, CP, HC or high
voltage status for CC
*/
func (m *MKS937B) SetPowerStatus(channel int, status bool) error {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("CP%d", channel)
	if status {
		return m.Set(command, "ON")
	}
	return m.Set(command, "OFF")
}

/*
Gets a gas sentivity for an Hot Cathode sensor on the desired channel
*/
func (m *MKS937B) GetGasSensitivy(channel int) (float64, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return 0, NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("SEN%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(response, 64)
}

/*
Sets a gas sensitivity for an Hot Cathode sensor on the desired channel

Valid range for sensivity is from 1.0 to 50.0
*/
func (m *MKS937B) SetGasSentivity(channel int, sensitivity float64) error {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	if sensitivity < 1.0 || 50.0 < sensitivity {
		return NewErrInvalidRangeExp(0.1, 50, sensitivity)
	}
	command := fmt.Sprintf("SEN%d", channel)
	return m.Set(command, fmt.Sprintf("%.1f", sensitivity))
}

/*
Gets Hot Cathode degas status
*/
func (m *MKS937B) GetDegasStatus(channel int) (bool, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return false, NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("DG%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return false, err
	}
	return response == "ON", nil
}

/*
Sets Hot Cathode degas status
*/
func (m *MKS937B) SetDegasStatus(channel int, status bool) error {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("DG%d", channel)
	if status {
		return m.Set(command, "ON")
	}
	return m.Set(command, "OFF")
}

/*
Get Hot Cathode degas time
*/
func (m *MKS937B) GetDegasTime(channel int) (int, error) {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return 0, NewErrInvalidChannelControl(channel)
	}
	command := fmt.Sprintf("DGT%d", channel)
	response, err := m.Query(command)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(response)
}

/*
Set Hot Cathode degas time
*/
func (m *MKS937B) SetDegasTime(channel int, time int) error {
	valid := []int{1, 3, 5}
	if !slices.Contains(valid, channel) {
		return NewErrInvalidChannelControl(channel)
	}
	if time < 5 && 240< time {
		return NewErrInvalidRangeExp(5, 240, float64(time))
	}
	
	command := fmt.Sprintf("DGT%d", channel)
	return m.Set(command, fmt.Sprint(time))
}