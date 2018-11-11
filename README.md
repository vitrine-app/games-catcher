# Vitrine games catcher

## About
This module is used to cache data being pulled from IGDB, in order to reduce the amount of queries.

When a game is asked, it will check in the database that the game is not already cached, and if it isn't it will query the IGDB web API,
format the result and store it into the database. Then, it will be returned as a JSON string.

This module is written in Go and is not made to work on is own, but rather to be compiled to a C dynamic library to be imported
by the server written in Python.

## Running
As it's made to be executed as library by the web server and not on it's own, the games catcher module embraces the Docker configuration
of the [Vitrine web server](https://github.com/vitrine-app/server), the best way to use it is to get the web server running.

However, it can be painful to recompile the library and to use the library from the server for every change, so another solution
is possible if you want to use the games catcher outside of the Docker containers.

You need to populate the `main` function in `games_catcher.go` (which is empty by default because not-needed for library export), then run:
```
make binary
```
This will build a `games_catcher` binary that you could execute.

However, some env variables are needed in order to communicate with the database.
You need to create an `.env` file at the root of the folder based on this structure:
```
MYSQL_HOST={Your MySQL host IP}
MYSQL_ROOT_PASSWORD={The root password for MySQL}
```
These variables are the ones used by the `db` container of the Vitrine server Docker configuration (see server [docker-compose.yml](https://github.com/vitrine-app/server/blob/master/docker-compose.yml) for more details)
Then, just run
```
sh run_binary.sh
```
