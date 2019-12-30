package mes_model

// ABOMDETAIL : ABOMDETAIL
type ABOMDETAIL struct {
	BOM_ID          string `bson:"BIN_ID"`
	MTRL_PRODUCT_ID string `bson:"MTRL_PRODUCT_ID"`
	PLAN_QTY        string `bson:"PLAN_QTY"`
	MTRL_UNIT       string `bson:"MTRL_UNIT"`
	PARENT_ID       string `bson:"PARENT_ID"`
	PARTG_ID        string `bson:"PARTG_ID"`
}

// ACALEND : ACAEND
type ACALEND struct {
	CAL_DATE       string `bson:"CAL_DATE"`
	HOLI_CODE      string `bson:"HOLI_CODE"`
	MFGSHOP        string `bson:"MFGSHOP"`
	MFGMONTH       string `bson:"MFGMONTH"`
	WEEKDAY        string `bson:"WEEKDAY"`
	FGWEEK         string `bson:"FGWEEK"`
	MFGQUTR        string `bson:"MFGQUTR"`
	SHIFT_1ST_TIME string `bson:"SHIFT_1ST_TIME"`
	SHIFT_2ND_TIME string `bson:"SHIFT_2ND_TIME`
	SHIFT_3RD_TIME string `bson:"SHIFT_3RD_TIME, omitemty"`
}

// ACARRHLD : ACARRHLD
type ACARRHLD struct {
	CRR_ID       string `bson:"CAL_DATE"`
	HLD_SEQ_NO   string `bson:"HLD_SEQ_NO"`
	HLD_USER     string `bson:"HLD_USER"`
	HLD_CATE     string `bson:"HLD_CATE"`
	HLD_CODE     string `bson:"HLD_CODE"`
	HLD_DATE     string `bson:"HLD_DATE"`
	HLD_TIME     string `bson:"HLD_TIME"`
	HLD_MFDT     string `bson:"HLD_MFDT"`
	HLD_MFWK     string `bson:"HLD_MFWK"`
	HLD_MFMN     string `bson:"HLD_MFMN"`
	HLD_MFSH     string `bson:"HLD_MFSH"`
	PLN_REL_DATE string `bson:"PLN_REL_DATE"`
	HLD_NOTE     string `bson:"HLD_NOTE"`
}

