// MIT License
//
// Copyright (c) 2021-2022 Josef 'veloc1ty' Stautner
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Simple golang Interface for Gen 3 Tesla Wallboxes
package teslaWallbox

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Represents basic information
type Version struct {
	// Wallbox firmware version
	FirmwareVersion string `json:"firmware_version"`

	// Wallbox part number
	PartNumber string `json:"part_number"`

	// Wallbox serial number
	SerialNumber string `json:"serial_number"`
}

// Represents the current vitals
type Vitals struct {
	// Is the contactor to mains power closed
	ContactorClosed bool `json:"contactor_closed"`

	// Is a vehicle connected
	VehicleConnected bool `json:"vehicle_connected"`

	// Session duration in seconds
	SessionDuration uint `json:"session_s"`

	// Energy dispensed in Watthours
	SessionEnergy float64 `json:"session_energy_wh"`

	// Grid voltage
	GridVoltage float64 `json:"grid_v"`

	// Grid frequency in hertz
	GridFrequency float64 `json:"grid_hz"`

	// Maximum vehicle current in amp
	VehicleCurrent float64 `json:"vehicle_current_a"`

	// Phase A current in amp
	PhaseACurrent float64 `json:"currentA_a"`

	// Phase B current in amp
	PhaseBCurrent float64 `json:"currentB_a"`

	// Phase C current in amp
	PhaseCCurrent float64 `json:"currentC_a"`

	// Neutral current in amp
	NeutralCurrent float64 `json:"currentN_a"`

	// Phase A voltage
	PhaseAVoltage float64 `json:"voltageA_v"`

	// Phase B voltage
	PhaseBVoltage float64 `json:"voltageB_v"`

	// Phase C voltage
	PhaseCVoltage float64 `json:"voltageC_v"`

	// Relay coil voltage
	RelayCoilVoltage float64 `json:"relay_coil_v"`

	// PCB / Mainboard temperature in celsius
	PcbTemperature float64 `json:"pcba_temp_c"`

	// Handle temperature in celsius
	HandleTemperature float64 `json:"handle_temp_c"`

	// Microcontroller temperature in celsius
	McuTemperature float64 `json:"mcu_temp_c"`

	// Wallbox uptime in seconds
	Uptime uint `json:"uptime_s"`

	// Voltage of proximity pin
	ProximityVoltage float64 `json:"prox_v"`

	// High voltage of control pilot pin
	PilotHighVoltage float64 `json:"pilot_high_v"`

	// Low voltage of control pilot pin
	PilotLowVoltage float64 `json:"pilot_low_v"`

	// Config state
	ConfigStatus uint `json:"config_status"`

	// EVSE state describes the current working mode
	EvseState uint `json:"evse_state"`

	// Currently encountered alerts
	Alerts []interface{} `json:"current_alerts"`
}

// Represents lifetime statistics
type LifetimeStats struct {
	// How often did the contactor close to connect the charger to mains power
	ContactorCycles uint `json:"contactor_cycles"`

	// How often did the contactor close while power is actively drawn
	ContactorCyclesLoaded uint `json:"contactor_cycles_loaded"`

	// No idea. Maybe button on handle press count?
	ConnectorCycles uint `json:"connector_cycles"`

	// How often was charging interrupted due to exceeded temperature
	ThermalFoldbacks uint `json:"thermal_foldbacks"`

	// Average startup temperature in celsius
	AverageStartupTemperature float64 `json:"avg_startup_temp"`

	// Started charging sessions
	StartedChargingSessions uint `json:"charge_starts"`

	// Dispensed energy in watthours
	DispensedEnergy uint `json:"energy_wh"`

	// Total uptime in seconds
	TotalUptime uint `json:"uptime_s"`

	// Total charging time in seconds
	ChargingTime uint `json:"charging_time_s"`
}

// Fetch current vital statistics from the given IP address
func FetchVitals(ip string) (*Vitals, error) {
	vitals := &Vitals{}
	return vitals, fetchFromWallbox(ip, "vitals", vitals)
}

// Fetch current lifetime statistics from the given IP address
func FetchLifetimeStats(ip string) (*LifetimeStats, error) {
	stats := &LifetimeStats{}
	return stats, fetchFromWallbox(ip, "lifetime", stats)
}

// Fetch current version information from the given IP address
func FetchVersion(ip string) (*Version, error) {
	version := &Version{}
	return version, fetchFromWallbox(ip, "version", version)
}

func fetchFromWallbox(ip, apiEndpoint string, destination interface{}) error {
	response, responseError := http.Get(fmt.Sprintf("http://%s/api/1/%s", ip, apiEndpoint))

	if responseError != nil {
		return responseError
	} else {
		defer response.Body.Close()
		return json.NewDecoder(response.Body).Decode(destination)
	}
}
