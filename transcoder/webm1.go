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
	webM, _  := ioutil.ReadAll(r)
	webMT := ""
	w.Header().Set("Content-Type", "video/webm")
	trans := new(transcoder.Transcoder)
	err := trans.Initialize( webM, webMT )
	trans.MediaFile().SetPreset("ultrafast")
	trans.MediaFile().SetQuality(20)
	done := trans.Run(false)
	
	w.Write(webMT)
	return nil
}
