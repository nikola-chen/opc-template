package generator

import (
	"encoding/json"
	"fmt"
	"opc-template/backend/overlord/ai"
	"os"
)

type Generator struct {
	AIClient ai.Client
}

func NewGenerator(client ai.Client) *Generator {
	return &Generator{
		AIClient: client,
	}
}

func (g *Generator) Run() error {
	// 1. Read Schema
	schemaBytes, err := os.ReadFile("design/schema.json")
	if err != nil {
		return fmt.Errorf("failed to read schema: %w", err)
	}

	var schema Schema
	if err := json.Unmarshal(schemaBytes, &schema); err != nil {
		return fmt.Errorf("failed to parse schema: %w", err)
	}

	fmt.Printf("Generating app: %s (v%s)\n", schema.AppName, schema.Version)

	// 2. Generate Backend
	if err := g.generateBackend(schema); err != nil {
		return fmt.Errorf("backend generation failed: %w", err)
	}

	// 3. Generate Frontend
	if err := g.generateFrontend(schema); err != nil {
		return fmt.Errorf("frontend generation failed: %w", err)
	}

	return nil
}

func (g *Generator) generateBackend(schema Schema) error {
	fmt.Println("-> Generating Backend (Go/Gin)...")

	prompt := fmt.Sprintf(`
You are an expert Go backend developer.
Generate a main.go file for a Go backend using Gin framework based on this schema:
%s

Requirements:
- Use "github.com/gin-gonic/gin"
- Implement valid structs for Models
- Implement simple in-memory CRUD handlers for each model
- The server should run on port 8080
- Output ONLY the Go code, no markdown code fences if possible, or inside a single code block.
`, mustMarshal(schema))

	req := ai.GenerateRequest{
		Prompt: prompt,
	}

	code, err := g.AIClient.Generate(req)
	if err != nil {
		return err
	}

	code = cleanCode(code)

	// Write to backend/main.go
	// Ensure directory exists
	os.MkdirAll("backend", 0755)
	return os.WriteFile("backend/main.go", []byte(code), 0644)
}

func (g *Generator) generateFrontend(schema Schema) error {
	// Default to 'wechat' if not specified
	frontendType := schema.Frontend.Type
	if frontendType == "" {
		frontendType = "wechat"
	}

	if frontendType == "wechat" {
		return g.generateWeChatFrontend(schema)
	} else if frontendType == "web" {
		return g.generateWebFrontend(schema)
	}

	return fmt.Errorf("unknown frontend type: %s", frontendType)
}

func (g *Generator) generateWebFrontend(schema Schema) error {
	fmt.Println("-> Generating Web Frontend (HTML/JS)...")

	prompt := fmt.Sprintf(`
You are an expert Frontend developer.
Generate a single index.html file that implements a UI for this schema:
%s

Requirements:
- Use standard HTML/JS (no framework build steps needed for this MV)
- Connect to the backend API at http://localhost:8080%s
- List items and allow adding new items
- Modern, clean UI style
- Output ONLY the HTML code.
`, mustMarshal(schema), schema.API.BasePath)

	req := ai.GenerateRequest{
		Prompt: prompt,
	}

	code, err := g.AIClient.Generate(req)
	if err != nil {
		return err
	}

	code = cleanCode(code)

	os.MkdirAll("frontend", 0755)
	return os.WriteFile("frontend/index.html", []byte(code), 0644)
}

func (g *Generator) generateWeChatFrontend(schema Schema) error {
	fmt.Println("-> Generating WeChat Mini Program...")

	// Create directory structure
	baseDir := "frontend/miniprogram"
	os.MkdirAll(baseDir+"/pages/index", 0755)

	// 1. Generate app.json
	appJson := `{
  "pages": [
    "pages/index/index"
  ],
  "window": {
    "backgroundTextStyle": "light",
    "navigationBarBackgroundColor": "#fff",
    "navigationBarTitleText": "` + schema.AppName + `",
    "navigationBarTextStyle": "black"
  },
  "style": "v2",
  "sitemapLocation": "sitemap.json"
}`
	os.WriteFile(baseDir+"/app.json", []byte(appJson), 0644)

	// 2. Generate project.config.json
	projectConfig := `{
  "miniprogramRoot": "miniprogram/",
  "projectname": "` + schema.AppName + `",
  "description": "AI Generated OPC Mini Program",
  "appid": "touristappid",
  "setting": {
    "urlCheck": false,
    "es6": true,
    "enhance": true,
    "postcss": true,
    "preloadBackgroundData": false,
    "minified": true,
    "newFeature": true,
    "coverView": true,
    "nodeModules": false,
    "autoAudits": false,
    "showShadowRootInWxmlPanel": true,
    "scopeDataCheck": false,
    "uglifyFileName": false,
    "checkInvalidSitemap": true
  },
  "compileType": "miniprogram",
  "libVersion": "2.19.4",
  "packOptions": {
    "ignore": [],
    "include": []
  }
}`
	// Note: project.config.json usually sits at the root of the frontend or project root,
	// but standard WeChat structure puts miniprogram code in miniprogram/
	// For simplicity, we put project.config.json in frontend/ so you can open frontend/ dir in IDE.
	os.WriteFile("frontend/project.config.json", []byte(projectConfig), 0644)

	// 3. Generate WXML via AI
	fmt.Println("   Generating WXML...")
	if err := g.genFile(schema, "WXML", "frontend/miniprogram/pages/index/index.wxml", `
Generate the WXML (WeChat XML) content for the main index page.
Display a list of data models and a form to add new ones.
Use standard WeChat Mini Program components (view, button, input, etc.).
Output ONLY the WXML code.
`); err != nil {
		return err
	}

	// 4. Generate JS via AI
	fmt.Println("   Generating JS...")
	if err := g.genFile(schema, "JS", "frontend/miniprogram/pages/index/index.js", `
Generate the JS content for the main index page.
- Page() data should include list of models
- onLoad: fetch data from http://localhost:8080`+schema.API.BasePath+`
- Implement function to add new item (POST request)
- Use wx.request
- Output ONLY the JS code.
`); err != nil {
		return err
	}

	// 5. Generate WXSS via AI
	fmt.Println("   Generating WXSS...")
	if err := g.genFile(schema, "WXSS", "frontend/miniprogram/pages/index/index.wxss", `
Generate the WXSS (WeChat Style Sheets) content for simple clean styling of the list and form.
Output ONLY the WXSS code.
`); err != nil {
		return err
	}

	return nil
}

func (g *Generator) genFile(schema Schema, fileType, path, promptInstruction string) error {
	prompt := fmt.Sprintf(`
You are an expert WeChat Mini Program developer.
Based on this schema:
%s

%s
`, mustMarshal(schema), promptInstruction)

	req := ai.GenerateRequest{
		Prompt: prompt,
	}

	code, err := g.AIClient.Generate(req)
	if err != nil {
		return err
	}

	code = cleanCode(code)
	return os.WriteFile(path, []byte(code), 0644)
}

func mustMarshal(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func cleanCode(raw string) string {
	if len(raw) > 3 && raw[:3] == "```" {
		for i, c := range raw {
			if c == '\n' {
				raw = raw[i+1:]
				break
			}
		}
		if len(raw) > 3 && raw[len(raw)-3:] == "```" {
			raw = raw[:len(raw)-3]
		}
	}
	// Also strip specific language blocks like ```xml or ```javascript
	if len(raw) > 0 && raw[0] == '`' {
		// fallback aggressive strip if the loop above missed (e.g. no newline immediately)
		// But usually the loop handles ```xml\n
	}
	return raw
}
