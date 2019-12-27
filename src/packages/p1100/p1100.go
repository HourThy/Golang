package p1100

import (
	"config"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "gopkg.in/goracle.v2"
	"gopkg.in/mgo.v2"
	bbson "gopkg.in/mgo.v2/bson"
)

// ABOMDETAIL : ABOMDETAIL
type ABOMDETAIL struct {
	BOM_ID          string `bson:"BIN_ID, omitempty"`
	MTRL_PRODUCT_ID string `bson:"MTRL_PRODUCT_ID, omitempty"`
	PLAN_QTY        string `bson:"PLAN_QTY, omitempty"`
	MTRL_UNIT       string `bson:"MTRL_UNIT, omitempty"`
	PARENT_ID       string `bson:"PARENT_ID, omitempty"`
	PARTG_ID        string `bson:"PARTG_ID, omitempty"`
}

// ACALEND : ACALEND
type ACALEND struct {
	CAL_DATE       string `bson:"CAL_DATE, omitempty"`
	HOLI_CODE      string `bson:"HOLI_CODE, omitempty"`
	MFGSHOP        string `bson:"MFGSHOP, omitempty"`
	MFGMONTH       string `bson:"MFGMONTH, omitempty"`
	WEEKDAY        string `bson:"WEEKDAY, omitempty"`
	FGWEEK         string `bson:"FGWEEK, omitempty"`
	MFGQUTR        string `bson:"MFGQUTR, omitempty"`
	SHIFT_1ST_TIME string `bson:"SHIFT_1ST_TIME, omitempty"`
	SHIFT_2ND_TIME string `bson:"SHIFT_2ND_TIME, omitempty"`
	SHIFT_3RD_TIME string `bson:"SHIFT_3RD_TIME, omitempty"`
}

// ACARRHLD : ACARRHLD
type ACARRHLD struct {
	CRR_ID       string `bson:"CAL_DATE, omitempty"`
	HLD_SEQ_NO   string `bson:"HLD_SEQ_NO, omitempty"`
	HLD_USER     string `bson:"HLD_USER, omitempty"`
	HLD_CATE     string `bson:"HLD_CATE, omitempty"`
	HLD_CODE     string `bson:"HLD_CODE, omitempty"`
	HLD_DATE     string `bson:"HLD_DATE, omitempty"`
	HLD_TIME     string `bson:"HLD_TIME, omitempty"`
	HLD_MFDT     string `bson:"HLD_MFDT, omitempty"`
	HLD_MFWK     string `bson:"HLD_MFWK, omitempty"`
	HLD_MFMN     string `bson:"HLD_MFMN, omitempty"`
	HLD_MFSH     string `bson:"HLD_MFSH, omitempty"`
	PLN_REL_DATE string `bson:"PLN_REL_DATE, omitempty"`
	HLD_NOTE     string `bson:"HLD_NOTE, omitempty"`
}

