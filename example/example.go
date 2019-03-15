package main

import (
	"fmt"
	"strings"
)

//go:generate oui -p main -o oui.go
func main(){
	mac := "902e1c000000"
	oui := strings.ToUpper(mac[:6])
	organization, ok := MacAndOrganization[oui]
	if ok != true{
		organization = "unknown"
	}
	fmt.Printf("%s : %s\n", mac, organization)
}
