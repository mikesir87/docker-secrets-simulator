package main

import (
    "os"
    "strings"
    "io/ioutil"
    "fmt"
    "sort"
)

func main() {
  // This MUST remain in alphabetical order for contains to work
  ignores := []string{ "GOLANG_VERSION", "GOPATH", "HOME", "HOSTNAME", "PATH", "no_proxy" }

  for _, e := range os.Environ() {
    pair := strings.Split(e, "=")
    if !contains(ignores, pair[0]) {
      contents := []byte(pair[1]);
      err := ioutil.WriteFile("/run/secrets/" + pair[0], contents, 0644)
      check(err) 
    }
  }
}

func contains(list []string, value string) bool {
	i := sort.SearchStrings(list, value)
	return i < len(list) && list[i] == value
}

func check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
