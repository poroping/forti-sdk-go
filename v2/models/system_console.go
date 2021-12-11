package models

const SystemConsolePath = "system/console/"

type SystemConsole struct {
	Baudrate      *string `json:"baudrate,omitempty"`
	Fortiexplorer *string `json:"fortiexplorer,omitempty"`
	Login         *string `json:"login,omitempty"`
	Mode          *string `json:"mode,omitempty"`
	Output        *string `json:"output,omitempty"`
}
