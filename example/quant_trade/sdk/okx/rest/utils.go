package rest

import (
	"example/quant_trade/sdk/okx/rest/responses"
	"fmt"
)

func checkBasic(err error, b *responses.Basic) error {
	if err != nil {
		return err
	}
	if b.Code != 0 {
		return fmt.Errorf("failed | Code = %d, Msg = %s", b.Code, b.Msg)
	}
	return nil
}

func S2M(i interface{}) map[string]string {
	m := make(map[string]string)
	j, _ := json.Marshal(i)
	_ = json.Unmarshal(j, &m)

	return m
}
