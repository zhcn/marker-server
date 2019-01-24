package proc

import (
	"encoding/json"
	"fmt"
	"market-server/server/common/types"
	"testing"
)

func TestMakeResp(t *testing.T) {
	var pos types.Position
	res, _ := GetNearShops(&pos, nil)
	data, _ := json.MarshalIndent(res, "", " ")
	fmt.Println(string(data))
}
