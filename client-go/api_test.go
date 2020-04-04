package client_go

import (
	"fmt"
	"os"
	"testing"
)

var (
	list map[string]*Contact
	err  error
)

// All testcases do subscribe and then all testcases do unsubscribe. The uuid of the created memberlist elements
// has to be overwritten by the uuid from the testcase.
func TestRegister(t *testing.T) {

	serviceUrl := os.Getenv("GCP_SERVICE_URL")
	if serviceUrl == "" {
		fmt.Printf("GCP_SERVICE_URL environment variable unset or missing")
		os.Exit(1)
	}

	var app *App
	app, err = Create()

	fmt.Printf("app: %v\n", app)
	app.Register()
	fmt.Printf("app: %v\n", app)

	//// Reset the Firestore collection
	//err = Reset(serviceUrl)
	//if err != nil {
	//	fmt.Printf("could not reset Firestore collection")
	//	os.Exit(1)
	//}
	//
	//// TestCases
	//var testcases = map[string]struct {
	//	member IpAddress
	//	uuid   string
	//}{
	//	"empty": {
	//		member: IpAddress{},
	//		uuid:   uuid.NewV4().String(),
	//	},
	//	"only name": {
	//		member: IpAddress{
	//			Name: "test",
	//		},
	//		uuid: uuid.NewV4().String(),
	//	},
	//	"only ip": {
	//		member: IpAddress{
	//			Ip: "test",
	//		},
	//		uuid: uuid.NewV4().String(),
	//	},
	//	"only port": {
	//		member: IpAddress{
	//			Port: "test",
	//		},
	//		uuid: uuid.NewV4().String(),
	//	},
	//	"only protocol": {
	//		member: IpAddress{
	//			Protocol: "test",
	//		},
	//		uuid: uuid.NewV4().String(),
	//	},
	//	"complete": {
	//		member: IpAddress{
	//			Name:     "test",
	//			Ip:       "127.0.0.1",
	//			Port:     "12345",
	//			Protocol: "tcp",
	//		},
	//		uuid: uuid.NewV4().String(),
	//	},
	//	"same member": {
	//		member: IpAddress{
	//			Name:     "test",
	//			Ip:       "127.0.0.1",
	//			Port:     "12345",
	//			Protocol: "tcp",
	//		},
	//		uuid: uuid.NewV4().String(),
	//	},
	//	"again same member": {
	//		member: IpAddress{
	//			Name:     "test",
	//			Ip:       "127.0.0.1",
	//			Port:     "12345",
	//			Protocol: "tcp",
	//		},
	//		uuid: uuid.NewV4().String(),
	//	},
	//}

	//var registeredMember map[string]*IpAddress = make(map[string]*IpAddress)
	//
	//numTc := 0
	//for name, tc := range testcases {
	//
	//	memberlist, err := Create(&tc.member)
	//	if err != nil {
	//		t.Errorf("unexpected error: %v", err)
	//	}
	//	memberlist.Uuid = tc.uuid
	//
	//	list, err = memberlist.Subscribe()
	//	if err != nil {
	//		t.Errorf("unexpected error: %v", err)
	//	}
	//
	//	registeredMember[memberlist.Uuid] = memberlist.Self
	//	numTc++
	//
	//	if len(list) != numTc {
	//		t.Errorf("unexpected number (%d) of members in result: %v\n", numTc, len(list))
	//	}
	//
	//	t.Run(name, func(t *testing.T) {
	//
	//		if (*list[memberlist.Uuid]) != tc.member {
	//			t.Errorf("unexpected result map[%q]: %v\n", memberlist.Uuid, list)
	//		}
	//	})
	//}
	//
	//// Unsubcribe one by one
	//i := 0
	//for name, tc := range testcases {
	//
	//	//fmt.Printf("registeredMember[%d]: %v\n", i, registeredMember[i])
	//
	//	memberlist, err := Create(&tc.member)
	//	if err != nil {
	//		t.Errorf("unexpected error: %v", err)
	//	}
	//	memberlist.Uuid = tc.uuid
	//
	//	err = memberlist.Unsubscribe()
	//	if err != nil {
	//		t.Errorf("unexpected error: %v", err)
	//	}
	//	numTc--
	//
	//	// List remaining members and check count
	//	t.Run(name, func(t *testing.T) {
	//		list, err := memberlist.List()
	//		if err != nil {
	//			t.Errorf("unexpected error: %v", err)
	//		}
	//		if len(list) != numTc {
	//			t.Errorf("unexpected number (%d) of members in result: %v\n", numTc, len(list))
	//		}
	//	})
	//	i++
	//}
}
