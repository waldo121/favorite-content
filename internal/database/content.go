package database

func InsertContent(url string) (int64, error) {
	db := GetConnection()
	defer db.Close()
	var id int64
	insertStatement := `INSERT INTO Items (uri) values ($1) RETURNING id;`
	err := db.QueryRow(insertStatement, url).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func SelectRandomContent() (uint64, string, error) {
	db := GetConnection()
	defer db.Close()
	var uri string
	var id uint64
	selectStatement := `SELECT id, uri FROM Items ORDER BY RANDOM() LIMIT 1;`
	err := db.QueryRow(selectStatement).Scan(&id, &uri)
	if err != nil {
		return 0, "", err
	}
	return id, uri, nil
}
