/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com> */

package main

import (
	"encoding/base64"
	"fmt"
	"github.com/webview/webview"
  "gotron/store"
	"os"
  "github.com/spf13/pflag"
)

type config struct {
  Name string
}

var cfg config

func init() {
  pflag.StringVarP(&cfg.Name, `name`, `n`, `Trevor`, `The name of the user to start with`)
  pflag.Parse()
}

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle(`GoTron Example`)
	w.SetSize(800, 478, webview.HintNone)
	index, err := Index()
	if err != nil {
		fmt.Fprintf(os.Stderr, `could not generate inlined index file: %v`, err)
		os.Exit(-1)
	}
  str := store.NewStore(w)
  str.State.Session.Username = cfg.Name

	w.Navigate(`data:text/html;base64,` + base64.StdEncoding.EncodeToString(index))
	w.Run()
}
