package types

import (
	"encoding/json"

	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
)

type SuccessfulAcknowledgement struct {
	Result []byte `json:"result"`
}

type FailedAcknowledgement struct {
	Error string `json:"error"`
}

var (
	_ ibcexported.Acknowledgement = (*SuccessfulAcknowledgement)(nil)
	_ ibcexported.Acknowledgement = (*FailedAcknowledgement)(nil)
)

func NewSuccessfulAcknowledgement(r string) SuccessfulAcknowledgement {
	return SuccessfulAcknowledgement{Result: []byte(r)}
}

func (SuccessfulAcknowledgement) Success() bool {
	return true
}

func (a SuccessfulAcknowledgement) Acknowledgement() []byte {
	if bz, err := json.Marshal(a); err != nil {
		panic(err)
	} else {
		return bz
	}
}

func NewFailedAcknowledgement(e string) FailedAcknowledgement {
	return FailedAcknowledgement{Error: e}
}

func (FailedAcknowledgement) Success() bool {
	return false
}

func (a FailedAcknowledgement) Acknowledgement() []byte {
	if bz, err := json.Marshal(a); err != nil {
		panic(err)
	} else {
		return bz
	}
}
