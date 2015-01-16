package pubacc

import (
//    "encoding/xml"
)

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
//    //add <!CdData[]> for *string_t
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

//func AddCdData(s **string_t) **string_t {
//    return s
//}
//msgtype:text,image,voice,video,location,link,event



type XmlMsgHead struct {
    ToUserName *string_t
    FromUserName *string_t
    CreateTime int64
    MsgType *string_t
    MsgId int64
}

type XmlMsgText struct {
    Content *string_t
}

type XmlMsgMedia struct {
    MediaId *string_t
}
type XmlMsgImage struct {
    PicUrl *string_t
}

type XmlMsgVoice struct {
    Format *string_t
}

type XmlMsgVoiceRest struct {
    Recognition *string_t
}

type XmlMsgVideo struct {
    ThumbMediaId *string_t
}

type XmlMsgLocation struct {
    Location_X *string_t
    Location_Y *string_t
    Scale *string_t
    Label *string_t
}

type XmlMsgLink struct {
    Title *string_t
    Description *string_t
    Url *string_t
}

//event:subscribe,SCAN,LOCATION,CLICK,VIEW
type XmlMsgEvent struct {
    Event *string_t
    EventKey *string_t
    Ticket *string_t
    XmlMsgEventLocation
}

type XmlMsgEventLocation struct {
    Latitude *string_t
    Longitude *string_t
    Precision *string_t
}

