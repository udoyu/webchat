package pubacc
import "encoding/xml"

type XmlMsg struct {
    XMLName xml.Name `xml:"xml"`
    XmlMsgHead
    XmlMsgText
    XmlMsgMedia
    XmlMsgImage
    XmlMsgVoice
    XmlMsgVoiceRest
    XmlMsgVideo
    XmlMsgLocation
    XmlMsgTitle
    XmlMsgLink
    XmlMsgEvent
}

//msgtype:text,image,voice,video,location,link,event
const (
    TEXT = "text"
    IMAGE = "image"
    VOICE = "voice"
    LOCATION = "location"
    LINK = "link"
    EVENT = "event"
    MUSIC = "music"
    NEWS = "news"
)

type XmlMsgHead struct {
    ToUserName *string
    FromUserName *string
    CreateTime *int64
    MsgType *string
    MsgId *int64
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

type XmlMsgTitle struct {
    Title *string
}

type XmlMsgLink struct {
    Description *string
    Url *string
}

//event:subscribe,SCAN,LOCATION,CLICK,VIEW
const (
    SUBSCRIBE = "subscribe"
    SCAN = "SCAN"
    LOCATION = "LOCATION"
    CLICK = "CLICK"
    VIEW = "VIEW"
)

type XmlMsgEvent struct {
    Event *string
    EventKey *string
    Ticket *string
    *XmlMsgEventLocation
}

type XmlMsgEventLocation struct {
    Latitude *string
    Longitude *string
    Precision *string
}

