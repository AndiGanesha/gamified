# Authentication

This is a service for processing user data get sign up and sign in

## Prerequisites

Before running the service, make sure you have the following installed on your machine:

- Go (version 1.20.5)
- Redis (version 6.0.16)
- SQL

## The More You Know
This application is used for currently local configuration, this program include main folder application, infrastructure, controller, service, and repository.

Mainly the unit test will be set to test controller service and repotiry.
1. application : main program that will initiate the others
2. infrastructure : the one that will be hold the ground from outside world
3. controller : the one that controll input
4. service : is where the most of logic reside
5. repository : the folder to read the database SQL
6. configuration : config for the database redis and others that can be easily change over time (ps: make sure your config is filled and same as your intended DB/Redis)

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/AndiGanesha/authentication.git
2. Navigate to the project directory:

   ```shell
   cd authentication
   go mod tidy
3. Build a project:

   ```shell
   make build
4. Create a database:

   ```shell
    $mysql -u root -p
     Enter password:

    mysql> create database authentication;
    mysql> use authentication;
    mysql> source ~authentication/db/CREATE_USER_TABLE.SQL;
5. Set up the configuration by creating a local.env file in the config directory. You can use the local.env.example file as a template.

    Run the service:

    ```shell
   make run
   ```
    The service will start running on the specified address and port. You can now make API requests to the service.

# API Documentation

The service exposes the following endpoints:
and for that 2 expoxes endpoints we need to have application/json body that include this contract below as an example.

    {
        "username" : "abc",
        "password" : "123sssdddd",
        "phone" : "+622222222" //Optional
    }
## Sign up
    {host}/sign_up (POST)
For signing up users/front-end needed to hit the exposes API above using the contract needed, if the data is verified and username not currently in use, this service will making a new user and insert it to the user database, as that if its succeess it will trying to generate JWT token and save it to the Redis for later use in the actual Apps. and for the response if all the process above OK (200) it will give JWT.Raw token, that will be used in the actual Apps.

    {host}/sign_in (POST)
For signing in users/front-end needed to hit the exposes API above using the contract needed without the optional, if the data is verified and username there in the Database, this service will try to generate JWT token and save it to the Redis for later use in the actual Apps. and for the response if all the process above OK (200) it will give JWT.Raw token, that will be used in the actual Apps.

# Development
## Running Tests

To run the unit tests, execute the following command:
```shell
make test
```
## Generating Mocks

To run the unit tests, execute the following command:
```shell
make mock -B
```
## Linting

To run the unit tests, execute the following command:
```shell
make lint
```