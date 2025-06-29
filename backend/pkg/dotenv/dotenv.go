package dotenv

import (
	"os"
	"strings"
)

// reads the provided file and sets the key, value in os variable i.e. os.SetEnv(key, value)
func ReadDotEnv(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		arr := strings.SplitN(line, "=", 2)
		if len(arr) != 2 {
			continue
		}

		key := strings.TrimSpace(arr[0])
		value := strings.TrimSpace(arr[1])

		err := os.Setenv(key, value)

		if err != nil {
			return err
		}
	}

	return nil
}
