package app

import (
	"rbac-demo/pkg/db"
)

var Global db.ShareDaoFactory

func NewGlobal(factory db.ShareDaoFactory) {
	Global = factory
}
