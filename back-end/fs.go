/* © 2022 Ben C. Forsberg <benfrsbrg@gmail.com> */

package main

//go:generate sh -c "cd ../front-end && npx react-scripts build"
//go:generate sh -c "rm -f ./assets/js/*.js"
//go:generate sh -c "cp ../front-end/build/static/js/*.js ./assets/js/"
//go:generate sh -c "rm -f ./assets/css/*.css"
//go:generate sh -c "cp ../front-end/build/static/css/*.css ./assets/css/"

import (
  "embed"
  "html/template"
  "fmt"
  "bytes"
  "path"
)

//go:embed assets/index.html
var index string

//go:embed assets/js/*.js
var scripts embed.FS

//go:embed assets/css/*.css
var styles embed.FS

var indexTemplate = template.Must(template.New("index.html").Parse(index))

type templateData struct {
  Scripts []template.JS
  Styles []template.CSS
}

func (td *templateData) ReadScripts(scripts embed.FS) error {
  entries, err := scripts.ReadDir(`assets/js`)
  if err != nil {
    return fmt.Errorf(`could not read js directory: %w`, err)
  }
  for _, dirEntry := range entries {
    info, err := dirEntry.Info()
    if err != nil {
      return fmt.Errorf(`could not get info on %q: %w`, dirEntry.Name(), err)
    }
    file, err := scripts.Open(path.Join(`assets/js`, info.Name()))
    if err != nil {
      return fmt.Errorf(`could not open %q: %w`, info.Name(), err)
    }
    defer file.Close()
    buff := new(bytes.Buffer)
    if _, err := buff.ReadFrom(file); err != nil {
      return fmt.Errorf(`could not read from %q: %w`, info.Name(), err)
    }
    td.Scripts = append(td.Scripts, template.JS(buff.String()))
  }
  return nil
}

func (td *templateData) ReadStyles(styles embed.FS) error {
  entries, err := styles.ReadDir(`assets/css`)
  if err != nil {
    return fmt.Errorf(`could not read css directory: %w`, err)
  }
  for _, dirEntry := range entries {
    info, err := dirEntry.Info()
    if err != nil {
      return fmt.Errorf(`could not get info on %q: %w`, dirEntry.Name(), err)
    }
    file, err := styles.Open(path.Join(`assets/css`, info.Name()))
    if err != nil {
      return fmt.Errorf(`could not open %q: %w`, info.Name(), err)
    }
    defer file.Close()
    buff := new(bytes.Buffer)
    if _, err := buff.ReadFrom(file); err != nil {
      return fmt.Errorf(`could not read from %q: %w`, info.Name(), err)
    }
    td.Styles = append(td.Styles, template.CSS(buff.String()))
  }
  return nil
}

func Index() ([]byte, error) {
  td := new(templateData)
  err := td.ReadScripts(scripts)
  if err != nil {
    return nil, fmt.Errorf(`could not read scripts: %w`, err)
  }
  err = td.ReadStyles(styles)
  if err != nil {
    return nil, fmt.Errorf(`could not read styles: %w`, err)
  }
  buff := new(bytes.Buffer)
  err = indexTemplate.Execute(buff, td)
  if err != nil {
    return nil, fmt.Errorf(`could not render template: %w`, err)
  }
  return buff.Bytes(), nil
}
