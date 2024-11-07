package main

import (
	"fmt"
	"os"
)

type Downtime struct {
    Id int
    AppName string
	Active bool
}

func main() {
	fmt.Println(len(os.Args), os.Args)

	if len(os.Args) < 2 {
		fmt.Println("Too few arguments. Need at least app name")
		return
	}


}

func datadogDowntime(app string, active bool) (error) {
	// List all downtimes
	// GET https://api.datadoghq.com/api/v2/downtime
	// 	resp, r, err := api.ListDowntimes(ctx, *datadogV2.NewListDowntimesOptionalParameters())

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error when calling `DowntimesApi.ListDowntimes`: %v\n", err)
	// 	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	// 	return err
	// }

	// responseContent, _ := json.MarshalIndent(resp, "", "  ")



	responseContent := []Downtime{
			"Id": 1234,
			"AppName": toto,
			"Active": true,
		  }


	// Check if app if already down
	for _, e range Downtime {

	}

		// if already in proper state
		return nil

	// if not 
} 