// ACARRIER : ACARRIER
type ACARRIER struct {
	CRR_ID              string `bson:"CRR_ID, omitempty"`
	CRR_STAT            string `bson:"CRR_STAT, omitempty"`
	VALID_FLG           string `bson:"VALID_FLG, omitempty"`
	CRR_MNT_RST_DATE    string `bson:"CRR_MNT_RST_DATE, omitempty"`
	GROUP_ID            string `bson:"GROUP_ID, omitempty"`
	PPBOX_ID            string `bson:"PPBOX_ID, omitempty"`
	REAL_EMP            string `bson:"REAL_EMP, omitempty"`
	SGR_ID              string `bson:"SGR_ID, omitempty"`
	LOT_ID              string `bson:"LOT_ID, omitempty"`
	SPLT_ID             string `bson:"SPLT_ID, omitempty"`
	PRODUCT_CATE        string `bson:"PRODUCT_CATE, omitempty"`
	PRODUCT_ID          string `bson:"PRODUCT_ID, omitempty"`
	EC_CODE             string `bson:"EC_CODE, omitempty"`
	NX_ROUTE_ID         string `bson:"NX_ROUTE_ID, omitempty"`
	NX_ROUTE_VER        string `bson:"NX_ROUTE_VER, omitempty"`
	NX_OPE_NO           string `bson:"NX_OPE_NO, omitempty"`
	X_OPE_ID            string `bson:"X_OPE_ID, omitempty"`
	NX_OPE_VER          string `bson:"NX_OPE_VER, omitempty"`
	NX_PROC_ID          string `bson:"NX_PROC_ID, omitempty"`
	NX_EQPTG_ID         string `bson:"NX_EQPTG_ID, omitempty"`
	NX_EQPT_ID          string `bson:"NX_EQPT_ID, omitempty"`
	NX_PEP_LVL          string `bson:"NX_PEP_LVL, omitempty"`
	NX_DEPT_CODE        string `bson:"NX_DEPT_CODE, omitempty"`
	NX_RECIPE_ID        string `bson:"NX_RECIPE_ID, omitempty"`
	EQPT_ID             string `bson:"EQPT_ID, omitempty"`
	EQPT_PORT_ID        string `bson:"EQPT_PORT_ID, omitempty"`
	HLD_USER            string `bson:"HLD_USER, omitempty"`
	HLD_CODE            string `bson:"HLD_CODE, omitempty"`
	HLD_DATE            string `bson:"HLD_DATE, omitempty"`
	HLD_TIME            string `bson:"HLD_TIME, omitempty"`
	HLD_MFDT            string `bson:"HLD_MFDT, omitempty"`
	HLD_MFWK            string `bson:"HLD_MFWK, omitempty"`
	HLD_MFMN            string `bson:"HLD_MFMN, omitempty"`
	HLD_MFSH            string `bson:"HLD_MFSH, omitempty"`
	PRTY                string `bson:"PRTY, omitempty"`
	LOGOF_DATE          string `bson:"LOGOF_DATE, omitempty"`
	LOGOF_TIME          string `bson:"LOGOF_TIME, omitempty"` //
	MAX_QRS_AVL_FLG     string `bson:"MAX_QRS_AVL_FLG, omitempty"`
	MAX_QRS_ROUTE_ID    string `bson:"MAX_QRS_ROUTE_ID, omitempty"`
	MAX_QRS_ROUTE_VER   string `bson:"MAX_QRS_ROUTE_VER, omitempty"`
	MAX_QRS_OPE_ID      string `bson:"MAX_QRS_OPE_ID, omitempty"`
	MAX_QRS_DATE        string `bson:"MAX_QRS_DATE, omitempty"`
	MAX_QRS_TIME        string `bson:"MAX_QRS_TIME, omitempty"`
	MIN_QRS_AVL_FLG     string `bson:"MIN_QRS_AVL_FLG, omitempty"`
	MIN_QRS_ROUTE_ID    string `bson:"MIN_QRS_ROUTE_ID, omitempty"`
	MIN_QRS_ROUTE_VER   string `bson:"MIN_QRS_ROUTE_VER, omitempty"`
	MIN_QRS_OPE_ID      string `bson:"MIN_QRS_OPE_ID, omitempty"`
	MIN_QRS_DATE        string `bson:"MIN_QRS_DATE, omitempty"`
	MIN_QRS_TIME        string `bson:"MIN_QRS_TIME, omitempty"`
	SHT_CNT             string `bson:"SHT_CNT, omitempty"`
	PNL_CNT             string `bson:"PNL_CNT, omitempty"`
	PLN_REL_DATE        string `bson:"PLN_REL_DATE, omitempty"`
	ABND_FLG            string `bson:"ABND_FLG, omitempty"`
	CRR_CLN_ROUTE_ID    string `bson:"CRR_CLN_ROUTE_ID, omitempty"`
	CRR_CLN_ROUTE_VER   string `bson:"CRR_CLN_ROUTE_VER, omitempty"`
	INIT_CLUP_FLG       string `bson:"INIT_CLUP_FLG, omitempty"`
	CLN_CRR_FLG         string `bson:"CLN_CRR_FLG, omitempty"`
	APPLY_EEN_ID        string `bson:"APPLY_EEN_ID, omitempty"`
	PR_MIX_FLG          string `bson:"PR_MIX_FLG, omitempty"`
	LOGOF_EQPT_ID       string `bson:"LOGOF_EQPT_ID, omitempty"`
	LOGOF_PORT_ID       string `bson:"LOGOF_PORT_ID, omitempty"`
	LOGOF_RECIPE_ID     string `bson:"LOGOF_RECIPE_ID, omitempty"`
	PLN_CMP_DATE        string `bson:"PLN_CMP_DATE, omitempty"`
	PLN_CMP_TIME        string `bson:"PLN_CMP_TIME, omitempty"`
	STB_SHOP            string `bson:"STB_SHOP, omitempty"`
	DEST_SHOP           string `bson:"DEST_SHOP, omitempty"`
	SHT_JUDGE           string `bson:"SHT_JUDGE, omitempty"`
	XF_CMD_FLG          string `bson:"XF_CMD_FLG, omitempty"`
	PRIORITY            string `bson:"PRIORITY, omitempty"`
	XFER_ID             string `bson:"XFER_ID, omitempty"`
	XFER_TYPE           string `bson:"XFER_TYPE, omitempty"`
	LN_EQPT_ID          string `bson:"LN_EQPT_ID, omitempty"`
	PLN_SEQ_NO          string `bson:"PLN_SEQ_NO, omitempty"`
	REQ_RCV_DATE        string `bson:"REQ_RCV_DATE, omitempty"`
	REQ_RCV_TIME        string `bson:"REQ_RCV_TIME, omitempty"`
	RSV_FLG             string `bson:"RSV_FLG, omitempty"`
	M_RSVD              string `bson:"M_RSVD, omitempty"`
	RSV_USER_ID         string `bson:"RSV_USER_ID, omitempty"`
	RSV_DATE            string `bson:"RSV_DATE, omitempty"`
	RSV_TIME            string `bson:"RSV_TIME, omitempty"`
	RSV_NO              string `bson:"RSV_NO, omitempty"`
	RSV_EQPT_ID         string `bson:"RSV_EQPT_ID, omitempty"`
	RSV_PORT_ID         string `bson:"RSV_PORT_ID, omitempty"`
	RSV_PATI_ID         string `bson:"RSV_PATI_ID, omitempty"`
	RSV_ZONE_ID         string `bson:"RSV_ZONE_ID, omitempty"`
	CRR_USE_CNT         string `bson:"CRR_USE_CNT, omitempty"`
	MAX_CRR_USE_CNT     string `bson:"MAX_CRR_USE_CNT, omitempty"`
	MAX_USE_OVER_FLG    string `bson:"MAX_USE_OVER_FLG, omitempty"`
	XFER_CLM_DATE       string `bson:"XFER_CLM_DATE, omitempty"`
	XFER_CLM_TIME       string `bson:"XFER_CLM_TIME, omitempty"`
	NO_XFER_FLG         string `bson:"NO_XFER_FLG, omitempty"`
	NO_XFER_TYP         string `bson:"NO_XFER_TYP, omitempty"`
	CLUP_FLG            string `bson:"CLUP_FLG, omitempty"`
	CLUP_DATE           string `bson:"CLUP_DATE, omitempty"`
	CLUP_TIME           string `bson:"CLUP_TIME, omitempty"` //
	RSTRCT_CLUP_DAY     string `bson:"RSTRCT_CLUP_DAY, omitempty"`
	SIZE_CODE           string `bson:"SIZE_CODE, omitempty"`
	OWNER_ID            string `bson:"OWNER_ID, omitempty"`
	SALES_ORDER         string `bson:"SALES_ORDER, omitempty"`
	ABNORMAL_FLG_1      string `bson:"ABNORMAL_FLG_1, omitempty"`
	CRR_CATE            string `bson:"CRR_CATE, omitempty"`
	CRR_SIZE            string `bson:"CRR_SIZE, omitempty"`
	CRR_SET_CODE        string `bson:"CRR_SET_CODE, omitempty"`
	CRR_MKR_ID          string `bson:"CRR_MKR_ID, omitempty"`
	CROSS_FLG           string `bson:"CROSS_FLG, omitempty"`
	CRR_RGST_DATE       string `bson:"CRR_RGST_DATE, omitempty"`
	CRR_RGST_TIME       string `bson:"CRR_RGST_TIME, omitempty"`
	CRR_OWN             string `bson:"CRR_OWN, omitempty"`
	POSITION            string `bson:"POSITION, omitempty"`
	PORT_ID             string `bson:"PORT_ID, omitempty"`
	PATI_ID             string `bson:"PATI_ID, omitempty"`
	ZONE_ID             string `bson:"ZONE_ID, omitempty"`
	SHELF_ID            string `bson:"POS_CHG_DATE, omitempty"`
	POS_CHG_DATE        string `bson:"POS_CHG_DATE, omitempty"`
	POS_CHG_TIME        string `bson:"POS_CHG_TIME, omitempty"`
	XF_CMD_STAT         string `bson:"XF_CMD_STAT, omitempty"`
	XFER_STAT           string `bson:"XFER_STAT, omitempty"`
	FR_EQPT_ID          string `bson:"FR_EQPT_ID, omitempty"`
	FR_PORT_ID          string `bson:"FR_PORT_ID, omitempty"`
	FR_PATI_ID          string `bson:"FR_PATI_ID, omitempty"` //
	FR_ZONE_ID          string `bson:"FR_ZONE_ID, omitempty"`
	TO_EQPT_ID          string `bson:"TO_EQPT_ID, omitempty"`
	TO_PORT_ID          string `bson:"TO_PORT_ID, omitempty"`
	TO_PATI_ID          string `bson:"TO_PATI_ID, omitempty"`
	TO_ZONE_ID          string `bson:"TO_ZONE_ID, omitempty"`
	PUSH_STOCKER_ID     string `bson:"PUSH_STOCKER_ID, omitempty"`
	PUSH_PATI_ID        string `bson:"PUSH_PATI_ID, omitempty"`
	PUSH_ZONE_ID        string `bson:"PUSH_ZONE_ID, omitempty"`
	LST_STC_IN_DATE     string `bson:"LST_STC_IN_DATE, omitempty"`
	LST_STC_IN_TIME     string `bson:"LST_STC_IN_TIME, omitempty"`
	SWEEP_DATE          string `bson:"SWEEP_DATE, omitempty"`
	SWEEP_TIME          string `bson:"SWEEP_TIME, omitempty"`
	XFER_RTN_CODE       string `bson:"XFER_RTN_CODE, omitempty"`
	LAST_XFR_CMD_RC     string `bson:"LAST_XFR_CMD_RC, omitempty"`
	LAST_LOC_RPT_RC     string `bson:"LAST_LOC_RPT_RC, omitempty"`
	CRRG_ID             string `bson:"CRRG_ID, omitempty"`
	STOCK_STAT          string `bson:"STOCK_STAT, omitempty"`
	MAX_QRK_AVL_FLG     string `bson:"MAX_QRK_AVL_FLG, omitempty"`
	MAX_QRK_ROUTE_ID    string `bson:"MAX_QRK_ROUTE_ID, omitempty"`
	MAX_QRK_ROUTE_VER   string `bson:"MAX_QRK_ROUTE_VER, omitempty"`
	MAX_QRK_OPE_ID      string `bson:"MAX_QRK_OPE_ID, omitempty"`
	MAX_QRK_DATE        string `bson:"MAX_QRK_DATE, omitempty"`
	MAX_QRK_TIME        string `bson:"MAX_QRK_TIME, omitempty"`
	SETTING_NO          string `bson:"SETTING_NO, omitempty"`
	CRR_RGST_OWN        string `bson:"CRR_RGST_OWN, omitempty"`
	STB_SAMPLING_FLG    string `bson:"STB_SAMPLING_FLG, omitempty"`
	FINISHED_DATE       string `bson:"FINISHED_DATE, omitempty"`
	FINISHED_TIME       string `bson:"FINISHED_TIME, omitempty"`
	REVW_SHT_JUDGE      string `bson:"REVW_SHT_JUDGE, omitempty"`
	MTRL_GRADE          string `bson:"MTRL_GRADE, omitempty"`
	AREA_CODE           string `bson:"AREA_CODE, omitempty"`
	MTRL_PRODUCT_ID     string `bson:"MTRL_PRODUCT_ID, omitempty"`
	CARTON_ID           string `bson:"CARTON_ID, omitempty"`
	MIX_OWNER           string `bson:"MIX_OWNER, omitempty"`
	REJ_EQPT_ID         string `bson:"REJ_EQPT_ID, omitempty"`
	REJ_PORT_ID         string `bson:"REJ_PORT_ID, omitempty"`
	REJ_PROC_ID         string `bson:"REJ_PROC_ID, omitempty"`
	UNIT                string `bson:"UNIT, omitempty"`
	STAGE_ID            string `bson:"STAGE_ID, omitempty"`
	WORDER_ID           string `bson:"WORDER_ID, omitempty"`
	XFER_MODE           string `bson:"XFER_MODE, omitempty"`
	ORDER_TYP           string `bson:"ORDER_TYP, omitempty"`
	MAX_QRS_OPE_NO      string `bson:"MAX_QRS_OPE_NO, omitempty"`
	MIN_QRS_OPE_NO      string `bson:"MIN_QRS_OPE_NO, omitempty"`
	MAX_QRK_OPE_NO      string `bson:"MAX_QRK_OPE_NO, omitempty"`
	ORG_LR_PRODUCT_ID   string `bson:"ORG_LR_PRODUCT_ID, omitempty"`
	ORG_LR_MTRL_PRODUCT string `bson:"ORG_LR_MTRL_PRODUCT, omitempty"`
	ORG_LR_LOT_ID       string `bson:"ORG_LR_LOT_ID, omitempty"`
	ORG_LR_SPLT_ID      string `bson:"ORG_LR_SPLT_ID, omitempty"`
	ORG_LR_SHT_CNT      string `bson:"ORG_LR_SHT_CNT, omitempty"`
	IFTHEN_ACT_FLG      string `bson:"IFTHEN_ACT_FLG, omitempty"`
}

