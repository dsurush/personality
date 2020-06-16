package settings

// Settings app settings
type Settings struct {
	AppParams               Params                  `json:"app"`
	PostgresMegafonDbParams PostgresMegafonDbParams `json:"postgresMegafondb"`
	PostgresHamsoyaDbParams PostgresHamsoyaDbParams `json:"postgresHamsoyadb"`
}

// Params contains params of server meta data
type Params struct {
	ServerName string `json:"serverName"`
	PortRun    int    `json:"portRun"`
	LogFile    string `json:"logFile"`
}

// PostgresMegafonDbParams conteins params of postgresql db server
type PostgresMegafonDbParams struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type PostgresHamsoyaDbParams struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}
