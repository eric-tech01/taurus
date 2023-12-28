package taurus

import (
	"fmt"
	"sync"

	slog "github.com/eric-tech01/simple-log"
	"github.com/eric-tech01/taurus/server"
	"github.com/fatih/color"
)

type Application struct {
	initOnce   sync.Once
	smu        *sync.RWMutex
	servers    []server.Server
	HideBanner bool
}

// Startup ..
func (app *Application) Startup(fns ...func() error) error {
	app.initialize()
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

// initialize application
func (app *Application) initialize() {
	app.initOnce.Do(func() {
		app.servers = make([]server.Server, 0)
		app.smu = &sync.RWMutex{}
		_ = app.printBanner()
	})
}

// printBanner init
func (app *Application) printBanner() error {
	if app.HideBanner {
		return nil
	}

	const banner = `
     _(_____)_
    |__.___.__|
        | |    
        | | 
        |_|       

 Welcome to taurus, starting application ...
`
	fmt.Println(color.GreenString(banner))
	return nil
}

// Serve start server
func (app *Application) Serve(s server.Server) error {
	app.smu.Lock()
	defer app.smu.Unlock()
	app.servers = append(app.servers, s)
	return nil
}

// Run run application
func (app *Application) Run() error {
	app.smu.Lock()
	defer app.smu.Unlock()
	wg := sync.WaitGroup{}
	for _, s := range app.servers {
		wg.Add(1)
		s := s
		go func() {
			defer wg.Done()
			if err := s.Start(); err != nil {
				slog.Error(err)
			}
		}()
	}
	wg.Wait()
	return nil

}
