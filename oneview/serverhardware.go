package oneview

import (
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
)

// GetAllServerHardwareDetails returns server hardware details
func GetAllServerHardwareDetails() (ov.ServerHardwareList, error) {
	ovc := GetOVClient().ovClient
	filters := []string{""}
	serverList, err := ovc.GetServerHardwareList(filters, "", "", "", "")
	if err != nil {
		return serverList, err
	}
	return serverList, nil
}

// GetServerHardwareByName returns server hardware by name
func GetServerHardwareByName(name string) (ov.ServerHardware, error) {
	ovc := GetOVClient().ovClient
	server, err := ovc.GetServerHardwareByName(name)
	if err != nil {
		return server, err
	}
	return server, nil
}

// UpdatePowerState toggle the power state of the given srever
func UpdatePowerState(name, newPowerState string) error {
	server, err := GetServerHardwareByName(name)
	if err != nil {
		return err
	}
	currentPowerState, err := server.GetPowerState()
	if err != nil {
		return err
	}
	if currentPowerState.String() == newPowerState {
		fmt.Printf("Server: %s is already in power state %s \n", name, newPowerState)
	} else if newPowerState == "On" {
		server.PowerOn()
	} else if newPowerState == "Off" {
		server.PowerOff()
	}

	return nil
}
