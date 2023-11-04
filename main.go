package main

import (
	"az-appservice/httpApi"
	"az-appservice/log"
	"crypto/tls"
	"fmt"
	"github.com/namsral/flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var prod = flag.Bool("prod", false, "Disable verbose logging")
var httpPortFlag = flag.Int("http port", 8001, "A port to accept http requests, https will be set to this port +100")
var apiKeyFileFlag = flag.String("apiKey", "api-key.pem", "A path to a file containing PEM private key for http API")
var apiCertFileFlag = flag.String("apiCert", "api-cert.pem", "A path to a file containing PEM certificate for http API")

func runHTTP(port int, server *http.Server, keyFile string, certFile string) {
	httpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Sugar.Errorf("Can not start http listening on port:%d\n%v", port, err)
		os.Exit(3)
	}

	log.Sugar.Infof("Start http listening :%d", port)
	//goland:noinspection GoUnhandledErrorResult
	go server.Serve(httpListener)

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)

	if err != nil {
		log.Sugar.Errorf("Loading certificate: %v", err)
	} else {
		port = port + 100
		tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}}
		httpsListener, err := tls.Listen("tcp", fmt.Sprintf(":%d", port), tlsCfg)
		if err != nil {
			log.Sugar.Errorf("Can not start http listening on port:%d\n%v", port, err)
			os.Exit(32)
		}
		log.Sugar.Infof("Start https listening :%d", port)
		//goland:noinspection GoUnhandledErrorResult
		go server.Serve(httpsListener)
	}
}

func mainWithDeferred() {
	//goland:noinspection GoUnhandledErrorResult
	defer log.Log.Sync()

	httpServer := httpApi.GetHttpServer()
	runHTTP(*httpPortFlag, httpServer, *apiKeyFileFlag, *apiCertFileFlag)

	time.Sleep(200 * time.Millisecond)
	log.Sugar.Info("Hit Ctrl-C to exit...")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	_ = <-sigs
	fmt.Println("\nGraceful exit")
	time.Sleep(1 * time.Second)
}

// just for testing CI, remove this comment asap
func main() {
	flag.Parse()
	log.Setup(*prod)
	mainWithDeferred()
}
