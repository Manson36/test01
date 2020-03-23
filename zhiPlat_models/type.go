package zhiPlat_models

import (
	"strconv"
	"strings"
	"time"
)

type JsonbNumArray []int64
type IDType int64

func (t IDType) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte(`""`), nil
	}

	return []byte(`"` + strconv.FormatInt(int64(t), 10) + `"`), nil
}

func (t *IDType) UnmarshalJSON(v []byte) error {
	vs := string(v)
	vs = strings.ReplaceAll(vs, `"`, ``)
	if vs == `` {
		*t = 0
		return nil
	}

	vi, err := strconv.ParseInt(vs, 10, 64)
	if err != nil {
		return err
	}

	*t = IDType(vi)
	return nil
}

type BasicInfo struct {
	CreatedAt *time.Time `json:"createdAt" pq_comment:"创建或注册时间"`
	UpdatedAt *time.Time `json:"updatedAt" pq_comment:"信息更新时间"`
	RemovedAt *time.Time `json:"removedAt" pq_comment:"移除时间"`
	Remove    bool       `json:"remove" pq_comment:"是否被移除"`
}
