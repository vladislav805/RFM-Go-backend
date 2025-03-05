#include <pthread.h>
#include <dlfcn.h>
#include <stdlib.h>
#include <string.h>

char *FM_LIBRARY_NAME = "fm_helium.so";
char *FM_LIBRARY_SYMBOL_NAME = "FM_HELIUM_LIB_INTERFACE";

enum helium_cmd_t {
    HCI_FM_HELIUM_SRCHMODE = 0x8000000 + 1,
    HCI_FM_HELIUM_SCANDWELL,
    HCI_FM_HELIUM_SRCHON,
    HCI_FM_HELIUM_STATE,
    HCI_FM_HELIUM_TRANSMIT_MODE,
    HCI_FM_HELIUM_RDSGROUP_MASK,
    HCI_FM_HELIUM_REGION,
    HCI_FM_HELIUM_SIGNAL_TH,
    HCI_FM_HELIUM_SRCH_PTY,
    HCI_FM_HELIUM_SRCH_PI,
    HCI_FM_HELIUM_SRCH_CNT,
    HCI_FM_HELIUM_EMPHASIS,
    HCI_FM_HELIUM_RDS_STD,
    HCI_FM_HELIUM_SPACING,
    HCI_FM_HELIUM_RDSON,
    HCI_FM_HELIUM_RDSGROUP_PROC,
    HCI_FM_HELIUM_LP_MODE,
    HCI_FM_HELIUM_ANTENNA,
    HCI_FM_HELIUM_RDSD_BUF,
    HCI_FM_HELIUM_PSALL,

    /*v4l2 Tx controls*/
    HCI_FM_HELIUM_IOVERC = 0x8000000 + 24,
    HCI_FM_HELIUM_INTDET,
    HCI_FM_HELIUM_MPX_DCC,
    HCI_FM_HELIUM_AF_JUMP,
    HCI_FM_HELIUM_RSSI_DELTA,
    HCI_FM_HELIUM_HLSI,

    /*Diagnostic commands*/
    HCI_FM_HELIUM_SOFT_MUTE,
    HCI_FM_HELIUM_RIVA_ACCS_ADDR,
    HCI_FM_HELIUM_RIVA_ACCS_LEN,
    HCI_FM_HELIUM_RIVA_PEEK,
    HCI_FM_HELIUM_RIVA_POKE,
    HCI_FM_HELIUM_SSBI_ACCS_ADDR,
    HCI_FM_HELIUM_SSBI_PEEK,
    HCI_FM_HELIUM_SSBI_POKE,
    HCI_FM_HELIUM_TX_TONE,
    HCI_FM_HELIUM_RDS_GRP_COUNTERS,
    HCI_FM_HELIUM_SET_NOTCH_FILTER, /* 0x8000028 */
    HCI_FM_HELIUM_SET_AUDIO_PATH,
    HCI_FM_HELIUM_DO_CALIBRATION,
    HCI_FM_HELIUM_SRCH_ALGORITHM,
    HCI_FM_HELIUM_GET_SINR,
    HCI_FM_HELIUM_INTF_LOW_THRESHOLD,
    HCI_FM_HELIUM_INTF_HIGH_THRESHOLD,
    HCI_FM_HELIUM_SINR_THRESHOLD,
    HCI_FM_HELIUM_SINR_SAMPLES,
    HCI_FM_HELIUM_SPUR_FREQ,
    HCI_FM_HELIUM_SPUR_FREQ_RMSSI,
    HCI_FM_HELIUM_SPUR_SELECTION,
    HCI_FM_HELIUM_UPDATE_SPUR_TABLE,
    HCI_FM_HELIUM_VALID_CHANNEL,
    HCI_FM_HELIUM_AF_RMSSI_TH,
    HCI_FM_HELIUM_AF_RMSSI_SAMPLES,
    HCI_FM_HELIUM_GOOD_CH_RMSSI_TH,
    HCI_FM_HELIUM_SRCHALGOTYPE,
    HCI_FM_HELIUM_CF0TH12,
    HCI_FM_HELIUM_SINRFIRSTSTAGE,
    HCI_FM_HELIUM_RMSSIFIRSTSTAGE,
    HCI_FM_HELIUM_RXREPEATCOUNT,
    HCI_FM_HELIUM_RSSI_TH,
    HCI_FM_HELIUM_AF_JUMP_RSSI_TH,
    HCI_FM_HELIUM_BLEND_SINRHI,
    HCI_FM_HELIUM_BLEND_RMSSIHI,
    HCI_FM_HELIUM_RDS_GRP_COUNTERS_EXT,
    HCI_FM_HELIUM_AGC_UCCTRL = 0x8000043, /* 0x8000043 */
    HCI_FM_HELIUM_AGC_GAIN_STATE,
    HCI_FM_HELIUM_ENABLE_LPF,

