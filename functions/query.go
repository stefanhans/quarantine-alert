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

// Query requests a new report from Firestore related to the requester ID.
func Query(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	//fmt.Printf("Body: %s\n", body)

	// Unmarshal request body
	bytes := []byte(string(body))
	var report Report
	err = json.Unmarshal(bytes, &report)
	if err != nil {
		fmt.Printf("Cannot unmarshall JSON input: %v\n", err)
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

	// Get document of requester
	docRef := client.Doc(fmt.Sprintf("contacts/%v", report.RequesterID))
	if docRef == nil {
		http.Error(w, fmt.Sprintf("failed to get reference of document %q: %s", report.RequesterID, err), http.StatusInternalServerError)
		return
	}

	// Get data of document
	var docsnap *firestore.DocumentSnapshot
	docsnap, err = docRef.Get(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get snapshot of document %q: %s", report.RequesterID, err), http.StatusInternalServerError)
		return
	}
	data := docsnap.Data()

	// Assert array of "nearby" contacts
	contactsArray, ok := data["nearby"].([]interface{})
	if !ok {
		http.Error(w, fmt.Sprintf("failed to assert %q: %s", "nearby", err), http.StatusInternalServerError)
		return
	}

	// Prepare report by searching in "nearby" for relevant "contagious" contacts, i.e.
	// - the day of contact was not too long before the day the contagion was noticed
	// - the day of contact was not too long after the day the contagion was noticed
	//
	// Loop over contacts
	var reportSlice []string
	for _, element := range contactsArray {

		// Assert contact
		contact, ok := element.(map[string]interface{})
		if !ok {
			http.Error(w, fmt.Sprintf("failed to assert %q contact map: %s", "nearby", err), http.StatusInternalServerError)
			return
		}

		// Get contact document
		docRef := client.Doc(fmt.Sprintf("contacts/%v", contact["ContactID"]))
		if docRef == nil {
			http.Error(w, fmt.Sprintf("cannot find contact document %q: %s", contact["ContactID"], err), http.StatusInternalServerError)
			return
		}

		// Get snapshot of contact document
		docSnap, err := docRef.Get(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get snapshot of contact document %q: %s", contact["ContactID"], err), http.StatusInternalServerError)
			return
		}

		// Get data of contact snapshot
		data := docSnap.Data()

		// Assert contact "contagious" status
		contagious, ok := data["contagious"].(bool)
		if !ok {
			http.Error(w, fmt.Sprintf("failed to assert %q status of contact: %s", "contagious", err), http.StatusInternalServerError)
			return
		}

		// Continue, if contact is not contagious
		if !contagious {
			continue
		}

		// Assert the time of contagion noticed
		timeContagionUpdated, ok := data["time-contagion-updated"].(time.Time)
		if !ok {
			http.Error(w, fmt.Sprintf("failed to assert %q: %s", "time-contagion-updated", err), http.StatusInternalServerError)
			return
		}

		// Estimate the first and last day of contagion calculated from today, i.e. you should query at least once a day
		earliestDayOfContagion := timeContagionUpdated.Add(time.Hour * -1 * time.Duration(daysOfBeforeContagionPeriod) * 24)
		lastDayOfContagion := timeContagionUpdated.Add(time.Hour * time.Duration(daysOfAfterContagionPeriod) * 24)

		// Assert the time of contact
		contactTime, ok := contact["ContactTime"].(time.Time)
		if !ok {
			http.Error(w, fmt.Sprintf("failed to assert %q: %s", "contact-time", err), http.StatusInternalServerError)
			return
		}

		// Continue, if time of contact was before earliest day of contagion
		if contactTime.Unix() < earliestDayOfContagion.Unix() {
			continue
		}

		// Continue, if time of contact was after latest day of contagion
		if lastDayOfContagion.Unix() < contactTime.Unix() {
			continue

		}

		// Append time of contagion noticed to the response
		reportSlice = append(reportSlice, fmt.Sprintf("%v", data["time-contagion-updated"]))

	}
	report.Report = reportSlice

	// Marshal report
	reportJson, err := json.Marshal(report)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(reportJson))

	// Response report
	_, err = fmt.Fprint(w, string(reportJson))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
	return
}

/*

requester: 				unique ID of requester
contagious: 			set to false as default [optional]
time-contagion-updated: set to current timestamp as default [optional]

curl localhost:8080/query -d '{"requester":"uBAJbYDJTBHOVqrceZur"}'

*/
