package main

import (
	"context"

	"github.com/Nathene/SyntheticMonitor/cmd"
	_ "modernc.org/sqlite"
)

func main() {
	cmd.Execute(context.Background())
}

// func main() {
// 	ctx := context.Background()
// 	dbFilename := "test.db" // Define the filename once

// 	// 1. Connect Ent to the disk file, ensuring FKs are enabled
// 	dsn := dbFilename + "?_pragma=foreign_keys(1)" // Append pragma to filename
// 	// Or use: dsn := "file:" + dbFilename + "?_pragma=foreign_keys(1)"
// 	client, err := ent.Open("sqlite", dsn)
// 	if err != nil {
// 		log.Fatalf("Ent Open failed: %v", err)
// 	}
// 	defer client.Close()
// 	log.Println("Ent client connected to", dbFilename, "(with FK Pragma)")

// 	// 2. Create schema in the disk file (should work now)
// 	if err := client.Schema.Create(ctx); err != nil {
// 		log.Fatalf("Ent Schema Create failed: %v", err)
// 	}
// 	log.Println("Ent schema created/verified in", dbFilename)

// 	// 3. Connect probesql to the SAME disk file
// 	// NOTE: probesql does NOT need the pragma in its connection string,
// 	// as the pragma affects the DB state for the connection pool used by Ent.
// 	// However, ensure probesql connects cleanly to the file.
// 	db, err := probesql.NewSQLite(ctx, dbFilename)
// 	if err != nil {
// 		log.Fatalf("probesql NewSQLite failed: %v", err)
// 	}
// 	defer db.Close()

// 	if err := db.Connect(); err != nil {
// 		log.Fatalf("probesql Connect failed: %v", err)
// 	}
// 	log.Println("probesql connected to", dbFilename)

// 	// 4. Query the disk file (table should exist now)
// 	_, err = db.Query("SELECT * FROM probes")
// 	if err != nil {
// 		log.Fatalf("probesql Query failed: %v", err)
// 	}
// }
