package conf

import (
	"os"
	"sync"

	"github.com/pelletier/go-toml"
)

type Configuration struct {
	config    *toml.Tree
	fileName  string
	mu        sync.RWMutex
	keyMap    *sync.Map
	onChanges []func(*Configuration)
	onLoadeds []func(*Configuration)
}

const (
	defaultKeyDelim = "."
)

// New constructs a new Configuration with provider.
func New() *Configuration {
	return &Configuration{
		keyMap:    &sync.Map{},
		onChanges: make([]func(*Configuration), 0),
		onLoadeds: make([]func(*Configuration), 0),
	}
}

func (c *Configuration) Load(filepath string) error {
	var err error
	c.config, err = toml.LoadFile(filepath)
	if err != nil {
		return err
	}
	c.fileName = filepath

	for _, loadHook := range c.onLoadeds {
		loadHook(c)
	}
	return nil
}

// Flush ...
func (c *Configuration) Flush() error {
	f, err := os.Create(c.fileName)
	if err != nil {
		return err
	}
	_, err = c.config.WriteTo(f)
	return err
}

// OnChange 注册change回调函数 TODO: 未实现
func (c *Configuration) OnChange(fn func(*Configuration)) {
	c.onChanges = append(c.onChanges, fn)
}

func (c *Configuration) OnLoaded(fn func(*Configuration)) {
	c.onLoadeds = append(c.onLoadeds, fn)
}
