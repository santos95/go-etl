import process 

import (
	"context"
	"database/sql"

)

func GetLasExecution(db *sql.DB) string {

	var endDate string 

	// query data
	rows, err := db.Query("SELECT TOP 1 endDate FROM dbo.ETL_LOG " +
	"WHERE etlState = 'END' ORDER BY id DESC;")

	if err != nil {
		log.Fatalf("Failed to scan the row: %v", err)
	}

	// iterates over rows
	for rows.Next() {

		if err := rows.Scan(&endDate);  err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
	}

	return endDate
}