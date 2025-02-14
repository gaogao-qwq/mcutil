package status_test

import (
	"context"
	"testing"

	"github.com/gaogao-qwq/mcutil/v4/status"
	"github.com/gaogao-qwq/mcutil/v4/util"
)

func TestBedrock(t *testing.T) {
	resp, err := status.Bedrock(context.Background(), "demo.mcstatus.io", util.DefaultBedrockPort)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", resp)
}
