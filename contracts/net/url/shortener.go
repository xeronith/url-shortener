package url

type ShortenerType int64

const (
	DefaultShortener ShortenerType = iota
)

type Shortener interface {
	BaseUrl() string
	UrlIsValid(url string) bool
	Shorten(url string) (string, error)
	Expand(url string) (string, error)
}
