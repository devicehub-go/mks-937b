/*
Author: Leonardo Rossi Leao
Created at: September 24rd, 2025
Last update: September 24rd, 2025
*/

package protocol

import (
	"errors"
	"fmt"
)

var (
	ErrNotConnected = errors.New("device not connected")
	ErrInvalidParameter = errors.New("invalid parameter")
)

type ErrInvalidAddress struct {
	Got int
}
func NewErrInvalidAddress(got int) *ErrInvalidAddress {
	return &ErrInvalidAddress{Got: got}
}
func (e *ErrInvalidAddress) Error() string {
	return fmt.Sprintf(
		"address must be an integer value between 1 and 254, got %d", 
		e.Got,
	)
}

type ErrUnexpectedReply struct {
	Sent string
	Got string
}
func NewErrUnexpectedReply(sent string, got string) *ErrUnexpectedReply {
	return &ErrUnexpectedReply{
		Sent: sent,
		Got: got,
	}
}
func (e *ErrUnexpectedReply) Error() string {
	return fmt.Sprintf(
		"not expected response, sent %s got %s",
		e.Sent, e.Got,
	)
}

type ErrUnexpectedAddress struct {
	Expected string
	Got string
}
func NewErrUnexpectedAddress(expected string, got string) *ErrUnexpectedAddress {
	return &ErrUnexpectedAddress{
		Expected: expected,
		Got: got,
	}
}
func (e *ErrUnexpectedAddress) Error() string {
	return fmt.Sprintf(
		"invalid received address, expected %s got %s",
		e.Expected, e.Got,
	)
}

type ErrUnexpectedParameter struct {
	Expected string
	Got string
}
func NewErrUnexpectedParamater(expected string, got string) *ErrUnexpectedParameter {
	return &ErrUnexpectedParameter{
		Expected: expected,
		Got: got,
	}
}
func (e *ErrUnexpectedParameter) Error() string {
	return fmt.Sprintf(
		"invalid received parameter, expected %s got %s",
		e.Expected, e.Got,
	)
}

type ErrInvalidChannel struct {
	MinChannel int
	MaxChannel int
	Channel int
}
func NewErrInvalidChannel(min int, max int, channel int) *ErrInvalidChannel {
	return &ErrInvalidChannel{
		MinChannel: min,
		MaxChannel: max,
		Channel: channel,
	}
}
func (e *ErrInvalidChannel) Error() string {
	return fmt.Sprintf(
		"channel must be an integer value between %d and %d, got %d",
		e.MinChannel, e.MaxChannel, e.Channel,
	)
}

type ErrInvalidChannelControl struct { Channel int }
func NewErrInvalidChannelControl(channel int) *ErrInvalidChannelControl {
	return &ErrInvalidChannelControl{ Channel: channel }
}
func (e *ErrInvalidChannelControl) Error() string {
	return fmt.Sprintf(
		"channel must be an integer value among 1, 3 or 5, got %d",
		e.Channel,
	)
}

type ErrInvalidBaudRate struct { Got int }
func NewErrInvalidBaudRate(got int) *ErrInvalidBaudRate {
	return &ErrInvalidBaudRate{ Got: got }
}
func (e *ErrInvalidBaudRate) Error() string {
	return fmt.Sprintf(
		"baud rate must be 9600, 19200, 38400, 57600 or 115200, got %d", 
		e.Got,
	)
}

type ErrInvalidParity struct { Got string }
func NewErrInvalidParity(got string) *ErrInvalidParity {
	return &ErrInvalidParity{Got: got}
}
func (e *ErrInvalidParity) Error() string {
	return fmt.Sprintf(
		"parity must be NONE, EVEN or ODD, got %s", 
		e.Got,
	)
}

type ErrInvalidUnit struct { Got string }
func NewErrInvalidUnit(got string) *ErrInvalidUnit {
	return &ErrInvalidUnit{Got: got}
}
func (e *ErrInvalidUnit) Error() string {
	return fmt.Sprintf(
		"unit must be Torr, MBAR, PASCAL or Micron, got %s", 
		e.Got,
	)
}

/* Control commands errors */

type ErrInvalidPRO struct { Got float64 }
func NewErrInvalidPRO(got float64) *ErrInvalidPRO {
	return &ErrInvalidPRO{ Got: got }
}
func (e *ErrInvalidPRO) Error() string {
	return fmt.Sprintf(
		"The protection target must be 0 (disabled) or between 1e-5 and 1e-2, got %f",
		e.Got,
	)
}

type ErrInvalidRangeExp struct {
	MinValue, MaxValue, Got float64
}
func NewErrInvalidRangeExp(min, max, got float64) *ErrInvalidRangeExp {
	return &ErrInvalidRangeExp{
		MinValue: min,
		MaxValue: max,
		Got: got,
	}
}
func (e *ErrInvalidRangeExp) Error() string {
	return fmt.Sprintf(
		"The target value must be between %.2E and %.2E, got %.2E",
		e.MinValue, e.MaxValue, e.Got,
	)
}

type ErrInvalidCSE struct { Got string }
func NewErrInvalidCSE(got string) *ErrInvalidCSE {
	return &ErrInvalidCSE{Got: got}
}
func (e *ErrInvalidCSE) Error() string {
	return fmt.Sprintf(
		"The target channel must be A1, A2, B1, B2, C1, C2 or OFF, got %s",
		e.Got,
	)
}

type ErrInvalidControlMode struct { Got string }
func NewErrInvalidControlMode(got string) *ErrInvalidControlMode {
	return &ErrInvalidControlMode{Got: got}
}
func (e *ErrInvalidControlMode) Error() string {
	return fmt.Sprintf(
		"The target mode must be AUTO, SAFE or OFF, got %s",
		e.Got,
	)
}

type ErrInvalidFilament struct { Got int }
func NewErrInvalidFilament(got int) *ErrInvalidFilament {
	return &ErrInvalidFilament{Got: got}
}
func (e *ErrInvalidFilament) Error() string {
	return fmt.Sprintf(
		"The filament must be 1 or 2, got %d",
		e.Got,
	)
}

type ErrInvalidEmissionCurrent struct { Got string }
func NewErrInvalidEmissionCurrent(got string) *ErrInvalidEmissionCurrent {
	return &ErrInvalidEmissionCurrent{Got: got}
}
func (e *ErrInvalidEmissionCurrent) Error() string {
	return fmt.Sprintf(
		"The emission current must be 20UA, 100UA, AUTO20 or AUTO100, got %s",
		e.Got,
	)
}