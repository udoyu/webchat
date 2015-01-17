package tools
import (
    "net/http"
    "bytes"
    "io/ioutil"
)

func HttpPost(url string, data []byte) ([]byte, error) {
    r, err := http.Post(url, globalPostType, bytes.NewReader(data))
    if err != nil {
        return nil,nil
    }
    defer r.Body.Close()
    return ioutil.ReadAll(r.Body)
}

func HttpGet(url string) ([]byte, error) {
    r, err := http.Get(url)
    if err != nil {
        return nil,nil
    }
    defer r.Body.Close()
    return ioutil.ReadAll(r.Body)
}
