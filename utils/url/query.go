package url

import (
	"net/url"
)

// AppendQueryParam appends a query parameter at the end of the given url.
func AppendQueryParam(originalURL, key, value string) (string, error) {
	urlObject, err := url.Parse(originalURL)
	if err != nil {
		return "", err
	}
	queryParams := urlObject.Query()
	queryParams.Add(key, value)
	urlObject.RawQuery = queryParams.Encode()

	return urlObject.String(), nil
}