// ACARRIER : ACARRIER
type ACARRIER struct {
	CRR_ID              string `bson:"CRR_ID"`
	CRR_STAT            string `bson:"CRR_STAT"`
	VALID_FLG           string `bson:"VALID_FLG"`
	CRR_MNT_RST_DATE    string `bson:"CRR_MNT_RST_DATE"`
	GROUP_ID            string `bson:"GROUP_ID"`
	PPBOX_ID            string `bson:"PPBOX_ID"`
	REAL_EMP            string `bson:"REAL_EMP"`
	SGR_ID              string `bson:"SGR_ID"`
	LOT_ID              string `bson:"LOT_ID"`
	SPLT_ID             string `bson:"SPLT_ID"`
	PRODUCT_CATE        string `bson:"PRODUCT_CATE"`
	PRODUCT_ID          string `bson:"PRODUCT_ID"`
	EC_CODE             string `bson:"EC_CODE"`
	NX_ROUTE_ID         string `bson:"NX_ROUTE_ID"`
	NX_ROUTE_VER        string `bson:"NX_ROUTE_VER"`
	NX_OPE_NO           string `bson:"NX_OPE_NO"`
	X_OPE_ID            string `bson:"X_OPE_ID"`
	NX_OPE_VER          string `bson:"NX_OPE_VER"`
	NX_PROC_ID          string `bson:"NX_PROC_ID"`
	NX_EQPTG_ID         string `bson:"NX_EQPTG_ID"`
	NX_EQPT_ID          string `bson:"NX_EQPT_ID"`
	NX_PEP_LVL          string `bson:"NX_PEP_LVL"`
	NX_DEPT_CODE        string `bson:"NX_DEPT_CODE"`
	NX_RECIPE_ID        string `bson:"NX_RECIPE_ID"`
	EQPT_ID             string `bson:"EQPT_ID"`
	EQPT_PORT_ID        string `bson:"EQPT_PORT_ID"`
	HLD_USER            string `bson:"HLD_USER"`
	HLD_CODE            string `bson:"HLD_CODE"`
	HLD_DATE            string `bson:"HLD_DATE"`
	HLD_TIME            string `bson:"HLD_TIME"`
	HLD_MFDT            string `bson:"HLD_MFDT"`
	HLD_MFWK            string `bson:"HLD_MFWK"`
	HLD_MFMN            string `bson:"HLD_MFMN"`
	HLD_MFSH            string `bson:"HLD_MFSH"`
	PRTY                string `bson:"PRTY"`
	LOGOF_DATE          string `bson:"LOGOF_DATE"`
	LOGOF_TIME          string `bson:"LOGOF_TIME"` //
	MAX_QRS_AVL_FLG     string `bson:"MAX_QRS_AVL_FLG"`
	MAX_QRS_ROUTE_ID    string `bson:"MAX_QRS_ROUTE_ID"`
	MAX_QRS_ROUTE_VER   string `bson:"MAX_QRS_ROUTE_VER"`
	MAX_QRS_OPE_ID      string `bson:"MAX_QRS_OPE_ID"`
	MAX_QRS_DATE        string `bson:"MAX_QRS_DATE"`
	MAX_QRS_TIME        string `bson:"MAX_QRS_TIME"`
	MIN_QRS_AVL_FLG     string `bson:"MIN_QRS_AVL_FLG"`
	MIN_QRS_ROUTE_ID    string `bson:"MIN_QRS_ROUTE_ID"`
	MIN_QRS_ROUTE_VER   string `bson:"MIN_QRS_ROUTE_VER"`
	MIN_QRS_OPE_ID      string `bson:"MIN_QRS_OPE_ID"`
	MIN_QRS_DATE        string `bson:"MIN_QRS_DATE"`
	MIN_QRS_TIME        string `bson:"MIN_QRS_TIME"`
	SHT_CNT             string `bson:"SHT_CNT"`
	PNL_CNT             string `bson:"PNL_CNT"`
	PLN_REL_DATE        string `bson:"PLN_REL_DATE"`
	ABND_FLG            string `bson:"ABND_FLG"`
	CRR_CLN_ROUTE_ID    string `bson:"CRR_CLN_ROUTE_ID"`
	CRR_CLN_ROUTE_VER   string `bson:"CRR_CLN_ROUTE_VER"`
	INIT_CLUP_FLG       string `bson:"INIT_CLUP_FLG"`
	CLN_CRR_FLG         string `bson:"CLN_CRR_FLG"`
	APPLY_EEN_ID        string `bson:"APPLY_EEN_ID"`
	PR_MIX_FLG          string `bson:"PR_MIX_FLG"`
	LOGOF_EQPT_ID       string `bson:"LOGOF_EQPT_ID"`
	LOGOF_PORT_ID       string `bson:"LOGOF_PORT_ID"`
	LOGOF_RECIPE_ID     string `bson:"LOGOF_RECIPE_ID"`
	PLN_CMP_DATE        string `bson:"PLN_CMP_DATE"`
	PLN_CMP_TIME        string `bson:"PLN_CMP_TIME"`
	STB_SHOP            string `bson:"STB_SHOP"`
	DEST_SHOP           string `bson:"DEST_SHOP"`
	SHT_JUDGE           string `bson:"SHT_JUDGE"`
	XF_CMD_FLG          string `bson:"XF_CMD_FLG"`
	PRIORITY            string `bson:"PRIORITY"`
	XFER_ID             string `bson:"XFER_ID"`
	XFER_TYPE           string `bson:"XFER_TYPE"`
	LN_EQPT_ID          string `bson:"LN_EQPT_ID"`
	PLN_SEQ_NO          string `bson:"PLN_SEQ_NO"`
	REQ_RCV_DATE        string `bson:"REQ_RCV_DATE"`
	REQ_RCV_TIME        string `bson:"REQ_RCV_TIME"`
	RSV_FLG             string `bson:"RSV_FLG"`
	M_RSVD              string `bson:"M_RSVD"`
	RSV_USER_ID         string `bson:"RSV_USER_ID"`
	RSV_DATE            string `bson:"RSV_DATE"`
	RSV_TIME            string `bson:"RSV_TIME"`
	RSV_NO              string `bson:"RSV_NO"`
	RSV_EQPT_ID         string `bson:"RSV_EQPT_ID"`
	RSV_PORT_ID         string `bson:"RSV_PORT_ID"`
	RSV_PATI_ID         string `bson:"RSV_PATI_ID"`
	RSV_ZONE_ID         string `bson:"RSV_ZONE_ID"`
	CRR_USE_CNT         string `bson:"CRR_USE_CNT"`
	MAX_CRR_USE_CNT     string `bson:"MAX_CRR_USE_CNT"`
	MAX_USE_OVER_FLG    string `bson:"MAX_USE_OVER_FLG"`
	XFER_CLM_DATE       string `bson:"XFER_CLM_DATE"`
	XFER_CLM_TIME       string `bson:"XFER_CLM_TIME"`
	NO_XFER_FLG         string `bson:"NO_XFER_FLG"`
	NO_XFER_TYP         string `bson:"NO_XFER_TYP"`
	CLUP_FLG            string `bson:"CLUP_FLG"`
	CLUP_DATE           string `bson:"CLUP_DATE"`
	CLUP_TIME           string `bson:"CLUP_TIME"` //
	RSTRCT_CLUP_DAY     string `bson:"RSTRCT_CLUP_DAY"`
	SIZE_CODE           string `bson:"SIZE_CODE"`
	OWNER_ID            string `bson:"OWNER_ID"`
	SALES_ORDER         string `bson:"SALES_ORDER"`
	ABNORMAL_FLG_1      string `bson:"ABNORMAL_FLG_1"`
	CRR_CATE            string `bson:"CRR_CATE"`
	CRR_SIZE            string `bson:"CRR_SIZE"`
	CRR_SET_CODE        string `bson:"CRR_SET_CODE"`
	CRR_MKR_ID          string `bson:"CRR_MKR_ID"`
	CROSS_FLG           string `bson:"CROSS_FLG"`
	CRR_RGST_DATE       string `bson:"CRR_RGST_DATE"`
	CRR_RGST_TIME       string `bson:"CRR_RGST_TIME"`
	CRR_OWN             string `bson:"CRR_OWN"`
	POSITION            string `bson:"POSITION"`
	PORT_ID             string `bson:"PORT_ID"`
	PATI_ID             string `bson:"PATI_ID"`
	ZONE_ID             string `bson:"ZONE_ID"`
	SHELF_ID            string `bson:"POS_CHG_DATE"`
	POS_CHG_DATE        string `bson:"POS_CHG_DATE"`
	POS_CHG_TIME        string `bson:"POS_CHG_TIME"`
	XF_CMD_STAT         string `bson:"XF_CMD_STAT"`
	XFER_STAT           string `bson:"XFER_STAT"`
	FR_EQPT_ID          string `bson:"FR_EQPT_ID"`
	FR_PORT_ID          string `bson:"FR_PORT_ID"`
	FR_PATI_ID          string `bson:"FR_PATI_ID"` //
	FR_ZONE_ID          string `bson:"FR_ZONE_ID"`
	TO_EQPT_ID          string `bson:"TO_EQPT_ID"`
	TO_PORT_ID          string `bson:"TO_PORT_ID"`
	TO_PATI_ID          string `bson:"TO_PATI_ID"`
	TO_ZONE_ID          string `bson:"TO_ZONE_ID"`
	PUSH_STOCKER_ID     string `bson:"PUSH_STOCKER_ID"`
	PUSH_PATI_ID        string `bson:"PUSH_PATI_ID"`
	PUSH_ZONE_ID        string `bson:"PUSH_ZONE_ID"`
	LST_STC_IN_DATE     string `bson:"LST_STC_IN_DATE"`
	LST_STC_IN_TIME     string `bson:"LST_STC_IN_TIME"`
	SWEEP_DATE          string `bson:"SWEEP_DATE"`
	SWEEP_TIME          string `bson:"SWEEP_TIME"`
	XFER_RTN_CODE       string `bson:"XFER_RTN_CODE"`
	LAST_XFR_CMD_RC     string `bson:"LAST_XFR_CMD_RC"`
	LAST_LOC_RPT_RC     string `bson:"LAST_LOC_RPT_RC"`
	CRRG_ID             string `bson:"CRRG_ID"`
	STOCK_STAT          string `bson:"STOCK_STAT"`
	MAX_QRK_AVL_FLG     string `bson:"MAX_QRK_AVL_FLG"`
	MAX_QRK_ROUTE_ID    string `bson:"MAX_QRK_ROUTE_ID"`
	MAX_QRK_ROUTE_VER   string `bson:"MAX_QRK_ROUTE_VER"`
	MAX_QRK_OPE_ID      string `bson:"MAX_QRK_OPE_ID"`
	MAX_QRK_DATE        string `bson:"MAX_QRK_DATE"`
	MAX_QRK_TIME        string `bson:"MAX_QRK_TIME"`
	SETTING_NO          string `bson:"SETTING_NO"`
	CRR_RGST_OWN        string `bson:"CRR_RGST_OWN"`
	STB_SAMPLING_FLG    string `bson:"STB_SAMPLING_FLG"`
	FINISHED_DATE       string `bson:"FINISHED_DATE"`
	FINISHED_TIME       string `bson:"FINISHED_TIME"`
	REVW_SHT_JUDGE      string `bson:"REVW_SHT_JUDGE"`
	MTRL_GRADE          string `bson:"MTRL_GRADE"`
	AREA_CODE           string `bson:"AREA_CODE"`
	MTRL_PRODUCT_ID     string `bson:"MTRL_PRODUCT_ID"`
	CARTON_ID           string `bson:"CARTON_ID"`
	MIX_OWNER           string `bson:"MIX_OWNER"`
	REJ_EQPT_ID         string `bson:"REJ_EQPT_ID"`
	REJ_PORT_ID         string `bson:"REJ_PORT_ID"`
	REJ_PROC_ID         string `bson:"REJ_PROC_ID"`
	UNIT                string `bson:"UNIT"`
	STAGE_ID            string `bson:"STAGE_ID"`
	WORDER_ID           string `bson:"WORDER_ID"`
	XFER_MODE           string `bson:"XFER_MODE"`
	ORDER_TYP           string `bson:"ORDER_TYP"`
	MAX_QRS_OPE_NO      string `bson:"MAX_QRS_OPE_NO"`
	MIN_QRS_OPE_NO      string `bson:"MIN_QRS_OPE_NO"`
	MAX_QRK_OPE_NO      string `bson:"MAX_QRK_OPE_NO"`
	ORG_LR_PRODUCT_ID   string `bson:"ORG_LR_PRODUCT_ID"`
	ORG_LR_MTRL_PRODUCT string `bson:"ORG_LR_MTRL_PRODUCT"`
	ORG_LR_LOT_ID       string `bson:"ORG_LR_LOT_ID"`
	ORG_LR_SPLT_ID      string `bson:"ORG_LR_SPLT_ID"`
	ORG_LR_SHT_CNT      string `bson:"ORG_LR_SHT_CNT"`
	IFTHEN_ACT_FLG      string `bson:"IFTHEN_ACT_FLG"`
}

