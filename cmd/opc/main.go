package main

import (
	"fmt"
	"opc-template/backend/overlord"
	"opc-template/backend/overlord/ai"
	"opc-template/pkg/generator"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	cmd := os.Args[1]

	// Init AI Client early if needed, or lazily.
	// For generate/heal we definitely need it.
	var aiClient ai.Client
	var err error

	if cmd == "generate" || cmd == "heal" {
		aiClient, err = ai.NewOpenAIClient()
		if err != nil {
			fmt.Printf("Warning: Failed to initialize AI client: %v\n", err)
		}
	}

	switch cmd {
	case "help", "--help", "-h":
		printHelp()
		os.Exit(0)

	case "explain":
		printExplain()
		os.Exit(0)

	case "design":
		fmt.Println("Design phase Check:")
		if _, err := os.Stat("design/schema.json"); os.IsNotExist(err) {
			fmt.Println("Creating default design/schema.json...")
			os.MkdirAll("design", 0755)
			defaultSchema := `{
  "appName": "MyNewApp",
  "version": "0.1.0",
  "models": [
    {
      "name": "Note",
      "fields": [
        { "name": "id", "type": "string", "primary": true },
        { "name": "content", "type": "string" }
      ]
    }
  ],
  "api": { "basePath": "/api/v1" },
  "frontend": { "type": "web" }
}`
			os.WriteFile("design/schema.json", []byte(defaultSchema), 0644)
			fmt.Println("Created design/schema.json. Please edit it to define your app.")
		} else {
			fmt.Println("Design schema exists at design/schema.json. generating code will use this schema.")
		}
		os.Exit(0)

	case "generate":
		fmt.Println("Generate phase:")
		if aiClient == nil {
			fmt.Println("Error: AI client required for generation. Check OPENAI_API_KEY.")
			os.Exit(1)
		}

		gen := generator.NewGenerator(aiClient)
		if err := gen.Run(); err != nil {
			fmt.Printf("Generation failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Generation complete.")
		os.Exit(0)

	case "run":
		fmt.Println("Run phase:")
		fmt.Println("- Start system")
		fmt.Println("- Write logs to runtime/logs/")
		os.Exit(0)

	case "heal":
		fmt.Println("Heal phase:")

		if aiClient == nil {
			fmt.Println("Overlord will run in degraded mode (human intervention likely required).")
		}

		o := overlord.Overlord{
			MaxRetry: 3,
			AIClient: aiClient,
		}
		result := o.Run()

		switch result {
		case overlord.Success:
			fmt.Println("System healed successfully")
			os.Exit(0)
		case overlord.NeedHuman:
			fmt.Println("Healing requires human intervention")
			os.Exit(2)
		default:
			os.Exit(1)
		}

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Print(`OPC CLI

Usage:
  opc <command>

Commands:
  design     Generate or update design schema
  generate   Generate code from schema
  run        Run the system
  heal       Try to heal system errors
  explain    Explain what each command does
  help       Show this help message
`)
}

func printExplain() {
	fmt.Print(`OPC Pipeline Explanation

design:
  - Prepare design/schema.json

generate:
  - Generate code based on schema
  - frontend/, backend/, infra/ may be overwritten

run:
  - Start the generated system
  - Logs go to runtime/logs/

heal:
  - Read logs
  - Try to fix errors
  - Retry run
  - If exit code = 2, human decision is required
`)
}
