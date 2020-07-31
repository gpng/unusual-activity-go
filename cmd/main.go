package main

import (
	"flag"

	d "github.com/gpng/unusual-volume-go/pkg/download"
)

func main() {
	monthsPtr := flag.Int("months", 5, "Number of months of historical volume to compare with (default 5)")
	cutoffPtr := flag.Int("cutoff", 8, "Standard deviation from mean before logging as unusual activity (default 8)")
	daysPtr := flag.Int("lastdays", 3, "Last number of trading days to check (default 3)")

	d.CheckAnomalies(*monthsPtr, *cutoffPtr, *daysPtr)
}
