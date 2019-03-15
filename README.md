# Faker

Faker是一个为您生成数据的工具。无论您需要填充数据库、以进行压力测试，还是匿名化从生产服务中获取的数据，Faker都将为您服务。

Faker requires go >= 1.9



## Installation

```sh
go get github.com/sereiner/faker
```


### `Base`


    
    RandomDigit()    //返回 0 到 9 之间随机数

    RandomDigitNotNull() //RandomDigitNotNull 返回 1 到 9 之间随机数

    RandomDigitNot(except int) //RandomDigitNot 返回一个 0 到 9 之间的随机数,但是不包括 except

    RandomFloat32() //RandomFloat32 返回一个 float32

    func RandomFloat64() //RandomFloat64 返回一个 float64

    RandInt(min, max int) //RandInt 返回 min 到 max 之间的随机数

    NumberBetween(min, max int) //NumberBetween 返回 min 到 max 之间的随机数

    RandomElements(a []string, count int) //RandomElements 从给定的数组中获取count个不重复的随机元素

    RandomElement(a []string) //RandomElement 从给定的数组中获取一个随机元素

    Numerify(format ...string) // Numerify 根据填充字符串,返回一个随机数,默认返回 5位

    Lexify(format ...string) // Lexify 根据填充字符串,返回一个随机字符串,默认返回 5位


### `Lorem`

    还没开始写

### `Person`

    Name(gender string) //随机返回一个人的名字,默认是男人的名字

    FirstName(gender string) //返回一个FirstName,默认还是男人的

    FirstNameMale() //返回一个男性的FirstName

    FirstNameFemale() //返回一个女性的FirstName

    LastName() //返回一个lastname

    Title(gender string) //返回一个title

    TitleMale()   //Mr.", "Dr.", "Prof.

    TitleFemale()  

    Suffix() //"Jr.", "Sr.", "I", "II",


### `Address`

    CityPrefix() //CityPrefix 城市前缀 

    State() //区

    StateAbbr()

    CitySuffix() //城市后缀

    StreetSuffix() //街道前缀

    BuildingNumber() //建筑物编号

    City()

    StreetName()  //街道名字

    StreetAddress() //街道详细地址

    Postcode() //邮编

    Address() //地址

    Country()  //国家
    
    Latitude() //纬度

    Longitude() //经度


### `PhoneNumber`

    PhoneNumber()             // 18312341234

### `Faker\Provider\en_US\Company`

    
    CatchPhrase()     // CatchPhrase 生成一个广告语

    Company()         // Company 生成一个公司名字

    CompanySuffix()   // CompanySuffix 生成一个公司后缀

    JobTitle()        // JobTitle 生成一个工作名称
	

### `Text`

    // RealText 通过马尔可夫链算法生成文本字符串
    RealText(size ...int) string

### `DateTime`

    还没写

### `Internet`

    
    Email()       // Email 生成一个 email

### `UserAgent`

    

### `Payment`



### `Color`

    

### `File`


### `Image`


### `Uuid`


### `Barcode`


### `Miscellaneous`



## License

Faker is released under the MIT Licence. See the bundled LICENSE file for details.
