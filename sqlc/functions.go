package sqlc

import (
	"github.com/relops/sqlc/meta"
)

func Count() IntField {
	return &intField{name: "*", fun: meta.Count}
}
