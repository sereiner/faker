package internet

var freeEmailDomain = []string{"gmail.com", "yahoo.com", "hotmail.com"}
var domainName = []string{"gmail.com", "yahoo.com", "hotmail.com"}
var tld = []string{"com", "com", "com", "com", "com", "com", "biz", "info", "net", "org"}

var userNameFormats = []string{
	"{{lastName}}.{{firstName}}",
	"{{firstName}}.{{lastName}}",
	"{{firstName}}##",
	"?{{lastName}}",
}
var emailFormats = []string{
	"{{userName}}@{{freeEmailDomain}}",
}
var urlFormats = []string{
	"http://www.{{domainName}}/",
	"http://{{domainName}}/",
	"http://www.{{domainName}}/{{slug}}",
	"http://www.{{domainName}}/{{slug}}",
	"https://www.{{domainName}}/{{slug}}",
	"http://www.{{domainName}}/{{slug}}.html",
	"http://{{domainName}}/{{slug}}",
	"http://{{domainName}}/{{slug}}",
	"http://{{domainName}}/{{slug}}.html",
	"https://{{domainName}}/{{slug}}.html",
}
