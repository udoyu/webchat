package pubacc

import (
    "encoding/xml"
)

type XmlMsg struct {
    XmlMsgHead
    XmlMsgText
    XmlMsgImage
    XmlMsgVoice
    XmlMsgVideo
    XmlMsgLocation
    XmlMsgLink
    XmlMsgEvent
}

func ParseMsg(buf []byte) (*XmlMsg, error) {
    msg := &XmlMsg{}
    err := xml.Unmarshal(buf, &msg)
    return msg, err
}

//msgtype:text,image,voice,video,location,link,event
type XmlMsgHead struct {
    ToUserName string
    FromUserName string
    CreateTime int64
    MsgType string
    MsgId int64
}

type XmlMsgText struct {
    Content string
}

type XmlMsgImage struct {
    MediaId string
    PicUrl string
}

type XmlMsgVoice struct {
    MediaId string
    Format string
}

type XmlMsgVideo struct {
    MediaId string
    ThumbMediaId string
}

type XmlMsgLocation struct {
    Location_X string
    Location_Y string
    Scale string
    Label string
}

type XmlMsgLink struct {
    Title string
    Description string
    Url string
}

//event:subscribe,SCAN,LOCATION,CLICK,VIEW
type XmlMsgEvent struct {
    Event string
    EventKey string
    Ticket string
    XmlMsgEventLocation
}

type XmlMsgEventLocation struct {
    Latitude string
    Longitude string
    Precision string
}

