package transcoder

import (
	"github.com/irmiller/compy/proxy"
	"github.com/chai2010/webp"
	"image/gif"
	"net/http"
	"io/ioutil"
	giftowebp "github.com/sizeofint/gif-to-webp"
)

type Gif struct{}

func (t *Gif) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
	//img, err := gif.Decode(r)
	gifBin, _  := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	//if SupportsWebP(headers) {
		w.Header().Set("Content-Type", "image/webp")
		options := webp.Options{
			Lossless: false,
			Quality: 10,
		}
		converter  := giftowebp.NewConverter()
		converter.WebPAnimEncoderOptions.SetKmin(9)
		converter.WebPAnimEncoderOptions.SetKmax(17)
		webpBin, err  := converter.Convert(gifBin)
		w.Write(webpBin)
		//if err = webp.Encode(w, img, &options); err != nil {
		//	return err
		//}
	//} //else {
		//if err = gif.Encode(w, img, nil); err != nil {
		//	return err
		//}
	//}
	return nil
}
