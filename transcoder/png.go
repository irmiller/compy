package transcoder

import (
	"github.com/irmiller/compy/proxy"
	"github.com/chai2010/webp"
	"image/png"
	"net/http"
	"log"
)

type Png struct{}

func (t *Png) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
	img, err := png.Decode(r)
	if err != nil {
		return err
	}
	log.Printf("logging")
	if SupportsWebP(headers) {
		w.Header().Set("Content-Type", "image/webp")
		options := webp.Options{
			Lossless: false,
			Quality: float32(10),
		}
		if err = webp.Encode(w, img, &options); err != nil {
			return err
		}
	} else {
		if err = png.Encode(w, img); err != nil {
			return err
		}
	}
	return nil
}
