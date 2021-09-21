package db

import (
	"encoding/json"
	"testing"

	"reflect"

	"io/ioutil"
	"os"
)

// TestReadCredentials calls loadConfigFile, checking
// for a valid return value.

func TestReadCredentials(t *testing.T) {

	const fakefilen = "fakefile.json"

	FakeCredential := PsqlCredentials{
		Host:     "fakehost",
		Port:     "fakeport",
		User:     "fakeus",
		Password: "fakepass",
		Dbname:   "fakename"}

	bytes, err := json.Marshal(FakeCredential)
	if err != nil {
		t.Fatalf("Error on marshal struct")
	}

	tmpfile, err := ioutil.TempFile("", fakefilen)
	if err != nil {
		t.Fatalf("Error on creating file")
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(bytes); err != nil {
		t.Fatalf("error on writing file")
	}

	expCred := loadConfigFile(tmpfile.Name())

	if reflect.DeepEqual(FakeCredential, expCred) == false {
		t.Fatalf("Error on reading json config")
	}
}
