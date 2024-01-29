package config

var (
	CORSAllowedMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	CORSAllowedOrigins = []string{
		"http://localhost:8080",
		"http://localhost:5400",
	}
	CORSAllowedHeader = []string{"Origin", "Content-Type"}
	CORSExposeHeader  = []string{"Content-Length"}
	PORTApp           = "5400"
	Host              = "127.0.0.1"
	User              = "root"
	Password          = "root"
	DBName            = "order"
	PORTDB            = "5432"
)
