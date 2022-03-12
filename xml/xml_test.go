package xml

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

type Books struct {
	XMLName xml.Name `xml:"books"`
	Nums    int      `xml:"nums,attr"`
	Book    []Book   `xml:"book"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	Name    string   `xml:"name,attr"`
	Author  string   `xml:"author"`
	Time    string   `xml:"time"`
	Types   []string `xml:"types>type"`
	Test    string   `xml:",any"`
}

func TestXMLCodec_Encode(t *testing.T) {
	codec := XMLCodec{}
	convey.Convey("TestXMLCodec_Encode", t, func() {
		convey.Convey("success", func() {
			data := &Books{
				XMLName: xml.Name{
					Local: "books",
				},
				Nums: 2,
				Book: []Book{
					{
						XMLName: xml.Name{
							Local: "book",
						},
						Name:   "思想",
						Author: "小张",
						Time:   "2018年1月20日",
						Types: []string{
							"教育",
							"历史",
						},
						Test: "我是多余的",
					},
					{
						XMLName: xml.Name{Local: "book"},
						Name:    "政治",
						Author:  "小王",
						Time:    "2018年1月20日",
						Types: []string{
							"科学",
							"人文",
						},
						Test: "我是多余的",
					},
				},
			}
			r, err := codec.Encode(data)
			assert.Equal(t, err, nil)
			b, err := ioutil.ReadAll(r)
			assert.Equal(t, err, nil)
			assert.Equal(t, string(b), `<books nums="2"><book name="思想"><author>小张</author><time>2018年1月20日</time><types><type>教育</type><type>历史</type></types><Test>我是多余的</Test></book><book name="政治"><author>小王</author><time>2018年1月20日</time><types><type>科学</type><type>人文</type></types><Test>我是多余的</Test></book></books>`)
		})
	})
}

func TestXMLCodec_Decode(t *testing.T) {
	codec := XMLCodec{}
	convey.Convey("TestXMLCodec_Decode", t, func() {
		convey.Convey("success", func() {
			data := `<?xml version="1.0" encoding="utf-8"?>
            <books nums="2">
                <book name="思想">
                    <author>小张</author>
                    <time>2018年1月20日</time>
                    <types>
                        <type>教育</type>
                        <type>历史</type>
                    </types>
                    <test>我是多余的</test>
                </book>
                <book name="政治">
                    <author>小王</author>
                    <time>2018年1月20日</time>
                    <types>
                        <type>科学</type>
                        <type>人文</type>
                    </types>
                    <test>我是多余的</test>
                </book>
            </books>`
			dst := &Books{}
			res := &Books{
				XMLName: xml.Name{
					Local: "books",
				},
				Nums: 2,
				Book: []Book{
					{
						XMLName: xml.Name{
							Local: "book",
						},
						Name:   "思想",
						Author: "小张",
						Time:   "2018年1月20日",
						Types: []string{
							"教育",
							"历史",
						},
						Test: "我是多余的",
					},
					{
						XMLName: xml.Name{Local: "book"},
						Name:    "政治",
						Author:  "小王",
						Time:    "2018年1月20日",
						Types: []string{
							"科学",
							"人文",
						},
						Test: "我是多余的",
					},
				},
			}

			err := codec.Decode(bytes.NewReader([]byte(data)), dst)
			assert.Equal(t, err, nil)
			assert.Equal(t, dst, res)
		})
	})
}
