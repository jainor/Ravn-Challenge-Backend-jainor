package db

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
  "sync"

	_ "github.com/lib/pq"
)

//PsqlCredentials stores the name of the env variables that stoers credentials to open a db connection
type psqlVarEnv struct {
	Host     string
	Port     string
	User    string
	Password string
	Dbname   string
}

//PsqlCredentials stores the name of the credentials to open a db connection
type PsqlCredentials struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

const configPath = "../../configs/postgresconfig.json"

// PsqlManager implements methods of interface DBManager
type PsqlManager struct {
	name        string
	credentials PsqlCredentials
}

//singleton pattern
var pm  *PsqlManager
var once sync.Once

// GetManager return the instance of the DBManager
func GetManager() *PsqlManager {
    once.Do(func() {
        pm = &PsqlManager{
            credentials : loadConfig()}
    })
    return pm
}

func loadConfigFile(filename string) PsqlCredentials {
	var config psqlVarEnv
	configFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return getVarsEnv(config)
}

func getVar(k string, v* string){
  *v = os.Getenv(k)
  if( *v == ""){
    log.Fatal("Environment var " + k + " not set!")
  }
}

func getVarsEnv(pvar psqlVarEnv) PsqlCredentials{
  var host, port, user, password, dbname string

  getVar(pvar.Host, &host)
  getVar(pvar.Port, &port)
  getVar(pvar.User, &user)
  getVar(pvar.Password, &password)
  getVar(pvar.Dbname, &dbname)

  return PsqlCredentials{
    Host:host,
    Port:port,
    User:user,
    Password:password,
    Dbname:dbname,
  }
}

func loadConfig() PsqlCredentials {
	return loadConfigFile(configPath)
}



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
/*
Query 1:

select * from authors order by date_of_birth limit 10;

Query 2:

select sum(salesi.item_price * salesi.quantity) from sale_items salesi
left join books b on salesi.book_id = b.id  left join authors a on b.author_id = a.id where a.name =  'Lorelai Gilmore';

Query 3:

select a.id, a.name,a.date_of_birth from sale_items salesi left join books b on salesi.book_id = b.id left join  authors a on b.author_id = a.id group by a.id order by sum(salesi.item_price * salesi.quantity) desc limit 10;


*/


//QueryStr returns the query string we are interested
func (Pm PsqlManager) QueryStr() string {
  return "select a.id, a.name,a.date_of_birth from sale_items salesi left join books b on salesi.book_id = b.id left join  authors a on b.author_id = a.id group by a.id order by sum(salesi.item_price * salesi.quantity) desc limit %d"
}

// Name returns the database manager name
func (Pm PsqlManager) Name() string {
	return "postgres"
}
