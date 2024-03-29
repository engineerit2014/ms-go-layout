# Project
This project is an example of layout using the Clean Architecture concepts.  

## Table of Contents

- [Running](#running)
    - [Dependencies](#dependencies)
    - [Environment Variables](#environment-variables)

## Running

This repository runs a simple main daemon (main) that implements a REST API for users. The daemon uses a postgres database to persist data and after creation uses migrations to create the user schema.

### Dependencies 

The only dependencies to run the services in this repository are:

- `docker`
- `docker-compose`

### Environment Variables

The program looks for the following environment variables:

- `DB_USER`: The postgres database username that gets used within the postgres connection
string (Default: `root`).
- `DB_PASS`: The postgres database password that gets used within the postgres connection
string (Default: `root`).
- `DB_NAME`: The postgres database name that gets used within the postgres connection string
(Default: `user`).
- `DB_HOST`: The postgres database host name that gets used within the postgres connection
string (Default `db`).
- `DB_PORT`: The postgres database port that gets used within the postgres connection string
(Default: `5432`).

If the environment variable has a supplied default and none are set within the context of the host
machine, then the default will be used.
 
To set any given environment variable, simply execute the following
pattern, replacing `[ENV_NAME]` with the name of the environment variable and `[ENV_VALUE]` with the
desired value of the environment variable: `export [ENV_NAME]=[ENV_VALUE]`. To unset any set environment
variable, simply execute the following pattern, replacing `[ENV_NAME]` with the name of the environment
variable: `unset [ENV_NAME]`.