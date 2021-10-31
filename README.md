# TeslaWallbox

Simple Golang interface for Gen 3 Tesla wallboxes.

The following things are provided:

- Fetch version information from the wallbox like serial number and firmware version
- Fetch lifetime stats like dispensed energy, total uptime, how often the contactors were closed
- Fetch vitals like voltage, current, temperatures etc

Planned features:

- Translation of EVSE states to meaningfull values
- Translation of config states to meaningfull values

## Install

Just like any other go module:

```
go get -u github.com/vlcty/TeslaWallbox
```

## Sample code

``` go
package main

import (
	"fmt"
	"github.com/vlcty/TeslaWallbox"
)

func main() {
	wallboxIP := "10.10.0.39"

	version, versionError := teslaWallbox.FetchVersion(wallboxIP)
	stats, statsError := teslaWallbox.FetchLifetimeStats(wallboxIP)
	vitals, vitalsError := teslaWallbox.FetchVitals(wallboxIP)

	if versionError != nil {
		panic(versionError)
	} else {
		fmt.Printf("Version:\n%+v\n", version)
	}

	if statsError != nil {
		panic(statsError)
	} else {
		fmt.Printf("Stats:\n%+v\n", stats)
	}

	if vitalsError != nil {
		panic(vitalsError)
	} else {
		fmt.Printf("Vitals:\n%+v\n", vitals)
	}
}
```

## Tested with

Firmware versions:

- 21.29.1
- 21.36.5

Only tested with the 3 phase electricity grid in the EU. Other grid systems may not work. Tesla can revoke access to the internal API at any given time.
