package kvdto

import (
	"time"

	"github.com/sunjin110/folio/golio/utils"
)

type Metadata struct {
	Expiration *int64 `json:"expiration,omitempty"`
}

func NewMetadata(expiration *time.Time) *Metadata {
	if expiration == nil {
		return &Metadata{}
	}
	return &Metadata{
		Expiration: utils.ToPointer(expiration.Unix()),
	}
}

type PathInput struct {
	AccountID   string
	NamespaceID string
	KeyName     string
}

type WriteKVPairWithMetadataOutput struct {
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
	Result   interface{}   `json:"result"`
	Success  bool          `json:"success"`
}
