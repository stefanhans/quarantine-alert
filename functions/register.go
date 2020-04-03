package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
)

// Register stores a new reporter in Firestore and returns the data including its ID.
func Register(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	//fmt.Printf("Body: %s\n", body)

	// Unmarshal request body
	bytes := []byte(string(body))
	var registration Registration
	err = json.Unmarshal(bytes, &registration)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot unmarshall JSON input: %s", err), http.StatusInternalServerError)
		return
	}

	// Get a Firestore client.
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create Firestore client: %s", err), http.StatusInternalServerError)
		return
	}

	// Close client when done.
	defer func() {
		_ = client.Close() // ignore error
	}()

	//fmt.Printf("registration.Since.Unix(): %v\n", registration.Since.Unix())

	// Set "time-contagion-updated" to current timestamp, if empty
	if registration.Since.IsZero() {
		registration.Since = time.Now()
	}

	// Register in new Firestore document
	docRef, _, err := client.Collection("contacts").Add(ctx, map[string]interface{}{
		"nearby":                 []Opposite{},
		"contagious":             registration.Contagious,
		"time-contagion-updated": registration.Since,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to register as Firestore document: %s", err), http.StatusInternalServerError)
		return
	}

	//fmt.Printf("docRef.ID: %s\n", docRef.ID)

	// Set reporter ID to the registration response
	registration.ReporterID = docRef.ID

	// Marshal registration
	registrationJson, err := json.Marshal(registration)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal dump: %s", err), http.StatusInternalServerError)
		return
	}
	//fmt.Println(string(registrationJson))

	// Response registration
	_, err = fmt.Fprint(w, string(registrationJson))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
	return
}

/*

reporter: 				unique ID will be created by Firestore [ignored]
contagious: 			set to false as default [optional]
time-contagion-updated: set to current timestamp as default [optional]

curl localhost:8080/register -d '{}'
curl localhost:8080/register -d '{"contagious":false}'
curl localhost:8080/register -d '{"contagious":true}'
curl localhost:8080/register -d '{"contagious":true,"time-contagion-updated":"2020-04-01T08:00:00+02:00"}'
curl localhost:8080/register -d '{"contagious":true,"time-contagion-updated":"2020-04-01T00:00:00Z"}'
curl localhost:8080/register -d '{"reporter":"to be ignored","contagious":false,"time-contagion-updated":"2020-04-01T08:46:36.649207+02:00"}'

*/
