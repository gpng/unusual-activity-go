package main

import (
	"flag"
	"os"

	d "github.com/gpng/unusual-volume-go/pkg/download"
)

func main() {
	monthsPtr := flag.Int("months", 5, "Number of months of historical volume to compare with")
	cutoffPtr := flag.Int("cutoff", 8, "Standard deviation from mean before logging as unusual activity")
	daysPtr := flag.Int("lastdays", 3, "Last number of trading days to check")
	helpPtr := flag.Bool("h", false, "Help Text")
	flag.Parse()

	if *helpPtr {
		flag.PrintDefaults()
		os.Exit(0)
	}

	d.CheckAnomalies(*monthsPtr, *cutoffPtr, *daysPtr)
}
