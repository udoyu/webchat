package pubacc

import (
    "encoding/xml"
    "strings"
)

func String(s string) *string{
    return &s
}

func Int64(v int64) *int64{
    return &v
}

func CDATAString(s string) *string {
    str := "<![CDATA[" + s +"]]>"
    return &str
}

var STRING_PTR reflect.Type
const CDATA_PREFIX_DEST = `><![CDATA[`
const CDATA_SUFFIX_DEST = `]]><`
const CDATA_PREFIX_SRC = `>&lt;![CDATA[`
const CDATA_SUFFIX_SRC = `]]&gt;<`

func init() {
    s := ""
    STRING_PTR = reflect.TypeOf(&s)
}

func Unmarshal (data []byte, v interface{}) error {
    return xml.Unmarshal(data, v)
}

func Marshal(v interface{}) ([]byte, error) {
    //add <!CdData[]> for *string
    values := reflect.ValueOf(v)
    for i := 0; i < values.NumField(); i++ {
        value := values.Field(i)
        if value.Type() == STRING_PTR {
            if value.Elem().CanSet() {
                value.Elem().SetString(*CDATAString(value.Elem().String()))
            }
        }
    }

    data, err := xml.Marshal(v)
    if err == nil {
        str := strings.Replace(string(data),CDATA_PREFIX_SRC, CDATA_PREFIX_DEST, -1)
        str = strings.Replace(str, CDATA_SUFFIX_SRC, CDATA_SUFFIX_DEST, -1)
        data = []byte(str)
    }
    return data, err
}


