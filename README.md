# HTTP API frontend with GRPC backend and Postgres DB

API exposes HTTP endpoint and communicates with PortService via GRPC
PortService uses Posgres for database

Dependencies are vendored and go.mod has local rewrite directives for offline development

## Running

### With docker compose

Exposed ports :8080 for ClientAPI and :1234 for PortDomainService
Put ports.json to `sampleData` folder for altrenative inputs

    $ docker-compose up api
    [...]

    $ curl http://localhost:8080/\?name\=UAIEV
    {"Key":"UAIEV","Name":"Kiev","City":"Kiev","Country":"Ukraine","Aliases":null,"Regions":null,"Latitude":30.5234,"Longitude":50.4501,"Province":"Kyiv city","LocationName":"Europe/Kiev","Unlocs":["UAIEV"],"Code":""}

### Individually

    # start postgres
    $ docker run --name pg -p 5432:5432 -e POSTGRES_PASSWORD=pw -d postgres
    ...

    // start domain service
    /repo/PortDomainService $ ./run.sh
    Connecting to db:  postgres://postgres:pw@localhost:5432/postgres?sslmode=disable
    sql.Open() ok
    Ping ok, DB connected
    Starting RPC server on  :1234

    // start client api
    /repo/ClientAPI $ ./run.sh
    Dialed  localhost:1234
    Starting http service on port  8080

    // run tests against domain service and api
    /repo/tests $ ./run.sh
    Dialed  localhost:1234
    Sent summary: count:1 
    Sent summary: count:1 
    PASS
    ok  github.com/tonky/shipping/tests 0.018s

### Testing

#### Integration tests via compose

 $ docker-compose up tests

Inserts and overwrites a record via GRPC, then queries it with HTTP to make sure updated version is there

#### Integration tests with local services

Run each service individually, as described in "Running -> Individually", then

    /repo/tests $ ./run.sh

#### ClientAPI test

    /repo/ClientAPI $ go test

#### PortDomainService test

    /repo/PortDomainService $ go test
