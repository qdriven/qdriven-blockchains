package model

import (
	"gin-vue-admin/global"
)

type PMetrics struct {
	global.GVA_MODEL
	TotalCount         uint64
	TotalExecutionTime int
	AvgTps             string

}
