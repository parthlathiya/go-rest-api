# REST APIs in Go

This repo contains sample Go application which has 2 APIs. one to create an user and another to get user. I have used gorill mux for request router.
For easy setup, I have added docker-compose config.

## Getting Started

### Prerequisites

Install Docker. see the [install instructions.](https://www.docker.com/get-started)

Below instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Installing

Clone repo in local

```
git clone https://github.com/parthlathiya/go-rest-api.git
```

### Build and run

Build services using docker-compose. Go to project's root dir where docker-compose.yml located and run below cmd.
It will build all services mentioned in docker-compose file.

```
docker-compose build
```

Now create container for go app and mysql db.

```
docker-compose up
```


Now you should have go app and mysql db up & running...

### Usage

Create User API

```
curl --location --request POST 'http://localhost:8092/user' --header 'Content-Type: application/json' --data-raw '{
    "email": "test@gmail.com",
    "first_name": "test_first_name",
    "last_name": "test_last_name"
}'
```

Get User API

```
curl --location --request GET 'http://localhost:8092/user?email=test@gmail.com' --header 'Content-Type: application/json'
```

### Other useful commands

Enter inside go app container(In docker-compose file, we have defined it as `web`)

```
docker-compose exec web sh
```

Connect with mysql inside container from host(your machine). Enter `test_password` as password

```
mysql -P 3312 --protocol=tcp -u test_user -p
```
