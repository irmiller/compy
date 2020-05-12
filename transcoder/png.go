package transcoder

import (
	"github.com/irmiller/compy/proxy"
	"github.com/chai2010/webp"
	"image/png"
	"net/http"
// 	"image/gif"
	"github.com/pixiv/go-libjpeg/jpeg"
// 	"log"
	"fmt"
	"io/ioutil"
	giftowebp "github.com/sizeofint/gif-to-webp"
)

type Png struct{}

func (t *Png) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
// 	img, err := ioutil.ReadAll(r)

	img, err := png.Decode(r)
// 	fmt.Println("Error 1 %s",err)
	if err != nil {
		err = nil
		img, err = jpeg.Decode(r, &jpeg.DecoderOptions{})
		fmt.Println("Error 2 %s",err)
		if err != nil {
			err = nil
			img, err := ioutil.ReadAll(r)
			fmt.Println("Error 3 %s",err)
// 			if err != nil {
// 				return err
// 			}
			w.Header().Set("Content-Type", "image/webp")
// 			options := webp.Options{
// 				Lossless: false,
// 				Quality: float32(25),
// 			}
			converter  := giftowebp.NewConverter()
			converter.WebPConfig.SetQuality(25)
			converter.WebPConfig.SetPreprocessing(2)
			webpBin, _  := converter.Convert(img)
			w.Write(webpBin)
// 			if err = webp.Encode(w, webpBin, &options); err != nil {
// 				return err
// 			}
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
