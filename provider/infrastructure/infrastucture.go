package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/dghubble/sling"
	_ "github.com/go-sql-driver/mysql" // Import mysql driver
	"github.com/hinha/api-box/provider"
	"github.com/hinha/api-box/provider/infrastructure/adapter"
	"os"
	"sync"
)

type Infrastructure struct {
	mysqlMutex  *sync.Once
	mysqlDB     *sql.DB
	mysqlConfig struct {
		username string
		password string
		host     string
		port     string
		dbname   string
	}
}

func Fabricate() (*Infrastructure, error) {
	i := &Infrastructure{
		mysqlMutex: &sync.Once{},
	}
	i.mysqlConfig.host = os.Getenv("MYSQL_HOST")
	i.mysqlConfig.username = os.Getenv("MYSQL_USERNAME")
	i.mysqlConfig.password = os.Getenv("MYSQL_PASSWORD")
	i.mysqlConfig.dbname = os.Getenv("MYSQL_DATABASE")
	i.mysqlConfig.port = os.Getenv("MYSQL_PORT")

	return i, nil
}

func (i *Infrastructure) Clone() {
	if i.mysqlDB != nil {
		_ = i.mysqlDB.Close()
	}
}

// MYSQL provide mysql interface
func (i *Infrastructure) MYSQL() (*sql.DB, error) {
	i.mysqlMutex.Do(func() {
		db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
			i.mysqlConfig.username,
			i.mysqlConfig.password,
			i.mysqlConfig.host,
			i.mysqlConfig.port,
			i.mysqlConfig.dbname))

		i.mysqlDB = db
	})

	return i.mysqlDB, nil
}

// Network provider
func (i *Infrastructure) GoogleOAuthNetwork() provider.GoogleOAuth {
	return adapter.AdapterGoogleNetwork(i.Sling())
}

// Sling return sling library used for network feature
func (i *Infrastructure) Sling() *sling.Sling {
	return sling.New()
}

func (i *Infrastructure) Close() {
	if i.mysqlDB != nil {
		_ = i.mysqlDB.Close()
	}
}

func (i *Infrastructure) DB() (provider.DB, error) {
	db, err := i.MYSQL()
	if err != nil {
		return nil, err
	}

	return adapter.AdaptSQL(db), nil
}
