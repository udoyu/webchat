package menu
import (
    "encoding/json"
    "github.com/udoyu/webchat/tools"
)


// https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
const MENU_URL = "https://api.weixin.qq.com/cgi-bin/menu/"
const ACCESS_TOKEN = "access_token="
const MENU_URL_CREATE = MENU_URL + "create?" + ACCESS_TOKEN
const MENU_URL_DELETE = MENU_URL + "delete?" + ACCESS_TOKEN
const MENU_URL_GET = MENU_URL + "get?" + ACCESS_TOKEN

type menu struct {
    accessToken string
    createUrl string
    deleteUrl string
    getUrl string
}

func NewMenu(accessToken string) *menu{
    m := &menu{AccessToken:accessToken}
    m.createUrl = MENU_URL_CREATE + accessToken
    m.deleteUrl = MENU_URL_DELETE + accessToken
    m.getUrl = MENU_URL_GET + accessToken
}

func (this *menu) Create (v *MenuSt) (*ErrorSt, error) {
    data, err := json.Marshal(v)
    if err != nil {
        return nil, err
    }
    data, err = tools.HttpPost(this.createUrl, data)
    if err != nil {
        return nil, err
    }
    errst := &ErrorSt{}
    err = json.Unmarshal(data, errst)
    return errst, err
}

func (this *menu) Delete() (*ErrorSt, error) {
    data, err := tools.HttpGet(this.deleteUrl, data)
    if err != nil {
        return nil, err
    }
    errst := &ErrorSt{}
    err = json.Unmarshal(data, errst)
    return errst, err
}

func (this *menu) Get() (*ErrorSt, error) {
    data, err := tools.HttpGet(this.getUrl, data)
    if err != nil {
        return nil, err
    }
    errst := &ErrorSt{}
    err = json.Unmarshal(data, errst)
    return errst, err
}

