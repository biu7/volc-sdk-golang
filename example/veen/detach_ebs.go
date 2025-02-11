package veen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volcengine/volc-sdk-golang/service/veen"
	"testing"
)

func DetachEbs(t *testing.T) {
	resp, err := veen.NewInstance(ak, sk).DetachEbs(&veen.DetachEbsReq{
		EbsIds: []string{"disk-t9p44586fn6cbs9", "disk-fz26fhklggjnbhj"},
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Printf("%+v", veen.ToJson(resp))
}
