package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	from = flag.String("from", "localhost:80", "The host of Hadoop nameNode ")
	port = flag.String("port", "8080", "The port of \"/metrics\"  output endpoint")
	path = flag.String("path", "/metrics", "Path of output endpoint")
)

func init() {
	log.Printf("initalizing")
}

func main() {
	flag.Parse()
	prometheus.MustRegister(NewHadoopCollector(map[string]string{*from: *from}))
	http.Handle(*path, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head><title>Hadoop Exporter</title></head>
            	<body>
            		<h1>Hadoop Exporter</h1>
            		<p><a href='` + *path + `'>Metrics</a></p>
            	</body>
			</html>`))
	})
	listenAddress := ":" + *port
	log.Printf("server listing at %v", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
