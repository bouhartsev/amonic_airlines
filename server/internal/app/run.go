package app

import "github.com/bouhartsev/amonic_airlines/server/internal/server"

func (a *application) Run() error {
	s, err := server.New(a.logger, a.cfg)

	if err != nil {
		return err
	}

	return s.Run()
}
