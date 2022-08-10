package internal

import (
	"bytes"
	"encoding/json"
	"github.com/go-njn/gonjn030-api-client/pkg/domain"
	"io"
	"net/url"
	"path"
	"strconv"
)

type HeaderName = string
type HeaderValue = string

const (
	ContentType   HeaderName = "Content-Type"
	Authorization HeaderName = "Authorization"
)

const (
	ApplicationJson HeaderValue = "application/json"
)

func joinPath(basePath string, id domain.ItemId) string {
	baseUrl, _ := url.Parse(basePath)
	baseUrl.Path = path.Join(baseUrl.Path, strconv.Itoa(int(id)))

	return baseUrl.String()
}

func asJSON(data any) io.Reader {
	if data == nil {
		return nil
	}

	s, _ := json.Marshal(data)

	return bytes.NewReader(s)
}
