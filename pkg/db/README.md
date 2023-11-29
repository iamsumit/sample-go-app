## Database

This package connects to a SQL based database, mysql for now. pgSQL will be added later on.

It returns `*sql.DB` instance which can be used to query to the configured database anytime in the application.

### Example

```go
sqlDB, err := db.New(db.Config{
  Type:     db.MySQL,
  Name:     configuration.MySQL.Name,
  User:     configuration.MySQL.User,
  Password: configuration.MySQL.Password,
  Port:     configuration.MySQL.Port,
  Host:     configuration.MySQL.Host,
})
if err != nil {
  return err
}

defer func() {
  _ = sqlDB.Close()
}()

rows, err := sqlDB.Query("SELECT id, name from users")
if err != nil {
  // Do something with error
}
defer rows.Close()

for rows.Next() {
  var (
    id   int64
    name string
  )
  if err := rows.Scan(&id, &name); err != nil {
    // Do something with the error.
  }

  log.Printf("id %d name is %s\n", id, name)
}
```
