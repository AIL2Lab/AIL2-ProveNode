package types

import "errors"

type StakingType uint8

const (
	ShortTerm StakingType = iota
	LongTerm
	Free
)

type NotifyType uint8

const (
	ContractRegister NotifyType = iota
	MachineRegister
	MachineUnregister
	MachineOnline
	MachineOffline
)

type ContractReportInfo struct {
	ProjectName string      `json:"project_name"`
	StakingType StakingType `json:"staking_type"`
	MachineId   string      `json:"machine_id"`
	ContainerId string      `json:"container_id" bson:"container_id"`
}

func (mr *ContractReportInfo) Validate() error {
	if mr.ProjectName == "" {
		return errors.New("empty project name")
	}
	if mr.MachineId == "" {
		return errors.New("empty machine id")
	}
	return nil
}
