package main

import (
  "io"
  "bufio"
  "os"
  "fmt"
  "encoding/json"
  "github.com/cbroglie/mustache"
  "gopkg.in/alecthomas/kingpin.v2"
)

var (
  templatePath = kingpin.Flag("template", "File path to the html template to use.").Short('t').Required().PlaceHolder("template.html").String()
  verbose = kingpin.Flag("verbose", "Use this to print errors.").Short('v').Bool()
)

func main () {
  kingpin.Parse()
  reader := bufio.NewReader(os.Stdin)
  template, err := mustache.ParseFile(*templatePath)
  if err != nil {
    perr(err)
    os.Exit(2)
  }

  for err == nil {
    str, err := reader.ReadString('\n')
    if err != nil {
      if err == io.EOF {
        os.Exit(0)
      }
      perr(err)
      os.Exit(3)
    }
    byt := []byte(str)
    var data map[string]interface{}
    if err = json.Unmarshal(byt, &data); err != nil {
      perr(err)
      os.Exit(3)
    }
    template.FRender(os.Stdout, data)
  }
}

func perr (err error) {
  if *verbose == true {
    fmt.Println(err)
  }
}
