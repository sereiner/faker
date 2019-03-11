package company

import (
	"bytes"
	"io"

	"github.com/sereiner/faker"
	"github.com/sereiner/faker/base"
	"github.com/valyala/fasttemplate"
)

var template *fasttemplate.Template

func init() {
	template = fasttemplate.New(base.RandomElement(formats), "${", "}")
}

// CatchPhrase 生成一个广告语
func CatchPhrase() string {
	return base.RandomElement(catchPhrase)
}

// Company 生成一个公司名字
func Company() string {

	buf := faker.BufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer faker.BufferPool.Put(buf)

	_, err := template.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
		switch tag {
		case "companyPrefix":
			return w.Write([]byte(base.RandomElement(companyPrefix)))
		case "companySuffix":
			return w.Write([]byte(base.RandomElement(companySuffix)))
		}
		return 0, nil
	})
	if err == nil {
		return buf.String()
	}
	return ""
}

// CompanySuffix 生成一个公司后缀
func CompanySuffix() string {
	return base.RandomElement(companySuffix)
}

// JobTitle 生成一个工作名称
func JobTitle() string {
	return base.RandomElement(jobTitleFormat)
}
