// +build go1.11

package bugs

import (
	"database/sql"
	"testing"
)

func TestStats(t *testing.T) {
	db, err := sql.Open("mysql", "root@/")
	if err != nil {
		t.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)

	t.Logf("%+v", db.Stats())

	for i := 0; i < 10; i++ {
		if _, err = db.Exec("SELECT 1"); err != nil {
			t.Fatal(err)
		}
	}

	t.Logf("%+v", db.Stats())
	if db.Stats().MaxIdleClosed > 5 {
		t.Errorf("expected MaxIdleClosed to be reasonable")
	}
}
