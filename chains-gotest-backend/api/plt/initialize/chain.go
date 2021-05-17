package initialize

import (
	"gin-vue-admin/chain/config"
	"gin-vue-admin/chain/core"
	chainData "gin-vue-admin/chain/model"
	"gin-vue-admin/chain/pkg/log"
	"gin-vue-admin/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"time"
)


func SetupTestCases(path string) {
	rand.Seed(time.Now().UnixNano())
	defer time.Sleep(time.Second)
	log.InitLog(1, log.Stdout)
	config.Init(path) //TODO: reading from setting
	core.Initialize()

}

func ChainDataBase(db *gorm.DB) {
	SetupTestCases(global.GVA_CONFIG.Local.TestConfig)
	err := db.AutoMigrate(
		chainData.GlobalParameter{},
		chainData.PaletteNode{},
		//EVENT Tables
		chainData.RewardEvent{},
		chainData.NFTEvent{},
		chainData.PLTTransferEvent{},
		chainData.PLTApprovalEvent{},
		chainData.TestAccount{},
		chainData.TestAccountBalance{},
		chainData.PMetrics{},
		chainData.StakingEvent{},
		chainData.CrossChainTx{},
	)
	if err != nil {
		global.GVA_LOG.Error("register blockchain table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register blockchain tables success")
}
