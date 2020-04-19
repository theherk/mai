package macaddress_test

//go:generate mockgen -destination=mocks/client/mock.go github.com/theherk/mai/pkg/macaddress Client

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/theherk/mai/pkg/macaddress"
	mock_client "github.com/theherk/mai/pkg/macaddress/mocks/client"
)

// equateErrorMessage reports errors to be equal if both are nil
// or both have the same message.
//
// This is taken from go-cmp documentation for the EquateErrors feature.
// By default this does not compare the strings of errors, due to the
// fact that errors can be anything.
var equateErrorMessage = cmp.Comparer(func(x, y error) bool {
	if x == nil || y == nil {
		return x == nil && y == nil
	}
	return x.Error() == y.Error()
})

func ExampleAPI_Get_noKey() {
	_, err := macaddress.API{Key: ""}.Get("test-query")
	fmt.Println(err)
	// output:
	// key empty; no call
}

func TestAPI_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// case: simple
	simpleClient := mock_client.NewMockClient(ctrl)
	simpleResB := []byte("test-info")
	simpleRes := new(http.Response)
	simpleResReadCloser := ioutil.NopCloser(bytes.NewBuffer(simpleResB))
	simpleRes.Body = simpleResReadCloser
	simpleClient.EXPECT().Do(gomock.Any()).Return(simpleRes, nil)

	// case: request err
	reqErrClient := mock_client.NewMockClient(ctrl)
	reqErrResB := []byte("test-info")
	reqErrRes := new(http.Response)
	reqErrResReadCloser := ioutil.NopCloser(bytes.NewBuffer(reqErrResB))
	reqErrRes.Body = reqErrResReadCloser
	reqErrClient.EXPECT().Do(gomock.Any()).Return(reqErrRes, errors.New("test-error"))

	tt := []struct {
		name   string
		client macaddress.Client
		info   string
		err    error
	}{
		{
			name:   "simple",
			client: simpleClient,
			info:   "test-info",
		},
		{
			name:   "request error",
			client: reqErrClient,
			err:    errors.New("test-error"),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			api := macaddress.API{Key: "test-key"}
			api.Client = tc.client
			info, err := api.Get("test-query")
			if !cmp.Equal(&err, &tc.err, equateErrorMessage) {
				t.Errorf("unexpected error; got: %v, want: %v", err, tc.err)
			}
			if info != tc.info {
				t.Errorf("unexpected info; got: %s, want: %s", info, tc.info)
			}
		})
	}
}
