package handler

import (
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"go.uber.org/zap"
)

func (h *Handler) PowerSystemBoardComputingModuleСompatibilityСheck(boardComputingModule model.BoardComputingModule, powerSystem model.CubeSatPowerSystem) bool {
	zap.L().Info("Power system board computing Module compatibility check")

	if boardComputingModule.MinSupplyVoltage.Float64 > powerSystem.MinSystemBusVoltageOutputChannels.Float64 {
		zap.L().Info("voltage of power system is too low")
		return false
	}

	if boardComputingModule.MaxSupplyVoltage.Float64 < powerSystem.MinSystemBusVoltageOutputChannels.Float64 {
		zap.L().Info("voltage of power system is too high")
		return false
	}

	if boardComputingModule.PowerConsumption.Float64 > powerSystem.CurrentOutputChannelsMax.Float64*powerSystem.MinSystemBusVoltageOutputChannels.Float64 {
		zap.L().Info("power supply system capacity is too low")
		return false
	}

	if boardComputingModule.DataBus.String != powerSystem.DataInterface.String {
		zap.L().Info("the systems have different interfaces")
		return false
	}

	return true
}

func (h *Handler) VhfAntennaSystemBoardComputingSystemСompatibilityСheck(boardComputingModule model.BoardComputingModule, vhfAntennasystem model.VHFAntennaSystem) bool {
	zap.L().Info("Vhf antenna system board computing Module compatibility check")

	if vhfAntennasystem.Name.String == "SXC-AUH-03" && boardComputingModule.Name.String == "SXC-MB-04-ADC" {
		return true
	}
	return false
}

func (h *Handler) VhfTransceiverBoardComputingSystemСompatibilityСheck(boardComputingModule model.BoardComputingModule, vhfTransceiver model.VHFTransceiver) bool {
	zap.L().Info("Vhf transceiver board computing Module compatibility check")

	if boardComputingModule.DataBus.String != vhfTransceiver.Interface.String {
		zap.L().Info("the systems have different interfaces")
		return false
	}

	return true
}