// ABOM : ABOM
type ABOM struct {
	BOM_ID      string `bson:"BIN_ID"`
	BOM_DSC     string `bson:"BOM_DSC"`
	ADDT_INFO_1 string `bson:"ADDT_INFO_1"`
	ADDT_INFO_2 string `bson:"ADDT_INFO_2"`
	ADDT_INFO_3 string `bson:"ADDT_INFO_3"`
}

// ABINJDGRL : ABINJDGRL
type ABINJDGRL struct {
	BIN_ID        string `bson:"BIN_ID"`
	JUDGE_RULE_ID string `bson:"JUDGE_RULE_ID"`
}

// ABINGRADE : ABINGRADE
type ABINGRADE struct {
	BIN_ID    string `bson:"BIN_ID"`
	BIN_NO    string `bson:"BIN_NO"`
	SHT_GRADE string `bson:"SHT_GRADE"`
	HOLD_FLG  string `bson:"HOLD_FLG"`
}

// ABILBDLN : ABILBDLN
type ABILBDLN struct {
	CODE      string `bson:"CODE"`
	BILBD_MSG string `bson:"BILBD_MSG"`
	BOARD_MSG string `bson:"BOARD_MSG"`
}

//ABAYAREA : ABAYAREA
type ABAYAREA struct {
	BAY_ID      string `bson:"BAY_ID"`
	OPI_FLG     string `bson:"OPI_FLG"`
	BAY_DSC     string `bson:"BAY_DSC"`
	ADDT_INFO_1 string `bson:"ADDT_INFO_1"`
	ADDT_INFO_2 string `bson:"ADDT_INFO_2"`
	ADDT_INFO_3 string `bson:"ADDT_INFO_3"`
}

