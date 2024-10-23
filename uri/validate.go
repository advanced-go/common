package uri

import (
	"errors"
	"fmt"
	"net/url"
)

const (
	AuthorityPath     = "authority"
	AuthorityRootPath = "/" + AuthorityPath
)

// ValidateURL - validate a URL
func ValidateURL(url *url.URL, authority string) (p *Parsed, err error) {
	if url == nil {
		return &Parsed{}, errors.New("error: URL is nil")
	}
	if len(authority) == 0 {
		return &Parsed{}, errors.New("error: authority is empty")
	}
	if url.Path == AuthorityRootPath {
		return &Parsed{Path: AuthorityPath}, nil
	}
	if url.RawQuery != "" {
		p = Uproot(url.Path + "?" + url.RawQuery)
	} else {
		p = Uproot(url.Path)
	}
	if !p.Valid {
		return p, p.Err
	}
	if p.Authority != authority {
		return p, errors.New(fmt.Sprintf("error: invalid URI, authority does not match: \"%v\" \"%v\"", url.Path, authority))
	}
	if len(p.Path) == 0 {
		return p, errors.New(fmt.Sprintf("error: invalid URI, path only contains an authority: \"%v\"", url.Path))
	}
	return p, nil
}
