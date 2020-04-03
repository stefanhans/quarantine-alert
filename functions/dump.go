package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// Dump requests all contacts from Firestore
func Dump(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	//fmt.Printf("Body: %s\n", body)

	// Unmarshal request body
	bytes := []byte(string(body))
	var format Format
	err = json.Unmarshal(bytes, &format)
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

	// Get all documents
	colRef := client.Collection("contacts")
	if colRef == nil {
		http.Error(w, fmt.Sprintf("cannot find collection %q", "contacts"), http.StatusInternalServerError)
		return
	}
	docRefs := colRef.Documents(ctx)

	// Iterate over documents and fill dump
	var contactsCollection = make(ContactsCollection)
	for {
		doc, err := docRefs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to iterate"), http.StatusInternalServerError)
			return
		}
		contactsCollection[doc.Ref.ID] = doc.Data()
	}

	// Marshal contacts according to format request
	var contactsJson []byte
	if format.Indent {
		contactsJson, err = json.MarshalIndent(contactsCollection, "", "  ")
	} else {
		contactsJson, err = json.Marshal(contactsCollection)
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal dump: %s", err), http.StatusInternalServerError)
		return
	}
	//fmt.Println(string(contactsJson))

	// Response contacts
	_, err = fmt.Fprint(w, string(contactsJson))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to write response: %s", err), http.StatusInternalServerError)
	}
	return
}

/*

indent: JSON data formatted with indent [default: false]

curl localhost:8080/dump -d '{}'
curl localhost:8080/dump -d '{"indent":true}'
curl localhost:8080/dump -d '{"indent":false}'

*/
