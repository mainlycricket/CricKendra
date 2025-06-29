I am building `CricKendra` out of my passion for cricket and programming.

### Current Progress

- A CSV parser - [`backend/cmd/csv-parser`](./backend/cmd/csv-parser/) - that extracts data from CSV files available on [cricsheet.org](https://cricsheet.org) and saves in the database.

- A main server - [`backend/cmd/main-server`](./backend/cmd/main-server/) - that provides matches, series, stats & admin endpoints

- A frontend application - [`frontend/`](./frontend/) - currently in the development

### Roadmap

- CSV Parser ✅
- Backend
  - Stats Tools: Batting, Bowling, Team ✅
  - Single Match & Single Series routes ✅
  - Admin Routes with Live Scoring System ✅
- Frontend
  - Single Match, Single Series & Single Player ✅
  - Stats Tool ⏳
  - Admin Panel with Live Scoring System ⏳
  - Navbar, Header, Footer fixes ⏳
- Series & Tournaments view with search ⏳
- All-round & Partnership in Stats Tools ⏳
- Content Management System ⏳

### Run the project

- Rename `dockerfiles/.env.example` to `dockerfiles/.env` and set the variables

```sh
    mv dockerfiles/.env.example dockerfiles/.env
    # set the variables
```

- Run the CSV parser to seed initial data

```sh
    docker compose --file=./dockerfiles/docker-compose-csv-parser.yml up --build
    # stop the container after you see the message "DB dumped, stop the process"
```

- Run the app locally

```sh
    docker compose --file=./dockerfiles/docker-compose-main-server.yml up --build
    # Access it at http://localhost:3000
```

### Resources

- Screenshots [here](https://www.linkedin.com/pulse/crickendra-screenshots-tushar-jain-vftjc)
- Postman Collection [here](https://documenter.getpostman.com/view/25403102/2sAYBREZ3x)
- DB Schema [here](https://dbdiagram.io/d/CricKendra-670bfc5697a66db9a3d0b44a)

### Acknowledgement

- The stats tool idea is inspired by [_ESPNCricinfo_](https://espncricinfo.com), but I don't know its implementation.
- I am working on this project only for learning purpose: designed the database schema & the developing it on my own.
