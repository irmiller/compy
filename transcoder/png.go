package transcoder

import (
	"github.com/irmiller/compy/proxy"
	"github.com/chai2010/webp"
	"image/png"
	"net/http"
// 	"image/gif"
	"github.com/pixiv/go-libjpeg/jpeg"
// 	"log"
// 	"fmt"
// 	"io/ioutil"
// 	"image/webp"
// 	giftowebp "github.com/sizeofint/gif-to-webp"
)

type Png struct{}

func (t *Png) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
// 	img, err := ioutil.ReadAll(r)

	img, err := png.Decode(r)
// 	fmt.Println("Error 1 %s",err)
	if err != nil {
		err = nil
		img, err = jpeg.Decode(r, &jpeg.DecoderOptions{})
		if err != nil {
			w.Write(r)
			return nil
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
