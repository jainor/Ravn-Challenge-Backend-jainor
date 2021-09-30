package messagequeue

import (
	"encoding/json"
	"testing"
	"time"

	"reflect"

	"io/ioutil"
	"os"

	dbent "internal/db/entities"
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

	const nameenv = "fake_name_var"
	const nameval = "fake_name"
	
	//set fake vars
	t.Setenv(hostenv, hostval)
	t.Setenv(portenv, portval)
	t.Setenv(userenv, userval)
	t.Setenv(passenv, passval)
	t.Setenv(nameenv, nameval)
	//

	FakeEnvVars := rabbitVarEnv{
		Host:     hostenv,
		Port:     portenv,
		User:     userenv,
		Password: passenv,
		Name:     nameenv}

	FakeCredential := RabbitCredentials{
		Host:     hostval,
		Port:     portval,
		User:     userval,
		Password: passval,
		Name:     nameval}

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

var exp []dbent.Author

func setup() {
	timeany, _ := time.Parse(time.RFC3339, "2012-01-02T15:04:05.000Z")
	exp = []dbent.Author{
		dbent.Author{
			Id:    0,
			Name:  "test",
			Dated: timeany,
		},
	}
}

type FakeModel struct {
}

func (fm FakeModel) QueryDB(n int) []dbent.Author {
	return exp
}

func TestChannelCreation(t *testing.T) {

	setup()

	// TODO create a rabbitmq instance for testing
	//  Not common interface for connection and fakeconnection in
}
