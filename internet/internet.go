// email                   // 'tkshlerin@collins.com'
// safeEmail               // 'king.alford@example.org'
// freeEmail               // 'bradley72@gmail.com'
// companyEmail            // 'russel.durward@mcdermott.org'
// freeEmailDomain         // 'yahoo.com'
// safeEmailDomain         // 'example.org'
// userName                // 'wade55'
// password                // 'k&|X+a45*2['
// domainName              // 'wolffdeckow.net'
// domainWord              // 'feeney'
// tld                     // 'biz'
// url                     // 'http://www.skilesdonnelly.biz/aut-accusantium-ut-architecto-sit-et.html'
// slug                    // 'aut-repellat-commodi-vel-itaque-nihil-id-saepe-nostrum'
// ipv4                    // '109.133.32.252'
// localIpv4               // '10.242.58.8'
// ipv6                    // '8e65:933d:22ee:a232:f1c1:2741:1f10:117c'
// macAddress              // '43:85:B7:08:10:CA'

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

// SafeEmail 生成一个安全的邮箱
func SafeEmail() string {
	return person.FirstNameMale() + SafeEmailDomain()
}

//SafeEmailDomain  生成一个安全的邮箱domain
func SafeEmailDomain() string {
	safeDomain := []string{"example.com", "example.org", "example.net"}
	return base.RandomElement(safeDomain)
}

//FreeEmailDomain 生成一个免费的邮箱domain
func FreeEmailDomain() string {
	return base.RandomElement(freeEmailDomain)
}

func DomainName() string {
	return base.RandomElement(domainName)
}

// Tld
func Tld() string {
	return base.RandomElement(tld)
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
		case "lastName":
			return w.Write([]byte(person.LastName()))
		case "domainName":
			return w.Write([]byte(base.RandomElement(domainName)))
		case "freeEmailDomain":
			return w.Write([]byte(base.RandomElement(freeEmailDomain)))
		}
		return 0, nil
	})
	if err == nil {
		return buf.String()
	}
	return ""
}
