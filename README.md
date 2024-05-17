# Application Documentation

## Requirement
1. Go version 1.20 above
2. Migrate, for database migration. https://github.com/golang-migrate/migrate
3. PostgresSQL Database or docker for dababase

## Setup
1. Git clone this repository
2. Install application with command : 
```bash 
make init-app 
```
3. Copy appplication config with command : 
```bash 
cp config-example.yml config.yml
```
4. Edit `config.yml` file based on your own configuration.
5. Execute migration with command : 
```bash 
migrate -path migrations -database "postgresql://username:password@host:port/databasename?sslmode=disable" -verbose up
```


## Run Application
1. Run with command : 
```bash 
make run
```
2. API url example : [GET] localhost:10001/food/menus 
