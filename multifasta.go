package main

import (
  "flag"
  "fmt"
  "io"
  "log"
  "os"
  "path/filepath"
  "time"
)

var (
  output_name string
)

func init() {
  flag.StringVar(&output_name, "output", output_name, "Output filename (- for STDOUT, do not specify for default, multifasta.txt)")
  flag.Parse()
}

func defaultOutput(inputPath string) string {
  dir := filepath.Dir(inputPath)
  t := time.Now()
  outfile := fmt.Sprintf("multifasta_output_%d%02d%02d-%02d%02d.txt",
    t.Year(),
    t.Month(),
    t.Day(),
    t.Hour(),
    t.Minute())
  return filepath.Join(dir,outfile)
}

func openOutput(outputPath string) *os.File {
  switch outputPath {
  case "-":
    return os.Stdout
  default:
    out, err := os.Create(outputPath)
    if nil != err {
      log.Fatal(err)
    }
    return out
  }
}

func main() {
  if len(flag.Args()) <= 0 {
    fmt.Fprintf(os.Stderr,"Missing one or more input file(s)\n")
    flag.Usage()
    os.Exit(1)
  }
  if "" == output_name {
    output_name = defaultOutput(flag.Args()[0])
  }
  out := openOutput(output_name)

  defer out.Close()

  for _, infile := range flag.Args() {
    basefile := filepath.Base(infile)
    extension := filepath.Ext(basefile)
    fmt.Fprintf(out,">%s exported from %s\n", basefile[:len(basefile)-len(extension)], basefile)
    in, err := os.Open(infile)
    if err != nil {
      log.Fatal(err)
    }
    io.Copy(out,in)
    in.Close()
    fmt.Fprintln(out,"")
  }
}
