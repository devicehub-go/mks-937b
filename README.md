# MKS 937B

A Go library for communicating with MKS 937B Multi-Sensor Vacuum Gauge Controllers. This library provides a comprehensive interface for controlling and monitoring vacuum gauges including Hot Cathode (HC), Cold Cathode (CC), Pirani (PR), and Capacitance Manometer (CM) sensors.

## Features

- **Multi-Protocol Support**: Communicate via Serial (RS-232/RS-485) or TCP/IP using the unified Unicomm interface
- **Complete Device Control**: Full support for all MKS 937B commands and parameters
- **Sensor Management**: Monitor and control Hot Cathode, Cold Cathode, Pirani, and Capacitance Manometer sensors
- **Thread-Safe Operations**: Built-in mutex protection for concurrent access
- **Configuration Management**: Set device parameters like address, baud rate, pressure units

## Installation

```bash
go get github.com/devicehub-go/mks-937b
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/devicehub-go/mks-937b"
    "github.com/devicehub-go/unicomm"
    "github.com/devicehub-go/unicomm/protocol/unicommserial"
    "go.bug.st/serial"
)

func main() {
    // Create MKS 937B instance with Serial communication
    device := mks937b.New(1, unicomm.UnicommOptions{
        Protocol: unicomm.Serial,
        Serial: unicommserial.SerialOptions{
            PortName:     "/dev/ttyUSB0",
            BaudRate:     9600,
            Parity:       serial.NoParity,
            DataBits:     8,
            StopBits:     serial.OneStopBit,
            ReadTimeout:  time.Second,
            WriteTimeout: time.Second,
            EndDelimiter: "",
        },
    })

    // Connect to device
    if err := device.Connect(); err != nil {
        panic(err)
    }
    defer device.Disconnect()

    // Read pressure from channel 1
    pressure, err := device.GetPressure(1)
    if err != nil {
        fmt.Printf("Error reading pressure: %v\n", err)
        return
    }

    fmt.Printf("Channel 1 Pressure: %.2E, Status: %s\n", pressure.Value, pressure.Status)

    // Read all channel pressures
    pressures, err := device.GetPressures()
    if err != nil {
        fmt.Printf("Error reading pressures: %v\n", err)
        return
    }

    for i, p := range pressures {
        fmt.Printf("Channel %d: %.2E (%s)\n", i+1, p.Value, p.Status)
    }
}
```

### TCP Communication

```go
device := mks937b.New(1, unicomm.UnicommOptions{
    Protocol: unicomm.TCP,
    TCP: unicommtcp.TCPOptions{
        Host:         "192.168.1.100",
        Port:         23,
        ReadTimeout:  time.Second,
        WriteTimeout: time.Second,
        EndDelimiter: "",
    },
})
```

## API Reference

### Constructor

#### `New(address int, options unicomm.UnicommOptions) *protocol.MKS937B`
Creates a new MKS 937B instance with the specified device address and communication options.

**Parameters:**
- `address`: Device address (1-254)
- `options`: Communication configuration (Serial or TCP)

### Connection Management

#### `Connect() error`
Establishes connection with the device. Validates address range and initializes communication.

#### `Disconnect() error`
Closes the connection with the device.

#### `IsConnected() bool`
Returns true if the device is connected and responsive.

### Low-Level Communication

#### `Query(command string) (string, error)`
Sends a query command to the device and returns the response value.

#### `Set(command string, parameter string) error`
Sets a parameter on the device using the specified command.

### Pressure Reading

#### `GetPressure(channel int) (PressureReading, error)`
Reads pressure from a specific channel (1-6).

**Returns:** `PressureReading` struct with `Value` (float64) and `Status` (string)

#### `GetPressures() ([]PressureReading, error)`
Reads pressures from all 6 channels simultaneously.

**Returns:** Array of 6 `PressureReading` structs

#### `GetPressureCombination(channel int) (PressureReading, error)`
Reads combination sensor pressure for channel 1 or 2.

### Device Configuration

#### `GetAddress() (int, error)`
Returns the current device address.

#### `SetAddress(address int) error`
Sets device address (1-254).

#### `GetBaudRate() (int, error)`
Returns the current baud rate setting.

#### `SetBaudRate(baudrate int) error`
Sets baud rate. Valid values: 9600, 19200, 38400, 57600, 115200.

#### `SetParity(parity string) error`
Sets parity setting. Valid values: "NONE", "EVEN", "ODD".

#### `GetDelayTime() (int, error)`
Returns RS485 communication delay time in milliseconds.

#### `SetDelayTime(delay int) error`
Sets RS485 communication delay time. Minimum 1ms, default 8ms.

#### `GetPressureUnit() (string, error)`
Returns the current pressure unit setting.

#### `SetPressureUnit(unit string) error`
Sets pressure unit. Valid values: "Torr", "MBAR", "PASCAL", "Micron".

### Sensor Control (Channels 1, 3, 5)

#### `GetPowerStatus(channel int) (bool, error)`
Returns power status for PR, CP, HC, or high voltage status for CC.

