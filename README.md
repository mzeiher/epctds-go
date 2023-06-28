# epctds-go EPC Data Tags Standard
This library helps with the handling of epc tags. 
It is possible to parse and create the hex representation you usually get from barcodes or RFID tags.

currently supported are sscc-96 and sgln-96 tags, but more will follow ;)

Installation
```bash
go get github.com/mzeiher/epctds-go 

```

Example:
```go
package main

import (
    "github.com/mzeiher/epctds-go"
)

func main() {
    tag, err := epctds.ParseEpcTagData("3118E511C46699F387000000")
    if err != nil {
        panic(err)
    }
    switch tag.(type) {
    case epctds.SGLN96:
        // do something
    case epctds.SSCC96:
        // do something
    }
}
```