# ecom
A rest API written in Go.


Migration CLI 
```
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Migration Create
```
migrate create -ext sql -dir db/migrations create_users_table
```