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
	// Check app name is present
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments. Need at least app name")
		return
	}

	// Set downtime on app
	err := datadogDowntime(os.Args[1], true)
	if err != nil {
		fmt.Println("Unable to set downtime on app")
		return
	}

	// Launch ArgiCD application sync
	err = argoCDSyncApp(os.Args[1])
	if err != nil {
		fmt.Println("Unable to launch ArgoCD application sync")
	}

	// Remove downtime anyway
	err = datadogDowntime(os.Args[1], false)
	if err != nil {
		fmt.Println("Unable to remove downtime on app")
		fmt.Println("Warning, downtime is still active")
		// Fallback can be to use schedule end for downtime
		// like maxDowntimeSeconds := 3600
	}


	// Remove downtime on app
	datadogDowntime(os.Args[1], false)
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
				// if already in proper state
				// return nil
				// else set downtime on app
	}
	return nil
} 

func argoCDSyncApp(app string) (err) {
	// POST https://cd.apps.argoproj.io/api/v1/applications/{name}/sync
	// if return 200 {return nil}
	// else {return err}
}