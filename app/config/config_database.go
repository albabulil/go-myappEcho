package config

const (
	DB_HOST            string     = "DB_HOST"
	DB_PORT            string     = "DB_PORT"
	DB_USER            string     = "DB_USER"
	DB_PASS            string     = "DB_PASS"
	DB_NAME            string     = "DB_NAME"
	DB_SSL_MODE        string     = "DB_SSL_MODE"
	DB_TIMEZONE        string     = "DB_TIMEZONE"
	QUERY_LOGIC_IN     QueryLogic = "IN"
	QUERY_LOGIC_AND    QueryLogic = "AND"
	QUERY_LOGIC_OR     QueryLogic = "OR"
	QUERY_LOGIC_NOT    QueryLogic = "NOT"
	QUERY_LOGIC_LIKE   QueryLogic = "LIKE"
	QUERY_LOGIC_NOT_IN QueryLogic = "NOT IN"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Pass     string
	DbName   string
	SSLMode  string
	Timezone string
}
