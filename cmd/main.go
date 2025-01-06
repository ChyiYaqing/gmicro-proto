package main

import (
	"fmt"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"os"

	"github.com/chyiyaqing/gmicro-proto/config"
	"github.com/chyiyaqing/gmicro-proto/openapiv2"
	"google.golang.org/grpc/grpclog"
)

func getOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(openapiv2.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	oaAddr := fmt.Sprintf("0.0.0.0:%d", config.GetApplicationHttpPort())
	oa := getOpenAPIHandler()
	oaServer := &http.Server{
		Addr: oaAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			oa.ServeHTTP(w, r)
		}),
	}
	log.Info("Serving Swagger UI on http://", oaAddr)
	if err := oaServer.ListenAndServe(); err != nil {
		log.Fatalln("Failed to listen:", err)
	}
}