#### `SetPowerStatus(channel int, status bool) error`
Controls power for PR, CP, HC, or high voltage for CC.

#### `GetSensorStatus(channel int) (string, error)`
Returns human-readable sensor status string.

#### `GetControlMode(channel int) (string, error)`
Returns current control mode.

#### `SetControlMode(channel int, mode string) error`
Sets control mode. Valid values: "AUTO", "SAFE", "OFF".

#### `GetControlChannelStatus(channel int) (string, error)`
Returns which control channel is assigned to the sensor.

#### `SetControlChannelStatus(channel int, target string) error`
Sets control channel assignment. Valid values: "A1", "A2", "B1", "B2", "C1", "C2", "OFF".

### Protection and Set Points

#### `GetProtectionTarget(channel int) (float64, error)`
Returns protection set point value.

#### `SetProtectionTarget(channel int, target float64) error`
Sets protection set point (1e-5 to 1e-2 Torr, or 0 to disable).

#### `GetTarget(channel int) (float64, error)`
Returns control set point value.

#### `SetTarget(channel int, target float64) error`
Sets control set point (5e-4 to 1e-2 Torr).

#### `GetHysterisesTarget(channel int) (float64, error)`
Returns hysteresis value.

#### `SetHysterisesTarget(channel int, target float64) error`
Sets hysteresis (1.2*CSP to 0.03 Torr).

#### `GetUpperControlStatus(channel int) (bool, error)`
Returns upper control set point status.

#### `SetUpperControlStatus(channel int, status bool) error`
Enables/disables upper control set point (extends range to 9.5e-1 Torr).

### Hot Cathode Control

#### `GetActiveFilament(channel int) (int, error)`
Returns active filament number (1 or 2).

#### `SetActiveFilament(channel int, filament int) error`
Sets active filament (1 or 2).

#### `GetEmissionCurrent(channel int) (string, error)`
Returns emission current setting.

#### `SetEmissionCurrent(channel int, current string) error`
Sets emission current. Valid values: "20UA", "100UA", "AUTO20", "AUTO100".

#### `GetHCGasCorrection(channel int) (float64, error)`
Returns Hot Cathode gas correction factor.

#### `SetHCGasCorrection(channel int, factor float64) error`
Sets Hot Cathode gas correction factor (0.1 to 50.0).

#### `GetGasSensitivy(channel int) (float64, error)`
Returns gas sensitivity value.

#### `SetGasSentivity(channel int, sensitivity float64) error`
Sets gas sensitivity (1.0 to 50.0).

#### `GetDegasStatus(channel int) (bool, error)`
Returns degas operation status.

#### `SetDegasStatus(channel int, status bool) error`
Starts/stops degas operation.

#### `GetDegasTime(channel int) (int, error)`
Returns degas time in seconds.

#### `SetDegasTime(channel int, time int) error`
Sets degas time (5 to 240 seconds).

#### `GetGasType(channel int) (string, error)`
Returns gas type setting.

#### `SetGasType(channel int, gas string) error`
Sets gas type. Valid values: "Nitrogen", "Argon", "Helium", "Custom".

### Cold Cathode Control

#### `GetCCGasCorrection(channel int) (float64, error)`
Returns Cold Cathode gas correction factor.

#### `SetUCGasCorrection(channel int, factor float64) error`
Sets Cold Cathode gas correction factor (0.1 to 10.0).

## Error Types

The library provides specific error types for detailed error handling:

- `ErrNotConnected`: Device not connected
- `ErrInvalidAddress`: Invalid device address (must be 1-254)
- `ErrInvalidChannelControl`: Invalid control channel (must be 1, 3, or 5)
- `ErrInvalidChannel`: Invalid channel number for specific operation
- `ErrInvalidBaudRate`: Invalid baud rate value
- `ErrInvalidParity`: Invalid parity setting
- `ErrInvalidUnit`: Invalid pressure unit
- `ErrInvalidRangeExp`: Value outside valid range
- `ErrInvalidPRO`: Invalid protection target value
- `ErrInvalidCSE`: Invalid control channel assignment
- `ErrInvalidControlMode`: Invalid control mode
- `ErrInvalidFilament`: Invalid filament number
- `ErrInvalidEmissionCurrent`: Invalid emission current setting
- `ErrInvalidGas`: Invalid gas type
- `ErrUnexpectedReply`: Unexpected device response
- `ErrUnexpectedAddress`: Wrong device address in response
- `ErrUnexpectedParameter`: Wrong parameter in response

## Thread Safety

All operations are thread-safe and protected by internal mutexes. Multiple goroutines can safely access the same device instance simultaneously.

```go
// Safe concurrent access
go func() {
    pressure, _ := device.GetPressure(1)
    fmt.Printf("Pressure: %.2E\n", pressure.Value)
}()

go func() {
    device.SetPowerStatus(1, true)
}()
```

## License

This project is authored by Leonardo Rossi Leao and was created on September 23rd, 2025.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.