// AAQLLEVEL : AAQLLEVEL
type AAQLLEVEL struct {
	AQL_LVL      string `bson:"AQL_LVL"`
	SEQ_NO       string `bson:"SEQ_NO"`
	AQL_FLG      string `bson:"AQL_FLG"`
	PNL_QTY_FROM string `bson:"PNL_QTY_FROM"`
	PNL_QTY_TO   string `bson:"PNL_QTY_TO"`
	SMP_QTY      string `bson:"SMP_QTY"`
	ACCEPT_QTY   string `bson:"ACCEPT_QTY"`
	UPD_USER_ID  string `bson:"UPD_USER_ID"`
	UPD_DATE     string `bson:"UPD_DATE"`
	UPD_TIME     string `bson:"UPD_TIME"`
}

// AALERT : AALERT
type AALERT struct {
	EQPT_ID          string `bson:"EQPT_ID"`
	RPT_DATE         string `bson:"RPT_DATE"`
	RPT_TIME         string `bson:"RPT_TIME"`
	ALRT_ID          string `bson:"ALRT_ID"`
	RPT_SOURCE       string `bson:"RPT_SOURCE"`
	ALRT_CODE        string `bson:"ALRT_CODE"`
	ALRT_LVL         string `bson:"ALRT_LVL"`
	ALERT_ON_OFF_FLG string `bson:"ALERT_ON_OFF_FLG"`
	ALRT_DSC         string `bson:"ALRT_DSC"`
	ALRT_COMMENT     string `bson:"ALRT_COMMENT"`
	CFM_USER_ID      string `bson:"CFM_USER_ID"`
	CFM_DATE         string `bson:"CFM_DATE"`
	CFM_TIME         string `bson:"CFM_TIME"`
	CFM_COMMENT      string `bson:"CFM_COMMENT"`
}

//ABIN : ABIN
type ABIN struct {
	BIN_ID      string `bson:"BIN_ID"`
	BIN_DSC     string `bson:"BIN_DSC"`
	ADDT_INFO_1 string `bson:"ADDT_INFO_1"`
	ADDT_INFO_2 string `bson:"ADDT_INFO_2"`
	ADDT_INFO_3 string `bson:"ADDT_INFO_3"`
}