// ABOM : ABOM
type ABOM struct {
	BOM_ID      string `bson:"BIN_ID, omitempty"`
	BOM_DSC     string `bson:"BOM_DSC, omitempty"`
	ADDT_INFO_1 string `bson:"ADDT_INFO_1, omitempty"`
	ADDT_INFO_2 string `bson:"ADDT_INFO_2, omitempty"`
	ADDT_INFO_3 string `bson:"ADDT_INFO_3, omitempty"`
}

// ABINJDGRL : ABINJDGRL
type ABINJDGRL struct {
	BIN_ID        string `bson:"BIN_ID, omitempty"`
	JUDGE_RULE_ID string `bson:"JUDGE_RULE_ID, omitempty"`
}

// ABINGRADE : ABINGRADE
type ABINGRADE struct {
	BIN_ID    string `bson:"BIN_ID, omitempty"`
	BIN_NO    string `bson:"BIN_NO, omitempty"`
	SHT_GRADE string `bson:"SHT_GRADE, omitempty"`
	HOLD_FLG  string `bson:"HOLD_FLG, omitempty"`
}

// ABILBDLN : ABILBDLN
type ABILBDLN struct {
	CODE      string `bson:"CODE, omitempty"`
	BILBD_MSG string `bson:"BILBD_MSG, omitempty"`
	BOARD_MSG string `bson:"BOARD_MSG, omitempty"`
}

//ABAYAREA : ABAYAREA
type ABAYAREA struct {
	BAY_ID      string `bson:"BAY_ID, omitempty"`
	OPI_FLG     string `bson:"OPI_FLG, omitempty"`
	BAY_DSC     string `bson:"BAY_DSC, omitempty"`
	ADDT_INFO_1 string `bson:"ADDT_INFO_1, omitempty"`
	ADDT_INFO_2 string `bson:"ADDT_INFO_2, omitempty"`
	ADDT_INFO_3 string `bson:"ADDT_INFO_3, omitempty"`
}

// AAQLLEVEL : AAQLLEVEL
type AAQLLEVEL struct {
	AQL_LVL      string `bson:"AQL_LVL, omitempty"`
	SEQ_NO       string `bson:"SEQ_NO, omitempty"`
	AQL_FLG      string `bson:"AQL_FLG, omitempty"`
	PNL_QTY_FROM string `bson:"PNL_QTY_FROM, omitempty"`
	PNL_QTY_TO   string `bson:"PNL_QTY_TO, omitempty"`
	SMP_QTY      string `bson:"SMP_QTY, omitempty"`
	ACCEPT_QTY   string `bson:"ACCEPT_QTY, omitempty"`
	UPD_USER_ID  string `bson:"UPD_USER_ID, omitempty"`
	UPD_DATE     string `bson:"UPD_DATE, omitempty"`
	UPD_TIME     string `bson:"UPD_TIME, omitempty"`
}

// AALERT : AALERT
type AALERT struct {
	EQPT_ID          string `bson:"EQPT_ID, omitempty"`
	RPT_DATE         string `bson:"RPT_DATE, omitempty"`
	RPT_TIME         string `bson:"RPT_TIME, omitempty"`
	ALRT_ID          string `bson:"ALRT_ID, omitempty"`
	RPT_SOURCE       string `bson:"RPT_SOURCE, omitempty"`
	ALRT_CODE        string `bson:"ALRT_CODE, omitempty"`
	ALRT_LVL         string `bson:"ALRT_LVL, omitempty"`
	ALERT_ON_OFF_FLG string `bson:"ALERT_ON_OFF_FLG, omitempty"`
	ALRT_DSC         string `bson:"ALRT_DSC, omitempty"`
	ALRT_COMMENT     string `bson:"ALRT_COMMENT, omitempty"`
	CFM_USER_ID      string `bson:"CFM_USER_ID, omitempty"`
	CFM_DATE         string `bson:"CFM_DATE, omitempty"`
	CFM_TIME         string `bson:"CFM_TIME, omitempty"`
	CFM_COMMENT      string `bson:"CFM_COMMENT, omitempty"`
}

//ABIN : ABIN
type ABIN struct {
	BIN_ID      string `bson:"BIN_ID, omitempty"`
	BIN_DSC     string `bson:"BIN_DSC, omitempty"`
	ADDT_INFO_1 string `bson:"ADDT_INFO_1, omitempty"`
	ADDT_INFO_2 string `bson:"ADDT_INFO_2, omitempty"`
	ADDT_INFO_3 string `bson:"ADDT_INFO_3, omitempty"`
}

