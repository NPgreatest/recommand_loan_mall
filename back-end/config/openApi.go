package config

type Keys struct {
	OpenApiBase   string `mapstructure:"open_api_base" json:"open_api_base" yaml:"open_api_base"`
	OpenApiKey    string `mapstructure:"open_api_key" json:"open_api_key" yaml:"open_api_key"`
	FineTuneModel string `mapstructure:"fine_tune_model" json:"fine_tune_model" yaml:"fine_tune_model"`
}