// AMTRLPRD : AMTRLPRD
type AMTRLPRD struct {
	MTRL_PRODUCT_ID   string `bson:"MTRL_PRODUCT_ID"`
	MTRL_PRODUCT_DSC  string `bson:"MTRL_PRODUCT_DSC"`
	MTRL_CATE         string `bson:"MTRL_CATE"`
	MTRL_MKR_ID       string `bson:"MTRL_MKR_ID"`
	HIS_FLG           string `bson:"HIS_FLG"`
	MTRL_LOT_FLG      string `bson:"MTRL_LOT_FLG"`
	VERIFY_FLG        string `bson:"VERIFY_FLG"`
	ROOT_MTRL_PRODUCT string `bson:"ROOT_MTRL_PRODUCT"`
	MTRL_TRNS_CATE    string `bson:"MTRL_TRNS_CATE"`
	MTRL_QTIM_CATE    string `bson:"MTRL_QTIM_CATE"`
	MTRL_MAX_RWK_CNT  string `bson:"MTRL_MAX_RWK_CNT"`
	TAR_CON_TYP       string `bson:"TAR_CON_TYP"`
	MTRL_USAGE_FLG    string `bson:"MTRL_USAGE_FLG"`
	MTRL_USAGE_CATE   string `bson:"MTRL_USAGE_CATE"`
	POSITION_FLG      string `bson:"POSITION_FLG"`
	POSITION_CATE     string `bson:"POSITION_CATE"`
	AX_ATT_CNT        string `bson:"AX_ATT_CNT"`
	CAPACITY          string `bson:"CAPACITY"`
	DFT_MTRL_STAT     string `bson:"DFT_MTRL_STAT"`
	ADDT_INFO_1       string `bson:"ADDT_INFO_1"`
	ADDT_INFO_2       string `bson:"ADDT_INFO_2"`
	ADDT_INFO_3       string `bson:"ADDT_INFO_3"`
	SUB_MTRL          string `bson:"SUB_MTRL"`
	MTRL_UNIT         string `bson:"MTRL_UNIT"`
	PLAN_QTY          string `bson:"PLAN_QTY"`
	MIN_INVENTORY     string `bson:"MIN_INVENTORY"`
}

//ACODE : ACODE
type ACODE struct {
	CODE_CATE string `bson:"CODE_CATE"`
	CODE      string `bson:"CODE"`
	CODE_EXT  string `bson:"CODE_EXT"`
	SUBITEM   string `bson:"SUBITEM"`
	CODE_DSC  string `bson:"CODE_DSC"`
	EXT_1     string `bson:"EXT_1"`
	EXT_2     string `bson:"EXT_2"`
	EXT_3     string `bson:"EXT_3"`
	EXT_4     string `bson:"EXT_4"`
	EXT_5     string `bson:"EXT_5"`
}

// APRDCT : APRDCT
type APRDCT struct {
	PRODUCT_ID    string `bson:"PRODUCT_ID"`
	PRODUCT_DSC   string `bson:"PRODUCT_DSC"`
	PRODUCT_DSC2  string `bson:"PRODUCT_DSC2"`
	PRODUCT_CATE  string `bson:"PRODUCT_CATE"`
	X_AXIS_CNT    string `bson:"X_AXIS_CNT"`
	Y_AXIS_CNT    string `bson:"Y_AXIS_CNT"`
	PNL_SHT_CNT   string `bson:"PNL_SHT_CNT"`
	CHARGE_CODE   string `bson:"CHARGE_CODE"`
	GROUP_PN      string `bson:"GROUP_PN"`
	GRADE_CODE    string `bson:"GRADE_CODE"`
	CONT_TYP      string `bson:"CONT_TYP"`
	MAX_RECYC_CNT string `bson:"MAX_RECYC_CNT"`
	LAYOUT_ID     string `bson:"LAYOUT_ID"`
	PTYPE         string `bson:"PTYPE"`
	ACCEPT_GRADE  string `bson:"ACCEPT_GRADE"`
	CUSTOMER_ID   string `bson:"CUSTOMER_ID"`
	VERI_CODE     string `bson:"VERI_CODE"`
	INTERNAL_CODE string `bson:"INTERNAL_CODE"`
	ADDT_INFO_1   string `bson:"ADDT_INFO_1"`
	ADDT_INFO_2   string `bson:"ADDT_INFO_2"`
	ADDT_INFO_3   string `bson:"ADDT_INFO_3"`
}

// AEQPT : AEQPT
type AEQPT struct {
	EQPT_ID           string `bson:"EQPT_ID"`
	EQPT_DSC          string `bson:"EQPT_DSC`
	SYS_SUFFIX        string `bson:"SYS_SUFFIX "`
	UNIT_TYP          string `bson:"UNIT_TYP "`
	ROOT_EQPT_ID      string `bson:"ROOT_EQPT_ID"`
	PRT_EQPT_IDv      string `bson:"PRT_EQPT_IDv"`
	EQPT_TYP          string `bson:"EQPT_TYP"`
	EQPT_SUBTYP       string `bson:"EQPT_SUBTYP"`
	EQPT_CATE         string `bson:"EQPT_CATE"`
	KANBAN            string `bson:"KANBAN"`
	BAY_ID            string `bson:"BAY_ID"`
	BCS_NODE          string `bson:"BCS_NODE"`
	TCS_NODE          string `bson:"TCS_NODE"`
	MAX_SHT_CNT       string `bson:"MAX_SHT_CNT"`
	MAX_SHT_TIME      string `bson:"MAX_SHT_TIME"`
	LOOR_CODE         string `bson:"LOOR_CODE"`
	CRR_CLN_ROUTE_ID  string `bson:"CRR_CLN_ROUTE_ID"`
	CRR_CLN_ROUTE_VER string `bson:"CRR_CLN_ROUTE_VER"`
	QPT_TRNS_CATE     string `bson:"QPT_TRNS_CATE"`
	BATCH_FLG         string `bson:"BATCH_FLG"`
	NL_STC_EQP_FLG    string `bson:"NL_STC_EQP_FLG"`
	CNCT_EQPT_ID      string `bson:"CNCT_EQPT_ID"`
	AREA_CODE         string `bson:"AREA_CODE"`
	INE_ID            string `bson:"INE_ID"`
	ADDT_INFO_1       string `bson:"ADDT_INFO_1"`
	ADDT_INFO_2       string `bson:"ADDT_INFO_2"`
	ADDT_INFO_3       string `bson:"ADDT_INFO_3"`
}