// AMTRLPRD : AMTRLPRD
type AMTRLPRD struct {
	MTRL_PRODUCT_ID   string `bson:"MTRL_PRODUCT_ID, omitempty"`
	MTRL_PRODUCT_DSC  string `bson:"MTRL_PRODUCT_DSC, omitempty"`
	MTRL_CATE         string `bson:"MTRL_CATE, omitempty"`
	MTRL_MKR_ID       string `bson:"MTRL_MKR_ID, omitempty"`
	HIS_FLG           string `bson:"HIS_FLG, omitempty"`
	MTRL_LOT_FLG      string `bson:"MTRL_LOT_FLG, omitempty"`
	VERIFY_FLG        string `bson:"VERIFY_FLG, omitempty"`
	ROOT_MTRL_PRODUCT string `bson:"ROOT_MTRL_PRODUCT, omitempty"`
	MTRL_TRNS_CATE    string `bson:"MTRL_TRNS_CATE, omitempty"`
	MTRL_QTIM_CATE    string `bson:"MTRL_QTIM_CATE, omitempty"`
	MTRL_MAX_RWK_CNT  string `bson:"MTRL_MAX_RWK_CNT, omitempty"`
	TAR_CON_TYP       string `bson:"TAR_CON_TYP, omitempty"`
	MTRL_USAGE_FLG    string `bson:"MTRL_USAGE_FLG, omitempty"`
	MTRL_USAGE_CATE   string `bson:"MTRL_USAGE_CATE, omitempty"`
	POSITION_FLG      string `bson:"POSITION_FLG, omitempty"`
	POSITION_CATE     string `bson:"POSITION_CATE, omitempty"`
	AX_ATT_CNT        string `bson:"AX_ATT_CNT, omitempty"`
	CAPACITY          string `bson:"CAPACITY, omitempty"`
	DFT_MTRL_STAT     string `bson:"DFT_MTRL_STAT, omitempty"`
	ADDT_INFO_1       string `bson:"ADDT_INFO_1, omitempty"`
	ADDT_INFO_2       string `bson:"ADDT_INFO_2, omitempty"`
	ADDT_INFO_3       string `bson:"ADDT_INFO_3, omitempty"`
	SUB_MTRL          string `bson:"SUB_MTRL, omitempty"`
	MTRL_UNIT         string `bson:"MTRL_UNIT, omitempty"`
	PLAN_QTY          string `bson:"PLAN_QTY, omitempty"`
	MIN_INVENTORY     string `bson:"MIN_INVENTORY, omitempty"`
}

//ACODE : ACODE
type ACODE struct {
	CODE_CATE string `bson:"CODE_CATE, omitempty"`
	CODE      string `bson:"CODE, omitempty"`
	CODE_EXT  string `bson:"CODE_EXT, omitempty"`
	SUBITEM   string `bson:"SUBITEM, omitempty"`
	CODE_DSC  string `bson:"CODE_DSC, omitempty"`
	EXT_1     string `bson:"EXT_1, omitempty"`
	EXT_2     string `bson:"EXT_2, omitempty"`
	EXT_3     string `bson:"EXT_3, omitempty"`
	EXT_4     string `bson:"EXT_4, omitempty"`
	EXT_5     string `bson:"EXT_5, omitempty"`
}

// APRDCT : APRDCT
type APRDCT struct {
	PRODUCT_ID    string `bson:"PRODUCT_ID, omitempty"`
	PRODUCT_DSC   string `bson:"PRODUCT_DSC, omitempty"`
	PRODUCT_DSC2  string `bson:"PRODUCT_DSC2, omitempty"`
	PRODUCT_CATE  string `bson:"PRODUCT_CATE, omitempty"`
	X_AXIS_CNT    string `bson:"X_AXIS_CNT, omitempty"`
	Y_AXIS_CNT    string `bson:"Y_AXIS_CNT, omitempty"`
	PNL_SHT_CNT   string `bson:"PNL_SHT_CNT, omitempty"`
	CHARGE_CODE   string `bson:"CHARGE_CODE, omitempty"`
	GROUP_PN      string `bson:"GROUP_PN, omitempty"`
	GRADE_CODE    string `bson:"GRADE_CODE, omitempty"`
	CONT_TYP      string `bson:"CONT_TYP, omitempty"`
	MAX_RECYC_CNT string `bson:"MAX_RECYC_CNT, omitempty"`
	LAYOUT_ID     string `bson:"LAYOUT_ID, omitempty"`
	PTYPE         string `bson:"PTYPE, omitempty"`
	ACCEPT_GRADE  string `bson:"ACCEPT_GRADE, omitempty"`
	CUSTOMER_ID   string `bson:"CUSTOMER_ID, omitempty"`
	VERI_CODE     string `bson:"VERI_CODE, omitempty"`
	INTERNAL_CODE string `bson:"INTERNAL_CODE, omitempty"`
	ADDT_INFO_1   string `bson:"ADDT_INFO_1, omitempty"`
	ADDT_INFO_2   string `bson:"ADDT_INFO_2, omitempty"`
	ADDT_INFO_3   string `bson:"ADDT_INFO_3, omitempty"`
}

// AEQPT : AEQPT
type AEQPT struct {
	EQPT_ID           string `bson:"EQPT_ID, omitempty"`
	EQPT_DSC          string `bson:"EQPT_DSC, omitempty"`
	SYS_SUFFIX        string `bson:"SYS_SUFFIX, omitempty"`
	UNIT_TYP          string `bson:"UNIT_TYP, omitempty"`
	ROOT_EQPT_ID      string `bson:"ROOT_EQPT_ID, omitempty"`
	PRT_EQPT_IDv      string `bson:"PRT_EQPT_IDv, omitempty"`
	EQPT_TYP          string `bson:"EQPT_TYP, omitempty"`
	EQPT_SUBTYP       string `bson:"EQPT_SUBTYP, omitempty"`
	EQPT_CATE         string `bson:"EQPT_CATE, omitempty"`
	KANBAN            string `bson:"KANBAN, omitempty"`
	BAY_ID            string `bson:"BAY_ID, omitempty"`
	BCS_NODE          string `bson:"BCS_NODE, omitempty"`
	TCS_NODE          string `bson:"TCS_NODE, omitempty"`
	MAX_SHT_CNT       string `bson:"MAX_SHT_CNT, omitempty"`
	MAX_SHT_TIME      string `bson:"MAX_SHT_TIME, omitempty"`
	LOOR_CODE         string `bson:"LOOR_CODE, omitempty"`
	CRR_CLN_ROUTE_ID  string `bson:"CRR_CLN_ROUTE_ID, omitempty"`
	CRR_CLN_ROUTE_VER string `bson:"CRR_CLN_ROUTE_VER, omitempty"`
	QPT_TRNS_CATE     string `bson:"QPT_TRNS_CATE, omitempty"`
	BATCH_FLG         string `bson:"BATCH_FLG, omitempty"`
	NL_STC_EQP_FLG    string `bson:"NL_STC_EQP_FLG, omitempty"`
	CNCT_EQPT_ID      string `bson:"CNCT_EQPT_ID, omitempty"`
	AREA_CODE         string `bson:"AREA_CODE, omitempty"`
	INE_ID            string `bson:"INE_ID, omitempty"`
	ADDT_INFO_1       string `bson:"ADDT_INFO_1, omitempty"`
	ADDT_INFO_2       string `bson:"ADDT_INFO_2, omitempty"`
	ADDT_INFO_3       string `bson:"ADDT_INFO_3, omitempty"`
}

// AEQPTG : AEQPTG
type AEQPTG struct {
	EQPTG_ID       string `bson:"EQPTG_ID, omitempty"`
	EQPTG_DSC      string `bson:"EQPTG_DSC, omitempty"`
	EQP_STK_FLG    string `bson:"EQP_STK_FLG, omitempty"`
	REP_EQPT_ID    string `bson:"REP_EQPT_ID, omitempty"`
	UB_REP_EQPT_ID string `bson:"UB_REP_EQPT_ID, omitempty"`
	PLURAL_DEST    string `bson:"PLURAL_DEST, omitempty"`
	DEF_RECIPE_ID  string `bson:"DEF_RECIPE_ID, omitempty"`
	STAT_CHG_FLG   string `bson:"STAT_CHG_FLG, omitempty"`
}

