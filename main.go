package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/scaleway/scaleway-cli/pkg/api"
)

func main() {
	orgToken := os.Getenv("SCALEWAY_ORG_TOKEN")
	token := os.Getenv("SCALEWAY_TOKEN")

	if strings.TrimSpace(orgToken) == "" || strings.TrimSpace(token) == "" {
		panic("required environmental variables are not set")
	}

	dynamicInventory := make(map[string][]string)

	getServers(dynamicInventory, token, orgToken)

	body, err := json.Marshal(dynamicInventory)
	if err != nil {
		panic("failed to marshal the dynamic inventory")
	}

	fmt.Println(string(body))
}

func getServers(dict map[string][]string, token, orgToken string) {
	disabledLoggerFunc := func(a *api.ScalewayAPI) {
		a.Logger = api.NewDisableLogger()
	}

	api, err := api.NewScalewayAPI(orgToken, token, "Scaleway Dynamic Inventory",
		"", disabledLoggerFunc)
	if err != nil {
		panic(fmt.Sprintf("failed to create API instance: %s", err))
	}

	servers, err := api.GetServers(true, 0)
	if err != nil {
		panic(fmt.Sprintf("failed to get servers: %s", err))
	}

	for _, server := range *servers {
		for _, tag := range server.Tags {
			if _, ok := dict[tag]; !ok {
				dict[tag] = make([]string, 0)
			}

			dict[tag] = append(dict[tag], server.PrivateAddress.IP)
		}
	}
}
