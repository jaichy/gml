package gml

const (
	elBr uint8 = iota + 100 // offset to avoid accidental overlap
	elHr
	elBase
	elLink
	elMeta
	elInput
	elWbr
	elCol
	elArea
	elEmbed
	elImg
	elParam
	elSource
	elTrack
	elDoctypeHTML
	elBasefont
	elBgsound
	elCommand
	elFrame
	elImage
	elIsindex
	elKeygen
	elMenuitem
	elNextid
	elSpacer
)

var voidElementNames = map[uint8][]byte{
	elBr:          []byte("br"),
	elHr:          []byte("hr"),
	elBase:        []byte("base"),
	elLink:        []byte("link"),
	elMeta:        []byte("meta"),
	elInput:       []byte("input"),
	elWbr:         []byte("wbr"),
	elCol:         []byte("col"),
	elArea:        []byte("area"),
	elEmbed:       []byte("embed"),
	elImg:         []byte("img"),
	elParam:       []byte("param"),
	elSource:      []byte("source"),
	elTrack:       []byte("track"),
	elDoctypeHTML: []byte("!DOCTYPE html"),
	elBasefont: []byte("basefont"),
	elBgsound:  []byte("bgsound"),
	elCommand:  []byte("command"),
	elFrame:    []byte("frame"),
	elImage:    []byte("image"),
	elIsindex:  []byte("isindex"),
	elKeygen:   []byte("keygen"),
	elMenuitem: []byte("menuitem"),
	elNextid:   []byte("nextid"),
	elSpacer:   []byte("spacer"),
}

// <br { attr }>
func Br() *gmlElement { return newGmlElement(elBr, true) }

// <hr { attr }>
func Hr() *gmlElement { return newGmlElement(elHr, true) }

// <base { attr }>
func Base() *gmlElement { return newGmlElement(elBase, true) }

// <link { attr }>
func Link() *gmlElement { return newGmlElement(elLink, true) }

// <meta { attr }>
func Meta() *gmlElement { return newGmlElement(elMeta, true) }

// <input { attr }>
func Input() *gmlElement { return newGmlElement(elInput, true) }

// <wbr { attr }>
func Wbr() *gmlElement { return newGmlElement(elWbr, true) }

// <col { attr }>
func Col() *gmlElement { return newGmlElement(elCol, true) }

// <area { attr }>
func Area() *gmlElement { return newGmlElement(elArea, true) }

// <embed { attr }>
func Embed() *gmlElement { return newGmlElement(elEmbed, true) }

// <img { attr }>
func Img() *gmlElement { return newGmlElement(elImg, true) }

// <param { attr }>
func Param() *gmlElement { return newGmlElement(elParam, true) }

// <source { attr }>
func Source() *gmlElement { return newGmlElement(elSource, true) }

// <track { attr }>
func Track() *gmlElement { return newGmlElement(elTrack, true) }

// <!DOCTYPE html>
func DocTypeHtml() *gmlElement {
	return newGmlElement(elDoctypeHTML, true)
}


func Basefont() *gmlElement { return newGmlElement(elBasefont, true) }
func Bgsound() *gmlElement  { return newGmlElement(elBgsound, true) }
func Command() *gmlElement  { return newGmlElement(elCommand, true) }
func Frame() *gmlElement    { return newGmlElement(elFrame, true) }
func Image() *gmlElement    { return newGmlElement(elImage, true) } 
func Isindex() *gmlElement  { return newGmlElement(elIsindex, true) }
func Keygen() *gmlElement   { return newGmlElement(elKeygen, true) }
func Menuitem() *gmlElement { return newGmlElement(elMenuitem, true) }
func Nextid() *gmlElement   { return newGmlElement(elNextid, true) }
func Spacer() *gmlElement   { return newGmlElement(elSpacer, true) }
