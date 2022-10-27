package main

import (
	"github.com/ethercflow/hookfs/hookfs"
	"net/http"
	"time"
)

func main() {
	h := &MyHook{
		dur: time.Second,
	}
	fs, _ := hookfs.NewHookFs("/tmp/b", "/tmp/a", h)
	fs.Serve()

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			h.dur, _ = time.ParseDuration(r.FormValue("dur"))
		})

		http.ListenAndServe("127.0.0.1:8080", nil)
	}()
}

type MyHookContext struct{}
type MyHook struct {
	dur time.Duration
}

func (h *MyHook) PreRead(path string, length int64, offset int64) (buf []byte, hooked bool, ctx hookfs.HookContext, err error) {
	time.Sleep(h.dur)
	return nil, false, MyHookContext{}, nil
}

func (h *MyHook) PostRead(realRetCode int32, realBuf []byte, prehookCtx hookfs.HookContext) (buf []byte, hooked bool, err error) {
	return realBuf, false, nil
}