    /*using private CIDs under userclass*/
    HCI_FM_HELIUM_AUDIO_MUTE = 0x980909,
    HCI_FM_HELIUM_READ_DEFAULT = 0x00980928,
    HCI_FM_HELIUM_WRITE_DEFAULT,
    HCI_FM_HELIUM_SET_CALIBRATION,
    HCI_FM_HELIUM_SET_SPURTABLE = 0x0098092D,
    HCI_FM_HELIUM_GET_SPUR_TBL  = 0x0098092E,
    HCI_FM_HELIUM_FREQ,
    HCI_FM_HELIUM_SEEK,
    HCI_FM_HELIUM_UPPER_BAND,
    HCI_FM_HELIUM_LOWER_BAND,
    HCI_FM_HELIUM_AUDIO_MODE,
    HCI_FM_HELIUM_RMSSI,
    HCI_FM_HELIUM_AF_ALGO,
    HCI_FM_HELIUM_AF_SINR_GD_CH_TH,
    HCI_FM_HELIUM_AF_SINR_TH,

    /*FM VSC command to enable/disable slimbus data port*/
    HCI_FM_HELIUM_AUDIO = 0x00980940,
};

typedef unsigned char uint8;
typedef unsigned int  uint32;
typedef unsigned char boolean;

typedef struct band_limit_freq
{
  uint32 lower_limit;
  uint32 upper_limit;
}band_limit_freq;

typedef struct fm_config_data
{
  uint8 band;
  uint8 emphasis;
  uint8 spacing;
  uint8 rds_system;
  band_limit_freq bandlimits;
  uint8 is_fm_tx_on;
}fm_config_data;

fm_config_data * fmconfig_ptr;
typedef void (*enb_result_cb)();
typedef void (*tune_rsp_cb)(int Freq);
typedef void (*seek_rsp_cb)(int Freq);
typedef void (*scan_rsp_cb)();
typedef void (*srch_list_rsp_cb)(uint16_t *scan_tbl);
typedef void (*stereo_mode_cb)(boolean status);
typedef void (*rds_avl_sts_cb)(boolean status);
typedef void (*af_list_cb)(uint16_t *af_list);
typedef void (*rt_cb)(char *rt);
typedef void (*ps_cb)(char *ps);
typedef void (*oda_cb)();
typedef void (*rt_plus_cb)(char *rt_plus);
typedef void (*ert_cb)(char *ert);
typedef void (*disable_cb)();
typedef void (*callback_thread_event)(unsigned int evt);
typedef void (*rds_grp_cntrs_cb)(char *rds_params);
typedef void (*rds_grp_cntrs_ext_cb)(char *rds_params);
typedef void (*fm_peek_cb)(char *peek_rsp);
typedef void (*fm_ssbi_peek_cb)(char *ssbi_peek_rsp);
typedef void (*fm_agc_gain_cb)(char *agc_gain_rsp);
typedef void (*fm_ch_det_th_cb)(char *ch_det_rsp);
typedef void (*fm_sig_thr_cb) (int val, int status);
typedef void (*fm_get_ch_det_thrs_cb) (int val, int status);
typedef void (*fm_def_data_rd_cb) (int val, int status);
typedef void (*fm_get_blnd_cb) (int val, int status);
typedef void (*fm_set_ch_det_thrs_cb) (int status);
typedef void (*fm_def_data_wrt_cb) (int status);
typedef void (*fm_set_blnd_cb) (int status);
typedef void (*fm_get_stn_prm_cb) (int val, int status);
typedef void (*fm_get_stn_dbg_prm_cb) (int val, int status);
typedef void (*fm_ecc_evt_cb)(char *ecc_rsp);
typedef void (*fm_enable_sb_cb) (int status);

