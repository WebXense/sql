# sql

This package provides a simple way to use [gorm](https://gorm.io/) in your project.

## Installation

```bash
go get -u github.com/WebXense/sql
```

## Packages

There are four packages of this project:

-   `sql` - The main package, which provides the statement builder, read and write operations.
-   `sql/stm` - The statement builder object. The statement builder in `sql` package is based on this package.
-   `sql/conn` - Provides the connection methods for MySQL and SQLite.
-   `sql/aes` - Provides the AES encryption and decryption methods.

## Connections

### MySQL

```go
db, err := conn.MySQL("127.0.0.1", "3306", "username", "password", "database", false)
```

### SQLite

File-path database:

```go
db, err := conn.SQLite("database.db", false)
```

In-memory database:

```go
db, err := conn.SQLiteInMemory(false)
```

## Query

### FindOne

```go
user, err := sql.FindOne[User](db, sql.Eq("id", 1))
```

### FindAll

```go
users, err := sql.FindAll[User](db, sql.Gt("id", 1), &sql.Pagination{
    Page: 1, // start from 1
    Size: 10,
}, &sql.Sort{
    SortBy: "id",
    Asc:    false,
})
```

### Count

```go
count, err := sql.Count[User](db, sql.Gt("id", 1))
```

## Write

### Create

```go
user := User{
    Name: "John",
    Age:  18,
}

created, err := sql.Create(db, &user)
```

### Update

```go
user := User{
    ID:   1,
    Name: "John",
    Age:  18,
}

updated, err := sql.Update(db, &user)
```

### Delete

```go
user := User{
    ID: 1,
}

deleted, err := sql.Delete(db, &user)
```

### DeleteBy

```go
err := sql.DeleteBy[User](db, sql.Eq("id", 1))
```

## Encryption

### Encrypt

```go
type User struct {
    ID    int64  `gorm:"primaryKey"`
    Name  string
    Age   int64
    Phone []byte
}

encryptKey := "1234567890123456"

user := User{
    ID:    1,
    Name:  "John",
    Age:   18,
    Phone: aes.Encrypt("51234567", encryptKey),
}
```

### Decrypt

```go
var phone string = aes.Decrypt(user.Phone, encryptKey)
```
