package swaggo

type SwaggerData struct {
	Swagger     string                           `json:"swagger"`
	Info        map[string]interface{}           `json:"info"`
	Paths       map[string]SwaggerApiData        `json:"paths"`
	Definitions map[string]SwaggerDefinitionData `json:"definitions"`
}

type SwaggerApiData map[string]SwaggerApiMethodData

type SwaggerApiMethodData struct {
	Description string                      `json:"description"`
	Consumes    []string                    `json:"consumes"`
	Produces    []string                    `json:"produces"`
	Tags        []string                    `json:"tags"`
	Summary     string                      `json:"summary"`
	Parameters  []SwaggerApiMethodParameter `json:"parameters"`
}

type SwaggerApiMethodParameter struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Name        string `json:"name"`
	In          string `json:"in"`
	Required    bool   `json:"required"`
}

type SwaggerDefinitionData struct {
	Type       string                                   `json:"type"`
	Properties map[string]SwaggerDefinitionPropertyData `json:"properties"`
}

type SwaggerDefinitionPropertyData struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}
