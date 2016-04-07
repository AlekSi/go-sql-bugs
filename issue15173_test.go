package bugs

import (
	"database/sql"
	"log"
	"testing"
)

type (
	Int64  int64
	Int32  int32
	String string
)

type User struct {
	ID   int64
	Age  int32
	Name string
}

type UserT struct {
	ID   Int64
	Age  Int32
	Name String
}

func TestDefaultParameterConverter(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE users (id integer, age integer, name varchar)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO users (id, age, name) VALUES (1, 29, "Alexey")`)
	if err != nil {
		log.Fatal(err)
	}

	// scan into predeclared types
	{
		var u User
		err = db.QueryRow("SELECT id, age, name FROM users").Scan(&u.ID, &u.Age, &u.Name)
		if err != nil {
			t.Error(err)
		}
		e := User{1, 29, "Alexey"}
		if u != e {
			t.Errorf("expected %+v, got %+v", e, u)
		}
	}

	// scan into types with predeclared underlying types
	{
		var ut UserT
		err = db.QueryRow("SELECT id, age, name FROM users").Scan(&ut.ID, &ut.Age, &ut.Name)
		if err != nil {
			t.Error(err)
		}
		e := UserT{1, 29, "Alexey"}
		if ut != e {
			t.Errorf("expected %+v, got %+v", e, ut)
		}
	}
}
