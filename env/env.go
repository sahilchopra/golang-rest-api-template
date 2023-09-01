package env

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	if GoEnv() == "development" {
		if err := godotenv.Load(fmt.Sprintf("%s/../.env.development", basepath)); err != nil {
			fmt.Println(err)
		}
	}
	if GoEnv() == "test" {
		if err := godotenv.Load(fmt.Sprintf("%s/../.env.test", basepath)); err != nil {
			fmt.Println(err)
		}
	}

}

// GoEnv returns the Go environment
func GoEnv() string {
	return get("GO_ENV", "development")
}

// LogLevel returns the log level
func LogLevel() string {
	return get("LOG_LEVEL", "info")
}

func get(name string, defaultVal string) string {
	val := os.Getenv(name)

	if len(val) > 0 {
		return val
	}
	return defaultVal
}

func getInt(name string, defaultVal int) int {
	val := os.Getenv(name)

	if len(val) == 0 {
		return defaultVal
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return intVal
}

func mustGet(name string) string {
	val, present := os.LookupEnv(name)
	if !present {
		panic(fmt.Sprintf("envvar %#v is not present", name))
	}
	return val
}

func getBool(name string, defaultVal bool) bool {
	val := os.Getenv(name)

	if len(val) == 0 {
		return defaultVal
	}

	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		panic(err)
	}
	return boolVal
}

// Postgresql

// PostgresUser returns Postgres User
func PostgresUser() string {
	return mustGet("POSTGRES_USER")
}

// PostgresPassword returns Postgres Password
func PostgresPassword() string {
	return mustGet("POSTGRES_PASSWORD")
}

// PostgresDBName returns Postgres DB Name
func PostgresDBName() string {
	return mustGet("POSTGRES_DB_NAME")
}

// PostgresHost returns Postgres Host
func PostgresHost() string {
	return mustGet("POSTGRES_HOST")
}

// PostgresPort returns Postgres Port
func PostgresPort() string {
	return mustGet("POSTGRES_PORT")
}
