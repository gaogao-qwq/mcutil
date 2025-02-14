package query_test

import (
	"context"
	"testing"

	"github.com/gaogao-qwq/mcutil/v4/query"
	"github.com/gaogao-qwq/mcutil/v4/util"
)

func TestFull(t *testing.T) {
	resp, err := query.Full(context.Background(), "demo.mcstatus.io", util.DefaultJavaPort)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", resp)
}
