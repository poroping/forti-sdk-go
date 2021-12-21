package models

const SystemSpeedTestServerPath = "system/speed-test-server/"

type SystemSpeedTestServer struct {
	Host      *[]SystemSpeedTestServerHost `json:"host,omitempty"`
	Name      *string                      `json:"name,omitempty"`
	Timestamp *float64                     `json:"timestamp,omitempty"`
}

type SystemSpeedTestServerHost struct {
	Id       *float64 `json:"id,omitempty"`
	Ip       *string  `json:"ip,omitempty"`
	Password *string  `json:"password,omitempty"`
	Port     *float64 `json:"port,omitempty"`
	User     *string  `json:"user,omitempty"`
}
