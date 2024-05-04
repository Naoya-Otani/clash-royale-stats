package event

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// EventHandler is an HTTP Cloud Function.
func EventHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		var err error
		db, err = initSocketConnectionPool()
		if err != nil {
			log.Printf("db connection error \n")
			fmt.Fprintf(w, err.Error())
			return
		}
		defer db.Close()
		log.Printf("db connection success \n")
		response, err := requestDailyRankingData()
		if err != nil {
			log.Fatal(err)
		}
		insertDailyRankingLog(*response)
	}
	if r.Method == http.MethodPost {
		io.WriteString(w, "This is a post request")
	}
}

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Printf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

// initSocketConnectionPool initializes a Unix socket connection pool for
// a Cloud SQL instance of MySQL.
func initSocketConnectionPool() (*sql.DB, error) {
	// [START cloud_sql_mysql_databasesql_create_socket]
	var (
		dbUser                 = mustGetenv("DB_USER")
		dbPwd                  = mustGetenv("DB_PASS")
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME")
		dbName                 = mustGetenv("DB_NAME")
	)

	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	// [START_EXCLUDE]
	configureConnectionPool(dbPool)
	// [END_EXCLUDE]

	return dbPool, nil
	// [END cloud_sql_mysql_databasesql_create_socket]
}

// configureConnectionPool sets database connection pool properties.
// For more information, see https://golang.org/pkg/database/sql
func configureConnectionPool(dbPool *sql.DB) {
	// [START cloud_sql_mysql_databasesql_limit]

	// Set maximum number of connections in idle connection pool.
	dbPool.SetMaxIdleConns(5)

	// Set maximum number of open connections to the database.
	dbPool.SetMaxOpenConns(7)

	// [END cloud_sql_mysql_databasesql_limit]

	// [START cloud_sql_mysql_databasesql_lifetime]

	// Set Maximum time (in seconds) that a connection can remain open.
	dbPool.SetConnMaxLifetime(1800)

	// [END cloud_sql_mysql_databasesql_lifetime]
}

var (
	myPlayerTag = "#RU2RQGU"
	myName = "早稲田大学第一文学部"
	myRank = 0
	myEloRating = 0
)


func requestDailyRankingData() (*RankingPlayersResponse, error) {
	client := &http.Client{}
	baseUrl := "https://proxy.royaleapi.dev/v1/locations/"
	countryCode := "57000122" // 日本
	limit := "1000"
	requestUrl := baseUrl + countryCode + "/pathoflegend/players?limit=" + limit
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	apiToken := os.Getenv("API_BEARER_TOKEN")
	req.Header.Add("Authorization", "Bearer "+apiToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}
	return parsePlayerData(body)
}

func parsePlayerData(body []byte) (*RankingPlayersResponse, error) {
	var response RankingPlayersResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("parsing player data: %w", err)
	}
	return &response, nil
}

func insertDailyRankingLog(rankingPlayersResponse RankingPlayersResponse) {
	var err error
	var isFound bool
	for _, player := range rankingPlayersResponse.Items {
		if player.Tag == myPlayerTag {
			query := "INSERT INTO daily_ranking_logs (`name`, `tag`, `rank`, `elo_rating`) VALUES (?, ?, ?, ?)"
			_, err = db.Exec(query, player.Name, player.Tag, player.Rank, player.EloRating)
			if err != nil {
				log.Printf("Error inserting data: %s\n", err)
				return
			}
			isFound = true
			log.Printf("Insert successful for player: %s\n", player.Name)
		}
	}
	if !isFound {
		log.Printf("Player not found: %s\n", myName)
			query := "INSERT INTO daily_ranking_logs (`name`, `tag`, `rank`, `elo_rating`) VALUES (?, ?, ?, ?)"
			_, err = db.Exec(query, myName, myPlayerTag, 0, myEloRating)
			if err != nil {
				log.Printf("Error inserting data: %s\n", err)
				return
			}
			log.Printf("Insert successful for player: %s\n", myName)
	}
}