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

// Contacted inserts a contact between reporters in Firestore.
func Contacted(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	//fmt.Printf("Body: %s\n", body)

	// Unmarshal request body
	bytes := []byte(string(body))
	var contact Contact
	err = json.Unmarshal(bytes, &contact)
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

	// Set "when" to current timestamp, if empty
	if contact.ContactTime.IsZero() {
		contact.ContactTime = time.Now()
	}

	// Round timestamp to day to avoid to many entries
	roundedDate := time.Date(contact.ContactTime.Year(),
		contact.ContactTime.Month(),
		contact.ContactTime.Day(),
		0, 0, 0, 0,
		time.UTC)

	// Get document of "reporter"
	docRef := client.Doc(fmt.Sprintf("contacts/%v", contact.ReporterID))
	if docRef == nil {
		http.Error(w, fmt.Sprintf("cannot find reporter %q: %s", contact.ReporterID, err), http.StatusInternalServerError)
		return
	}

	// Update "nearby" with "contact" and "when"
	var opposite Opposite
	opposite = Opposite{contact.ContactID, roundedDate}
	_, err = docRef.Update(ctx, []firestore.Update{
		{Path: "nearby", Value: firestore.ArrayUnion(opposite)},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update %q of %q: %s", "nearby", "reporter", err), http.StatusInternalServerError)
		return
	}

	// Get document of "contact"
	docRef = client.Doc(fmt.Sprintf("contacts/%v", contact.ContactID))
	if docRef == nil {
		http.Error(w, fmt.Sprintf("cannot find contact %q: %s", contact.ContactID, err), http.StatusInternalServerError)
		return
	}

	// Update "nearby" with "contact" and "when"
	opposite = Opposite{contact.ReporterID, roundedDate}
	_, err = docRef.Update(ctx, []firestore.Update{
		{Path: "nearby", Value: firestore.ArrayUnion(opposite)},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update %q of %q: %s", "nearby", "contact", err), http.StatusInternalServerError)
		return
	}

	// Marshal contact
	contactJson, err := json.Marshal(contact)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(contactJson))

	// Response contact
	_, err = fmt.Fprint(w, string(contactJson))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
	return
}

/*

reporter: 	unique ID of reporting app has to match Firestore's Document ID
contact: 	unique ID of contacted app
when: 		set to current timestamp as default [optional]

curl localhost:8080/contacted -d '{"reporter":"B3AQeB8acvRzK9FXtPxP","contact":"HlUv5MfjBhbfvEtNfATR"}'
curl localhost:8080/contacted -d '{"reporter":"B3AQeB8acvRzK9FXtPxP","contact":"TEnodwRdDGSrmVANiCW4"}'
curl localhost:8080/contacted -d '{"reporter":"B3AQeB8acvRzK9FXtPxP","contact":"zJD3j7lFzaTmw4AprfIl"}'

*/
