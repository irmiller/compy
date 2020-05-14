package transcoder

import (
	"github.com/irmiller/compy/proxy"
// 	"github.com/chai2010/webp"
// 	"github.com/pixiv/go-libjpeg/jpeg"
	"net/http"
// 	"image/png"
// 	"strconv"
	"io/ioutil"
// 	"testing"
	"sync"
	"github.com/xfrr/goffmpeg/transcoder"
)
type Webm struct{}

func (t *Webm) Transcode(w *proxy.ResponseWriter, r *proxy.ResponseReader, headers http.Header) error {
// 	err := webpbin.NewCWebP().
// 		Quality(80).
// 		Input(r).
// 		Output(w).
// 		Run()
// 	if err != nil {
// 		return err
// 	}
	cmdName := "ffmpeg"
	args := []string{
		"-hide_banner",
		"-re",
		"-i",
		"pipe:0,
		"-preset",
		"superfast",
		"-c:v",
		"h264",
		"-crf",
		"0",
		"-c",
		"copy",
		"-f", "rawvideo", "-",
	}
	cmd := exec.Command(cmdName, args...)
	cmd.Stdin = bytes.NewReader(io.ReadAll(r))
	
	err2 := cmd.Start()
	
	var buf bytes.Buffer
	n, err := io.Copy(&buf, cmd.StdoutPipe())
	err = cmd.Wait()
	data := buf.Bytes()
	
	
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
