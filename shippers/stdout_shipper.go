package shippers

import "fmt"
import "github.com/josegonzalez/metricsd/structs"
import "github.com/vaughan0/go-ini"

type StdoutShipper struct{}

func (shipper *StdoutShipper) Setup(_ ini.File) {
}

func (shipper *StdoutShipper) Ship(logs structs.MetricSlice) error {
	for _, item := range logs {
		serialized := item.ToJson()
		fmt.Printf("%s\n", string(serialized))
	}

	return nil
}
