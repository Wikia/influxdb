package api

import (
	libhttp "net/http"
)

type RequestContext struct {
	remoteAddr            string
}

func NewEmptyRequestContext() *RequestContext {
	requestContext := &RequestContext{
		remoteAddr:            "",
	}

	return requestContext
}

func NewRequestContext(r *libhttp.Request) *RequestContext {
	requestContext := &RequestContext{
		remoteAddr:            r.RemoteAddr,
	}

	return requestContext
}

func (self RequestContext) RemoteAddr() string {
	return self.remoteAddr
}
