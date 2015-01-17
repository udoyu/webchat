package pubacc
import "encoding/xml"

type XmlMsgS struct {
    XMLName xml.Name `xml:"xml"`
    *XmlMsgHeadS
    *XmlMsgTextS
    *XmlMsgImageS
    *XmlMsgVoiceS
    *XmlMsgVideoS
    *XmlMsgMusicS
    *XmlMsgArticleS
}

type XmlMsgHeadS XmlMsgHead
type XmlMsgTextS XmlMsgText
type XmlMsgImages struct {
    Image *XmlMsgMedia
}
type XmlMsgVoiceS struct {
    voice *XmlMsgMedia
}
type XmlMsgVideoS struct {
    video *XmlMsgMedia
}
type XmlMsgMusicS struct {
    Music *XmlMsgMusic
}
type XmlMsgArticleS struct {
    ArticleCount *int64
    Articles *XmlMsgArticleItemsS
}

type XmlMsgMusic struct {
    Title *string
    Description *string
    MusicUrl *string
    HQMusicUrl *string
    ThumbMediaId *string
}

type XmlMsgArticleItemsS struct {
    []Item *XmlMsgArtItemS `xml:"item`
}

type XmlMsgArtItemS struct {
    Title *string
    Description *string
    PicUrl *string
    Url *string
}

//回复文本消息
func RespText(head *XmlMsgHeadS, content *XmlMsgTextS) (string,error) {
    head.MsgType = String(TEXT)
    msg := &XmlMsgS{XmlMsgHeadS:head, XmlMsgTextS:content}
    data, err := Marshal(msg)
    return string(data), err
}

//回复图片消息
func RespImage(head *XmlMsgHeadS, image *XmlMsgImageS) (string, error) {
    head.MsgType = String(IMAGE)
    msg := &XmlMsgS{XmlMsgHeadS:head, XmlMsgImageS:image}
    data, err := Marshal(msg)
    return string(data), err
}

//回复语音消息
func RespVoice(head *XmlMsgHeadS, voice *XmlMsgVoiceS) (string, error) {
    head.MsgType = String(IMAGE)
    msg := &XmlMsgS{XmlMsgHeadS:head, XmlMsgVoiceS:voice}
    data, err := Marshal(msg)
    return string(data), err
}

//回复视频消息
func RespVideo(head *XmlMsgHeadS, video *XmlMsgVideoS) (string, error) {
    head.MsgType = String(IMAGE)
    msg := &XmlMsgS{XmlMsgHeadS:head, XmlMsgVideoS:video}
    data, err := Marshal(msg)
    return string(data), err
}

//回复音乐消息
func RespMusic(head *XmlMsgHeadS, music *XmlMsgMusicS) (string, error) {
    head.MsgType = String(MUSIC)
    msg := &XmlMsgS{XmlMsgHeadS:head, XmlMsgMusicS:music}
    data, err := Marshal(msg)
    return string(data), err
}

//回复图文消息
func RespNews(head *XmlMsgHeadS, artcile *XmlMsgArticleS) (string, error) {
    head.MsgType = String(NEWS)
    msg := &XmlMsgS{XmlMsgHeadS:head, XmlMsgArticleS:artcile}
    data, err := Marshal(msg)
    return string(data), err
}
