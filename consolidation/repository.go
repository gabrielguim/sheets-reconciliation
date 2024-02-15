package consolidation

import (
	"sheets-reconciliation/commons"
)

type AreaRepository interface {
	Read(areaName string) commons.Area
}
