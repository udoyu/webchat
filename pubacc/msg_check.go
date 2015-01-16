package pubacc

import (
    "sort"
    "crypto/sha1"
)

type RawQuerySt struct {
    Signature string
    Timestamp string
    Nonce     string
    Echostr   string
}

func (this *RawQuerySt) MsgCheck(token string) bool {
    strarr := []string{token, this.Timestamp, this.Nonce}
    sort.Strings(strarr)
    signature := strarr[0] + strarr[1] + strarr[2]
    signature = string(sha1.Sum([]byte(signature)))
    return signature == this.Signature
}
