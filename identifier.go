package pizzameeting

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
)

type Identifier struct {
	id string
}

func (id *Identifier) ID() string {
	if id.id == "" {
		byteId := make([]byte, 12)
		rand.Read(byteId)
		id.id = hex.EncodeToString(byteId)
	}
	return id.id
}

func (id *Identifier) MarshalJSON() ([]byte, error) {
	var ret struct {
		Id string `json:"Id"`
	}
	ret.Id = id.ID()
	return json.Marshal(ret)
}

func (id *Identifier) UnmarshalJSON(raw []byte) error {
	var ret struct {
		Id string `json:"Id"`
	}
	err := json.Unmarshal(raw, ret)
	if err != nil {
		return err
	}

	id.id = ret.Id
	return nil
}