// AEQPTG : AEQPTG
type AEQPTG struct {
	EQPTG_ID       string `bson:"EQPTG_ID"`
	EQPTG_DSC      string `bson:"EQPTG_DSC"`
	EQP_STK_FLG    string `bson:"EQP_STK_FLG"`
	REP_EQPT_ID    string `bson:"REP_EQPT_ID"`
	UB_REP_EQPT_ID string `bson:"UB_REP_EQPT_ID"`
	PLURAL_DEST    string `bson:"PLURAL_DEST"`
	DEF_RECIPE_ID  string `bson:"DEF_RECIPE_ID"`
	STAT_CHG_FLG   string `bson:"STAT_CHG_FLG"`
}

// AEQPTGR : AEQPTGR
type AEQPTGR struct {
	EQPT_ID       string `bson:"EQPT_ID"`
	EQPTG_ID      string `bson:"EQPTG_ID"`
	ASSIGN_VALID  string `bson:"ASSIGN_VALID"`
	OPI_FLG       string `bson:"OPI_FLG"`
	PH_RECIPE_FLG string `bson:"PH_RECIPE_FLG"`
	PH_RECIPE_ID  string `bson:"PH_RECIPE_ID"`
	PRCP_EFF_FLG  string `bson:"PRCP_EFF_FLG"`
	EQPT_PRTY     string `bson:"EQPT_PRTY"`
	EQPTG_PRTY    string `bson:"EQPTG_PRTY"`
}

type AEQPTPT struct {
	EQPT_ID          string `bson:"EQPT_ID"`
	PORT_ID          string `bson:"PORT_ID"`
	PORT_TYP         string `bson:"PORT_TYP"`
	PORT_SUB_TYP     string `bson:"PORT_SUB_TYP"`
	ASM_PORT_TYP     string `bson:"ASM_PORT_TYP"`
	PORT_STAT        string `bson:"PORT_STAT"`
	PORT_STAT_DSC    string `bson:"PORT_STAT_DSC"`
	PORT_DATE        string `bson:"PORT_DATE"`
	PORT_TIME        string `bson:"PORT_TIME"`
	AGV_MODE         string `bson:"AGV_MODE"`
	SUP_STOCK_DR_FLG string `bson:"SUP_STOCK_DR_FLG"`
	SUP_STOCKER_ID   string `bson:"SUP_STOCKER_ID"`
	RET_STOCK_DR_FLG string `bson:"RET_STOCK_DR_FLG"`
	RET_STOCKER_ID   string `bson:"RET_STOCKER_ID"`
	XFER_ID          string `bson:"XFER_ID"`
	XFER_NODE        string `bson:"XFER_NODE"`
	ALLOW_STOCK_OUT  string `bson:"ALLOW_STOCK_OUT"`
	FORCE_EMPTY_CRR  string `bson:"FORCE_EMPTY_CRR"`
	PI_LOT_SEL_TYP   string `bson:"PI_LOT_SEL_TYP"`
	XFER_FAIL        string `bson:"XFER_FAIL"`
	XFER_FAIL_TO_SUP string `bson:"XFER_FAIL_TO_SUP"`
	XFER_AVAIL_FLG   string `bson:"XFER_AVAIL_FLG"`
	LAST_XFER_DATE   string `bson:"LAST_XFER_DATE"`
	LAST_XFER_TIME   string `bson:"LAST_XFER_TIME"`
	INL_STOCKER_ID   string `bson:"INL_STOCKER_ID"`
	PORT_ENABLE_FLG  string `bson:"PORT_ENABLE_FLG"`
	MAP_FLG          string `bson:"MAP_FLG"`
	LAST_PUSH_DATE   string `bson:"LAST_PUSH_DATE"`
	LAST_PUSH_TIME   string `bson:"LAST_PUSH_TIME"`
	PAIR_FLG         string `bson:"PAIR_FLG"`
	RPT_SHT_ID       string `bson:"RPT_SHT_ID"`
	CRR_SET_CODE     string `bson:"CRR_SET_CODE"`
	NG_TYPE          string `bson:"NG_TYPE"`
	SRT_SETTING_NO   string `bson:"SRT_SETTING_NO"`
	PARTIAL_FULL_FLG string `bson:"PARTIAL_FULL_FLG"`
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
	MOVE_IN_WEIGHT  uint64 `bson:"MOVE_IN_WEIGHT"`
	MOVE_OUT_WEIGHT uint64 `bson:"MOVE_OUT_WEIGHT"`
	PLAN_OPT_WEIGHT uint64 `bson:"PLAN_OPT_WEIGHT"`
	SHT_CNT         int64  `bson:"SHT_CNT"`
	CR_SHT_CNT      int64  `bson:"CR_SHT_CNT"`
	FIT_EQPTS       string `bson:"FIT_EQPTS"`
}

