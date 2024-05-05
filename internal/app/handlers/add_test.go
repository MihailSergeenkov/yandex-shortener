package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddHandler(t *testing.T) {
	type request struct {
		method string
		body   string
	}

	type want struct {
		code int
	}
	tests := []struct {
		name    string
		request request
		want    want
	}{
		{
			name: "success add url",
			request: request{
				method: http.MethodPost,
				body:   "https://ya.ru/some",
			},
			want: want{
				code: 201,
			},
		},
		{
			name: "when does not Post method",
			request: request{
				method: http.MethodGet,
				body:   "https://ya.ru/some",
			},
			want: want{
				code: 400,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(test.request.method, "/", nil)
			w := httptest.NewRecorder()
			AddHandler(w, request)

			res := w.Result()

			assert.Equal(t, test.want.code, res.StatusCode)

			if test.want.code == 201 {
				defer res.Body.Close()
				resBody, err := io.ReadAll(res.Body)

				require.NoError(t, err)
				assert.NotEmpty(t, resBody)
			}

		})
	}
}
