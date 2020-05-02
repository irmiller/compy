package transcoder

import (
	"github.com/irmiller/compy/proxy"
	//"github.com/chai2010/webp"
	//"image/gif"
	"net/http"
	"io/ioutil"
	giftowebp "github.com/sizeofint/gif-to-webp"
)

type Gif struct{}

func (t *Gif) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
	gifBin, _  := ioutil.ReadAll(r)
	w.Header().Set("Content-Type", "image/webp")

	converter  := giftowebp.NewConverter()
	converter.WebPConfig.SetEmulateJpegSize(1)
	converter.WebPConfig.SetQuality(20)
	converter.WebPConfig.SetThreadLevel(2)
	converter.WebPAnimEncoderOptions.SetKmin(9)
	converter.WebPAnimEncoderOptions.SetKmax(17)
	webpBin, _  := converter.Convert(gifBin)
	w.Write(webpBin)

	return nil
}
