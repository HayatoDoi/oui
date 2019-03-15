# oui
When you executed `go generate`, this library download the latest oui list from <http://standards-oui.ieee.org>.
## Usage

### downloaad library
```bash
$ go get github.com/HayatoDoi/oui
```

### example
#### code
`main.go`
```go:main.go
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
```
#### make oui.go
this take about 10-20 seconds.
```bash
$ go generate
```

#### exec
```bash
$ go run main.go oui.go
```

#### out put
```
$ go run example.go oui.go 
902e1c000000 : Intel Corporate
```

## Licence
These codes are licensed under MIT.

## Author
[HayatoDoi](https://github.com/HayatoDoi)