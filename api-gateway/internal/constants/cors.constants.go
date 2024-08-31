package constants

var CORSHeaders = map[string]string{
	"Access-Control-Allow-Origin":  "*",
	"Access-Control-Allow-Methods": "POST, PUT, DELETE, OPTIONS",
	"Access-Control-Allow-Headers": "Content-Type, Authorization",
}

var AllowedMethods = []string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"OPTIONS",
}

var AllowedOrigins = []string{
	"*",
}