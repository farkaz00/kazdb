package kazdb

import (
	"strings"

	"github.com/farkaz00/kazconfig"
	"github.com/farkaz00/kazmongo"
	"github.com/farkaz00/kazsql"
)

//DbConnection represents a generic connection to a database engine
type DbConnection interface {
	GetConnString() string
	Close()
}

//NewDbConnection returns a connection object used to connect to the specific database engine
func NewDbConnection(s *kazconfig.Settings) *DbConnection {

	var dbconn DbConnection

	switch dbtype := strings.ToUpper(s.Get("dbtype")); dbtype {
	case "MONGO", "MONGODB":
		dbconn = kazmongo.NewMongoConnection(s)
	case "MYSQL":
		dbconn = kazsql.NewMySQLConnection(s)
	default:
		dbconn = nil
	}

	return &dbconn
}
