# GrindLists

A webapp to help gamers keep track of materials collected during grinding
sessions, keeping track of progress toward a goal (i.e. a quest completion
or item to craft).

Currently, the easiest way to run this if you wish to contribute is to use 
docker (docker and make required).

## Run with Docker and Make (Recommended)

Requirements
    - Docker & Docker Compose
    - Make
    - Go (optional: need it installed locally if you want to take advantage of the Golang VSCode extension)

1. Create a `secrets` folder in the root directory of the project and create
three files, `postgres_db`, `postgres_passwd`, `postgres_user`, with the corresponding
values inside.

2. Create a `config.yaml` file in the `api` folder based on the provided `config.example.yaml`,
substituting your own `secret_key` and the values from step 1 for `db_conn`.
Make sure that `host` is `db` if you are using the docker setup. Eventually
this will be changed so if your config does not provide one, the secret files will
be used instead, in combination with some other defaults for the docker setup.

3. Add the following entry to your hosts file.
```
127.0.0.1 api.local traefik.api.local
```
This will allow you to access the api at `api.local` on your computer, such as
in the graphql playground hosted by the api itself, or hoppscotch.io.
You can also access traefik's dashboard at `traefik.api.local`

4. Run `make api`. This will start up containers for the api, redis, and postgresql,
all behind a traefik reverse proxy. If you added the host file entries in
step 3, you can now access the graphql playground in your browser at `api.local`,
query the graphql api at `api.local/graphql`, and access traefik's dashboard at
`traefik.api.local` You can choose different domains if you wish, but you will
need to edit the `docker-compose.yml` file.
5. (Optional) Run `make debug-api` in order to run the api in debug mode. This
starts a delve instance inside the container, and the provided launch.json
provides an easy way to attach to the delve instance in vscode via the
"Launch Remote" configuration. So you can set breakpoints and make requests that
trigger them and view the values of variables easily.

Eventually you will be able to run `make all` in order to launch the api *and*
the client, but right now the client

## Run Without Docker (Not Recommended)

Requirements
    - PostgresQL Server Running with a database already created for the api
    - A Redis Server (currently a hard requirement, will be fixed later on)
    - Go
    - NPM or Yarn

More Detailed steps will be coming later, but basically you will need to 
Build the api with Go (or run it with go run), build the Client with Yarn (or run it with
yarn dev), start up and configure a Postgres instance, with the database and users already
set up (but not seeded), start up a redis server, provide the config.yaml file
to the api and the .env file for the client, run the `migrate up` command from the
api (either previously built or with `go run`), and run the `serve` command of the 
api. Then host the frontend somewhere (most likely locally with yarn dev).

As you can tell this is complicated and is why we recommend the docker configuration.

## All available Make commands

Requirements
    - Docker & Docker Compose
    - Make

The following commands can be run via `make` to do the following actions:

| Command | Description |
| ------- | ----------- |
| all     | Starts up all services for the app (Currently equivalent to `make api`, as the client has barely any progress on it so far).
| api     | Start all the necessary services for the api |
| debug-api | Start all the necessary services for the api in debug mode |
| tidy    | Clean up the unused dependencies from the api containers go.mod file |
| exec    | Run followed by a username, service name, and command to run that command in the specified container as the specified user (i.e. `make exec user=grindlists service=grindlists cmd="migrate create -n migration_name"` to create a new migration). The username and service can be omitted (they are root and api, respectfully). |
| test-api | Run all the tests for the api (there currently are none, I need to catch up on that!) |
| debug-db | Opens a connection to the db container in pgcli |
| dump     | Dumps the current database data to `./api/scripts/backup.sql` |
| down     | Tears down all containers (equivalent to docker-compose down) |