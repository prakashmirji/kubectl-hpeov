package oneview

import (
	"github.com/HewlettPackard/oneview-golang/ov"
)

// GetAllServerTemplateDetails returns all server template details
func GetAllServerTemplateDetails() (ov.ServerProfileList, error) {
	ovc := GetOVClient().ovClient
	sort := "name:asc"
	sptList, err := ovc.GetProfileTemplates("", "", "", sort, "")
	if err != nil {
		return sptList, err
	}
	return sptList, nil
}

// GetServerTemplateByName returns server template by name
func GetServerTemplateByName(name string) (ov.ServerProfile, error) {
	ovc := GetOVClient().ovClient
	spt, err := ovc.GetProfileTemplateByName(name)
	if err != nil {
		return spt, err
	}
	return spt, nil
}
