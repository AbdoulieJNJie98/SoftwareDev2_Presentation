# Software Development 2 Persentation code
This project uses a Golang backend service with a dockerized postgreSQL database, and postman to send and test the APIs

## Table of Summary


- [Prerquisties](#prerquisties)
- [APIs](#apis)
- [PostgreSQL](#postgresql)

## Prerquisties

Install the following softwares with their corresponding versions to run this application

- [Golang](https://go.dev/) (v1.20.6)
- [Docker](https://www.docker.com/) (v24.0.6)

### Additional Softwares Used

- [Postman](https://www.postman.com/)

  Postman is used for testing the web service's APIs; the current version of Postman used is Version 10.19.7

- [Docker Desktop](https://www.docker.com/products/docker-desktop/)

  Docker Desktop is used to manage the container and image; Docker Desktop is not neccsary since this application donwloads the docker package. The current version of Docker Desktop used is v4.23.0

## APIs

  The following API descriptions consist of:
  
  HTTP request types, specific paths/endpoints, summaries of what the API does, possible responses from the API, descriptions of the responses, and the content type of the responses.

### POST /videoGame Add a new video game
  
  200 - Request was successful

  400 - Request was invalid
  
  application/json

### PATCH /videoGame Edit an existing video game
  200 - Request was successful

  400 - Request was invalid
  
  application/json

### Delete /videoGame Delete an existing video game

  200 - Request was successful

  400 - Request was invalid
  
  application/json

## PostgreSQL

  PostgreSQL is used for storing video game titles that consist of the game's title, platform, genre, and price. 
  In order to ensure the project is connected to the database, run the following command first:
  ```bash
 docker pull postgres:13.11
  ```
  Next, enter the following command
 ```bash
 docker run -p 5432:5432 --name some-postgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -d postgres:13.11
  ```
 The first command pulls the postgres docker image from the docker hub, and the second command runs the actual image with the proper enviromental varaibles
