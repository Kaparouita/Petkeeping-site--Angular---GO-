# Assignment 3 - HY-359

# Authors: Stivaktakis Giorgos (AM: 4300)

## Description

This project implements a web application with a backend written in Go and a frontend developed using Angular.  
The application utilizes Postgres as the database for storing data. The backend handles the business logic  
and provides RESTful APIs for communication with the frontend. The frontend is responsible for presenting  
the user interface and interacting with the backend APIs.

## About

The main functionality of the application is to act as a pet keeper service. The system allows owners to find  
and connect with keepers who can take care of their pets when they are unable to do so.

## Installation

### Backend

The backend is written in Go so you will need to have Go installed on your system.
I use docker to init the database.  
Navigate to the backend directory and run the following command to install the dependencies and start the server:

```bash
docker-compose up -d
go run main.go
```

### Frontend

The frontend is written in Angular so you will need to have Node.js and npm installed on your system.  
Navigate to the frontend directory and run the following command to install the dependencies and start the server:

```bash
npm install
ng serve
```

## Missing

Maps dosent work  
After a user change his fields the checks for the fields are not working