// AEQPTGR : AEQPTGR
type AEQPTGR struct {
	EQPT_ID       string `bson:"EQPT_ID, omitempty"`
	EQPTG_ID      string `bson:"EQPTG_ID, omitempty"`
	ASSIGN_VALID  string `bson:"ASSIGN_VALID, omitempty"`
	OPI_FLG       string `bson:"OPI_FLG, omitempty"`
	PH_RECIPE_FLG string `bson:"PH_RECIPE_FLG, omitempty"`
	PH_RECIPE_ID  string `bson:"PH_RECIPE_ID, omitempty"`
	PRCP_EFF_FLG  string `bson:"PRCP_EFF_FLG, omitempty"`
	EQPT_PRTY     string `bson:"EQPT_PRTY, omitempty"`
	EQPTG_PRTY    string `bson:"EQPTG_PRTY, omitempty"`
}

type AEQPTPT struct {
	EQPT_ID          string `bson:"EQPT_ID, omitempty"`
	PORT_ID          string `bson:"PORT_ID, omitempty"`
	PORT_TYP         string `bson:"PORT_TYP, omitempty"`
	PORT_SUB_TYP     string `bson:"PORT_SUB_TYP, omitempty"`
	ASM_PORT_TYP     string `bson:"ASM_PORT_TYP, omitempty"`
	PORT_STAT        string `bson:"PORT_STAT, omitempty"`
	PORT_STAT_DSC    string `bson:"PORT_STAT_DSC, omitempty"`
	PORT_DATE        string `bson:"PORT_DATE, omitempty"`
	PORT_TIME        string `bson:"PORT_TIME, omitempty"`
	AGV_MODE         string `bson:"AGV_MODE, omitempty"`
	SUP_STOCK_DR_FLG string `bson:"SUP_STOCK_DR_FLG, omitempty"`
	SUP_STOCKER_ID   string `bson:"SUP_STOCKER_ID, omitempty"`
	RET_STOCK_DR_FLG string `bson:"RET_STOCK_DR_FLG, omitempty"`
	RET_STOCKER_ID   string `bson:"RET_STOCKER_ID, omitempty"`
	XFER_ID          string `bson:"XFER_ID, omitempty"`
	XFER_NODE        string `bson:"XFER_NODE, omitempty"`
	ALLOW_STOCK_OUT  string `bson:"ALLOW_STOCK_OUT, omitempty"`
	FORCE_EMPTY_CRR  string `bson:"FORCE_EMPTY_CRR, omitempty"`
	PI_LOT_SEL_TYP   string `bson:"PI_LOT_SEL_TYP, omitempty"`
	XFER_FAIL        string `bson:"XFER_FAIL, omitempty"`
	XFER_FAIL_TO_SUP string `bson:"XFER_FAIL_TO_SUP, omitempty"`
	XFER_AVAIL_FLG   string `bson:"XFER_AVAIL_FLG, omitempty"`
	LAST_XFER_DATE   string `bson:"LAST_XFER_DATE, omitempty"`
	LAST_XFER_TIME   string `bson:"LAST_XFER_TIME, omitempty"`
	INL_STOCKER_ID   string `bson:"INL_STOCKER_ID, omitempty"`
	PORT_ENABLE_FLG  string `bson:"PORT_ENABLE_FLG, omitempty"`
	MAP_FLG          string `bson:"MAP_FLG, omitempty"`
	LAST_PUSH_DATE   string `bson:"LAST_PUSH_DATE, omitempty"`
	LAST_PUSH_TIME   string `bson:"LAST_PUSH_TIME, omitempty"`
	PAIR_FLG         string `bson:"PAIR_FLG, omitempty"`
	RPT_SHT_ID       string `bson:"RPT_SHT_ID, omitempty"`
	CRR_SET_CODE     string `bson:"CRR_SET_CODE, omitempty"`
	NG_TYPE          string `bson:"NG_TYPE, omitempty"`
	SRT_SETTING_NO   string `bson:"SRT_SETTING_NO, omitempty"`
	PARTIAL_FULL_FLG string `bson:"PARTIAL_FULL_FLG, omitempty"`
}

// AEQPTRESV : AEQPTRESV
type AEQPTRESV struct {
	LOT_ID          string `bson:"LOT_ID"`
	NX_OPE_NO       string `bson:"NX_OPE_NO"`
	NX_OPE_ID       string `bson:"NX_OPE_ID"`
	NX_OPE_VER      string `bson:"NX_OPE_VER"`
	SPLT_ID         string `bson:"SPLT_ID"`
	RESV_EQPT_ID    string `bson:"RESV_EQPT_ID"`
	LOT_STAT        string `bson:"LOT_STAT"`
	RUN_FLAG        string `bson:"RUN_FLAG"`
	RESV_DATE       string `bson:"RESV_DATE"`
	RESV_SHIFT_SEQ  string `bson:"RESV_SHIFT_SEQ"`
	RESV_COMMENT    string `bson:"RESV_COMMENT"`
	CLAIM_DATE      string `bson:"CLAIM_DATE"`
	CLAIM_TIME      string `bson:"CLAIM_TIME"`
	CLAIM_USER      string `bson:"CLAIM_USER"`
	MOVE_IN_WEIGHT  string `bson:"MOVE_IN_WEIGHT"`
	MOVE_OUT_WEIGHT string `bson:"MOVE_OUT_WEIGHT"`
	PLAN_OPT_WEIGHT string `bson:"PLAN_OPT_WEIGHT"`
	SHT_CNT         string `bson:"SHT_CNT"`
	CR_SHT_CNT      string `bson:"CR_SHT_CNT"`
	FIT_EQPTS       string `bson:"FIT_EQPTS"`
}

// APARAM : APARAM
type APARAM struct {
	EQPT_ID        string `bson:"EQPT_ID, omitempty"`
	PRODUCT_ID     string `bson:"PRODUCT_IDs, omitempty"`
	EC_CODE        string `bson:"EC_CODE, omitempty"`
	ROUTE_ID       string `bson:"ROUTE_ID, omitempty"`
	ROUTE_VER      string `bson:"ROUTE_VER, omitempty"`
	OPE_ID         string `bson:"OPE_ID, omitempty"`
	OPE_VER        string `bson:"OPE_VER, omitempty"`
	RETICLE_SET_ID string `bson:"RETICLE_SET_ID, omitempty"`
	OLD_SET_ID     string `bson:"OLD_SET_ID, omitempty"`
	UNIT_MODE      string `bson:"UNIT_MODE, omitempty"`
	RECIPE_ID      string `bson:"RECIPE_ID, omitempty"`
	CTIVE_FLAG     string `bson:"CTIVE_FLAG, omitempty"`
	SKIP_FLG       string `bson:"SKIP_FLG, omitempty"`
	MES_ID         string `bson:"MES_ID, omitempty"`
	RR_SET_CODE    string `bson:"RR_SET_CODE, omitempty"`
	BOM_ID         string `bson:"BOM_ID, omitempty"`
	BIN_ID         string `bson:"BIN_ID, omitempty"`
	CHG_FROM       string `bson:"CHG_FROM, omitempty"`
	CHG_USER       string `bson:"CHG_USER, omitempty"`
	CHG_DATE       string `bson:"CHG_DATE, omitempty"`
	HG_TIME        string `bson:"HG_TIME, omitempty"`
	PROBE_FLAG     string `bson:"PROBE_FLAG, omitempty"`
}

// APARAMSPEC : APARAMSPEC
type APARAMSPEC struct {
	EQPT_ID    string `bson:"EQPT_ID, omitempty"`
	PRODUCT_ID string `bson:"PRODUCT_ID, omitempty"`
	EC_CODE    string `bson:"EC_CODE, omitempty"`
	OUTE_ID    string `bson:"OUTE_ID, omitempty"`
	ROUTE_VER  string `bson:"ROUTE_VER, omitempty"`
	OPE_ID     string `bson:"OPE_ID, omitempty"`
	OPE_VER    string `bson:"OPE_VER, omitempty"`
	S_NAME     string `bson:"S_NAME, omitempty"`
	S_VALUE    string `bson:"S_VALUE, omitempty"`
	S_DESC     string `bson:"S_DESC, omitempty"`
}

