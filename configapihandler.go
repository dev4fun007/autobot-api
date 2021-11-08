package api

import (
	"context"
	"github.com/dev4fun007/autobot-common"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type ConfigApiHandler struct {
	ctx              context.Context
	apiService       common.ApiService
	strategyTypeList []common.StrategyType
}

func NewConfigApiHandler(ctx context.Context, apiService common.ApiService, strategyTypeList []common.StrategyType) ConfigApiHandler {
	return ConfigApiHandler{
		ctx:              ctx,
		apiService:       apiService,
		strategyTypeList: strategyTypeList,
	}
}

func (h ConfigApiHandler) SaveConfig(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	strategyType, err := ValidateStrategyType(h.strategyTypeList, params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	// implementor should convert config to correct type
	err = h.apiService.CreateStrategy(h.ctx, bodyBytes, strategyType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (h ConfigApiHandler) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name, err := ValidateName(params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	strategyType, err := ValidateStrategyType(h.strategyTypeList, params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	// implementor should convert config to correct type
	err = h.apiService.UpdateStrategy(h.ctx, name, bodyBytes, strategyType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (h ConfigApiHandler) GetConfigByNameAndType(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name, err := ValidateName(params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	strategyType, err := ValidateStrategyType(h.strategyTypeList, params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	bytes, err := h.apiService.GetStrategyByNameAndType(h.ctx, name, strategyType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	} else {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bytes)
	}
}

func (h ConfigApiHandler) GetConfigByType(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	strategyType, err := ValidateStrategyType(h.strategyTypeList, params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	bytes, err := h.apiService.GetStrategiesByType(h.ctx, strategyType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	} else {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bytes)
	}
}

func (h ConfigApiHandler) DeleteConfigByNameAndType(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name, err := ValidateName(params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	strategyType, err := ValidateStrategyType(h.strategyTypeList, params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = h.apiService.DeleteStrategyByNameAndType(h.ctx, name, strategyType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
