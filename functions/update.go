package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
)

// Update edits contagious status of a reporter in Firestore.
func Update(w http.ResponseWriter, r *http.Request) {

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
		http.Error(w, fmt.Sprintf("cannot unmarshal JSON input: %s", err), http.StatusInternalServerError)
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

	// Get document to be updated
	docRef := client.Doc(fmt.Sprintf("contacts/%v", registration.ReporterID))
	if docRef == nil {
		http.Error(w, fmt.Sprintf("failed to get Firestore document %q: %s", registration.ReporterID, err), http.StatusInternalServerError)
		return
	}

	// Update "contagious" as bool
	_, err = docRef.Update(ctx, []firestore.Update{
		{Path: "contagious", Value: registration.Contagious},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update %q: %s", "contagious", err), http.StatusInternalServerError)
		return
	}

	// Set "time-contagion-updated" to current timestamp, if empty
	if registration.timeContagionUpdated.IsZero() {
		registration.timeContagionUpdated = time.Now()
	}

	// Update "time-contagion-updated" as timestamp
	_, err = docRef.Update(ctx, []firestore.Update{
		{Path: "time-contagion-updated", Value: registration.timeContagionUpdated},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update %q: %s", "time-contagion-updated", err), http.StatusInternalServerError)
		return
	}

	// Marshall updated data
	registrationJson, err := json.Marshal(registration)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(registrationJson))

	// Response updated data
	_, err = fmt.Fprint(w, string(registrationJson))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
	return
}

/*

reporter: 				unique ID has to match Firestore's Document ID
contagious: 			to be updated
time-contagion-updated: set to current timestamp as default [optional]

curl localhost:8080/update -d '{"reporter":"B3AQeB8acvRzK9FXtPxP","contagious":true}'
curl localhost:8080/update -d '{"reporter":"55fr4CFD1a1IvjIQReAd","contagious":true,"time-contagion-updated":"2020-04-01T08:46:36.649207+02:00"}'
curl localhost:8080/update -d '{"reporter":"55fr4CFD1a1IvjIQReAd","contagious":true,"time-contagion-updated":"2020-04-01T08:00:00+02:00"}'
curl localhost:8080/update -d '{"reporter":"55fr4CFD1a1IvjIQReAd","contagious":true,"time-contagion-updated":"2020-04-01T00:00:00Z"}'

*/
