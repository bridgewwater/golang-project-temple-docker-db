# test

## github.com/stretchr/testify

[github.com/stretchr/testify](https://github.com/stretchr/testify)

```bash
GO111MODULE=on go mod edit -require='github.com/stretchr/testify@v1.4.0'
GO111MODULE=on go mod vendor
```

fast use

```go
import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

  // assert equality
  assert.Equal(t, 123, 123, "they should be equal")

  // assert inequality
  assert.NotEqual(t, 123, 456, "they should not be equal")

  // assert for nil (good for errors)
  assert.Nil(t, object)

  // assert for not nil (good when you expect something)
  if assert.NotNil(t, object) {

    // now we know that object isn't nil, we are safe to make
    // further assertions without causing any errors
    assert.Equal(t, "Something", object.Value)

  }
}
```
## github.com/bxcodec/faker faker

[github.com/bxcodec/faker](https://github.com/bxcodec/faker)

```go
import (
    "github.com/bxcodec/faker/v3"
)
type SomeStructWithTags struct {
	Email              string  `faker:"email"`
	DomainName         string  `faker:"domain_name"`
	IPV4               string  `faker:"ipv4"`
	IPV6               string  `faker:"ipv6"`
	Password           string  `faker:"password"`
	PhoneNumber        string  `faker:"phone_number"`
	MacAddress         string  `faker:"mac_address"`
	URL                string  `faker:"url"`

	UserName           string  `faker:"username"`
	TitleMale          string  `faker:"title_male"`
	TitleFemale        string  `faker:"title_female"`
	FirstName          string  `faker:"first_name"`
	FirstNameMale      string  `faker:"first_name_male"`
	FirstNameFemale    string  `faker:"first_name_female"`
	LastName           string  `faker:"last_name"`
	Name               string  `faker:"name"`

	UnixTime           int64   `faker:"unix_time"`
	Date               string  `faker:"date"`
	Time               string  `faker:"time"`
	MonthName          string  `faker:"month_name"`
	Year               string  `faker:"year"`
	DayOfWeek          string  `faker:"day_of_week"`
	DayOfMonth         string  `faker:"day_of_month"`
	Timestamp          string  `faker:"timestamp"`
	Century            string  `faker:"century"`
	TimeZone           string  `faker:"timezone"`

	Word               string  `faker:"word"`
	Sentence           string  `faker:"sentence"`
	ASString []string          `faker:"len=50"`
	SString  string            `faker:"len=25"`

	UUIDHypenated      string  `faker:"uuid_hyphenated"`
	UUID               string  `faker:"uuid_digit"`

	Skip               string  `faker:"-"`
}
func TestSomething(t *testing.T) {
	_ = faker.SetRandomMapAndSliceSize(20) // Random generated map or array size wont exceed 20... 
	a := SomeStruct{}
	_ = faker.FakeData(&a)
	fmt.Printf("%+v", a)
}
```

demo:
- [example_with_tags_test.go](https://github.com/bxcodec/faker/blob/master/example_with_tags_test.go)
- [example_with_tags_lenbounds_test.go](https://github.com/bxcodec/faker/blob/master/example_with_tags_lenbounds_test.go)