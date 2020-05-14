package transcoder

import (
	"github.com/irmiller/compy/proxy"
// 	"github.com/chai2010/webp"
// 	"github.com/pixiv/go-libjpeg/jpeg"
	"net/http"
// 	"image/png"
// 	"strconv"
// 	"io/ioutil"
	"github.com/nickalie/go-webpbin"
)

func (t *image) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
	err := webpbin.NewCWebP().
		Quality(80).
		Input(r).
		Output(w).
		Run()
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "image/webp")
// 	options := webp.Options{
// 		Lossless: false,
// 		Quality:  float32(25),
// 	}
// 	if err = webp.Encode(w, img, &options); err != nil {
// 		return err
// 	}
// 	} else {
// 		if err = jpeg.Encode(w, img, encOptions); err != nil {
// 			return err
// 		}
// 	}
	return nil
}
