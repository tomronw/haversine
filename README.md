Package to compare 2 sets of lats and lons, then see if they are in a certain radius (km) of each other.
float32 is used for the lat and lon values.

## Installation

```bash
go get github.com/tomronw/haversine
```


## Usage

```go
package main

import (
	"fmt"
	"github.com/tomronw/haversine"
)

func main() {
	// Create a new set of lats and lons
	lats1 := 53.4975
	lons1 := -2.2862
	lats2 := 53.4801
	lons2 := -2.1342

	// compare the 2 sets of lats and lons
	h := haversine.IsInRadius(lats1, lons1, lats2, lons2, 50.0)
	// h is a bool
	if h {
		fmt.Println("The 2 sets of lats and lons are in the radius")
	} else {
		fmt.Println("The 2 sets of lats and lons are not in the radius")
	}

}
```