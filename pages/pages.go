package pages

import (
	"fmt"
	"os"
	"strings"
)

func Welcome(name string) (string, error) {
	filePath := "./pages/welcome/welcome.html"

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("no se pudo leer el archivo HTML: %w", err)
	}

	html := strings.Replace(string(content), "${{name}}", name, -1)

	return html, nil
}
