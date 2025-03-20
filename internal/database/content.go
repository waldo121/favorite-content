package database

func InsertContent(videoId string, platform string) (int64, error) {
	db := GetConnection()
	defer db.Close()
	var id int64
	insertStatement := `INSERT INTO Items (video_id, platform) values ($1, $2) RETURNING id;`
	err := db.QueryRow(insertStatement, videoId, platform).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func SelectRandomContent() (uint64, string, string, error) {
	db := GetConnection()
	defer db.Close()
	var videoId string
	var id uint64
	var platform string
	selectStatement := `SELECT id, video_id, platform FROM Items ORDER BY RANDOM() LIMIT 1;`
	err := db.QueryRow(selectStatement).Scan(&id, &videoId, &platform)
	if err != nil {
		return 0, "", "", err
	}
	return id, videoId, platform, nil
}
