# Faker

Faker是一个为您生成数据的工具。无论您需要填充数据库、以进行压力测试，还是匿名化从生产服务中获取的数据，Faker都将为您服务。

Faker is heavily inspired by php's.

Faker requires go >= 1.11

[![Monthly Downloads](https://poser.pugx.org/fzaninotto/faker/d/monthly.png)](https://packagist.org/packages/fzaninotto/faker) [![Build Status](https://travis-ci.org/fzaninotto/Faker.svg?branch=master)](https://travis-ci.org/fzaninotto/Faker) [![SensioLabsInsight](https://insight.sensiolabs.com/projects/eceb78a9-38d4-4ad5-8b6b-b52f323e3549/mini.png)](https://insight.sensiolabs.com/projects/eceb78a9-38d4-4ad5-8b6b-b52f323e3549)

# Table of Contents

- [Installation](#installation)
- [Basic Usage](#basic-usage)
- [Formatters](#formatters)
	- [Base](#fakerproviderbase)
	- [Lorem Ipsum Text](#fakerproviderlorem)
	- [Person](#fakerprovideren_usperson)
	- [Address](#fakerprovideren_usaddress)
	- [Phone Number](#fakerprovideren_usphonenumber)
	- [Company](#fakerprovideren_uscompany)
	- [Real Text](#fakerprovideren_ustext)
	- [Date and Time](#fakerproviderdatetime)
	- [Internet](#fakerproviderinternet)
	- [User Agent](#fakerprovideruseragent)
	- [Payment](#fakerproviderpayment)
	- [Color](#fakerprovidercolor)
	- [File](#fakerproviderfile)
	- [Image](#fakerproviderimage)
	- [Uuid](#fakerprovideruuid)
	- [Barcode](#fakerproviderbarcode)
	- [Miscellaneous](#fakerprovidermiscellaneous)
	- [Biased](#fakerproviderbiased)
	- [Html Lorem](#fakerproviderhtmllorem)
- [Modifiers](#modifiers)
- [Localization](#localization)
- [Populating Entities Using an ORM or an ODM](#populating-entities-using-an-orm-or-an-odm)
- [Seeding the Generator](#seeding-the-generator)
- [Faker Internals: Understanding Providers](#faker-internals-understanding-providers)
- [Real Life Usage](#real-life-usage)
- [Language specific formatters](#language-specific-formatters)
- [Third-Party Libraries Extending/Based On Faker](#third-party-libraries-extendingbased-on-faker)
- [License](#license)


## Installation

```sh
go get github.com/sereiner/faker
```

## Basic Usage

Use `Faker\Factory::create()` to create and initialize a faker generator, which can generate data by accessing properties named after the type of data you want.

```php
<?php
// require the Faker autoloader
require_once '/path/to/Faker/src/autoload.php';
// alternatively, use another PSR-0 compliant autoloader (like the Symfony2 ClassLoader for instance)

// use the factory to create a Faker\Generator instance
$faker = Faker\Factory::create();

// generate data by accessing properties
echo $faker->name;
  // 'Lucy Cechtelar';
echo $faker->address;
  // "426 Jordy Lodge
  // Cartwrightshire, SC 88120-6700"
echo $faker->text;
  // Dolores sit sint laboriosam dolorem culpa et autem. Beatae nam sunt fugit
  // et sit et mollitia sed.
  // Fuga deserunt tempora facere magni omnis. Omnis quia temporibus laudantium
  // sit minima sint.
```

Even if this example shows a property access, each call to `$faker->name` yields a different (random) result. This is because Faker uses `__get()` magic, and forwards `Faker\Generator->$property` calls to `Faker\Generator->format($property)`.

```php
<?php
for ($i=0; $i < 10; $i++) {
  echo $faker->name, "\n";
}
  // Adaline Reichel
  // Dr. Santa Prosacco DVM
  // Noemy Vandervort V
  // Lexi O'Conner
  // Gracie Weber
  // Roscoe Johns
  // Emmett Lebsack
  // Keegan Thiel
  // Wellington Koelpin II
  // Ms. Karley Kiehn V
```

**Tip**: For a quick generation of fake data, you can also use Faker as a command line tool thanks to [faker-cli](https://github.com/bit3/faker-cli).

## Formatters

Each of the generator properties (like `name`, `address`, and `lorem`) are called "formatters". A faker generator has many of them, packaged in "providers". Here is a list of the bundled formatters in the default locale.

### `Faker\Provider\Base`

    randomDigit             // 7
    randomDigitNotNull      // 5
    randomNumber($nbDigits = NULL, $strict = false) // 79907610
    randomFloat($nbMaxDecimals = NULL, $min = 0, $max = NULL) // 48.8932
    numberBetween($min = 1000, $max = 9000) // 8567
    randomLetter            // 'b'
    // returns randomly ordered subsequence of a provided array
    randomElements($array = array ('a','b','c'), $count = 1) // array('c')
    randomElement($array = array ('a','b','c')) // 'b'
    shuffle('hello, world') // 'rlo,h eoldlw'
    shuffle(array(1, 2, 3)) // array(2, 1, 3)
    numerify('Hello ###') // 'Hello 609'
    lexify('Hello ???') // 'Hello wgt'
    bothify('Hello ##??') // 'Hello 42jz'
    asciify('Hello ***') // 'Hello R6+'
    regexify('[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}'); // sm0@y8k96a.ej

### `Faker\Provider\Lorem`

    word                                             // 'aut'
    words($nb = 3, $asText = false)                  // array('porro', 'sed', 'magni')
    sentence($nbWords = 6, $variableNbWords = true)  // 'Sit vitae voluptas sint non voluptates.'
    sentences($nb = 3, $asText = false)              // array('Optio quos qui illo error.', 'Laborum vero a officia id corporis.', 'Saepe provident esse hic eligendi.')
    paragraph($nbSentences = 3, $variableNbSentences = true) // 'Ut ab voluptas sed a nam. Sint autem inventore aut officia aut aut blanditiis. Ducimus eos odit amet et est ut eum.'
    paragraphs($nb = 3, $asText = false)             // array('Quidem ut sunt et quidem est accusamus aut. Fuga est placeat rerum ut. Enim ex eveniet facere sunt.', 'Aut nam et eum architecto fugit repellendus illo. Qui ex esse veritatis.', 'Possimus omnis aut incidunt sunt. Asperiores incidunt iure sequi cum culpa rem. Rerum exercitationem est rem.')
    text($maxNbChars = 200)                          // 'Fuga totam reiciendis qui architecto fugiat nemo. Consequatur recusandae qui cupiditate eos quod.'

### `Faker\Provider\en_US\Person`

    title($gender = null|'male'|'female')     // 'Ms.'
    titleMale                                 // 'Mr.'
    titleFemale                               // 'Ms.'
    suffix                                    // 'Jr.'
    name($gender = null|'male'|'female')      // 'Dr. Zane Stroman'
    firstName($gender = null|'male'|'female') // 'Maynard'
    firstNameMale                             // 'Maynard'
    firstNameFemale                           // 'Rachel'
    lastName                                  // 'Zulauf'

### `Faker\Provider\en_US\Address`

    cityPrefix                          // 'Lake'
    secondaryAddress                    // 'Suite 961'
    state                               // 'NewMexico'
    stateAbbr                           // 'OH'
    citySuffix                          // 'borough'
    streetSuffix                        // 'Keys'
    buildingNumber                      // '484'
    city                                // 'West Judge'
    streetName                          // 'Keegan Trail'
    streetAddress                       // '439 Karley Loaf Suite 897'
    postcode                            // '17916'
    address                             // '8888 Cummings Vista Apt. 101, Susanbury, NY 95473'
    country                             // 'Falkland Islands (Malvinas)'
    latitude($min = -90, $max = 90)     // 77.147489
    longitude($min = -180, $max = 180)  // 86.211205

### `Faker\Provider\en_US\PhoneNumber`

    phoneNumber             // '201-886-0269 x3767'
    tollFreePhoneNumber     // '(888) 937-7238'
    e164PhoneNumber     // '+27113456789'

### `Faker\Provider\en_US\Company`

    catchPhrase             // 'Monitored regional contingency'
    bs                      // 'e-enable robust architectures'
    company                 // 'Bogan-Treutel'
    companySuffix           // 'and Sons'
    jobTitle                // 'Cashier'

### `Faker\Provider\en_US\Text`

    realText($maxNbChars = 200, $indexSize = 2) // "And yet I wish you could manage it?) 'And what are they made of?' Alice asked in a shrill, passionate voice. 'Would YOU like cats if you were never even spoke to Time!' 'Perhaps not,' Alice replied."

### `Faker\Provider\DateTime`

    unixTime($max = 'now')                // 58781813
    dateTime($max = 'now', $timezone = null) // DateTime('2008-04-25 08:37:17', 'UTC')
    dateTimeAD($max = 'now', $timezone = null) // DateTime('1800-04-29 20:38:49', 'Europe/Paris')
    iso8601($max = 'now')                 // '1978-12-09T10:10:29+0000'
    date($format = 'Y-m-d', $max = 'now') // '1979-06-09'
    time($format = 'H:i:s', $max = 'now') // '20:49:42'
    dateTimeBetween($startDate = '-30 years', $endDate = 'now', $timezone = null) // DateTime('2003-03-15 02:00:49', 'Africa/Lagos')
    dateTimeInInterval($startDate = '-30 years', $interval = '+ 5 days', $timezone = null) // DateTime('2003-03-15 02:00:49', 'Antartica/Vostok')
    dateTimeThisCentury($max = 'now', $timezone = null)     // DateTime('1915-05-30 19:28:21', 'UTC')
    dateTimeThisDecade($max = 'now', $timezone = null)      // DateTime('2007-05-29 22:30:48', 'Europe/Paris')
    dateTimeThisYear($max = 'now', $timezone = null)        // DateTime('2011-02-27 20:52:14', 'Africa/Lagos')
    dateTimeThisMonth($max = 'now', $timezone = null)       // DateTime('2011-10-23 13:46:23', 'Antarctica/Vostok')
    amPm($max = 'now')                    // 'pm'
    dayOfMonth($max = 'now')              // '04'
    dayOfWeek($max = 'now')               // 'Friday'
    month($max = 'now')                   // '06'
    monthName($max = 'now')               // 'January'
    year($max = 'now')                    // '1993'
    century                               // 'VI'
    timezone                              // 'Europe/Paris'

Methods accepting a `$timezone` argument default to `date_default_timezone_get()`. You can pass a custom timezone string to each method, or define a custom timezone for all time methods at once using `$faker::setDefaultTimezone($timezone)`.

### `Faker\Provider\Internet`

    email                   // 'tkshlerin@collins.com'
    safeEmail               // 'king.alford@example.org'
    freeEmail               // 'bradley72@gmail.com'
    companyEmail            // 'russel.durward@mcdermott.org'
    freeEmailDomain         // 'yahoo.com'
    safeEmailDomain         // 'example.org'
    userName                // 'wade55'
    password                // 'k&|X+a45*2['
    domainName              // 'wolffdeckow.net'
    domainWord              // 'feeney'
    tld                     // 'biz'
    url                     // 'http://www.skilesdonnelly.biz/aut-accusantium-ut-architecto-sit-et.html'
    slug                    // 'aut-repellat-commodi-vel-itaque-nihil-id-saepe-nostrum'
    ipv4                    // '109.133.32.252'
    localIpv4               // '10.242.58.8'
    ipv6                    // '8e65:933d:22ee:a232:f1c1:2741:1f10:117c'
    macAddress              // '43:85:B7:08:10:CA'

### `Faker\Provider\UserAgent`

    userAgent              // 'Mozilla/5.0 (Windows CE) AppleWebKit/5350 (KHTML, like Gecko) Chrome/13.0.888.0 Safari/5350'
    chrome                 // 'Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5) AppleWebKit/5312 (KHTML, like Gecko) Chrome/14.0.894.0 Safari/5312'
    firefox                // 'Mozilla/5.0 (X11; Linuxi686; rv:7.0) Gecko/20101231 Firefox/3.6'
    safari                 // 'Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_7_1 rv:3.0; en-US) AppleWebKit/534.11.3 (KHTML, like Gecko) Version/4.0 Safari/534.11.3'
    opera                  // 'Opera/8.25 (Windows NT 5.1; en-US) Presto/2.9.188 Version/10.00'
    internetExplorer       // 'Mozilla/5.0 (compatible; MSIE 7.0; Windows 98; Win 9x 4.90; Trident/3.0)'

### `Faker\Provider\Payment`

    creditCardType          // 'MasterCard'
    creditCardNumber        // '4485480221084675'
    creditCardExpirationDate // 04/13
    creditCardExpirationDateString // '04/13'
    creditCardDetails       // array('MasterCard', '4485480221084675', 'Aleksander Nowak', '04/13')
    // Generates a random IBAN. Set $countryCode to null for a random country
    iban($countryCode)      // 'IT31A8497112740YZ575DJ28BP4'
    swiftBicNumber          // 'RZTIAT22263'

### `Faker\Provider\Color`

    hexcolor               // '#fa3cc2'
    rgbcolor               // '0,255,122'
    rgbColorAsArray        // array(0,255,122)
    rgbCssColor            // 'rgb(0,255,122)'
    safeColorName          // 'fuchsia'
    colorName              // 'Gainsbor'

### `Faker\Provider\File`

    fileExtension          // 'avi'
    mimeType               // 'video/x-msvideo'
    // Copy a random file from the source to the target directory and returns the fullpath or filename
    file($sourceDir = '/tmp', $targetDir = '/tmp') // '/path/to/targetDir/13b73edae8443990be1aa8f1a483bc27.jpg'
    file($sourceDir, $targetDir, false) // '13b73edae8443990be1aa8f1a483bc27.jpg'

### `Faker\Provider\Image`

    // Image generation provided by LoremPixel (http://lorempixel.com/)
    imageUrl($width = 640, $height = 480) // 'http://lorempixel.com/640/480/'
    imageUrl($width, $height, 'cats')     // 'http://lorempixel.com/800/600/cats/'
    imageUrl($width, $height, 'cats', true, 'Faker') // 'http://lorempixel.com/800/400/cats/Faker'
    imageUrl($width, $height, 'cats', true, 'Faker', true) // 'http://lorempixel.com/grey/800/400/cats/Faker/' Monochrome image
    image($dir = '/tmp', $width = 640, $height = 480) // '/tmp/13b73edae8443990be1aa8f1a483bc27.jpg'
    image($dir, $width, $height, 'cats')  // 'tmp/13b73edae8443990be1aa8f1a483bc27.jpg' it's a cat!
    image($dir, $width, $height, 'cats', false) // '13b73edae8443990be1aa8f1a483bc27.jpg' it's a filename without path
    image($dir, $width, $height, 'cats', true, false) // it's a no randomize images (default: `true`)
    image($dir, $width, $height, 'cats', true, true, 'Faker') // 'tmp/13b73edae8443990be1aa8f1a483bc27.jpg' it's a cat with 'Faker' text. Default, `null`.

### `Faker\Provider\Uuid`

    uuid                   // '7e57d004-2b97-0e7a-b45f-5387367791cd'

### `Faker\Provider\Barcode`

    ean13          // '4006381333931'
    ean8           // '73513537'
    isbn13         // '9790404436093'
    isbn10         // '4881416324'

### `Faker\Provider\Miscellaneous`

    boolean // false
    boolean($chanceOfGettingTrue = 50) // true
    md5           // 'de99a620c50f2990e87144735cd357e7'
    sha1          // 'f08e7f04ca1a413807ebc47551a40a20a0b4de5c'
    sha256        // '0061e4c60dac5c1d82db0135a42e00c89ae3a333e7c26485321f24348c7e98a5'
    locale        // en_UK
    countryCode   // UK
    languageCode  // en
    currencyCode  // EUR
    emoji         // 😁

### `Faker\Provider\Biased`

    // get a random number between 10 and 20,
    // with more chances to be close to 20
    biasedNumberBetween($min = 10, $max = 20, $function = 'sqrt')



## License

Faker is released under the MIT Licence. See the bundled LICENSE file for details.