typedef struct {
    size_t  size;
    enb_result_cb  enabled_cb;
    tune_rsp_cb tune_cb;
    seek_rsp_cb  seek_cmpl_cb;
    scan_rsp_cb  scan_next_cb;
    srch_list_rsp_cb  srch_list_cb;
    stereo_mode_cb  stereo_status_cb;
    rds_avl_sts_cb  rds_avail_status_cb;
    af_list_cb  af_list_update_cb;
    rt_cb  rt_update_cb;
    ps_cb  ps_update_cb;
    oda_cb  oda_update_cb;
    rt_plus_cb  rt_plus_update_cb;
    ert_cb  ert_update_cb;
    disable_cb  disabled_cb;
    rds_grp_cntrs_cb rds_grp_cntrs_rsp_cb;
    rds_grp_cntrs_ext_cb rds_grp_cntrs_ext_rsp_cb;
    fm_peek_cb fm_peek_rsp_cb;
    fm_ssbi_peek_cb fm_ssbi_peek_rsp_cb;
    fm_agc_gain_cb fm_agc_gain_rsp_cb;
    fm_ch_det_th_cb fm_ch_det_th_rsp_cb;
    fm_ecc_evt_cb fm_ext_country_code_cb;
    callback_thread_event thread_evt_cb;
    fm_sig_thr_cb fm_get_sig_thres_cb;
    fm_get_ch_det_thrs_cb fm_get_ch_det_thr_cb;
    fm_def_data_rd_cb fm_def_data_read_cb;
    fm_get_blnd_cb fm_get_blend_cb;
    fm_set_ch_det_thrs_cb fm_set_ch_det_thr_cb;
    fm_def_data_wrt_cb fm_def_data_write_cb;
    fm_set_blnd_cb fm_set_blend_cb;
    fm_get_stn_prm_cb fm_get_station_param_cb;
    fm_get_stn_dbg_prm_cb fm_get_station_debug_param_cb;
    fm_enable_sb_cb fm_enable_slimbus_cb;
} fm_vendor_callbacks_t;

typedef struct {
    int (*hal_init)(fm_vendor_callbacks_t *p_cb);
    int (*set_fm_ctrl)(int ioctl, int val);
    int (*get_fm_ctrl) (int ioctl, int* val);
} fm_interface_t;