// APARAM : APARAM
type APARAM struct {
	EQPT_ID        string `bson:"EQPT_ID"`
	PRODUCT_ID     string `bson:"PRODUCT_IDs"`
	EC_CODE        string `bson:"EC_CODE"`
	ROUTE_ID       string `bson:"ROUTE_ID"`
	ROUTE_VER      string `bson:"ROUTE_VER"`
	OPE_ID         string `bson:"OPE_ID"`
	OPE_VER        string `bson:"OPE_VER"`
	RETICLE_SET_ID string `bson:"RETICLE_SET_ID"`
	OLD_SET_ID     string `bson:"OLD_SET_ID"`
	UNIT_MODE      string `bson:"UNIT_MODE"`
	RECIPE_ID      string `bson:"RECIPE_ID"`
	CTIVE_FLAG     string `bson:"CTIVE_FLAG"`
	SKIP_FLG       string `bson:"SKIP_FLG"`
	MES_ID         string `bson:"MES_ID"`
	RR_SET_CODE    string `bson:"RR_SET_CODE"`
	BOM_ID         string `bson:"BOM_ID"`
	BIN_ID         string `bson:"BIN_ID"`
	CHG_FROM       string `bson:"CHG_FROM"`
	CHG_USER       string `bson:"CHG_USER"`
	CHG_DATE       string `bson:"CHG_DATE"`
	HG_TIME        string `bson:"HG_TIME"`
	PROBE_FLAG     string `bson:"PROBE_FLAG"`
}

// APARAMSPEC : APARAMSPEC
type APARAMSPEC struct {
	EQPT_ID    string `bson:"EQPT_ID"`
	PRODUCT_ID string `bson:"PRODUCT_ID"`
	EC_CODE    string `bson:"EC_CODE"`
	OUTE_ID    string `bson:"OUTE_ID"`
	ROUTE_VER  string `bson:"ROUTE_VER"`
	OPE_ID     string `bson:"OPE_ID"`
	OPE_VER    string `bson:"OPE_VER"`
	S_NAME     string `bson:"S_NAME"`
	S_VALUE    string `bson:"S_VALUE"`
	S_DESC     string `bson:"S_DESC"`
}

// AOPER : AOPER
type AOPER struct {
	OPE_ID          string `bson:"OPE_ID"`
	OPE_VER         string `bson:"OPE_VER"`
	OPE_DSC         string `bson:"OPE_DSC"`
	PROC_ID         string `bson:"PROC_ID"`
	EQPTG_ID        string `bson:"EQPTG_ID"`
	PEP_LVL         string `bson:"PEP_LVL"`
	DEPT_CODE       string `bson:"DEPT_CODE"`
	STD_LD_TIME     string `bson:"STD_LD_TIME"`
	MAN_OPE_TIME    string `bson:"MAN_OPE_TIME"`
	UP_LOAD_ID      string `bson:"UP_LOAD_ID"`
	OWN_LOAD_ID     string `bson:"OWN_LOAD_ID"`
	CRR_CLN_FLG     string `bson:"CRR_CLN_FLG"`
	RECIPE_RULE_FLG string `bson:"RECIPE_RULE_FLG"`
	REP_UNIT        string `bson:"REP_UNIT"`
	FC_BANK_ID      string `bson:"FC_BANK_ID"`
	LD_PFC_BANK_ID  string `bson:"LD_PFC_BANK_ID"`
	PTT_CHK_FLG     string `bson:"PTT_CHK_FLG"`
	TEST_REP_TYP    string `bson:"TEST_REP_TYP"`
	FC_TYP          string `bson:"FC_TYP"`
	USE_PFC_FLG     string `bson:"USE_PFC_FLG"`
	QRS_RST_FLG     string `bson:"QRS_RST_FLG"`
	METAL_FLG_TYP   string `bson:"METAL_FLG_TYP"`
	COR_UPLDID      string `bson:"COR_UPLDID"`
	COR_DNLDID      string `bson:"COR_DNLDID"`
	EL_UPLDID       string `bson:"EL_UPLDID"`
	DEL_DNLDID      string `bson:"DEL_DNLDID"`
	ADDT_INFO_1     string `bson:"DDT_INFO_2"`
	ADDT_INFO_2     string `bson:"DDT_INFO_2"`
	ADDT_INFO_3     string `bson:"ADDT_INFO_3"`
}

// AOPER_ACTIVE : AOPER_ACTIVE
type AOPER_ACTIVE struct {
	OPE_ID      string `bson:"OPE_ID"`
	OPE_VER     string `bson:"OPE_VER"`
	UPDATE_USER string `bson:"UPDATE_USER"`
	UPDATE_TIME string `bson:"UPDATE_TIME"`
}

