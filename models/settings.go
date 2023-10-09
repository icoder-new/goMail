package models

type Settings struct {
	AppParams      Params
	PostgresParams PostgresSettings
	STMPParams     SMTPSettings
}

type Params struct {
	PortRun       string
	LogInfo       string
	LogError      string
	LogDebug      string
	LogWarning    string
	LogMaxSize    int
	LogMaxBackups int
	LogMaxAge     int
	LogCompress   bool
}

type PostgresSettings struct {
	User     string
	Password string
	Server   string
	Port     string
	Database string
	SSLMode  string
}

type SMTPSettings struct {
	Server   string
	Port     string
	Username string
	Password string
}
