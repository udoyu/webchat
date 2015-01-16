package pubacc

type string_t string

func NewStringT(s string_t) *string_t {
    return &s
}

func (this *string_t) StringT() string_t {
    return *this
}

func (this *string_t) SetStringT(s string_t) {
    this = &s
}


