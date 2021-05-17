package service

import (
	"gin-vue-admin/chain/core"
	"gin-vue-admin/chain/pkg/log"
)

type ValidatorStaking struct {
	Validator      string
	StakingAccount string
	StakingBalance string
}

func GetAllValidatorStaking() []*ValidatorStaking {
	validators := core.Client.GetEffectiveValidators("latest")
	validatorStakingList := make([]*ValidatorStaking, 0)
	for i, validator := range validators {
		log.Info("this is {} validator,address is {}", i, validator.String())
		stakingAccount, _ := core.Client.GetStakeAccount(validator, "latest")
		stakingBalance, _ := core.Client.BalanceOf(stakingAccount, "latest")
		validatorStaking := &ValidatorStaking{
			Validator:      validator.String(),
			StakingAccount: stakingAccount.String(),
			StakingBalance: stakingBalance.String(),
		}
		validatorStakingList = append(validatorStakingList, validatorStaking)
	}
	return validatorStakingList
}
