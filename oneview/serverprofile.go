package oneview

import (
	"github.com/HewlettPackard/oneview-golang/ov"
)

// GetAllServerProfileDetails returns server profile details
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
