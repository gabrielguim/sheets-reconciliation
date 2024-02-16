package consolidation

import (
	"sheets-reconciliation/internal/commons"
)

type AreaRepository interface {
	Read(areaName string) commons.Area
}
