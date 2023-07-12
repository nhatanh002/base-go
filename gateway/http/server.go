package http

import (
	"base-go/application"
	"base-go/gateway/http/controllers"
	"fmt"
	"net/http"

	config "base-go/common/config"

	"github.com/labstack/echo"
)

func NewHttpServer(cnf *config.Config, engine http.Handler) *http.Server {
	//tlsConfig := &tls.Config{
	//	MinVersion:               tls.VersionTLS12,
	//	PreferServerCipherSuites: true,
	//}

	//tlsConfig.Certificates = make([]tls.Certificate, 1)
	//var err error
	//if tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(cnf.X509CertFile, cnf.X509KeyFile); err != nil {
	//	logger.Get().Err(err).Msg("Unable to load TLS certificates, aborting...")
	//	panic(err)
	//}

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cnf.HttpConfig.Host, cnf.HttpConfig.Port),
		Handler: engine,
		//TLSConfig: tlsConfig,
		//ReadTimeout:    readTimeout,
		//WriteTimeout:   writeTimeout,
		//MaxHeaderBytes: maxHeaderBytes,
	}
}

func EchoRouter(cnf *config.Config, app *application.App) http.Handler {
	e := echo.New()

	// middlewares config goes here...

	// subrouters/controllers mounting
	catsController := controllers.NewCatsController(app.Cats)
	catsController.Mount(e)

	return e
}
