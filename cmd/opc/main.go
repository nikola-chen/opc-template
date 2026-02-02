package main

import (
	"fmt"
	"opc-template/backend/overlord"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "help", "--help", "-h":
		printHelp()
		os.Exit(0)

	case "explain":
		printExplain()
		os.Exit(0)

	case "design":
		fmt.Println("Design phase:")
		fmt.Println("- Generate or update design/schema.json")
		os.Exit(0)

	case "generate":
		fmt.Println("Generate phase:")
		fmt.Println("- Read design/schema.json")
		fmt.Println("- Generate frontend/, backend/, infra/")
		os.Exit(0)

	case "run":
		fmt.Println("Run phase:")
		fmt.Println("- Start system")
		fmt.Println("- Write logs to runtime/logs/")
		os.Exit(0)

	case "heal":
		fmt.Println("Heal phase:")
		o := overlord.Overlord{MaxRetry: 3}
		result := o.Run()

		switch result {
		case overlord.Success:
			fmt.Println("System healed successfully")
			os.Exit(0)
		case overlord.NeedHuman:
			fmt.Println("Healing requires human intervention")
			os.Exit(2) //明确：需要人工介入
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
	fmt.Println(`OPC CLI

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
	fmt.Println(`OPC Pipeline Explanation

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
