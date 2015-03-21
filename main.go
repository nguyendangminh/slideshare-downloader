package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
)

func main() {
    if len(os.Args) == 1 {
        fmt.Println("Enter URL of the slide")
        os.Exit(1)
    }

    url := os.Args[1]
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }

    var pattern = regexp.MustCompile("data-full(.*?)alt")
    loc := pattern.FindStringSubmatch(string(body))
    fmt.Println(loc)
}
