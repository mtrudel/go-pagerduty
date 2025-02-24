package pagerduty

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestAutomationActionsSidecarRunnerGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/automation_actions/runners/01DA2MLYN0J5EFC1LKWXUKDDKT", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{ "runner": { "id": "01DA2MLYN0J5EFC1LKWXUKDDKT", "name": "us-west-2 prod sidecar runner", "summary": "us-west-2 prod sidecar runner", "type": "runner", "description": "us-west-2 prod sidecar runner provisioned by SRE", "creation_time": "2022-10-21T19:42:52.127369Z", "runner_type": "sidecar", "status": "Configured", "teams": [ { "id": "PQ9K7I8", "type": "team_reference" } ], "privileges": { "permissions": [ "read" ] } } }`))
	})

	resp, _, err := client.AutomationActionsRunner.Get("01DA2MLYN0J5EFC1LKWXUKDDKT")
	if err != nil {
		t.Fatal(err)
	}

	permissions_read := "read"
	description := "us-west-2 prod sidecar runner provisioned by SRE"

	want := &AutomationActionsRunner{
		ID:             "01DA2MLYN0J5EFC1LKWXUKDDKT",
		Name:           "us-west-2 prod sidecar runner",
		Summary:        "us-west-2 prod sidecar runner",
		Description:    &description,
		CreationTime:   "2022-10-21T19:42:52.127369Z",
		LastSeenTime:   nil,
		RunnerType:     "sidecar",
		Type:           "runner",
		RunbookBaseUri: nil,
		RunbookApiKey:  nil,
		Teams: []*TeamReference{
			{
				Type: "team_reference",
				ID:   "PQ9K7I8",
			},
		},
		Privileges: &AutomationActionsPrivileges{
			Permissions: []*string{&permissions_read},
		},
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsRunbookRunnerGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/automation_actions/runners/01DA2MLYN0J5EFC1LKWXUKDDKT", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write([]byte(`{ "runner": { "id": "01DA2MLYN0J5EFC1LKWXUKDDKT", "name": "us-west-2 prod sidecar runner", "summary": "us-west-2 prod sidecar runner", "type": "runner", "description": "us-west-2 prod sidecar runner provisioned by SRE", "creation_time": "2022-10-21T19:42:52.127369Z", "last_seen": "2022-10-21T19:42:53.123456Z", "runner_type": "runbook", "status": "Configured", "runbook_base_uri": "acme.prod" } }`))
	})

	resp, _, err := client.AutomationActionsRunner.Get("01DA2MLYN0J5EFC1LKWXUKDDKT")
	if err != nil {
		t.Fatal(err)
	}

	last_seen := "2022-10-21T19:42:53.123456Z"
	runbook_base_uri := "acme.prod"
	description := "us-west-2 prod sidecar runner provisioned by SRE"
	want := &AutomationActionsRunner{
		ID:             "01DA2MLYN0J5EFC1LKWXUKDDKT",
		Name:           "us-west-2 prod sidecar runner",
		Summary:        "us-west-2 prod sidecar runner",
		Description:    &description,
		CreationTime:   "2022-10-21T19:42:52.127369Z",
		LastSeenTime:   &last_seen,
		RunnerType:     "runbook",
		Type:           "runner",
		RunbookBaseUri: &runbook_base_uri,
		RunbookApiKey:  nil,
		Teams:          nil,
		Privileges:     nil,
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsRunnerCreate(t *testing.T) {
	setup()
	defer teardown()
	description := "us-west-2 prod sidecar runner provisioned by SRE"
	input := &AutomationActionsRunner{
		Name:        "us-west-2 prod sidecar runner",
		Description: &description,
		RunnerType:  "sidecar",
	}

	mux.HandleFunc("/automation_actions/runners", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(AutomationActionsRunnerPayload)
		json.NewDecoder(r.Body).Decode(v)
		if !reflect.DeepEqual(v.Runner, input) {
			t.Errorf("Request body = %+v, want %+v", v.Runner, input)
		}
		w.Write([]byte(`{ "runner": { "id": "01DA2MLYN0J5EFC1LKWXUKDDKT", "name": "us-west-2 prod sidecar runner", "type": "runner", "description": "us-west-2 prod sidecar runner provisioned by SRE", "creation_time": "2022-10-21T19:42:52.127369Z", "runner_type": "sidecar", "status": "Configured" } }`))
	})

	resp, _, err := client.AutomationActionsRunner.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	want := &AutomationActionsRunner{
		ID:           "01DA2MLYN0J5EFC1LKWXUKDDKT",
		Name:         "us-west-2 prod sidecar runner",
		Description:  &description,
		CreationTime: "2022-10-21T19:42:52.127369Z",
		RunnerType:   "sidecar",
		Type:         "runner",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp, want)
	}
}

func TestAutomationActionsRunnerDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/automation_actions/runners/01DA2MLYN0J5EFC1LKWXUKDDKT", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	if _, err := client.AutomationActionsRunner.Delete("01DA2MLYN0J5EFC1LKWXUKDDKT"); err != nil {
		t.Fatal(err)
	}
}
