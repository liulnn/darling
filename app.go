package darling

import (
	"fmt"
	"net/http"
	"time"
)

type App struct {
	Handlers *ControllerRegister
	Server   *http.Server
}

func NewApp() *App {
	cr := NewControllerRegister()
	app := &App{Handlers: cr, Server: &http.Server{}}
	return app
}

func (app *App) Run(host string, port int) {
	app.Server.Addr = fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("%s:%d\n", host, port)
	app.Server.Handler = app.Handlers

	endRunning := make(chan bool, 1)

	go func() {
		err := app.Server.ListenAndServe()
		if err != nil {
			time.Sleep(100 * time.Microsecond)
			endRunning <- true
			return
		}
	}()
	<-endRunning
}