// AOPER : AOPER
type AOPER struct {
	OPE_ID          string `bson:"OPE_ID, omitempty"`
	OPE_VER         string `bson:"OPE_VER, omitempty"`
	OPE_DSC         string `bson:"OPE_DSC, omitempty"`
	PROC_ID         string `bson:"PROC_ID, omitempty"`
	EQPTG_ID        string `bson:"EQPTG_ID, omitempty"`
	PEP_LVL         string `bson:"PEP_LVL, omitempty"`
	DEPT_CODE       string `bson:"DEPT_CODE, omitempty"`
	STD_LD_TIME     string `bson:"STD_LD_TIME, omitempty"`
	MAN_OPE_TIME    string `bson:"MAN_OPE_TIME, omitempty"`
	UP_LOAD_ID      string `bson:"UP_LOAD_ID, omitempty"`
	OWN_LOAD_ID     string `bson:"OWN_LOAD_ID, omitempty"`
	CRR_CLN_FLG     string `bson:"CRR_CLN_FLG, omitempty"`
	RECIPE_RULE_FLG string `bson:"RECIPE_RULE_FLG, omitempty"`
	REP_UNIT        string `bson:"REP_UNIT, omitempty"`
	FC_BANK_ID      string `bson:"FC_BANK_ID, omitempty"`
	LD_PFC_BANK_ID  string `bson:"LD_PFC_BANK_ID, omitempty"`
	PTT_CHK_FLG     string `bson:"PTT_CHK_FLG, omitempty"`
	TEST_REP_TYP    string `bson:"TEST_REP_TYP, omitempty"`
	FC_TYP          string `bson:"FC_TYP, omitempty"`
	USE_PFC_FLG     string `bson:"USE_PFC_FLG, omitempty"`
	QRS_RST_FLG     string `bson:"QRS_RST_FLG, omitempty"`
	METAL_FLG_TYP   string `bson:"METAL_FLG_TYP, omitempty"`
	COR_UPLDID      string `bson:"COR_UPLDID, omitempty"`
	COR_DNLDID      string `bson:"COR_DNLDID, omitempty"`
	EL_UPLDID       string `bson:"EL_UPLDID, omitempty"`
	DEL_DNLDID      string `bson:"DEL_DNLDID, omitempty"`
	ADDT_INFO_1     string `bson:"DDT_INFO_2, omitempty"`
	ADDT_INFO_2     string `bson:"DDT_INFO_2, omitempty"`
	ADDT_INFO_3     string `bson:"ADDT_INFO_3, omitempty"`
}

// AOPER_ACTIVE : AOPER_ACTIVE
type AOPER_ACTIVE struct {
	OPE_ID      string `bson:"OPE_ID, omitempty"`
	OPE_VER     string `bson:"OPE_VER, omitempty"`
	UPDATE_USER string `bson:"UPDATE_USER, omitempty"`
	UPDATE_TIME string `bson:"UPDATE_TIME, omitempty"`
}

// AOPER_INSTRUCT : AOPER_INSTRUCT
type AOPER_INSTRUCT struct {
	OPE_ID      string `bson:"OPE_ID, omitempty"`
	OPE_VER     string `bson:"OPE_VER, omitempty"`
	INSTRUCT_01 string `bson:"INSTRUCT_01, omitempty"`
	INSTRUCT_02 string `bson:"INSTRUCT_02, omitempty"`
	INSTRUCT_03 string `bson:"INSTRUCT_03, omitempty"`
	INSTRUCT_04 string `bson:"INSTRUCT_04, omitempty"`
	INSTRUCT_05 string `bson:"INSTRUCT_05, omitempty"`
	INSTRUCT_06 string `bson:"INSTRUCT_06, omitempty"`
	INSTRUCT_07 string `bson:"INSTRUCT_07, omitempty"`
	INSTRUCT_08 string `bson:"INSTRUCT_08, omitempty"`
	INSTRUCT_09 string `bson:"INSTRUCT_09, omitempty"`
	INSTRUCT_10 string `bson:"INSTRUCT_10, omitempty"`
	INSTRUCT_11 string `bson:"INSTRUCT_11, omitempty"`
	INSTRUCT_12 string `bson:"INSTRUCT_12, omitempty"`
	INSTRUCT_13 string `bson:"INSTRUCT_13, omitempty"`
	INSTRUCT_14 string `bson:"INSTRUCT_14, omitempty"`
	INSTRUCT_15 string `bson:"INSTRUCT_15, omitempty"`
}

// APRDCTMT : APRDCTMT
type APRDCTMT struct {
	PRODUCT_ID      string `bson:"PRODUCT_ID, omitempty"`
	MTRL_PRODUCT_ID string `bson:"MTRL_PRODUCT_ID, omitempty"`
	OPE_ID          string `bson:"OPE_ID, omitempty"`
	LIMIT_FLG       string `bson:"LIMIT_FLG, omitempty"`
	PARTG_ID        string `bson:"PARTG_ID, omitempty"`
}

// APRDCT_ACTIVE : APRDCT_ACTIVE
type APRDCT_ACTIVE struct {
	PRODUCT_ID  string `bson:"PRODUCT_ID, omitempty"`
	EC_CODE     string `bson:"EC_CODE, omitempty"`
	UPDATE_USER string `bson:"UPDATE_USER, omitempty"`
	UPDATE_TIME string `bson:"UPDATE_TIME, omitempty"`
}

// APRDCTRT : APRDCTRT
type APRDCTRT struct {
	PRODUCT_ID string `bson:"PRODUCT_ID, omitempty"`
	EC_CODE    string `bson:"EC_CODE, omitempty"`
	EC_DSC     string `bson:"EC_DSC, omitempty"`
	ROUTE_ID   string `bson:"ROUTE_ID, omitempty"`
	ROUTE_VER  string `bson:"ROUTE_VER, omitempty"`
	BM_PATTERN string `bson:"BM_PATTERN, omitempty"`
}

// AROUTE_ACTIVE : AROUTE_ACTIVE
type AROUTE_ACTIVE struct {
	ROUTE_ID    string `bson:"ROUTE_ID, omitempty"`
	ROUTE_VER   string `bson:"ROUTE_VER, omitempty"`
	UPDATE_USER string `bson:"UPDATE_USER, omitempty"`
	UPDATE_TIME string `bson:"UPDATE_TIME, omitempty"`
}

// AROUTEI : AROUTEI
type AROUTEI struct {
	ROUTE_ID     string `bson:"ROUTE_ID, omitempty"`
	ROUTE_VER    string `bson:"ROUTE_VER, omitempty"`
	ROUTE_DSC    string `bson:"ROUTE_DSC, omitempty"`
	ROUTE_CATE   string `bson:"ROUTE_CATE, omitempty"`
	FIRST_OPE_NO string `bson:"FIRST_OPE_NO, omitempty"`
	STR_BANK_ID  string `bson:"STR_BANK_ID, omitempty"`
	END_BANK_ID  string `bson:"END_BANK_ID, omitempty"`
	MAX_RWK_CNT  string `bson:"MAX_RWK_CNT, omitempty"`
	RECYC_CNT_UP string `bson:"RECYC_CNT_UP, omitempty"`
	ADDT_INFO_1  string `bson:"ADDT_INFO_1, omitempty"`
	ADDT_INFO_2  string `bson:"ADDT_INFO_2, omitempty"`
	ADDT_INFO_3  string `bson:"ADDT_INFO_3, omitempty"`
}

