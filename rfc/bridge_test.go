package rfc

import (
	"github.com/darkweak/souin/cache/surrogate"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/darkweak/souin/cache/providers"
	"github.com/darkweak/souin/cache/types"
	"github.com/darkweak/souin/cache/ykeys"
	"github.com/darkweak/souin/errors"
	"github.com/darkweak/souin/tests"
)

func commonInitializer() (*http.Request, types.AbstractProviderInterface, *VaryTransport) {
	c := tests.MockConfiguration(tests.BaseConfiguration)
	prs := providers.InitializeProvider(c)
	r := httptest.NewRequest("GET", "http://domain.com/testing", nil)
	httptest.NewRecorder()

	return r, prs, NewTransport(prs, ykeys.InitializeYKeys(c.Ykeys), surrogate.InitializeSurrogate(c))
}

func TestCachedResponse_WithUpdate(t *testing.T) {
	r, c, v := commonInitializer()
	res, err := CachedResponse(c, r, GetCacheKey(r), v, true)

	if err != nil {
		errors.GenerateError(t, "CachedResponse cannot throw error")
	}

	if res != nil {
		errors.GenerateError(t, "Result from cached response should be a valid response")
	}
}

func TestCachedResponse_WithoutUpdate(t *testing.T) {
	r, c, v := commonInitializer()
	key := GetCacheKey(r)
	res, err := CachedResponse(c, r, key, v, false)

	if err != nil {
		errors.GenerateError(t, "CachedResponse cannot throw error")
	}

	if res != nil {
		errors.GenerateError(t, "Result from cached response should be a valid response")
	}
}
