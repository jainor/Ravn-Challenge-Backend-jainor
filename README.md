### Ravn Challenge Backend

## Description

Simple API endpoint that returns the list of the top authors

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#sql-queries">SQL Queries</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#Deploy-with-Google-Kubernetes-Engine"> Deploy with Google Kubernetes Engine</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Project developed as a challenge for RAVN.

The deployment of the project was done in GKE.

The current link of the endpoint is:

[http://34.125.185.134/authors?count=3](http://34.125.185.134/authors?count=3)

34.125.185.134 is the External IP of the API service. We also made external the management tool of the messagebroker

[http://34.125.13.114:15672](http://34.125.13.114:15672)


Use `guest` in both fields (user and password).

<!-- SQL QUERIES -->
## SQL Queries

The sql queries are commented in the last part of the archive [https://github.com/jainor/Ravn-Challenge-Backend-jainor/blob/main/internal/db/psqlmanager.go] (https://github.com/jainor/Ravn-Challenge-Backend-jainor/blob/main/internal/db/psqlmanager.go)



### Built With

* [Golang](https://golang.org/)
* [Postgres](https://www.postgresql.org/)
* [RabbitMQ](https://www.rabbitmq.com/)
* [Gin-Conic](https://github.com/gin-gonic/gin)
* [Docker](https://www.docker.com/)




<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.

* git
* docker
  
### Config

The name of environment variables  for database and messagebroker (RabbitMQ) credentials are storage in `config` directory.
The values should be setted before, read `.env` file to see an example.


Script for the database is storage in `scripts` directory.

### Installation

We show how we generate images and test them. Moreover, we also show to push the images into [Docker Hub Containing Library](https://hub.docker.com/)
We only push the images that corresponds to the API endpoint and workers.

1. Clone the repo
   ```sh
   git clone https://github.com/jainor/Ravn-Challenge-Backend-jainor
   ```
2. Build containers
   ```sh
   docker-compose build
   ```
3. Running containers
   ```sh
   docker-compose up
   ```


Once the generated images are running we have 4 different services running

1. Database
2. MessageBroker (Rabbit)
3. API endpoint (Publish messages and receives query results via RPC callbacks)
4. Workers (Consuming messages) 

To access psql inside Database container
   ```sh
   docker ps
   docker exec -it <container_id_for_db>  psql -U <user> -w <database_name>
   ```

To monitor Messagebroker (GUI) open `http://localhost:15672` in the browser and use guest (docker-compose config) as user and password 


![](images/screenshot.png)

To use the Api endpoint, execute

 ```sh
   curl localhost:8080/authors?count=<number>
   ```

Once we test the local deployment works correctly.
We create the images for the API endpoint and workers.
Please create an account in [Docker Hub Containing Library](https://hub.docker.com/) to push the image.

 ```sh
   docker ps
   docker container commit  <ContainerID>  jainor/challenge:latest
   docker login
   docker push jainor/challenge:latest
   ```

Now, any developer can use your image!!

<!-- USAGE EXAMPLES -->
## Deploy with Google Kubernetes Engine

Download kubectl and gcloutoon your computer.

Before deploy our services in GKE, we must use your gmail account in [Google Cloud](https://cloud.google.com), we must use the free trial to use GKE.
We have to activate GKE option.

We start by configuring gcloud locally and creating a project
 ```sh
gcloud init
gcloud  projects create jainor-ravn --name="jainor ravn"
   ```

create a cluster and configure credentials for kubectl
 ```sh
gcloud  container clusters create cluster-jainor
gcloud  container clusters get-credentials  cluster-jainor
   ```

Make sure your region is configured in init, or setup on  [Google Cloud](https://cloud.google.com), to create the cluster with the instruccion above.

We are ready to work with kubectl in our cluster!!!

First, check the current cluster

 ```sh
kubectl config current-context
   ```

Now we have to setup the registry of Docker hub in kubernetes (to use images),

 ```sh
kubectl create secret docker-registry regcred --docker-server=https://index.docker.io/ --docker-username=<user> --docker-password=<password> --docker-email=<email>

First, we create the configmaps and secrets for envinroment variables.
In the case of secrets, the value of env vars should be generated in the following way:
 ```sh
 echo -n '<value>' | base64
   ```
And then, copy into secrets.yaml. After that, we create configmaps and environment 

 ```sh
kubectl apply -f k8s/secrets.yaml
kubectl apply -f k8s/configmap.yaml

Later, we create the deployments in folder `k8s`

 ```sh
kubectl apply -f k8s/endpoint.yaml
kubectl apply -f k8s/rabbitmq.yaml
kubectl apply -f k8s/worker.yaml
kubectl apply -f k8s/persistentvolume.yaml
kubectl apply -f k8s/postgres.yaml
   ```

check the state of the deployments:

 ```sh
kubectl get deploy -o wide
   ```

and the status of the pods:

 ```sh
kubectl get pods -o wide
   ```

Note that you must wait some minutes (30 minutes or less) due to the claim of persistent volume.
Before this, only the endpoint and messagebroker (rabbitmq) deploys will be ready.

To see logs:

 ```sh
kubectl logs <Name-Pod>
  ```

Open Postgres database manager and create database

 ```sh
   kubectl exec -it <Name-Pod-Postgres-Dep> -- psql -U <user>
   # \psql create database <database>
   kubectl exec -it <Name-Pod-Postgres-Dep> -- psql -U <user> -d <database>
   ## create tables and populate database manually with scripts/scriptdb.sql
   ```

Now, we create the services for our deployments:

 ```sh
kubectl expose deployment/messagebroker --type=LoadBalancer --name=management --port 15672 --target-port 15672
kubectl expose deployment/messagebroker --type=LoadBalancer --name=messagebroker --port 5672 --target-port 5672
kubectl expose deployment/worker --type=ClusterIP --port 5672 --target-port 5672
kubectl expose deployment/endpoint --type=LoadBalancer --port 80 --target-port 8080
   ```

The service for the Postgres deployment was created in the previous step.

Now our API endpoint is on the web!!!

<!-- LICENSE -->
## License

Distributed under the Mozilla Public License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Jainor Cardenas - jainorcardenas@gmail.com

Project Link: [https://github.com/jainor/Ravn-Challenge-Backend-jainor](https://github.com/jainor/Ravn-Challenge-Backend-jainor)



<!-- ACKNOWLEDGEMENTS -->
## Acknowledgments

* [Golang spec](https://golang.org/ref/spec)
* [RabbitMQ tutorials](https://www.rabbitmq.com/getstarted.html)
* [Docker Reference](https://docs.docker.com/reference/)