// AROUTE : AROUTE
type AROUTE struct {
	ROUTE_ID        string `bson:"ROUTE_ID, omitempty"`
	ROUTE_VER       string `bson:"ROUTE_VER, omitempty"`
	CR_OPE_NO       string `bson:"CR_OPE_NO, omitempty"`
	PV_OPE_NO       string `bson:"PV_OPE_NO, omitempty"`
	NX_OPE_NO       string `bson:"NX_OPE_NO, omitempty"`
	MODULE_ID       string `bson:"MODULE_ID, omitempty"`
	MODULE_VER      string `bson:"MODULE_VER, omitempty"`
	CR_OPE_ID       string `bson:"CR_OPE_ID, omitempty"`
	CR_OPE_VER      string `bson:"CR_OPE_VER, omitempty"`
	RWK_RST_FLG     string `bson:"RWK_RST_FLG, omitempty"`
	RWK_AVL_FLG     string `bson:"RWK_AVL_FLG, omitempty"`
	MPROC_FLG       string `bson:"MPROC_FLG, omitempty"`
	MPROC_ID        string `bson:"MPROC_ID, omitempty"`
	WIP_BANK_FLG    string `bson:"WIP_BANK_FLG, omitempty"`
	DEF_WIP_BANK_ID string `bson:"DEF_WIP_BANK_ID, omitempty"`
	SHP_BANK_FLG    string `bson:"SHP_BANK_FLG, omitempty"`
	STAGE_ID        string `bson:"STAGE_ID, omitempty"`
	READ_GLASS_FLG  string `bson:"READ_GLASS_FLG, omitempty"`
	ERP_WK_CENTER   string `bson:"ERP_WK_CENTER, omitempty"`
}

// PostProCate : PostProCate
func PostProCate(w http.ResponseWriter, r *http.Request) {
	var count int
	if r.Method == "POST" {
		count++
		fmt.Println(count)
		valType := r.FormValue("subitem")
		arrVal := strings.Split(valType, "-")
		var code string
		var subitem string
		for i := 0; i < len(arrVal); i++ {
			code = arrVal[0]
			subitem = arrVal[1]
		}
		usr := r.FormValue("usr")
		indat := r.FormValue("indat")
		indat = strings.Replace(indat, "-", "", 64)
		intime := r.FormValue("intime")
		lot_id := code + subitem + indat + "000" + strconv.Itoa(count)
		fmt.Println(lot_id)
		fmt.Println(valType + "\t" + usr + "\t" + indat + "\t" + intime)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("Connected to MongoDb!")
		collection := client.Database("KV2MongoDB").Collection("APRDCT")
		var results []*APRDCT

		projection := bson.D{
			{"PRODUCT_ID", 1},
			{"PRODUCT_DSC", 1},
			{"_id", 0},
		}

		cur, err := collection.Find(
			context.Background(),
			bson.D{
				{"PRODUCT_CATE", code},
			},
			options.Find().SetProjection(projection),
		)
		// cur, err := collection.Find(context.TODO(),
		// 	bson.D{{{PRODUCT_CATE: "Z"}, {PRODUCT_ID: 1, PRODUCT_DSC: 1}}}, findOptions)
		if err != nil {
			println("Find data error")
		}
		defer cur.Close(ctx)
		//fmt.Println(cur)
		for cur.Next(context.TODO()) {
			var elem APRDCT
			err := cur.Decode(&elem)
			if err != nil {
				panic(err)
			}
			results = append(results, &elem)
		}
		if err := cur.Err(); err != nil {
			panic(err)
		}
		//fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			panic("Err!")
		}
		//fmt.Printf("%s\n", )
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

// GetProCate : GetProCate
func GetProCate(w http.ResponseWriter, r *http.Request) {
	// MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDb!")
	collection := client.Database("KV2MongoDB").Collection("ACODE")
	findOptions := options.Find()
	var results []*ACODE
	cur, err := collection.Find(context.TODO(),
		// bson.D{{
		// 	"SUBITEM",
		// 	bson.D{{
		// 		"$in",
		// 		bson.A{"V", "D", "G", "RUBB", "Z"},
		// 	}},
		// }}, findOptions)
		bson.D{{"CODE_CATE", "PCAT"}}, findOptions)
	if err != nil {
		println("Find data error")
	}
	defer cur.Close(ctx)
	for cur.Next(context.TODO()) {
		var elem ACODE
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		strings.TrimSpace(elem.SUBITEM)
		results = append(results, &elem)

	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	defer cur.Close(ctx)
	//fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)
	//fmt.Println("method:", r.Metho)
	if r.Method == "GET" {
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			panic("Err!")
		}
		//fmt.Printf("%s\n", )
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Write(b)
	}
}

var prodID string
var prodCate string
var endDate string
var startDate string

// PostProID : PostProID
func PostProID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		prodID = r.FormValue("prodID")
		prodCate = r.FormValue("prodCate")
		endDate = r.FormValue("endDate")
		startDate = r.FormValue("startDate")
		fmt.Println(prodID)
		fmt.Println(prodCate)
		fmt.Println(endDate)
		fmt.Println(startDate)

		db, err := config.GetMongoDB()
		if err != nil {
			panic(err)
		} else {
			var results []AEQPTRESV
			err2 := db.C("AEQPTRESV").Find(
				bson.M{
					"LOT_ID": bbson.RegEx{
						Pattern: "^" + prodCate,
						Options: "i",
					},
					"NX_OPE_ID": prodID,
					"RESV_DATE": bson.M{
						"$gte": startDate,
						"$lte": endDate,
					}}).Select(
				bson.M{
					"LOT_ID":          1,
					"NX_OPE_ID":       1,
					"RESV_EQPT_ID":    1,
					"LOT_STAT":        1,
					"RESV_DATE":       1,
					"RESV_SHIFT_SEQ":  1,
					"RESV_COMMENT":    1,
					"CLAIM_DATE":      1,
					"CLAIM_TIME":      1,
					"CLAIM_USER":      1,
					"PLAN_OPT_WEIGHT": 1,
					"SHT_CNT":         1,
				}).All(&results)
			if err2 != nil {
				panic(err)
			} else {
				b, err := json.MarshalIndent(results, "", "\t")
				if err != nil {
					panic("Err!")
				}
				//fmt.Printf("%s\n", )
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
			}
		}
	}
}

