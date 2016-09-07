package main

import (
    "bytes"
    "strconv"
    "net/http"
    "fmt"
    "io/ioutil"
    "regexp"
    "os"
    "bufio"
    "sync"
)

const (
    GeocodePattern = `\(\{lng:([0-9\.]+),lat:([0-9\.]+)\}\)`
    Length = 1000
)

// given a slice of urls
func main() {
    baseUrl := "http://www.dianping.com/shop/"
    initialId := 58013110

    // open a file
    f, err := os.Create("./out")
    if err != nil {
        panic(err)
    }
    w := bufio.NewWriter(f)
    defer f.Close()

    var wg sync.WaitGroup
    wg.Add(Length)
    mutex := &sync.Mutex{}

    for i := 0; i < Length; i ++ {
        go func(index int) {
            defer wg.Done()

            var urlWriter bytes.Buffer
            urlWriter.WriteString(baseUrl)
            urlWriter.WriteString(strconv.Itoa(initialId + index))
            url := urlWriter.String()

            fmt.Sprintf("Crawling %s\n", url)
            bodyString := Crawler(url)

            lng, lat, found := Parse(bodyString)
            if found {
                mutex.Lock()
                if _, err := w.WriteString(fmt.Sprintf("url: %s, lat: %s, lng: %s\n", url, lat, lng)); err != nil {
                    panic(err)
                }
                f.Sync()
                w.Flush()
                mutex.Unlock()
            }
        }(i)
    }

    wg.Wait()
}

func Crawler(url string) (string) {
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }

    // read body
    defer resp.Body.Close()
    byt, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    return string(byt)
}

func Parse(content string) (string, string, bool) {
    r, err := regexp.Compile(GeocodePattern)
    if err != nil {
        panic(err)
    }
    res := r.FindStringSubmatch(content)
    if len(res) > 2 {
        return res[1], res[2], true
    }

    return "", "", false
}

func Write(url, lng, lat string) {

}