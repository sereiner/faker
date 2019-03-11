package internet

import (
	"bytes"
	"io"

	"github.com/sereiner/faker/person"

	"github.com/sereiner/faker"
	"github.com/sereiner/faker/base"
	"github.com/valyala/fasttemplate"
)

var template *fasttemplate.Template

func init() {
	template = fasttemplate.New(base.RandomElement(emailFormats), "{{", "}}")
}

// Email 生成一个 email
func Email() string {
	format := base.RandomElement(emailFormats)
	return formatString(format)
}

func formatString(format string) string {
	buf := faker.BufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer faker.BufferPool.Put(buf)
	template.Reset(format, "{{", "}}")
	_, err := template.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
		switch tag {
		case "userName":
			return w.Write([]byte(person.LastName()))
		case "domainName":
			return w.Write([]byte(base.RandomElement(freeEmailDomain)))
		}
		return 0, nil
	})
	if err == nil {
		return buf.String()
	}
	return ""
}
