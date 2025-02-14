package status_test

import (
	"context"
	"testing"

	"github.com/gaogao-qwq/mcutil/v4/status"
	"github.com/gaogao-qwq/mcutil/v4/util"
)

func TestModernRaw(t *testing.T) {
	resp, err := status.ModernRaw(context.Background(), "demo.mcstatus.io", util.DefaultJavaPort)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", resp)
}
