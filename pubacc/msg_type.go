package pubacc

import (
//    "encoding/xml"
)

func String(s string) *string{
    return &s
}

func CDATAString(s string) *string {
    str := "<![CDATA[" + s +" ]]>"
    return &str
}

type XmlMsg struct {
    XmlMsgHead
    *XmlMsgText
    *XmlMsgMedia
    *XmlMsgImage
    *XmlMsgVoice
    *XmlMsgVoiceRest
    *XmlMsgVideo
    *XmlMsgLocation
    *XmlMsgLink
    *XmlMsgEvent
}

//var STRING_PTR reflect.Type
//
//func init() {
//    s := ""
//    STRING_PTR = reflect.TypeOf(&s)
//}
//
//func Unmarshal (data []byte v interface{}) error {
//    return xml.Unmarshal(data, v)
//}
//
//func Marshal(v interface{}) ([]byte, error) {
//    //add <!CdData[]> for *string
//    values := reflect.ValueOf(v)
//    for i := 0; i < values.NumField(); i++ {
//        value := values.Field(i)
//        if value.Type() == STRING_PTR {
//            if v.Elem().CanSet() {
//                v.Elem().SetString(*AddCdData(v))
//            }
//        }
//    }
//}

//func AddCdData(s **string) **string {
//    return s
//}
//msgtype:text,image,voice,video,location,link,event



type XmlMsgHead struct {
    ToUserName *string
    FromUserName *string
    CreateTime int64
    MsgType *string
    MsgId int64
}

type XmlMsgText struct {
    Content *string
}

type XmlMsgMedia struct {
    MediaId *string
}
type XmlMsgImage struct {
    PicUrl *string
}

type XmlMsgVoice struct {
    Format *string
}

type XmlMsgVoiceRest struct {
    Recognition *string
}

type XmlMsgVideo struct {
    ThumbMediaId *string
}

type XmlMsgLocation struct {
    Location_X *string
    Location_Y *string
    Scale *string
    Label *string
}

type XmlMsgLink struct {
    Title *string
    Description *string
    Url *string
}

//event:subscribe,SCAN,LOCATION,CLICK,VIEW
type XmlMsgEvent struct {
    Event *string
    EventKey *string
    Ticket *string
    XmlMsgEventLocation
}

type XmlMsgEventLocation struct {
    Latitude *string
    Longitude *string
    Precision *string
}

