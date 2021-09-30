package messagequeue

import (
	"fmt"
	"os"
	"sync"

	"encoding/json"
	"log"
	"math/rand"
	"strconv"

	_ "github.com/lib/pq"

	amqp "github.com/rabbitmq/amqp091-go"

	idb "internal/db"
)

//var conn *amqp.Connection

// rabbitVarEnv stores the names of the variables that stores the credentials for rabbit
type rabbitVarEnv struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}


// RabbitCredentials stores Rabbitmq credentials for connections
type RabbitCredentials struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

const configPath = "../../configs/rabbitconfig.json"

func loadConfig() RabbitCredentials {
	return loadConfigFile(configPath)
}
func loadConfigFile(filename string) RabbitCredentials {
	var config rabbitVarEnv
	configFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return getVarEnv(config)
}

func getVar(k string, v* string){
  *v = os.Getenv(k)
  if( *v == ""){
    log.Fatal("Environment var " + k + " not set!")
    }
}

func getVarEnv(pvar rabbitVarEnv) RabbitCredentials{
  var host, port, user, password, name string

  getVar(pvar.Host, &host)
  getVar(pvar.Port, &port)
  getVar(pvar.User, &user)
  getVar(pvar.Password, &password)
  getVar(pvar.Name, &name)

  return RabbitCredentials{
    Host:host,
    Port:port,
    User:user,
    Password:password,
    Name:name,
  }
}

var cred RabbitCredentials
var once sync.Once 

func GetManager() *RabbitCredentials {
    once.Do(func() {
	      cred = loadConfig()
    })
    return &cred
}

func (rc RabbitCredentials) strConn() string {
	connstr := fmt.Sprintf("%s://%s:%s@%s:%s/",
		rc.Host, rc.User, rc.Password, rc.Name, rc.Port)
	return connstr
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

// SendMessage sends a message to Rabbitmq server and returns a result via RPC
func SendMessage(n int) ([]byte, error) {

	conn, err := amqp.Dial(GetManager().strConn())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	return sendMessageToConn(n, conn)
}

func sendMessageToConn(n int, conn *amqp.Connection) ([]byte, error) {

	corrId := randomString(32)

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	err = ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})

	if err != nil {
		return make([]byte, 0), err
	}

	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			//res, err := strconv.Atoi(string(d.Body))
			res := d.Body
			failOnError(err, "Failed to convert body to integer")

			return res, nil
		}
	}
	return make([]byte, 0), nil
}

// ResponseMessage responds the messages waiting
func ResponseMessage(m idb.ModelQ) {

	conn, err := amqp.Dial(GetManager().strConn())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	responseMessageToConn(m, conn)
}

func responseMessageToConn(m idb.ModelQ, conn *amqp.Connection) {
	log.Println(" [*] Awaiting RPC requests")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to convert body to integer")

			//log.Println(" [.] QueryDB(%d)", n)
			response := m.QueryDB(n)
			fmt.Println(response)

			jsonbyte, err := json.Marshal(response)

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          jsonbyte,
				})
			failOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Println(" [*] Awaiting RPC requests")
	<-forever

}
