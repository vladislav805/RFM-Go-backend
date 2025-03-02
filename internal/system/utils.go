package system

func IsSMDTransportLayer() bool {
	res, err := GetProp("ro.qualcomm.bt.hci_transport")

	if err != nil {
		return false
	}

	return res == "smd"
}
