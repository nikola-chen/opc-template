package generator

type Schema struct {
	AppName  string   `json:"appName"`
	Version  string   `json:"version"`
	Models   []Model  `json:"models"`
	API      API      `json:"api"`
	Frontend Frontend `json:"frontend,omitempty"`
}

type Frontend struct {
	Type string `json:"type"` // "wechat" or "web"
}

type Model struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Primary bool   `json:"primary,omitempty"`
}

type API struct {
	BasePath string `json:"basePath"`
}
