package db

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

//PsqlCredentials stores the credentials to open a db connection
type PsqlCredentials struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

const configPath = "../../configs/postgresconfig.json"

//const configPath = "./configs/postgresconfig.json"

// PsqlManager implements methods of interface DBManager
type PsqlManager struct {
	name        string
	credentials PsqlCredentials
}

var Pm = PsqlManager{credentials: loadConfig()}

func GetManager() DBManager {
	return Pm
}

func loadConfigFile(filename string) PsqlCredentials {
	var config PsqlCredentials
	configFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func loadConfig() PsqlCredentials {
	return loadConfigFile(configPath)
}

/*

func GetManager() *PsqlManager {
    sync.once.Do(func() {

        pm = &PsqlCredentials{}
    })
    return pm
}
*/


// ConnectionStr returns the connection string to connect with db
func (Pm PsqlManager) ConnectionStr() string {
	return fmt.Sprintf(
		"host=%s "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s sslmode=disable",
		Pm.credentials.Host, Pm.credentials.Port, Pm.credentials.User, Pm.credentials.Password, Pm.credentials.Dbname)
}

//QueryStr returns the query string we are interested
func (Pm PsqlManager) QueryStr() string {
	return "select * from authors order by date_of_birth limit %d"
}

// Name returns the database manager name
func (Pm PsqlManager) Name() string {
	return "postgres"
}
