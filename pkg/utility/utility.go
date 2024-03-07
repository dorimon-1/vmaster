package utility

import (
	"strings"
)

func SplitArguments(args []string) map[string]string {
	arguments := make(map[string]string)

	for _, arg := range args {
		key, value := splitKeyValue(arg)
		arguments[key] = value
	}

	return arguments
}

func splitKeyValue(arg string) (string, string) {
	parts := strings.SplitN(arg, "=", 2) // Split at the first "="
	if len(parts) == 2 {
		return parts[0], parts[1] // Correctly return the key and value
	}
	// Handle the case where "=" is not found or the input is malformed
	return arg, ""
}
