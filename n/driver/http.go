package driver

import "github.com/b2wdigital/fxstack/n/channel/http"

type HTTP interface {
	Start(route ...*http.Route) error
}
