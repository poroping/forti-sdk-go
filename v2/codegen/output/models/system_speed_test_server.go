package models

const SystemSpeedTestServerPath = "system/speed-test-server/"

type SystemSpeedTestServer struct {
	Host      []SystemSpeedTestServerHost `json:"host,omitempty"`
	Name      *string                     `json:"name,omitempty"`
	Timestamp *int64                      `json:"timestamp,omitempty"`
}

type SystemSpeedTestServerHost struct {
	Id       *int64  `json:"id,omitempty"`
	Ip       *string `json:"ip,omitempty"`
	Password *string `json:"password,omitempty"`
	Port     *int64  `json:"port,omitempty"`
	User     *string `json:"user,omitempty"`
}
