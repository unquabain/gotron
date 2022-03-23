/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com> */

package main

import (
	"encoding/base64"
	"fmt"
	"github.com/webview/webview"
  "gotron/back-end/store"
	"os"
)

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
	store.NewStore(w)

	w.Navigate(`data:text/html;base64,` + base64.StdEncoding.EncodeToString(index))
	w.Run()
}
