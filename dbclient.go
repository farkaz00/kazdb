package kazdb

import (
	"strings"

	"github.com/farkaz00/kazconfig"
	"github.com/farkaz00/kazmongo"
	"github.com/farkaz00/kazsql"
)

//DbClient generic interface for database clients based on the kaz standard
type DbClient interface {
	SelectOne(table string, selector interface{}, result interface{}) error
	Select(table string, selector interface{}, result interface{}) error
	Insert(table string, values interface{}) error
	Update(table string, selector interface{}, values interface{}) error
	Delete(table string, selector interface{}) error
	Close()
}

//NewDbClient factory method which returns a database client based on the dbtype
func NewDbClient(s *kazconfig.Settings, connection DbConnection) (DbClient, error) {
	var dbclient DbClient
	var err error
	dbtype := s.Get("dbtype")
	dbName := s.Get("dbname")
	switch dbtype = strings.ToUpper(dbtype); dbtype {
	case "MONGO", "MONGODB":
		dbclient, err = kazmongo.NewMongoClient(connection.(*kazmongo.MongoConnection), dbName)
	case "MYSQL":
		conn := connection.(*kazsql.MySQLConnection)
		dbclient, err = kazsql.NewMySQLClient(conn)
	default:
		dbclient = nil
	}
	return dbclient, err
}
