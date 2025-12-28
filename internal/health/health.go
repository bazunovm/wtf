package health

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"wtf/internal/ai"
)

func Run() {
	fmt.Println("WTF health check\n")

	checkBinary()
	checkRules()
	checkAPIKey()
	checkAI()

	fmt.Println("\nHealth check completed")
}

func ok(msg string) {
	fmt.Println("✅", msg)
}

func warn(msg string) {
	fmt.Println("⚠️", msg)
}

func fail(msg string) {
	fmt.Println("❌", msg)
}

func checkBinary() {
	path, err := os.Executable()
	if err != nil {
		fail("Cannot determine WTF binary location")
		return
	}
	ok("Binary found: " + path)
}

func checkRules() {
	paths := []string{
		"./rules.yaml",
		"/usr/local/share/wtf/rules.yaml",
		filepath.Join(os.Getenv("HOME"), ".config/wtf/rules.yaml"),
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			ok("rules.yaml found at " + p)
			return
		}
	}

	warn("rules.yaml not found (AI fallback will be used)")
}

func checkAPIKey() {
	key := os.Getenv("WTF_AI_API_KEY")
	if key == "" {
		warn("WTF_AI_API_KEY not set")
		return
	}
	ok("WTF_AI_API_KEY is set")
}

func checkAI() {
	key := ai.LoadAPIKey()
	//key := os.Getenv("WTF_AI_API_KEY")
	if key == "" {
		return
	}

	client := &http.Client{Timeout: 3 * time.Second}
	req, err := http.NewRequest(
		"GET",
		"https://generativelanguage.googleapis.com/v1/models?key="+key,
		nil,
	)
	if err != nil {
		warn("Failed to create AI request")
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		warn("Cannot reach Gemini API")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		ok("Gemini API reachable")
	} else {
		warn("Gemini API returned status " + resp.Status)
	}
}
