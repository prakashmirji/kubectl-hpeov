package oneview

import (
	"os"
	"strconv"
	"sync"

	"github.com/HewlettPackard/oneview-golang/ov"
)

// ClientOV struct holds ov client
type ClientOV struct {
	ovClient *ov.OVClient
}

var clientov *ov.OVClient
var instance *ClientOV
var once sync.Once

// GetOVClient returns the ov client
func GetOVClient() *ClientOV {
	once.Do(
		func() {
			apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
			instance = &ClientOV{
				ovClient: clientov.NewOVClient(
					os.Getenv("ONEVIEW_OV_USER"),
					os.Getenv("ONEVIEW_OV_PASSWORD"),
					os.Getenv("ONEVIEW_OV_DOMAIN"),
					os.Getenv("ONEVIEW_OV_ENDPOINT"),
					false,
					apiversion,
					"*",
				),
			}
		},
	)
	return instance
}
