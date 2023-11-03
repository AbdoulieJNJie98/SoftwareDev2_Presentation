# Software Development 2 Persentation code
This project uses a Golang backend service with a dockerized postgreSQL database, and postman to send and test the APIs

## Table of Summary

-[Prerquisties](#prerquisties)
-[APIs](#apis)
-[Docker](#docker)
-[PostgresSQL](#postgresql)

## Prerquisties

Install the following softwares with their corresponding versions to run this application

- [Golang](https://go.dev/) (v1.20.6)
- [Docker](https://www.docker.com/) (v24.0.6)

### Additional Softwares Used

- [Postman](https://www.postman.com/)

Postman is used for testing the web service's APIs; the current version of Postman used is Version 10.19.7

-[Docker Desktop](https://www.docker.com/products/docker-desktop/)

Docker Desktop is used to manage the container and image; Docker Desktop is not neccsary since this application donwloads the docker package. The current version of Docker Desktop used is v4.23.0

## APIs


# Command used to install package that connects the applicaiton to the database
```bash
go get github.com/jackc/pgx/v5
```

