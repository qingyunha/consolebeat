// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
}

var DefaultConfig = Config{}
