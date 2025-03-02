package env

import "os"

var IsVerbose bool

func init() {
	value, exists := os.LookupEnv("VERBOSE")

	IsVerbose = exists && value != ""
}
