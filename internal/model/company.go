package model

// RoundConfig holds the system prompt instructions for a specific interview round.
type RoundConfig struct {
	Title        string `yaml:"title"`
	Instructions string `yaml:"instructions"`
}

// CompanyProfile defines how a company conducts its interviews.
type CompanyProfile struct {
	Name            string                 `yaml:"name"`
	DisplayName     string                 `yaml:"display_name"`
	RoleDescription string                 `yaml:"role_description"`
	Rounds          map[int]*RoundConfig   `yaml:"rounds"`
	EvaluationRubric string               `yaml:"evaluation_rubric"`
}
