package helium

import "unsafe"

type BtVendorInterface struct {
	/** Set to sizeof(bt_vndor_interface_t) */
	Size uint32

	/**
	 * Caller will open the interface and pass in the callback routines
	 * to the implemenation of this interface.
	 */
	init func(p_cb interface{}, local_bdaddr uint32) uint32

	/**  Vendor specific operations */
	op func(opcode HeliumOpCode, param unsafe.Pointer) int

	/** Closes the interface */
	cleanup func()
}

type HeliumOpCode uint32

/** Vendor specific operations OPCODE */
const (
	/*  [operation]
	 *      Power on or off the BT Controller.
	 *  [input param]
	 *      A pointer to int type with content of bt_vendor_power_state_t.
	 *      Typecasting conversion: (int *) param.
	 *  [return]
	 *      0 - default, don't care.
	 *  [callback]
	 *      None.
	 */
	BT_VND_OP_POWER_CTRL HeliumOpCode = iota

	/*  [operation]
	 *      Perform any vendor specific initialization or configuration
	 *      on the BT Controller. This is called before stack initialization.
	 *  [input param]
	 *      None.
	 *  [return]
	 *      0 - default, don't care.
	 *  [callback]
	 *      Must call fwcfg_cb to notify the stack of the completion of vendor
	 *      specific initialization once it has been done.
	 */
	BT_VND_OP_FW_CFG

	/*  [operation]
	 *      Perform any vendor specific SCO/PCM configuration on the BT Controller.
	 *      This is called after stack initialization.
	 *  [input param]
	 *      None.
	 *  [return]
	 *      0 - default, don't care.
	 *  [callback]
	 *      Must call scocfg_cb to notify the stack of the completion of vendor
	 *      specific SCO configuration once it has been done.
	 */
	BT_VND_OP_SCO_CFG

	/*  [operation]
	 *      Open UART port on where the BT Controller is attached.
	 *      This is called before stack initialization.
	 *  [input param]
	 *      A pointer to int array type for open file descriptors.
	 *      The mapping of HCI channel to fd slot in the int array is given in
	 *      bt_vendor_hci_channels_t.
	 *      And, it requires the vendor lib to fill up the content before returning
	 *      the call.
	 *      Typecasting conversion: (int (*)[]) param.
	 *  [return]
	 *      Numbers of opened file descriptors.
	 *      Valid number:
	 *          1 - CMD/EVT/ACL-In/ACL-Out via the same fd (e.g. UART)
	 *          2 - CMD/EVT on one fd, and ACL-In/ACL-Out on the other fd
	 *          4 - CMD, EVT, ACL-In, ACL-Out are on their individual fd
	 *  [callback]
	 *      None.
	 */
	BT_VND_OP_USERIAL_OPEN

	/*  [operation]
	 *      Close the previously opened UART port.
	 *  [input param]
	 *      None.
	 *  [return]
	 *      0 - default, don't care.
	 *  [callback]
	 *      None.
	 */
	BT_VND_OP_USERIAL_CLOSE

	/*  [operation]
	 *      Get the LPM idle timeout in milliseconds.
	 *      The stack uses this information to launch a timer delay before it
	 *      attempts to de-assert LPM WAKE signal once downstream HCI packet
	 *      has been delivered.
	 *  [input param]
	 *      A pointer to uint32_t type which is passed in by the stack. And, it
	 *      requires the vendor lib to fill up the content before returning
	 *      the call.
	 *      Typecasting conversion: (uint32_t *) param.
	 *  [return]
	 *      0 - default, don't care.
	 *  [callback]
	 *      None.
	 */
	BT_VND_OP_GET_LPM_IDLE_TIMEOUT

	/*  [operation]
	 *      Enable or disable LPM mode on BT Controller.
	 *  [input param]
	 *      A pointer to uint8_t type with content of bt_vendor_lpm_mode_t.
	 *      Typecasting conversion: (uint8_t *) param.
	 *  [return]
	 *      0 - default, don't care.
	 *  [callback]
	 *      Must call lpm_cb to notify the stack of the completion of LPM
	 *      disable/enable process once it has been done.
	 */
	BT_VND_OP_LPM_SET_MODE

	/*  [operation]
	 *      Assert or Deassert LPM WAKE on BT Controller.
	 *  [input param]
	 *      A pointer to uint8_t type with content of bt_vendor_lpm_wake_state_t.
	 *      Typecasting conversion: (uint8_t *) param.
	 *  [return]
	 *      0 - default, don't care.
	 *  [callback]
	 *      None.
	 */
	BT_VND_OP_LPM_WAKE_SET_STATE

	/*  [operation]
	 *      Perform any vendor specific commands related to audio state changes.
	 *  [input param]
	 *      a pointer to bt_vendor_op_audio_state_t indicating what audio state is
	 *      set.
	 *  [return]
	 *      0 - default, don't care.
	 *  [callback]
	 *      None.
	 */
	BT_VND_OP_SET_AUDIO_STATE

	/*  [operation]
	 *      The epilog call to the vendor module so that it can perform any
	 *      vendor-specific processes (e.g. send a HCI_RESET to BT Controller)
	 *      before the caller calls for cleanup().
	 *  [input param]
	 *      None.
	 *  [return]
	 *      0 - default, don't care.
	 *  [callback]
	 *      Must call epilog_cb to notify the stack of the completion of vendor
	 *      specific epilog process once it has been done.
	 */
	BT_VND_OP_EPILOG
)
