Possible error codes found in source code:

````
            r = (0, a.defineMessages)({
                internalFault: { id: "gen3wc_alert_name_internal_fault", defaultMessage: "Internal fault, charging disabled" },
                groundFault: { id: "gen3wc_alert_name_ground_fault", defaultMessage: "Ground fault due to unsafe current path, charging disabled" },
                highGroundResistance: { id: "gen3wc_alert_name_high_ground_resistance", defaultMessage: "High ground resistance detected, charging disabled" },
                highTemp: { id: "gen3wc_alert_name_high_temp", defaultMessage: "High temperature detected; charging limited or disabled" },
                overVoltage: { id: "gen3wc_alert_name_over_voltage", defaultMessage: "Overvoltage or poor grid quality detected, charging disabled" },
                overCurrent: { id: "gen3wc_alert_name_over_current", defaultMessage: "Vehicle overcurrent detected" },
            }),
            o = (0, a.defineMessages)({
                internalFault: {
                    id: "gen3wc_alert_note_internal_fault",
                    defaultMessage: "Turn the circuit breaker off, wait 5 seconds, and turn it back on.{br}If solid red light remains, document part number and serial number, then contact Tesla.",
                },
                groundFault: {
                    id: "gen3wc_alert_note_ground_fault",
                    defaultMessage:
                        "Inspect the handle, cable, Wall Connector, and vehicle charge port for damage or signs of water ingress.{br}Have an electrician check that earth ground is not directly connected to a conductor wire in the branch circuit.",
                },
                highGroundResistance: {
                    id: "gen3wc_alert_note_high_ground_resistance",
                    defaultMessage:
                        "Verify that the Wall Connector is properly grounded. The Ground connection must be bonded in the upstream power supply for proper operation.{br}Check all physical connections, including the wirebox terminals, electrical panel(s), and junction boxes. In residential power supplies, check the bond between Ground and neutral at the main panel.{br}If connected to a step-down transformer, contact the transformer's manufacturer for direction on how to bond the ground connection.",
                },
                highTemp: {
                    id: "gen3wc_alert_note_high_temp",
                    defaultMessage:
                        "Check the faceplate and cable handle for excessive warmth.{br}Have an electrician remove the Wall Connector from the wirebox and verify that the conductors used are sized correctly and that the terminal block is torqued to specification.",
                },
                overVoltage: {
                    id: "gen3wc_alert_note_over_voltage",
                    defaultMessage:
                        "Verify that the power supply is nominal 200-240 volts.{br}If the issue persists, have an electrician remove the Wall Connector from the wirebox and confirm that voltage readings are as expected at the terminal block using a multimeter.{br}Record the voltage readings for the following: L1 to L2/N, L1 to Ground, L2/N to Ground.",
                },
                overCurrent: {
                    id: "gen3wc_alert_note_over_current",
                    defaultMessage:
                        "Reduce the vehicle's charge current setting.{br}If the issue persists and the attached vehicle is manufactured by Tesla, record the vehicle's VIN and approximate time of the fault and contact Tesla.{br}If the vehicle is not manufactured by Tesla, contact the vehicle's manufacturer.",
                },
            });
        (t.alertIdToNameMessages = {
            [i.WCAlertID.CC_f001_gndMonIntrptLineSide]: r.highGroundResistance,
            [i.WCAlertID.CC_f003_CCIDTripped]: r.groundFault,
            [i.WCAlertID.CC_f004_CCIDSelfTestFault]: r.internalFault,
            [i.WCAlertID.CC_f006_inputOverCurrent]: r.overCurrent,
            [i.WCAlertID.CC_f007_inputOverVoltage]: r.overVoltage,
            [i.WCAlertID.CC_f010_contactorWelded]: r.internalFault,
            [i.WCAlertID.CC_f011_ambientOT]: r.highTemp,
            [i.WCAlertID.CC_f012_wallPlugOT]: r.highTemp,
            [i.WCAlertID.CC_f013_vehConnOT]: r.highTemp,
            [i.WCAlertID.CC_f014_mcuSelfTestFault]: r.internalFault,
            [i.WCAlertID.CC_f019_proxDisconnected]: r.internalFault,
            [i.WCAlertID.CC_f025_evseTemp]: r.highTemp,
            [i.WCAlertID.CC_f026_wallPlugTemp]: r.highTemp,
            [i.WCAlertID.CC_f027_vehicleHandleTemp]: r.highTemp,
            [i.WCAlertID.CC_f031_pllLockLost]: r.overVoltage,
            [i.WCAlertID.CC_f032_meteringFailure]: r.internalFault,
            [i.WCAlertID.CC_f033_vRefOutOfRange]: r.internalFault,
            [i.WCAlertID.CC_f036_contactorStuckOpen]: r.internalFault,
            [i.WCAlertID.CC_f037_internalModuleWatchdogExpired]: r.internalFault,
            [i.WCAlertID.CC_f039_internalModuleOverTemp]: r.highTemp,
        }),
            (t.alertIdToNoteMessages = {
                [i.WCAlertID.CC_f001_gndMonIntrptLineSide]: o.highGroundResistance,
                [i.WCAlertID.CC_f003_CCIDTripped]: o.groundFault,
                [i.WCAlertID.CC_f004_CCIDSelfTestFault]: o.internalFault,
                [i.WCAlertID.CC_f006_inputOverCurrent]: o.overCurrent,
                [i.WCAlertID.CC_f007_inputOverVoltage]: o.overVoltage,
                [i.WCAlertID.CC_f010_contactorWelded]: o.internalFault,
                [i.WCAlertID.CC_f011_ambientOT]: o.highTemp,
                [i.WCAlertID.CC_f012_wallPlugOT]: o.highTemp,
                [i.WCAlertID.CC_f013_vehConnOT]: o.highTemp,
                [i.WCAlertID.CC_f014_mcuSelfTestFault]: o.internalFault,
                [i.WCAlertID.CC_f019_proxDisconnected]: o.internalFault,
                [i.WCAlertID.CC_f025_evseTemp]: o.highTemp,
                [i.WCAlertID.CC_f026_wallPlugTemp]: o.highTemp,
                [i.WCAlertID.CC_f027_vehicleHandleTemp]: o.highTemp,
                [i.WCAlertID.CC_f031_pllLockLost]: o.overVoltage,
                [i.WCAlertID.CC_f032_meteringFailure]: o.internalFault,
                [i.WCAlertID.CC_f033_vRefOutOfRange]: o.internalFault,
                [i.WCAlertID.CC_f036_contactorStuckOpen]: o.internalFault,
                [i.WCAlertID.CC_f037_internalModuleWatchdogExpired]: o.internalFault,
                [i.WCAlertID.CC_f039_internalModuleOverTemp]: o.highTemp,
            });
    },
    function (e, t, n) {
```