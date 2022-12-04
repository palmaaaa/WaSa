package database

// SetName is an example that shows you how to execute insert/update
func (db *appdbimpl) SetName(name string) error {
	_, err := db.c.Exec("INSERT INTO example_table (id, name) VALUES (1, ?)", name)
	return err
}

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetName() (string, error) {
	var name string
	err := db.c.QueryRow("SELECT name FROM example_table WHERE id=1").Scan(&name)
	return name, err
}
