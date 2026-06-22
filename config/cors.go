package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"http://localhost:5174",
	"belatihan-production-33a2.up.railway.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