// GetLineID : GetLineID
func GetLineID(w http.ResponseWriter, r *http.Request) {
	// MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDb!")
	collection := client.Database("KV2MongoDB").Collection("ACODE")
	findOptions := options.Find()
	var results []*ACODE
	cur, err := collection.Find(context.TODO(),
		// bson.D{{
		// 	"SUBITEM",
		// 	bson.D{{
		// 		"$in",
		// 		bson.A{"V", "D", "G", "RUBB", "Z"},
		// 	}},
		// }}, findOptions)
		bson.D{{"CODE_CATE", "LNID"}}, findOptions)
	if err != nil {
		println("Find data error")
	}
	defer cur.Close(ctx)
	for cur.Next(context.TODO()) {
		var elem ACODE
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		//strings.TrimSpace(elem.SUBITEM)
		results = append(results, &elem)

	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	defer cur.Close(ctx)
	//fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)
	//fmt.Println("method:", r.Metho)
	if r.Method == "GET" {
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			panic("Err!")
		}
		//fmt.Printf("%s\n", )
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

// PostLineID : PostLineID
func PostLineID(w http.ResponseWriter, r *http.Request) {
	var count int
	if r.Method == "POST" {
		count++
		//fmt.Println(coun)

		//fmt.Println("PROD ID: ", prodI)
		valType := r.FormValue("subitem")
		arrVal := strings.Split(valType, "-")
		var code string
		var subitem string
		for i := 0; i < len(arrVal); i++ {
			code = arrVal[0]
			subitem = arrVal[1]
		}
		usr := r.FormValue("usr")
		indat := r.FormValue("indat")
		indat = strings.Replace(indat, "-", "", 64)
		intime := r.FormValue("intime")
		lot_id := code + subitem + indat + "000" + strconv.Itoa(count)
		fmt.Println(lot_id)
		fmt.Println(valType + "\t" + usr + "\t" + indat + "\t" + intime)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("Connected to MongoDb!")
		collection := client.Database("KV2MongoDB").Collection("APRDCT")
		//findOptions := options.Find()
		var results []*APRDCT

		projection := bson.D{
			{"PRODUCT_ID", 1},
			{"PRODUCT_DSC", 1},
			{"_id", 0},
		}

		cur, err := collection.Find(
			context.Background(),
			bson.D{
				{"PRODUCT_CATE", code},
			},
			options.Find().SetProjection(projection),
		)
		// cur, err := collection.Find(context.TODO(),
		// 	bson.D{{{PRODUCT_CATE: "Z"}, {PRODUCT_ID: 1, PRODUCT_DSC: 1}}}, findOptions)
		if err != nil {
			println("Find data error")
		}
		defer cur.Close(ctx)
		//fmt.Println(cur)
		for cur.Next(context.TODO()) {
			var elem APRDCT
			err := cur.Decode(&elem)
			if err != nil {
				panic(err)
			}
			results = append(results, &elem)
		}
		if err := cur.Err(); err != nil {
			panic(err)
		}
		//fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			panic("Err!")
		}
		//fmt.Printf("%s\n", b)
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

// GetOwnerID : GetOwnerID
func GetOwnerID(w http.ResponseWriter, r *http.Request) {
	// MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDb!")
	collection := client.Database("KV2MongoDB").Collection("ACODE")
	findOptions := options.Find()
	var results []*ACODE
	cur, err := collection.Find(context.TODO(),
		bson.D{{"CODE_CATE", "OWNR"}}, findOptions)
	if err != nil {
		println("Find data error")
	}
	defer cur.Close(ctx)
	for cur.Next(context.TODO()) {
		var elem ACODE
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		//strings.TrimSpace(elem.SUBITEM)
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
	defer cur.Close(ctx)
	//fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	//fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			panic("Err!")
		}
		//fmt.Printf("%s\n", b)
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

// GetMaxLOTID : GetMaxLOTID
func GetMaxLOTID(w http.ResponseWriter, r *http.Request) {
	// MongoDB
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB"))
	// if err != nil {
	// 	panic(err)
	// }
	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Connected to MongoDb!")
	// collection := client.Database("KV2MongoDB").Collection("AEQPTRESV")
	// var results []*AEQPTRESV
	// // projection := []bson.M{
	// // 	bson.M{"$match": bson.M{"RESV_DATE": "2019-12-12T00:00:00+07:00"}},
	// // 	bson.M{"$group": bson.M{"_id": bson.M{"LOT_ID": "$LOT_ID"}}},
	// // 	bson.M{"$project": bson.M{"LOT_ID": bson.M{"LOT_ID": "$_id.LOT_ID"}, "_id": 0}},
	// // 	bson.M{"$sort": bson.M{"LOT_ID": 1}},
	// // }
	// pipeline := mongo.Pipeline{
	// 	{{"$match", bson.D{
	// 		{"RESV_DATE", "2019-12-12T00:00:00+07:00"},
	// 	}}},
	// 	{{"$group", bson.D{
	// 		{"_id", bson.D{
	// 			{"LOT_ID", "$LOT_ID"}}},
	// 	}}},
	// 	{{"$project", bson.D{
	// 		{"LOT_ID", bson.D{
	// 			{"LOT_ID", "$_id.LOT_ID"}}},
	// 		{"_id", 0}}}},
	// 	{{"$sort", bson.D{
	// 		{"LOT_ID", 1},
	// 	}}},
	// }
	// cur, err := collection.Aggregate(context.Background(), pipeline)
	// if err != nil {
	// 	println("Find data error")
	// }
	// //fmt.Println(pipeline)
	// for cur.Next(context.Background()) {
	// 	var elem AEQPTRESV
	// 	// err := cur.Decode(&elem)
	// 	// if err != nil {
	// 	// 	panic(err)
	// 	// }
	// 	fmt.Println(elem)
	// 	results = append(results, &elem)
	// }
	// if r.Method == "GET" {
	// 	b, err := json.MarshalIndent(results, "", "\t")
	// 	if err != nil {
	// 		panic("Err!")
	// 	}
	// 	fmt.Printf("%s\n", b)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(b)
	// }
	session, err := mgo.Dial("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("KV2MongoDB").C("AEQPTRESV")
	pipeline := []bson.M{
		bson.M{"$match": bson.M{"NX_OPE_ID": prodID, "LOT_ID": bbson.RegEx{Pattern: "^" + prodCate, Options: "i"}, "RESV_DATE": bson.M{"$gte": startDate, "$lte": endDate}}},
		bson.M{"$group": bson.M{"_id": "", "MAX(LOT_ID)": bson.M{"$max": "$LOT_ID"}}},
		bson.M{"$project": bson.M{"MAX(LOT_ID)": "$MAX(LOT_ID)", "_id": 0}},
	}
	pipe := collection.Pipe(pipeline)
	resp := []bson.M{}
	err = pipe.All(&resp)
	if err != nil {
		fmt.Println("Errored: %#v \n", err)
	}
	b, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// InsertAEQPTRESV : InsertAEQPTRESV
func InsertAEQPTRESV(w http.ResponseWriter, r *http.Request) {
	LOT_ID := r.FormValue("LOT_ID")
	NX_OPE_NO := "";
	NX_OPE_ID := r.FormValue("NX_OPE_ID")
	NX_OPE_VER := ""
	SPLT_ID := ""
	RESV_EQPT_ID := ""
	LOT_STAT := ""
	RUN_FLAG := ""
	RESV_DATE := r.FormValue("RESV_DATE")
	RESV_SHIFT_SEQ := ""
	RESV_COMMENT := ""
	CLAIM_DATE := r.FormValue("CLAIM_DATE")
	CLAIM_TIME := r.FormValue("CLAIM_TIME")
	CLAIM_USER := r.FormValue("CLAIM_USER")
	MOVE_IN_WEIGHT := ""
	MOVE_OUT_WEIGHT := ""
	PLAN_OPT_WEIGHT := r.FormValue("PLAN_OPT_WEIGHT")
	SHT_CNT := ""
	CR_SHT_CNT := ""
	FIT_EQPTS := ""
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin12345@198.1.1.106:27017/?authSource=KV2MongoDB")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		println(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		println(err)
	}

	collection := client.Database("KV2MongoDB").Collection("AEQPTRESV")
	filter := bson.D{
		{"LOT_ID", LOT_ID},
		{"NX_OPE_NO", NX_OPE_NO},
		{"NX_OPE_ID", NX_OPE_ID},
		{"NX_OPE_VER", NX_OPE_VER},
		{"SPLT_ID", SPLT_ID},
		{"RESV_EQPT_ID", RESV_EQPT_ID},
		{"LOT_STAT", LOT_STAT},
		{"RUN_FLAG", RUN_FLAG},
		{"RESV_DATE", RESV_DATE},
		{"RESV_SHIFT_SEQ", RESV_SHIFT_SEQ},
		{"RESV_COMMENT", RESV_COMMENT},
		{"CLAIM_DATE", CLAIM_DATE},
		{"CLAIM_TIME", CLAIM_TIME},
		{"CLAIM_USER", CLAIM_USER},
		{"MOVE_IN_WEIGHT", MOVE_IN_WEIGHT},
		{"MOVE_IN_WEIGHT", MOVE_OUT_WEIGHT},
		{"PLAN_OPT_WEIGHT", PLAN_OPT_WEIGHT},
		{"SHT_CNT", SHT_CNT},
		{"CR_SHT_CNT", CR_SHT_CNT},
		{"FIT_EQPTS", FIT_EQPTS},
	}
	insertResult, err := collection.InsertOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	println("Inserted to database.")
	err = client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connecting to MongoDb closed.")
	fmt.Println("Inserted many document: ", insertResult.InsertedID)
}
