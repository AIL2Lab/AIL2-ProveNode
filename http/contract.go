package http

import (
	"context"
	"errors"
	ohttp "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	"AIL2-ProveNode/db"
	"AIL2-ProveNode/ail2"
	"AIL2-ProveNode/log"
	"AIL2-ProveNode/types"
)

var ReportContractTimeout = 60 * time.Second

func RegisterMachine(ctx *gin.Context) {
	rsp := types.BaseHttpResponse{
		Code:    0,
		Message: "ok",
	}

	var req types.ContractReportInfo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Code = int(types.ErrCodeParse)
		rsp.Message = types.ErrCodeParse.String()
		ctx.JSON(ohttp.StatusBadRequest, rsp)
		return
	}
	if err := req.Validate(); err != nil {
		rsp.Code = int(types.ErrCodeParam)
		rsp.Message = err.Error()
		ctx.JSON(ohttp.StatusBadRequest, rsp)
		return
	}

	ctx1, cancel := context.WithTimeout(ctx.Request.Context(), ReportContractTimeout)
	defer cancel()

	hash, err := dbc.DbcChain.Report(ctx1, types.MachineRegister, req.StakingType, req.ProjectName, req.MachineId)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": req}).Errorf("machine register failed: %v with hash %v", err, hash)
		rsp.Code = int(types.ErrCodeDbcChain)
		rsp.Message = err.Error()
		ctx.JSON(ohttp.StatusInternalServerError, rsp)
		return
	}
	log.Log.WithFields(logrus.Fields{"machine": req}).Info("machine register success with hash ", hash)

	mi, err := db.MDB.GetMachineInfo(
		ctx1,
		types.MachineKey{
			MachineId:   req.MachineId,
			Project:     req.ProjectName,
			ContainerId: req.ContainerId,
		},
	)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			if err := db.MDB.RegisterMachine(
				ctx1,
				types.MachineKey{
					MachineId:   req.MachineId,
					Project:     req.ProjectName,
					ContainerId: req.ContainerId,
				},
				req.StakingType,
			); err != nil {
				log.Log.WithFields(logrus.Fields{"machine": req}).Errorf("machine register failed: %v when insert database", err)
				rsp.Code = int(types.ErrCodeDatabase)
				rsp.Message = err.Error()
				ctx.JSON(ohttp.StatusInternalServerError, rsp)
				return
			}
		} else {
			log.Log.WithFields(logrus.Fields{"machine": req}).Errorf("machine register failed: %v when query database", err)
			rsp.Code = int(types.ErrCodeDatabase)
			rsp.Message = err.Error()
			ctx.JSON(ohttp.StatusInternalServerError, rsp)
			return
		}
	} else if !mi.RegisterTime.IsZero() {
		rsp.Message = "already registed"
	}

	ctx.JSON(ohttp.StatusOK, rsp)
}

func UnregisterMachine(ctx *gin.Context, unregisterNotify func(machine types.MachineKey, stakingType types.StakingType)) {
	rsp := types.BaseHttpResponse{
		Code:    0,
		Message: "ok",
	}

	var req types.ContractReportInfo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Code = int(types.ErrCodeParse)
		rsp.Message = types.ErrCodeParse.String()
		ctx.JSON(ohttp.StatusBadRequest, rsp)
		return
	}
	if err := req.Validate(); err != nil {
		rsp.Code = int(types.ErrCodeParam)
		rsp.Message = err.Error()
		ctx.JSON(ohttp.StatusBadRequest, rsp)
		return
	}

	ctx1, cancel := context.WithTimeout(ctx.Request.Context(), ReportContractTimeout)
	defer cancel()

	hash, err := dbc.DbcChain.Report(ctx1, types.MachineUnregister, req.StakingType, req.ProjectName, req.MachineId)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": req}).Errorf("machine unregister failed: %v with hash %v", err, hash)
		rsp.Code = int(types.ErrCodeDbcChain)
		rsp.Message = err.Error()
		ctx.JSON(ohttp.StatusInternalServerError, rsp)
		return
	}
	log.Log.WithFields(logrus.Fields{"machine": req}).Info("machine unregister success with hash ", hash)

	unregisterNotify(
		types.MachineKey{
			MachineId:   req.MachineId,
			Project:     req.ProjectName,
			ContainerId: req.ContainerId,
		},
		req.StakingType,
	)

	if err := db.MDB.UnregisterMachine(ctx1, types.MachineKey{
		MachineId:   req.MachineId,
		Project:     req.ProjectName,
		ContainerId: req.ContainerId,
	}); err != nil {
		log.Log.WithFields(logrus.Fields{"machine": req}).Errorf("machine unregister failed: %v when delete database", err)
		rsp.Code = int(types.ErrCodeDatabase)
		rsp.Message = err.Error()
		ctx.JSON(ohttp.StatusInternalServerError, rsp)
		return
	}
	ctx.JSON(ohttp.StatusOK, rsp)
}

func OnlineMachine(ctx *gin.Context) {
	rsp := types.BaseHttpResponse{
		Code:    0,
		Message: "ok",
	}

	var req types.ContractReportInfo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Code = int(types.ErrCodeParse)
		rsp.Message = types.ErrCodeParse.String()
		ctx.JSON(ohttp.StatusBadRequest, rsp)
		return
	}
	if err := req.Validate(); err != nil {
		rsp.Code = int(types.ErrCodeParam)
		rsp.Message = err.Error()
		ctx.JSON(ohttp.StatusBadRequest, rsp)
		return
	}

	ctx1, cancel := context.WithTimeout(ctx.Request.Context(), ReportContractTimeout)
	defer cancel()

	hash, err := dbc.DbcChain.Report(ctx1, types.MachineOnline, req.StakingType, req.ProjectName, req.MachineId)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": req}).Errorf("machine online failed: %v with hash %v", err, hash)
		rsp.Code = int(types.ErrCodeDbcChain)
		rsp.Message = err.Error()
		ctx.JSON(ohttp.StatusInternalServerError, rsp)
		return
	}
	log.Log.WithFields(logrus.Fields{"machine": req}).Info("machine online success with hash ", hash)
	ctx.JSON(ohttp.StatusOK, rsp)
}

func OfflineMachine(ctx *gin.Context) {
	rsp := types.BaseHttpResponse{
		Code:    0,
		Message: "ok",
	}

	var req types.ContractReportInfo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		rsp.Code = int(types.ErrCodeParse)
		rsp.Message = types.ErrCodeParse.String()
		ctx.JSON(ohttp.StatusBadRequest, rsp)
		return
	}
	if err := req.Validate(); err != nil {
		rsp.Code = int(types.ErrCodeParam)
		rsp.Message = err.Error()
		ctx.JSON(ohttp.StatusBadRequest, rsp)
		return
	}

	ctx1, cancel := context.WithTimeout(ctx.Request.Context(), ReportContractTimeout)
	defer cancel()

	hash, err := dbc.DbcChain.Report(ctx1, types.MachineOffline, req.StakingType, req.ProjectName, req.MachineId)
	if err != nil {
		log.Log.WithFields(logrus.Fields{"machine": req}).Errorf("machine offline failed: %v with hash %v", err, hash)
		rsp.Code = int(types.ErrCodeDbcChain)
		rsp.Message = err.Error()
		ctx.JSON(ohttp.StatusInternalServerError, rsp)
		return
	}
	log.Log.WithFields(logrus.Fields{"machine": req}).Info("machine offline success with hash ", hash)
	ctx.JSON(ohttp.StatusOK, rsp)
}
