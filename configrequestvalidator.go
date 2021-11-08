package api

import (
	"errors"
	"github.com/dev4fun007/autobot-common"
)

var (
	NameNotFound         = errors.New("name path parameter not found")
	StrategyTypeInvalid  = errors.New("strategy type not supported")
	StrategyTypeNotFound = errors.New("name path parameter not found")
)

func ValidateStrategyType(strategyTypeList []common.StrategyType, params map[string]string) (common.StrategyType, error) {
	strategyTypeString := params[StrategyTypePathParam]
	if strategyTypeString == "" {
		return common.InvalidStrategy, StrategyTypeNotFound
	}

	found := false
	var strategyType common.StrategyType
	for _, val := range strategyTypeList {
		if string(val) == strategyTypeString && val != common.InvalidStrategy {
			strategyType = val
			found = true
			break
		}
	}
	if !found {
		return common.InvalidStrategy, StrategyTypeInvalid
	}

	return strategyType, nil
}

func ValidateName(params map[string]string) (string, error) {
	name := params[ConfigNamePathParam]
	if name == "" {
		return "", NameNotFound
	}
	return name, nil
}