void fm_enabled_cb() {
    printf("fm_enabled_cb\n");
}
void fm_tune_cb(int Freq) {
    printf("fm_tune_cb %d\n", Freq);
}
void fm_seek_cmpl_cb(int Freq) {}
void fm_scan_next_cb() {}
void fm_srch_list_cb(uint16_t *scan_tbl) {}
void fm_stereo_status_cb(boolean stereo) {}
void fm_rds_avail_status_cb(boolean rds_avl) {}
void fm_rt_update_cb(char *rt) {
    printf("fm_rt_update_cb %s\n", rt);
}
void fm_ps_update_cb(char *buf) {}
void fm_oda_update_cb() {}
void fm_af_list_update_cb(uint16_t *af_list){}
void fm_rt_plus_update_cb(char *rt_plus) {}
void fm_ert_update_cb(char *ert) {}
void fm_disabled_cb() {}
void rds_grp_cntrs_rsp_cb(char *rds_grp_cntr_buff){}
void rds_grp_cntrs_ext_rsp_cb(char * rds_grp_cntr_buff){}
void fm_peek_rsp_cb(char *peek_rsp) {}
void fm_ssbi_peek_rsp_cb(char *ssbi_peek_rsp){}
void fm_agc_gain_rsp_cb(char *agc_gain_Resp){}
void fm_ch_det_th_rsp_cb(char *ch_det_rsp){}
static void fm_ext_country_code_cb(char *ecc_rsp){}
void fm_thread_evt_cb(unsigned int event){}
static void fm_get_sig_thres_cb(int val, int status){}
static void fm_get_ch_det_thr_cb(int val, int status){}
static void fm_def_data_read_cb(int val, int status){}
static void fm_get_blend_cb(int val, int status){}
static void fm_set_ch_det_thr_cb(int status){}
static void fm_def_data_write_cb(int status){}
static void fm_set_blend_cb(int status){}
static void fm_get_station_param_cb(int val, int status){}
static void fm_get_station_debug_param_cb(int val, int status){}
static void fm_enable_slimbus_cb(int status){}

void *lib_handle;
fm_interface_t *vendor_interface;
static   fm_vendor_callbacks_t fm_callbacks = {
    sizeof(fm_callbacks),
    fm_enabled_cb,
    fm_tune_cb,
    fm_seek_cmpl_cb,
    fm_scan_next_cb,
    fm_srch_list_cb,
    fm_stereo_status_cb,
    fm_rds_avail_status_cb,
    fm_af_list_update_cb,
    fm_rt_update_cb,
    fm_ps_update_cb,
    fm_oda_update_cb,
    fm_rt_plus_update_cb,
    fm_ert_update_cb,
    fm_disabled_cb,
    rds_grp_cntrs_rsp_cb,
    rds_grp_cntrs_ext_rsp_cb,
    fm_peek_rsp_cb,
    fm_ssbi_peek_rsp_cb,
    fm_agc_gain_rsp_cb,
    fm_ch_det_th_rsp_cb,
    fm_ext_country_code_cb,
    fm_thread_evt_cb,
    fm_get_sig_thres_cb,
    fm_get_ch_det_thr_cb,
    fm_def_data_read_cb,
    fm_get_blend_cb,
    fm_set_ch_det_thr_cb,
    fm_def_data_write_cb,
    fm_set_blend_cb,
    fm_get_station_param_cb,
    fm_get_station_debug_param_cb,
    fm_enable_slimbus_cb
};

void* bg(void * _) {
    while(1);
}

int main() {
    printf("1\n");

    lib_handle = dlopen(FM_LIBRARY_NAME, RTLD_NOW);

    if (!lib_handle) {
        printf("!lib_handle");
        lib_handle = NULL;
        return 1;
    }

    printf("2\n");

    vendor_interface = (fm_interface_t *)dlsym(lib_handle, FM_LIBRARY_SYMBOL_NAME);

    if (!vendor_interface) {
        printf("!vendor_interface");
        vendor_interface = NULL;
        return 1;
    }

    printf("3\n");

    int status = vendor_interface->hal_init(&fm_callbacks);

    printf("status = %d\n", status);

    int state = vendor_interface->set_fm_ctrl(HCI_FM_HELIUM_STATE, 1);
    int mute = vendor_interface->set_fm_ctrl(HCI_FM_HELIUM_AUDIO_MUTE, 0);
    int freq = vendor_interface->set_fm_ctrl(HCI_FM_HELIUM_FREQ, 87500);

    printf("state = %d, mute = %d, freq = %d\n", state, mute, freq);


    pthread_t bg_thread;
    int s = pthread_create(&bg_thread, NULL, &bg, NULL);


    return 0;
}
