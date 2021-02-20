## Simple Hub Management System

A simple API for Hub management

### Requirements

- Golang version ~1.14
- Docker

### How to run

- Build and run:
    - Update config to connect your database in `.env`.
    - Build app to executable file with command `go build -o app ./cmd/api/hub-api`.
    - Run your app `./app`.
- Use Docker:
    - Just run `docker-compose up` in project root folder.
- Check url `http://127.0.0.1:8080/`.

### Folder struct

```
.
├── adapter
│   └── repository        # Contains data layer logic codes
├── cmd
│   └── api
│       └── hub-api       # Store codes for API handlers
├── db                    # Database initial files
├── docker-compose.yml
├── domain                # Containts interface adapters
├── infra
│   └── postgres.go       # Init code for DB connection
├── pkg                   # Some helper funcs for running service
│   ├── core
│   └── utils
└── usecase               # Contains application specific business rules
```

## API Requests

| Endpoint        | HTTP Method           | Description       |
| --------------- | :---------------------: | :-----------------: |
| `/users` | `POST` | `Create user` |
| `/users/{{user_id}}` | `GET` | `Get user by id` |
| `/users`   | `GET` | `Get list user` |
| `/teams`| `POST` | `Create team` |
| `/teams/{{team_id}}`| `GET` | `Get team by id` |
| `/teams`| `GET` | `Get list team`  |
| `/teams/{{team_id}}/users`| `GET` | `Get list user in team` |
| `/teams/{{team_id}}/add-users`| `PUT` | `Add users to team` |
| `/hubs` | `POST` | `Create hub` |
| `/hubs/{{hub_id}}` | `GET` | `Get hub by id` |
| `/hubs`   | `GET` | `Get list hub` |
| `/hubs/{{hub_id}}/teams`   | `GET` | `Get list team in hub` |

## Test endpoints API using curl

- #### Creating new user

`Request`
```bash
curl --location --request POST 'http://127.0.0.1:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "role": "customer",
    "email": "email1@iamdzung.com"
}'
```

`Response`
```json
{
  "data": {
    "id": 9,
    "role": "customer",
    "email": "email1@iamdzung.com"
  },
  "status": "success"
}
```

- #### Creating new hub

`Request`
```bash
curl --location --request POST 'http://127.0.0.1:8080/hubs' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Hub ABC",
    "lat": 11.123,
    "long": 34.455
}'
```

`Response`
```json
{
  "data": {
    "id": 10,
    "name": "Hub ABC",
    "lat": 11.123,
    "long": 34.455
  },
  "status": "success"
}
```

- #### Creating new team

`Request`
```bash
curl --location --request POST 'http://127.0.0.1:8080/teams' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Hub ABC",
    "type": "pro",
    "hub_id": 10
}'
```

`Response`
```json
{
  "data": {
    "id": 5,
    "name": "Hub ABC",
    "type": "pro",
    "hub_id": 10
  },
  "status": "success"
}
```

- #### Add users to a team

`Request`
```bash
curl --location --request PUT 'http://127.0.0.1:8080/teams/4/add-users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_ids": [
        1, 2
    ]
}'
```

`Response`
```json
{
  "status": "success"
}
```

- #### Find hubs

`Request`
```bash
curl --location --request GET 'http://127.0.0.1:8080/hubs?name=ABC'
```

`Response`
```json
{
  "data": [
    {
      "id": 9,
      "name": "Hub ABC",
      "lat": 11.123,
      "long": 34.455
    },
    {
      "id": 10,
      "name": "Hub ABC",
      "lat": 11.123,
      "long": 34.455
    }
  ],
  "status": "success"
}
```

- #### Find teams

`Request`
```bash
curl --location --request GET 'http://127.0.0.1:8080/teams?name=Red&type=pro,cool&hub_id=2'
```

`Response`
```json
{
  "data": [
    {
      "id": 1,
      "name": "Team Red",
      "type": "cool",
      "hub_id": 2
    }
  ],
  "status": "success"
}
```