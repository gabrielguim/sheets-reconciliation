package repository

import (
	"sheets-reconciliation/commons"
)

type DataRepository interface {
	Read(areaName string) commons.Area
}
