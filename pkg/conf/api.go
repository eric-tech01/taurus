package conf

import "github.com/spf13/cast"

var defaultConfiguration = New()

// TODO: 未实现 OnChange 注册change回调函数
func OnChange(fn func(*Configuration)) {
	defaultConfiguration.OnChange(fn)
}

func OnLoaded(fn func(*Configuration)) {
	defaultConfiguration.OnLoaded(fn)
}

// LoadFromDataSource load configuration from data source
// if data source supports dynamic config, a monitor goroutinue
// would be
func Load(filepath string) error {
	return defaultConfiguration.Load(filepath)
}

// Reset resets all to default settings.
func Reset() {
	defaultConfiguration = New()
}

// Get returns an interface. For a specific value use one of the Get____ methods.
func Get(key string) interface{} {
	return defaultConfiguration.config.Get(key)
}

// Exists returns whether key exists
func Exists(key string) bool {
	return defaultConfiguration.config.Get(key) != nil
}

// Set set config value for key
func Set(key string, val interface{}) {
	defaultConfiguration.config.Set(key, val)
}

// Flush conf to file
func Flush() error {
	return defaultConfiguration.Flush()
}

// GetString returns the value associated with the key as a string with default defaultConfiguration.
func GetString(key string) string {
	return cast.ToString(Get(key))
}

// GetInt returns the value associated with the key as an integer with default defaultConfiguration.
func GetInt(key string) int {
	return cast.ToInt(Get(key))
}

// GetStringMap returns the value associated with the key as a map of interfaces with default defaultConfiguration.
func GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(Get(key))
}
