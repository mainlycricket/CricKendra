I am building `CricKendra` out of my passion for cricket and programming.

### Current Progress

- A CSV parser - [`cmd/csv-parser`](./cmd/csv-parser) - that extracts data from CSV files available on [cricsheet.org](https://cricsheet.org) and saves in the database.

- A main server - [`cmd/main-server`](./cmd/main-server) - that provides basic endpoints, currently focused on stats engine.

### Roadmap

- CSV Parser ✅
- Stats Tools:
    - Batting Stats ✅
    - Bowling Stats ✅
    - Teams Stats ✅
- Single Match & Single Series routes ✅
- Live Scoring System ✅
- Single Player, Series, Tournaments Stats Tool ✅
- Frontend ⏳
- All-round & Partnership in Stats Tools ⏳
- Content Management System ⏳

### Run the project

- Stop the localhost postgresql service
```sh
    sudo systemctl stop postgresql
```

- Rename `dockerfiles/.env.example` to `dockerfiles/.env` and set the variables
```sh
    mv dockerfiles/.env.example dockerfiles/.env
```

- Run the CSV parser to seed initial data
```sh
    docker compose --file=./dockerfiles/docker-compose-csv-parser.yml up --build 
```

- Stop the CSV parser container after you see the message "DB dumped, stop the process"

- Run the main-server to start the REST API server
```sh
    docker compose --file=./dockerfiles/docker-compose-main-server.yml up --build 
```


### Resources

- Postman Collection [here](https://documenter.getpostman.com/view/25403102/2sAYBREZ3x)
- DB Schema [here](https://dbdiagram.io/d/CricKendra-670bfc5697a66db9a3d0b44a)
