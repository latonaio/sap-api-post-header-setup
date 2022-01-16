package sap_api_post_header_setup

import (
	"fmt"
	"testing"
)

type option struct {
}

func (o option) User() string {
	return "XXXXXXX"
}

func (o option) Pass() string {
	return "XXXXXXXX"
}

func (o option) RefreshTokenURL() string {
	return "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

func (o option) RetryMax() int {
	return 1
}

func (o option) RetryInterval() int {
	return 1000
}

func Test_a(t *testing.T) {
	c := NewSAPPostClientWithOption(option{})
	res, err := c.POST(
		"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		map[string]string{
			"sap-client": "XXX",
		},
		`
{
	"Product": "",
	"IndustrySector": "M",
	"ProductType": "FERT",
	"BaseUnit": "PC",
	"to_Description": {
		"results": [
			{
				"Product": "",
				"Language": "EN",
				"ProductDescription": "Test Material"
			}
		]
	}
}`,
	)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", res.Body)
}
