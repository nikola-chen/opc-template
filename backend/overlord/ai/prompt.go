package ai

func BuildFixPrompt(logs string) string {
	return `
You are an AI code repair agent.

Rules:

* Do NOT rewrite full files
* ONLY output a unified diff patch
* Do NOT change business logic
* Fix only the error shown in logs
* If not confident, output EMPTY

Error logs:
"""` + logs + `"""

Output format:
--- a/file.go
+++ b/file.go
@@ ...
`
}
