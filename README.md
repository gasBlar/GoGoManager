## GO GO MANAGER!!!!!!!

## Installation

There are a few tools that you need to install to run the project.
So make sure you have the following tools installed on your machine.

- [Migrate (for DB migrations)](https://github.com/golang-migrate/migrate/tree/v4.17.0/cmd/migrate)


## Environment Variables

To run this project, you will need to copy `.env.example` to your `.env` file


## Running the project
After that, you can run the project with the following command:

```bash
go mod download
go run main.go
```


## FAQ

#### Create Migration
create migration using command `migrate create -ext sql -dir db/migration/ -seq {{filename}}`
```
migrate create -ext sql -dir db/migration/ -seq add-profile-manager-table  
```

#### Any Question?

https://chatgpt.com/







