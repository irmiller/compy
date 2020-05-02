package transcoder

import (
	"github.com/irmiller/compy/proxy"
	//"github.com/chai2010/webp"
	//"image/gif"
	"net/http"
	"io/ioutil"
    	"github.com/xfrr/goffmpeg/transcoder"
)

type WebM struct{}

func (t *WebM) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
	w.Header().Set("Content-Type", "video/webm")
	trans := new(transcoder.Transcoder)
	err := trans.InitializeEmptyTranscoder()
	
	in, err := trans.CreateInputPipe()
	webM = in.Read(ioutil.ReadAll(r))
	
	r, err := trans.CreateOutputPipe("webm")
	
	trans.MediaFile().SetPreset("ultrafast")
	trans.MediaFile().SetQuality(20)
	
	done := trans.Run(false)
	
	webMT, err := ioutil.ReadAll(r)
	
	w.Write(webMT)
	return nil
}
