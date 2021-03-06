package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync/atomic"

	"github.com/irmiller/compy/proxy"
	tc "github.com/irmiller/compy/transcoder"
)

var (
	host  = flag.String("host", ":8888", "<host:port>")
	cert  = flag.String("cert", "", "proxy cert path")
	key   = flag.String("key", "", "proxy cert key path")
	ca    = flag.String("ca", "/home/ubuntu/keys/ca.crt", "CA path")
	caKey = flag.String("cakey", "/home/ubuntu/keys/ca.key", "CA key path")
	user  = flag.String("user", "", "proxy user name")
	pass  = flag.String("pass", "", "proxy password")

	brotli = flag.Int("brotli", 11, "Brotli compression level (0-11)")
	jpeg   = flag.Int("jpeg", 25, "jpeg quality (1-100, 0 to disable)")
	gif    = flag.Bool("gif", true, "transcode gifs into static images")
	gzip   = flag.Int("gzip", 9, "gzip compression level (0-9)")
	png    = flag.Bool("png", true, "transcode png")
// 	webm   = flag.Bool("webm", true, "transcode webm")
	minify = flag.Bool("minify", true, "minify css/html/js - WARNING: tends to break the web")
)

func main() {
	flag.Parse()

	p := proxy.New(*host, *cert)

	if (*ca == "") != (*caKey == "") {
		log.Fatalln("must specify both CA certificate and key")
	}

	if (*cert == "") != (*key == "") {
		log.Fatalln("must specify both certificate and key")
	}

	if *ca != "" {
		if err := p.EnableMitm(*ca, *caKey); err != nil {
			fmt.Println("not using mitm:", err)
		}
	}

	// TODO: require cert and key?
	if (*user == "") != (*pass == "") {
		log.Fatalln("must specify both user and pass")
	} else {
		p.SetAuthentication(*user, *pass)
	}

	if *jpeg != 0 {
		p.AddTranscoder("image/jpeg", tc.NewJpeg(*jpeg))
		p.AddTranscoder("image/jpg", tc.NewJpeg(*jpeg))
	}
	if *gif {
		p.AddTranscoder("image/gif", &tc.Gif{})
	}
	if *png {
		p.AddTranscoder("image/png", &tc.Png{})
	}
// 	if *webm {
// 		p.AddTranscoder("video/webm", &tc.Webm{})
// 	}

	var ttc proxy.Transcoder
	if *minify {
		ttc = &tc.Zip{tc.NewMinifier(), *brotli, *gzip, false}
	} else {
		ttc = &tc.Zip{&tc.Identity{}, *brotli, *gzip, true}
	}

	p.AddTranscoder("text/css", ttc)
	p.AddTranscoder("text/html", ttc)
	p.AddTranscoder("text/javascript", ttc)
	p.AddTranscoder("application/javascript", ttc)
	p.AddTranscoder("application/x-javascript", ttc)
	p.AddTranscoder("application/json", ttc)

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			read := atomic.LoadUint64(&p.ReadCount)
			written := atomic.LoadUint64(&p.WriteCount)
			log.Printf("Quit: %dMB : %dMB (%3.1f%%)", written, read, float64(written)/float64(read)*100)
			os.Exit(0)
		}
	}()

	log.Printf("Listening on %s", *host)

	var err error
	if *cert != "" {
		err = p.StartTLS(*host, *cert, *key)
	} else {
		err = p.Start(*host)
	}
	log.Fatalln(err)
}
