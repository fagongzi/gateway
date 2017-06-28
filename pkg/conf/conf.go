package conf

// Conf config struct
type Conf struct {
	LogLevel string `json:"-"`

	Addr    string `json:"addr"`
	MgrAddr string `json:"mgrAddr"`

	RegistryAddr string `json:"registryAddr"`
	Prefix       string `json:"prefix"`

	Filers []*FilterSpec `json:"filers"`

	// Maximum number of connections which may be established to server
	MaxConns int `json:"maxConns"`
	// MaxConnDuration Keep-alive connections are closed after this duration.
	MaxConnDuration int `json:"maxConnDuration"`
	// MaxIdleConnDuration Idle keep-alive connections are closed after this duration.
	MaxIdleConnDuration int `json:"maxIdleConnDuration"`
	// ReadBufferSize Per-connection buffer size for responses' reading.
	ReadBufferSize int `json:"readBufferSize"`
	// WriteBufferSize Per-connection buffer size for requests' writing.
	WriteBufferSize int `json:"writeBufferSize"`
	// ReadTimeout Maximum duration for full response reading (including body).
	ReadTimeout int `json:"readTimeout"`
	// WriteTimeout Maximum duration for full request writing (including body).
	WriteTimeout int `json:"writeTimeout"`
	// MaxResponseBodySize Maximum response body size.
	MaxResponseBodySize int `json:"maxResponseBodySize"`

	// EnablePPROF enable pprof
	EnablePPROF bool `json:"enablePPROF"`
	// PPROFAddr pprof addr
	PPROFAddr string `json:"pprofAddr,omitempty"`
}

// FilterSpec filter spec
type FilterSpec struct {
	Name               string `json:"name"`
	External           bool   `json:"external,omitempty"`
	ExternalPluginFile string `json:"externalPluginFile,omitempty"`
}
