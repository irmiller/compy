package transcoder

import (
	"github.com/irmiller/compy/proxy"
	"github.com/chai2010/webp"
	"image/png"
	"net/http"
	"image/gif"
	"github.com/pixiv/go-libjpeg/jpeg"
// 	"log"
// 	"fmt"
)

type Png struct{}

func (t *Png) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
	img, err := png.Decode(r)
// 	fmt.Errorf("Error 1 %s",err)
	if err != nil {
		err = nil
		img, err = jpeg.Decode(r, &jpeg.DecoderOptions{})
		if err != nil {
			err = nil
			img, err = gif.Decode(r)
			if err != nil {
				return err
			}
		}
	}
	if SupportsWebP(headers) {
		w.Header().Set("Content-Type", "image/webp")
		options := webp.Options{
			Lossless: false,
			Quality: float32(25),
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
