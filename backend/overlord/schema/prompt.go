package schema

func BuildSchemaFixPrompt(schema string, logs string) string {
	return `
You are a software architect AI.

Rules:
- You must NOT change business meaning
- You must ONLY modify schema.json
- You must output a JSON Patch (RFC 6902)
- If unsure, output empty array []

Current schema:
"""` + schema + `"""

Error logs:
"""` + logs + `"""

Output:
[
  { "op": "add|replace|remove", "path": "/...", "value": ... }
]
`
}
