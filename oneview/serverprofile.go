package oneview

import (
	"github.com/HewlettPackard/oneview-golang/ov"
)

// GetAllServerProfileDetails returns all server profile details
func GetAllServerProfileDetails() (ov.ServerProfileList, error) {
	ovc := GetOVClient().ovClient
	sort := ""
	spList, err := ovc.GetProfiles("", "", "", sort, "")
	if err != nil {
		return spList, err
	}
	return spList, nil
}

// GetServerProfileByName returns server profile by name
func GetServerProfileByName(name string) (ov.ServerProfile, error) {
	ovc := GetOVClient().ovClient
	sp, err := ovc.GetProfileByName(name)
	if err != nil {
		return sp, err
	}
	return sp, nil
}

// CreateServerProfile creates server profile using template
func CreateServerProfile(sptName, spName string) error {

	var svrNameMatched ov.ServerHardware
	ovc := GetOVClient().ovClient
	serverName := ov.ServerHardwareType{}
	spt, err := ovc.GetProfileTemplateByName(sptName)

	if err != nil {
		return err
	}

	serverList, err := GetAllServerHardwareDetails()
	if err != nil {
		return err
	}

	hwName, err := ovc.GetServerHardwareTypeByUri(spt.ServerHardwareTypeURI)
	if err != nil {
		return err
	}

	for idx := range serverList.Members {
		serverName, _ = ovc.GetServerHardwareTypeByUri(serverList.Members[idx].ServerHardwareTypeURI)
		if serverName.Name == hwName.Name {
			svrNameMatched = serverList.Members[idx]
		}
	}

	err = ovc.CreateProfileFromTemplate(spName, spt, svrNameMatched)
	if err != nil {
		return err
	}
	return nil
}

// DeleteServerProfile delete the server profile
func DeleteServerProfile(spName string) error {

	ovc := GetOVClient().ovClient
	err := ovc.DeleteProfile(spName)

	if err != nil {
		return err
	}

	return nil
}
