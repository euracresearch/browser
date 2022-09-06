// Copyright 2021 Eurac Research. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package influx

// publicAllowed is the list of allowed measurements for the public role.
var publicAllowed = []string{
	"Air_RH_Avg",
	"Air_T_Avg",
	"Air_P_Avg",
	"Air_P_Std",
	"Wind_Speed_Avg",
	"Wind_Speed_Max",
	"Wind_Gust_Max",
	"Wind_Gust_Dir",
	"Wind_Dir",
	"Wind_Dir_Avg",
	"Wind_Speed",
	"NR_Up_SW_Avg",
	"SR_Avg",
	"Precip_Tot",
	"Precip_CS_Tot",
	"Precip_A_Tot",
	"Precip_C_Tot",
	"Precip_RT_NRT_Tot",
	"Snow_Height",
	"Snow_Height_Avg",
}

// maintenance is a list of measurement names only intressting for technicians.
var maintenance = []string{
	"RECORD",
	"Batt_V_Avg",
	"Batt_V_Std",
	"Log_T_Avg",
	"Log_T_Std",
	"Log_Samples_Tot",
	"Batt_V_Min",
	"Box_T_Ref_Avg",
	"Box_T_Ref_Std",
	"Box_T_Ref_bk_avg",
	"Box_T_Ref_bk_Std",
	"Air_T_old_Avg",
	"Air_T_old_Std",
	"Air_P0_Avg",
	"Bar_T_Avg",
	"Bar_T_Std",
	"Bar_Q_Avg",
	"Bar_Q_Std",
	"Wind_Fail_AxisX_Tot",
	"Wind_Fail_AxisXY_Tot",
	"Wind_Fail_AxisY_Tot",
	"Wind_Fail_MaxGain_Tot",
	"Wind_Fail_NoNewData_Tot",
	"Wind_Fail_NVM_Tot",
	"Wind_Fail_ROM_Tot",
	"Wind_Samples_Tot",
	"NR_Dn_Body_T_Avg",
	"NR_Dn_Body_T_Std",
	"NR_Up_Body_T_Avg",
	"NR_Up_Body_T_Std",
	"NR_Dn_LW0_Avg",
	"NR_Dn_LW0_Std",
	"NR_Body_T_Avg",
	"NR_Body_T_Std",
	"NR_PT100_Rs_Avg",
	"NR_PT100_Rs_Std",
	"NR_Up_LW0_Avg",
	"NR_Up_LW0_Std",
	"SR_Old_Avg",
	"SR_Old_Std",
	"NDVI_Dn_Tilt_Avg",
	"NDVI_Up_Tilt_Avg",
	"NDVI_Dn_Tilt_1_Avg",
	"NDVI_Up_Tilt_1_Avg",
	"PRI_Dn_Tilt_Avg",
	"PRI_Up_Tilt_Avg",
	"PRI_Dn_Tilt_1_Avg",
	"PRI_Up_Tilt_1_Avg",
	"Precip_Bucket_Level_NRT_Lt",
	"Precip_Bucket_Level_NRT_Perc",
	"Precip_Bucket_Level_RT_Lt",
	"Precip_ElectronicUnit_T_Avg",
	"Precip_ElectronicUnit_T_Std",
	"Precip_Heater_Status",
	"Precip_LoadCell_T_Avg",
	"Precip_LoadCell_T_Std",
	"Precip_NRT_Cum",
	"Precip_Pluvio_Status",
	"Precip_Pluvio_V_Avg",
	"Precip_Pluvio_V_Std",
	"Precip_Rim_T_Avg",
	"Precip_Rim_T_Std",
	"Snow_Dist",
	"Snow_Dist_Std",
	"Snow_Dist0",
	"Snow_Dist0_Std",
	"Snow_Quality",
	"Snow_Quality_Std",
	"Soil_Surf_T_mV_Avg",
	"Soil_Surf_T_mV_Std",
	"Soil_Surf_Body_T_Avg",
	"Soil_Surf_Body_T_Std",
	"Soil_Surf_T0_Avg",
	"Soil_Surf_T0_Std",
	"SWC_Wave_PA_02_Avg",
	"SWC_Wave_PA_02_Std",
	"SWC_Wave_PA_05_Avg",
	"SWC_Wave_PA_05_Std",
	"SWC_Wave_PA_20_Avg",
	"SWC_Wave_PA_20_Std",
	"SWC_Wave_PA_50_Avg",
	"SWC_Wave_PA_50_Std",
	"SWC_Wave_PA_05_1_Avg",
	"SWC_Wave_PA_05_1_Std",
	"SWC_Wave_PA_40_Avg",
	"SWC_Wave_PA_40_Std",
	"SWC_Wave_VR_02_Avg",
	"SWC_Wave_VR_02_Std",
	"SWC_Wave_VR_05_Avg",
	"SWC_Wave_VR_05_Std",
	"SWC_Wave_VR_20_Avg",
	"SWC_Wave_VR_20_Std",
	"SWC_Wave_VR_50_Avg",
	"SWC_Wave_VR_50_Std",
	"SWC_Wave_VR_05_1_Avg",
	"SWC_Wave_VR_05_1_Std",
	"SWC_Wave_VR_40_Avg",
	"SWC_Wave_VR_40_Std",
	"SWC_Wave_PA_A_02_Avg",
	"SWC_Wave_PA_A_02_Std",
	"SWC_Wave_PA_A_05_Avg",
	"SWC_Wave_PA_A_05_Std",
	"SWC_Wave_PA_A_20_Avg",
	"SWC_Wave_PA_A_20_Std",
	"SWC_Wave_PA_A_50_Avg",
	"SWC_Wave_PA_A_50_Std",
	"SWC_Wave_PA_B_02_Avg",
	"SWC_Wave_PA_B_02_Std",
	"SWC_Wave_PA_B_05_Avg",
	"SWC_Wave_PA_B_05_Std",
	"SWC_Wave_PA_B_20_Avg",
	"SWC_Wave_PA_B_20_Std",
	"SWC_Wave_PA_C_02_Avg",
	"SWC_Wave_PA_C_02_Std",
	"SWC_Wave_PA_C_05_Avg",
	"SWC_Wave_PA_C_05_Std",
	"SWC_Wave_PA_C_20_Avg",
	"SWC_Wave_PA_C_20_Std",
	"SWC_Wave_PA_C_50_Avg",
	"SWC_Wave_PA_C_50_Std",
	"SWC_Wave_VR_A_02_Avg",
	"SWC_Wave_VR_A_02_Std",
	"SWC_Wave_VR_A_05_Avg",
	"SWC_Wave_VR_A_05_Std",
	"SWC_Wave_VR_A_20_Avg",
	"SWC_Wave_VR_A_20_Std",
	"SWC_Wave_VR_A_50_Avg",
	"SWC_Wave_VR_A_50_Std",
	"SWC_Wave_VR_B_02_Avg",
	"SWC_Wave_VR_B_02_Std",
	"SWC_Wave_VR_B_05_Avg",
	"SWC_Wave_VR_B_05_Std",
	"SWC_Wave_VR_B_20_Avg",
	"SWC_Wave_VR_B_20_Std",
	"SWC_Wave_VR_C_02_Avg",
	"SWC_Wave_VR_C_02_Std",
	"SWC_Wave_VR_C_05_Avg",
	"SWC_Wave_VR_C_05_Std",
	"SWC_Wave_VR_C_20_Avg",
	"SWC_Wave_VR_C_20_Std",
	"SWC_Wave_VR_C_50_Avg",
	"SWC_Wave_VR_C_50_Std",
	"LWmV_Max",
	"LWmV_Min",
	"LWmV_Std",
}
