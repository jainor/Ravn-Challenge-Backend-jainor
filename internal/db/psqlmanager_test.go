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

	const hostenv = "fake_host_var"
	const hostval = "fake_host"

	const portenv = "fake_port_var"
	const portval = "fake_port"

	const userenv = "fake_user_var"
	const userval = "fake_user"

	const passenv = "fake_pass_var"
	const passval = "fake_pass"

	const dbnameenv = "fake_db_var"
	const dbnameval = "fake_db"
	
	//set fake vars
	t.Setenv(hostenv, hostval)
	t.Setenv(portenv, portval)
	t.Setenv(userenv, userval)
	t.Setenv(passenv, passval)
	t.Setenv(dbnameenv, dbnameval)
	//

	FakeEnvVars := psqlVarEnv{
		Host:     hostenv,
		Port:     portenv,
		User:     userenv,
		Password: passenv,
		Dbname:   dbnameenv}

	FakeCredential := PsqlCredentials{
		Host:     hostval,
		Port:     portval,
		User:     userval,
		Password: passval,
		Dbname:   dbnameval}

	bytes, err := json.Marshal(FakeEnvVars)
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
