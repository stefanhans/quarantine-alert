package client_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// App is the core struct for all relevant data
type App struct {
	ServiceUrl string
	Registration
}

func (app *App) String() string {

	// Marshal registration
	out, err := json.Marshal(app)
	if err != nil {
		return fmt.Sprintf("failed to marshal app: %s", err)
	}

	return string(out)
}

// Create returns the registered app data
func Create() (*App, error) {

	serviceUrl := os.Getenv("GCP_SERVICE_URL")
	if serviceUrl == "" {
		return nil, fmt.Errorf("GCP_SERVICE_URL environment variable unset or missing")
	}

	return &App{
		ServiceUrl: serviceUrl,
	}, nil
}

// Register stores a new member and returns the current memberlist
func (app *App) Register() error {

	// Send request to service
	res, err := http.Post(app.ServiceUrl+"/register",
		"application/json",
		strings.NewReader(fmt.Sprintf("{}")))

	fmt.Printf("failed to register app (%v): %v\n", app, err)
	if err != nil {
		return fmt.Errorf("failed to register app (%v): %v\n", app, err)
	}

	// Read response body in JSON
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to read response of registration (%q): %v\n", app.ServiceUrl, err)
	}
	fmt.Printf("Body: %s\n", body)

	// Unmarshall the registration
	err = json.Unmarshal(body, &app.Registration)
	if err != nil {
		return fmt.Errorf("failed to unmarshall response of memberlist subscription (%v): %v\n", res.Proto, err)
	}

	return nil
}
