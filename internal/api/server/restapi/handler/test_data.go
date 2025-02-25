package handler

import (
	"database/sql"
	"fmt"

	"github.com/Maksim646/space_vpx_satellite/internal/api/definition"
	"github.com/Maksim646/space_vpx_satellite/internal/api/server/restapi/api"
	"github.com/Maksim646/space_vpx_satellite/internal/model"
	"github.com/Maksim646/space_vpx_satellite/pkg/useful"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

func (h *Handler) LoadTestDataHandler(req api.LoadTestDataParams, principal *definition.Principal) middleware.Responder {
	zap.L().Info("test data request")
	ctx := req.HTTPRequest.Context()

	cubeSatFrame := model.CubeSatFrame{
		Height:                  sql.NullFloat64{Float64: 113.5, Valid: true},
		Name:                    sql.NullString{String: "SXC-F1U-02", Valid: true},
		Length:                  sql.NullFloat64{Float64: 100, Valid: true},
		Width:                   sql.NullFloat64{Float64: 100, Valid: true},
		Weight:                  sql.NullFloat64{Float64: 132, Valid: true},
		Size:                    1,
		Interface:               sql.NullString{String: "PC/104", Valid: true},
		OperatingTemperatureMin: sql.NullInt64{Int64: -40, Valid: true},
		OperatingTemperatureMax: sql.NullInt64{Int64: 85, Valid: true},
		MechanicalVibration:     sql.NullInt64{Int64: 12, Valid: true},
		MechanicalShock:         sql.NullInt64{Int64: 50, Valid: true},
		Link:                    sql.NullString{String: "https://sputnix.ru/ru/priboryi/pribory-cubesat/korpus-1u", Valid: true},
	}

	_, err := h.cubeSatFrameUsecase.CreateCubeSatFrame(ctx, cubeSatFrame)
	if err != nil {
		fmt.Println("error create frame", err)
		return api.NewLoadTestDataInternalServerError()
	}

	boardComputeringModule := model.BoardComputingModule{
		Name:                    sql.NullString{String: "SXC-MB-04-ADC", Valid: true},
		Length:                  sql.NullFloat64{Float64: 86.2, Valid: true},
		Width:                   sql.NullFloat64{Float64: 93.6, Valid: true},
		Height:                  sql.NullFloat64{Float64: 14, Valid: true},
		Weight:                  sql.NullFloat64{Float64: 55, Valid: true},
		SupplyVoltage:           sql.NullFloat64{Float64: 14, Valid: true},
		PowerConsumption:        sql.NullFloat64{Float64: 0.9, Valid: true},
		MinOperatingTemperature: sql.NullFloat64{Float64: -40, Valid: true},
		MaxOperatingTemperature: sql.NullFloat64{Float64: 85, Valid: true},
		MechanicalVibration:     sql.NullInt64{Int64: 12, Valid: true},
		MechanicalShock:         sql.NullInt64{Int64: 50, Valid: true},
	}

	_, err = h.cubeSatBoardComputingModuleUsecase.CreateBoardComputingModule(ctx, boardComputeringModule)
	if err != nil {
		fmt.Println("error create board computing module", err)
		return api.NewLoadTestDataInternalServerError()
	}

	powerSystem := model.CubeSatPowerSystem{
		Name:                           sql.NullString{String: "SXC-PSS-03", Valid: true},
		Length:                         sql.NullFloat64{Float64: 96, Valid: true},
		Width:                          sql.NullFloat64{Float64: 89, Valid: true},
		Height:                         sql.NullFloat64{Float64: 14, Valid: true},
		Weight:                         sql.NullFloat64{Float64: 58, Valid: true},
		SolarPanelChannels:             sql.NullInt64{Int64: 3, Valid: true},
		SolarPanelsType:                sql.NullString{String: "GaAs, Si", Valid: true},
		SolarPanelVoltageMin:           sql.NullFloat64{Float64: 3, Valid: true},
		SolarPanelVoltageMax:           sql.NullFloat64{Float64: 6, Valid: true},
		SolarPanelCurrentPerChannelMax: sql.NullFloat64{Float64: 1500, Valid: true},
		TotalCurrentOfSolarPanelsMax:   sql.NullFloat64{Float64: 3000, Valid: true},
		OutputChannels:                 sql.NullInt64{Int64: 4, Valid: true},
		SystemBusVoltageSolarPanels:    sql.NullFloat64{Float64: 8, Valid: true},
		SystemBusVoltageOutputChannels: sql.NullFloat64{Float64: 8, Valid: true},
		CurrentOutputChannelsMax:       sql.NullFloat64{Float64: 1500, Valid: true},
		TotalOutputCurrent:             sql.NullFloat64{Float64: 5000, Valid: true},
		DataInterface:                  sql.NullString{String: "CAN2.0B, Battery UART.", Valid: true},
		MaxOperatingTemperature:        sql.NullFloat64{Float64: 85, Valid: true},
		MinOperatingTemperature:        sql.NullFloat64{Float64: -40, Valid: true},
		MechanicalVibration:            sql.NullInt64{Int64: 12, Valid: true},
		MechanicalShock:                sql.NullInt64{Int64: 50, Valid: true},
	}

	_, err = h.cubeSatPowerSystemUsecase.CreatePowerSystem(ctx, powerSystem)
	if err != nil {
		fmt.Println("error create power system", err)
		return api.NewLoadTestDataInternalServerError()
	}

	solarPanel := model.CubeSatSolarPanelSide{

		Name:                    sql.NullString{String: "SXC-SGE-03", Valid: true},
		Length:                  sql.NullFloat64{Float64: 98, Valid: true},
		Width:                   sql.NullFloat64{Float64: 82.6, Valid: true},
		Weight:                  sql.NullFloat64{Float64: 32, Valid: true},
		MinOperatingTemperature: sql.NullFloat64{Float64: -40, Valid: true},
		MaxOperatingTemperature: sql.NullFloat64{Float64: 85, Valid: true},
		MechanicalVibration:     sql.NullInt64{Int64: 12, Valid: true},
		MechanicalShock:         sql.NullInt64{Int64: 50, Valid: true},
		CoilArea:                sql.NullFloat64{Float64: 1.9, Valid: true},
		CoilResistance:          sql.NullInt64{Int64: 200, Valid: true},
		Efficiency:              sql.NullInt64{Int64: 28, Valid: true},
		Imp:                     sql.NullFloat64{Float64: 480, Valid: true},
		Interface:               sql.NullString{String: "Hirose GT8E", Valid: true},
		Isc:                     sql.NullFloat64{Float64: 500, Valid: true},
		Vmp:                     sql.NullFloat64{Float64: 4.7, Valid: true},
		Voc:                     sql.NullFloat64{Float64: 5.3, Valid: true},
		Height:                  sql.NullFloat64{Float64: 8.6, Valid: true},
	}

	_, err = h.cubeSatSolarPanelSideUsecase.CreateSolarPanelSide(ctx, solarPanel)
	if err != nil {
		fmt.Println("error create solar panel side", err)
		return api.NewLoadTestDataInternalServerError()
	}

	solarPanelTop := model.SolarPanelTop{
		Name:                    sql.NullString{String: "SXC-SGE-03", Valid: true},
		Length:                  sql.NullFloat64{Float64: 98, Valid: true},
		Width:                   sql.NullFloat64{Float64: 82.6, Valid: true},
		Height:                  sql.NullFloat64{Float64: 8.6, Valid: true},
		Weight:                  sql.NullFloat64{Float64: 32, Valid: true},
		MinOperatingTemperature: sql.NullFloat64{Float64: -40, Valid: true},
		MaxOperatingTemperature: sql.NullFloat64{Float64: 85, Valid: true},
		MechanicalVibration:     sql.NullFloat64{Float64: 12, Valid: true},
		MechanicalShock:         sql.NullFloat64{Float64: 50, Valid: true},
		CoilArea:                sql.NullFloat64{Float64: 1.9, Valid: true},
		CoilResistance:          sql.NullFloat64{Float64: 200, Valid: true},
		Efficiency:              sql.NullFloat64{Float64: 28, Valid: true},
		Imp:                     sql.NullFloat64{Float64: 480, Valid: true},
		Interface:               sql.NullString{String: "Hirose GT8E", Valid: true},
		Isc:                     sql.NullFloat64{Float64: 500, Valid: true},
		Vmp:                     sql.NullFloat64{Float64: 4.7, Valid: true},
		Voc:                     sql.NullFloat64{Float64: 5.3, Valid: true},
	}

	_, err = h.cubeSatSolarPanelTopUsecase.CreateSolarPanelTop(ctx, solarPanelTop)
	if err != nil {
		fmt.Println("error create solar panel top", err)
		return api.NewLoadTestDataInternalServerError()
	}

	vhfAntennaSystem := model.VHFAntennaSystem{
		Name:                    sql.NullString{String: "SXC-AUH-03", Valid: true},
		Length:                  sql.NullFloat64{Float64: 107, Valid: true},
		Width:                   sql.NullFloat64{Float64: 107, Valid: true},
		Height:                  sql.NullFloat64{Float64: 95, Valid: true},
		Weight:                  sql.NullFloat64{Float64: 60, Valid: true},
		Interface:               sql.NullString{String: " Hirose GT8E", Valid: true},
		Frequency:               sql.NullFloat64{Float64: 435, Valid: true},
		MaxOperatingTemperature: sql.NullFloat64{Float64: 105, Valid: true},
		MinOperatingTemperature: sql.NullFloat64{Float64: -40, Valid: true},
		MechanicalVibration:     sql.NullInt64{Int64: 12, Valid: true},
		MechanicalShock:         sql.NullInt64{Int64: 50, Valid: true},
	}

	_, err = h.cubeSatVhfAntennaSystemUsecase.CreateVHFAntennaSystem(ctx, vhfAntennaSystem)
	if err != nil {
		fmt.Println("error create vhf antenna system", err)
		return api.NewLoadTestDataInternalServerError()
	}

	cubeSatVHFTransceiver := model.VHFTransceiver{
		Name:                    sql.NullString{String: "SXC-UHF-02", Valid: true},
		Length:                  sql.NullFloat64{Float64: 87, Valid: true},
		Width:                   sql.NullFloat64{Float64: 93, Valid: true},
		Height:                  sql.NullFloat64{Float64: 13, Valid: true},
		Weight:                  sql.NullFloat64{Float64: 43, Valid: true},
		SupplyVoltage:           sql.NullFloat64{Float64: 14, Valid: true},
		PowerConsumption:        sql.NullFloat64{Float64: 3.5, Valid: true},
		Interface:               sql.NullString{String: "CAN 2.0B, UART", Valid: true},
		OperatingFrequency:      sql.NullFloat64{Float64: 438, Valid: true},
		OutputPower:             sql.NullFloat64{Float64: 30, Valid: true},
		SensitivityReceiver:     sql.NullFloat64{Float64: -119, Valid: true},
		MaxOperatingTemperature: sql.NullFloat64{Float64: 85, Valid: true},
		MinOperatingTemperature: sql.NullFloat64{Float64: -40, Valid: true},
		MechanicalVibration:     sql.NullInt64{Int64: 12, Valid: true},
		MechanicalShock:         sql.NullInt64{Int64: 50, Valid: true},
	}

	_, err = h.cubeSatVhfTransceiverUsecase.CreateVHFTransceiver(ctx, cubeSatVHFTransceiver)
	if err != nil {
		fmt.Println("error create vhf transceiver", err)
		return api.NewLoadTestDataInternalServerError()
	}

	return api.NewLoadTestDataOK().WithPayload(&definition.Error{
		Message: useful.StrPtr("test data was loaded"),
	})
}
