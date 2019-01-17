package proc

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMakeResp(t *testing.T) {
	res := makeTestResp()
	data, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(data))
}