// AOPER_INSTRUCT : AOPER_INSTRUCT
type AOPER_INSTRUCT struct {
	OPE_ID      string `bson:"OPE_ID"`
	OPE_VER     string `bson:"OPE_VER"`
	INSTRUCT_01 string `bson:"INSTRUCT_01"`
	INSTRUCT_02 string `bson:"INSTRUCT_02"`
	INSTRUCT_03 string `bson:"INSTRUCT_03"`
	INSTRUCT_04 string `bson:"INSTRUCT_04"`
	INSTRUCT_05 string `bson:"INSTRUCT_05"`
	INSTRUCT_06 string `bson:"INSTRUCT_06"`
	INSTRUCT_07 string `bson:"INSTRUCT_07"`
	INSTRUCT_08 string `bson:"INSTRUCT_08"`
	INSTRUCT_09 string `bson:"INSTRUCT_09"`
	INSTRUCT_10 string `bson:"INSTRUCT_10"`
	INSTRUCT_11 string `bson:"INSTRUCT_11"`
	INSTRUCT_12 string `bson:"INSTRUCT_12"`
	INSTRUCT_13 string `bson:"INSTRUCT_13"`
	INSTRUCT_14 string `bson:"INSTRUCT_14"`
	INSTRUCT_15 string `bson:"INSTRUCT_15"`
}

// APRDCTMT : APRDCTMT
type APRDCTMT struct {
	PRODUCT_ID      string `bson:"PRODUCT_ID"`
	MTRL_PRODUCT_ID string `bson:"MTRL_PRODUCT_ID"`
	OPE_ID          string `bson:"OPE_ID"`
	LIMIT_FLG       string `bson:"LIMIT_FLG"`
	PARTG_ID        string `bson:"PARTG_ID"`
}

// APRDCT_ACTIVE : APRDCT_ACTIVE
type APRDCT_ACTIVE struct {
	PRODUCT_ID  string `bson:"PRODUCT_ID"`
	EC_CODE     string `bson:"EC_CODE"`
	UPDATE_USER string `bson:"UPDATE_USER"`
	UPDATE_TIME string `bson:"UPDATE_TIME"`
}

// APRDCTRT : APRDCTRT
type APRDCTRT struct {
	PRODUCT_ID string `bson:"PRODUCT_ID"`
	EC_CODE    string `bson:"EC_CODE"`
	EC_DSC     string `bson:"EC_DSC"`
	ROUTE_ID   string `bson:"ROUTE_ID"`
	ROUTE_VER  string `bson:"ROUTE_VER"`
	BM_PATTERN string `bson:"BM_PATTERN"`
}

// AROUTE_ACTIVE : AROUTE_ACTIVE
type AROUTE_ACTIVE struct {
	ROUTE_ID    string `bson:"ROUTE_ID"`
	ROUTE_VER   string `bson:"ROUTE_VER"`
	UPDATE_USER string `bson:"UPDATE_USER"`
	UPDATE_TIME string `bson:"UPDATE_TIME"`
}

// AROUTEI : AROUTEI
type AROUTEI struct {
	ROUTE_ID     string `bson:"ROUTE_ID"`
	ROUTE_VER    string `bson:"ROUTE_VER"`
	ROUTE_DSC    string `bson:"ROUTE_DSC"`
	ROUTE_CATE   string `bson:"ROUTE_CATE"`
	FIRST_OPE_NO string `bson:"FIRST_OPE_NO"`
	STR_BANK_ID  string `bson:"STR_BANK_ID"`
	END_BANK_ID  string `bson:"END_BANK_ID"`
	MAX_RWK_CNT  string `bson:"MAX_RWK_CNT"`
	RECYC_CNT_UP string `bson:"RECYC_CNT_UP"`
	ADDT_INFO_1  string `bson:"ADDT_INFO_1"`
	ADDT_INFO_2  string `bson:"ADDT_INFO_2"`
	ADDT_INFO_3  string `bson:"ADDT_INFO_3"`
}

// AROUTE : AROUTE
type AROUTE struct {
	ROUTE_ID        string `bson:"ROUTE_ID"`
	ROUTE_VER       string `bson:"ROUTE_VER"`
	CR_OPE_NO       string `bson:"CR_OPE_NO"`
	PV_OPE_NO       string `bson:"PV_OPE_NO"`
	NX_OPE_NO       string `bson:"NX_OPE_NO"`
	MODULE_ID       string `bson:"MODULE_ID"`
	MODULE_VER      string `bson:"MODULE_VER"`
	CR_OPE_ID       string `bson:"CR_OPE_ID"`
	CR_OPE_VER      string `bson:"CR_OPE_VER"`
	RWK_RST_FLG     string `bson:"RWK_RST_FLG"`
	RWK_AVL_FLG     string `bson:"RWK_AVL_FLG"`
	MPROC_FLG       string `bson:"MPROC_FLG"`
	MPROC_ID        string `bson:"MPROC_ID"`
	WIP_BANK_FLG    string `bson:"WIP_BANK_FLG"`
	DEF_WIP_BANK_ID string `bson:"DEF_WIP_BANK_ID"`
	SHP_BANK_FLG    string `bson:"SHP_BANK_FLG"`
	STAGE_ID        string `bson:"STAGE_ID"`
	READ_GLASS_FLG  string `bson:"READ_GLASS_FLG"`
	ERP_WK_CENTER   string `bson:"ERP_WK_CENTER"`
}
