package transcoder

import (
	"github.com/irmiller/compy/proxy"
	//"github.com/chai2010/webp"
	//"image/gif"
	"net/http"
	"io"
	"io/ioutil"
	"bytes"
	"os/exec"
	"log"
	
)

type Webm struct{}

func (t *Webm) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
	w.Header().Set("Content-Type", "video/webm")
	
	log.Printf("video found")
	webmBin, _ := ioutil.ReadAll(r)
		
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-c:av copy", "-b:v 1000","-v","-")
	cmd.Stdin = bytes.NewReader(webmBin)
	
	wr, wwr := io.Pipe()
	
	cmd.Stdout = wwr
	
	cmd.Start()
			    
	webMT, _ := ioutil.ReadAll(wr)
	
	w.Write(webMT)
	return nil
}
