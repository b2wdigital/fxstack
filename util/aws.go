package util

import (
	"fmt"

	"github.com/b2wdigital/fxstack/config"
)

//GetURL gets URL based on default aws configs, resource name and service as the pattern:
//  https://\<service\>.\<region\>.amazonaws.com/\<account_number\>/\<resource_name\>
func GetAwsUrl(resourceName, service string) string {
	region := GetStringConfigOrPanic(config.AwsRegion)
	accountNumber := GetStringConfigOrPanic(config.AwsAccountNumber)
	return fmt.Sprintf("https://%s.%s.amazonaws.com/%s/%s", service, region, accountNumber, resourceName)
}
