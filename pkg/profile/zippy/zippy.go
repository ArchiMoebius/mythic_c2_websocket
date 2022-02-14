package zippy

/*
import (
	"ArchiMoebius/mythic_c2_websocket/pkg/logger"
	"context"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// MailpipeAPI struct for API interaction
type MailpipeAPI struct {
	AttachmentStoragePath string
	Certificate           []tls.Certificate
	CAPool                *x509.CertPool
	Host                  string
	Port                  int
	Listener              net.Listener
	Server                *http.Server
	Pid                   int
	MTLS                  bool
	Verbose               bool
	Debug                 bool
	SessionSecret         [64]byte
}

// New - create and return an API instance
func New(mailpipeAPIConfig apiConfig.MailpipeAPIConfig) (MailpipeAPI, error) {

	if mailpipeAPIConfig.Debug {
		logger.Log("[D] Entering MailpipeAPI.New function")
	}

	_, err := redis.New(mailpipeAPIConfig.Redis)

	if err != nil {
		return MailpipeAPI{}, err
	}

	sessionSecret := [64]byte{}
	_, err = io.ReadFull(rand.Reader, sessionSecret[:])

	if err != nil {
		log.Fatal(err)
	}

	MailpipeAPI := MailpipeAPI{
		AttachmentStoragePath: mailpipeAPIConfig.Common.AttachmentStoragePath,
		Certificate:           mailpipeAPIConfig.Certificate,
		CAPool:                mailpipeAPIConfig.CertificatePool,
		Host:                  mailpipeAPIConfig.Common.Host,
		Port:                  mailpipeAPIConfig.Common.Port,
		Listener:              nil,
		Server:                nil,
		Pid:                   os.Getpid(),
		MTLS:                  mailpipeAPIConfig.Common.MTLS,
		Verbose:               mailpipeAPIConfig.Verbose,
		Debug:                 mailpipeAPIConfig.Debug,
		SessionSecret:         sessionSecret,
	}

	for _, transport := range mailpipeAPIConfig.Transports {
		config, err := transport.Data.MarshalString()

		if err != nil {
			logger.Log(fmt.Sprintf("[!] Failed to MarshalString for Transport: %s", transport.Title))
			continue
		}

		err = redis.AddSiteWideTransport(transport.Title, config)

		if err != nil {
			logger.Log(fmt.Sprintf("[!] Failed to AddSiteWideTransport for Transport: %s", transport.Title))
			continue
		}

		logger.Log(fmt.Sprintf("[+] Site Wide Transport Added: %s", transport.Title))
	}

	return MailpipeAPI, nil
}

// Run starts the mailpipe api listener
func (mailpipeAPI *MailpipeAPI) Run() error {

	if mailpipeAPI.Debug {
		logger.Log("[D] Entering MailpipeAPI.Run function")
	}

	router := setup(mailpipeAPI)

	addr := fmt.Sprintf("%s:%d", mailpipeAPI.Host, mailpipeAPI.Port)

	mailpipeAPI.Server = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         addr,
		Handler:      router,
	}

	if mailpipeAPI.Certificate != nil {
		mailpipeAPI.Server.ReadTimeout = time.Minute
		mailpipeAPI.Server.WriteTimeout = time.Minute
		mailpipeAPI.Server.TLSConfig = &tls.Config{
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			Renegotiation:            tls.RenegotiateNever,
			MinVersion:               tls.VersionTLS12,
			MaxVersion:               tls.VersionTLS13,
			CipherSuites: []uint16{
				tls.TLS_AES_128_GCM_SHA256,
				tls.TLS_AES_256_GCM_SHA384,
				tls.TLS_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
			},
			Certificates: mailpipeAPI.Certificate,
		}

		if mailpipeAPI.CAPool != nil {
			mailpipeAPI.Server.TLSConfig.ClientCAs = mailpipeAPI.CAPool
			mailpipeAPI.Server.TLSConfig.RootCAs = mailpipeAPI.CAPool
		}

		if mailpipeAPI.MTLS {
			mailpipeAPI.Server.TLSConfig.ClientAuth = tls.RequireAndVerifyClientCert
		}

		logger.Log("[+] API configured to utilize TLS")

		lsnr, err := tls.Listen("tcp", addr, mailpipeAPI.Server.TLSConfig)

		if err != nil {
			logger.Log(err.Error())
			return err
		}

		mailpipeAPI.Listener = lsnr

	} else {
		lsnr, err := net.Listen("tcp", addr)

		if err != nil {
			logger.Log(err.Error())
			return err
		}

		mailpipeAPI.Listener = lsnr
	}

	if mailpipeAPI.Verbose {
		logger.Log(fmt.Sprintf("[+] API Listening at %s", addr))
	}

	go func() {
		if err := mailpipeAPI.Server.Serve(mailpipeAPI.Listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}() // TODO: add chan. for error notification ?...

	return nil
}

// Shutdown stops the mailpipe api listener
func (mailpipeAPI *MailpipeAPI) Shutdown(ctx context.Context) error {

	if mailpipeAPI.Debug {
		logger.Log("[D] Entering MailpipeAPI.Shutdown function")
	}

	if err := mailpipeAPI.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
		return err
	}

	return nil
}
*/
