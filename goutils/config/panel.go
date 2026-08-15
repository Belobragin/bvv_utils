package config

type WfPanelConfig struct {
	ApiServerBaseUrl string `conf:"env:API_SERVER_BASE_URL"`
	UserBaseUrl      string `conf:"env:USER_SERVICE_BASE_URL"`
	AuthBaseUrl      string `conf:"env:AUTH_SERVICE_BASE_URL"`
}
