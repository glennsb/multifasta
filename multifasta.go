// Quick app to basically concatentate a number of input files
//
// Copyright (c) 2014 Stuart Glenn, Oklahoma Medical Research Foundation
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
  "flag"
  "fmt"
  "io"
  "log"
  "os"
  "path/filepath"
  "sort"
  "strings"
  "time"
)

var (
  output_name string
)

func init() {
  flag.StringVar(&output_name, "output", output_name, "Output filename (- for STDOUT, do not specify for default, multifasta_output_TIMESTAMP.fasta)")
  flag.Parse()
}

func defaultOutput(inputPath string) string {
  dir := filepath.Dir(inputPath)
  t := time.Now()
  outfile := fmt.Sprintf("multifasta_output_%d%02d%02d-%02d%02d%02d.fasta",
    t.Year(),
    t.Month(),
    t.Day(),
    t.Hour(),
    t.Minute(),
    t.Second())
  return filepath.Join(dir, outfile)
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
    fmt.Fprintf(os.Stderr, "Missing one or more input file(s)\n")
    flag.Usage()
    os.Exit(1)
  }
  input_paths := flag.Args()
  sort.Strings(input_paths)
  if "" == output_name {
    output_name = defaultOutput(input_paths[0])
  }
  out := openOutput(output_name)

  defer out.Close()

  for index, infile := range input_paths {
    basefile := filepath.Base(infile)
    extension := filepath.Ext(basefile)
    underscore := strings.Index(basefile, "_")
    if underscore < 0 {
      underscore = 0
    } else {
      underscore += 1
    }
    // CRLF is used as this is intended to be used by windows user in Notepad.exe typically
    fmt.Fprintf(out, ">%s exported from %s\r\n", basefile[underscore:len(basefile)-len(extension)], basefile)
    in, err := os.Open(infile)
    if err != nil {
      log.Fatal(err)
    }
    io.Copy(out, in)
    in.Close()
    if index+1 < len(input_paths) {
      fmt.Fprintf(out, "\r\n")
    }
  }
}
