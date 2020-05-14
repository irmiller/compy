package transcoder

import (
	"github.com/irmiller/compy/proxy"
// 	"github.com/chai2010/webp"
// 	"github.com/pixiv/go-libjpeg/jpeg"
	"net/http"
// 	"image/png"
// 	"strconv"
// 	"io/ioutil"
	"github.com/xfrr/goffmpeg/transcode"
)

func (t *Webm) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
// 	err := webpbin.NewCWebP().
// 		Quality(80).
// 		Input(r).
// 		Output(w).
// 		Run()
// 	if err != nil {
// 		return err
// 	}
	trans := new(transcoder.Transcoder)
	err := trans.InitializeEmptyTranscoder()
	
// 	trans.MediaFile().SetFrameRate(70)
// 	trans.MediaFile().SetPreset("veryfast")
	
	w_pipe, err := trans.CreateInputPipe()
	r = w_pipe
	r_pipe, err := trans.CreateOutputPipe("mkv")
	
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		data, err := ioutil.ReadAll(r_pipe)
		r_pipe.Close()
		wg.Done()
	}()
	go func() {
		w_pipe.Close()
	}()
	done := trans.Run(false)
	err = <-done

	wg.Wait()
	
	w.Write(data)
	w.Header().Set("Content-Type", "video/mastroka")
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
