// Code generated by "stringer -type=TransField"; DO NOT EDIT.

package gno

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TRANS_ROOT-0]
	_ = x[TRANS_BINARY_LEFT-1]
	_ = x[TRANS_BINARY_RIGHT-2]
	_ = x[TRANS_CALL_FUNC-3]
	_ = x[TRANS_CALL_ARG-4]
	_ = x[TRANS_INDEX_X-5]
	_ = x[TRANS_INDEX_INDEX-6]
	_ = x[TRANS_SELECTOR_X-7]
	_ = x[TRANS_SLICE_X-8]
	_ = x[TRANS_SLICE_LOW-9]
	_ = x[TRANS_SLICE_HIGH-10]
	_ = x[TRANS_SLICE_MAX-11]
	_ = x[TRANS_STAR_X-12]
	_ = x[TRANS_REF_X-13]
	_ = x[TRANS_TYPEASSERT_X-14]
	_ = x[TRANS_TYPEASSERT_TYPE-15]
	_ = x[TRANS_UNARY_X-16]
	_ = x[TRANS_COMPOSITE_TYPE-17]
	_ = x[TRANS_COMPOSITE_KEY-18]
	_ = x[TRANS_COMPOSITE_VALUE-19]
	_ = x[TRANS_FUNCLIT_TYPE-20]
	_ = x[TRANS_FUNCLIT_BODY-21]
	_ = x[TRANS_FIELDTYPE_TYPE-22]
	_ = x[TRANS_FIELDTYPE_TAG-23]
	_ = x[TRANS_ARRAYTYPE_LEN-24]
	_ = x[TRANS_ARRAYTYPE_ELT-25]
	_ = x[TRANS_SLICETYPE_ELT-26]
	_ = x[TRANS_INTERFACETYPE_METHOD-27]
	_ = x[TRANS_CHANTYPE_VALUE-28]
	_ = x[TRANS_FUNCTYPE_PARAM-29]
	_ = x[TRANS_FUNCTYPE_RESULT-30]
	_ = x[TRANS_MAPTYPE_KEY-31]
	_ = x[TRANS_MAPTYPE_VALUE-32]
	_ = x[TRANS_STRUCTTYPE_FIELD-33]
	_ = x[TRANS_MAYBENATIVETYPE_TYPE-34]
	_ = x[TRANS_ASSIGN_LHS-35]
	_ = x[TRANS_ASSIGN_RHS-36]
	_ = x[TRANS_BLOCK_BODY-37]
	_ = x[TRANS_DECL_BODY-38]
	_ = x[TRANS_DEFER_CALL-39]
	_ = x[TRANS_EXPR_X-40]
	_ = x[TRANS_FOR_INIT-41]
	_ = x[TRANS_FOR_COND-42]
	_ = x[TRANS_FOR_POST-43]
	_ = x[TRANS_FOR_BODY-44]
	_ = x[TRANS_GO_CALL-45]
	_ = x[TRANS_IF_INIT-46]
	_ = x[TRANS_IF_COND-47]
	_ = x[TRANS_IF_BODY-48]
	_ = x[TRANS_IF_ELSE-49]
	_ = x[TRANS_IF_CASE_BODY-50]
	_ = x[TRANS_INCDEC_X-51]
	_ = x[TRANS_RANGE_X-52]
	_ = x[TRANS_RANGE_KEY-53]
	_ = x[TRANS_RANGE_VALUE-54]
	_ = x[TRANS_RANGE_BODY-55]
	_ = x[TRANS_RETURN_RESULT-56]
	_ = x[TRANS_PANIC_EXCEPTION-57]
	_ = x[TRANS_SELECT_CASE-58]
	_ = x[TRANS_SELECTCASE_COMM-59]
	_ = x[TRANS_SELECTCASE_BODY-60]
	_ = x[TRANS_SEND_CHAN-61]
	_ = x[TRANS_SEND_VALUE-62]
	_ = x[TRANS_SWITCH_INIT-63]
	_ = x[TRANS_SWITCH_X-64]
	_ = x[TRANS_SWITCH_CASE-65]
	_ = x[TRANS_SWITCHCASE_CASE-66]
	_ = x[TRANS_SWITCHCASE_BODY-67]
	_ = x[TRANS_FUNC_RECV-68]
	_ = x[TRANS_FUNC_TYPE-69]
	_ = x[TRANS_FUNC_BODY-70]
	_ = x[TRANS_IMPORT_PATH-71]
	_ = x[TRANS_CONST_TYPE-72]
	_ = x[TRANS_CONST_VALUE-73]
	_ = x[TRANS_VAR_TYPE-74]
	_ = x[TRANS_VAR_VALUE-75]
	_ = x[TRANS_TYPE_TYPE-76]
	_ = x[TRANS_FILE_BODY-77]
}

const _TransField_name = "TRANS_ROOTTRANS_BINARY_LEFTTRANS_BINARY_RIGHTTRANS_CALL_FUNCTRANS_CALL_ARGTRANS_INDEX_XTRANS_INDEX_INDEXTRANS_SELECTOR_XTRANS_SLICE_XTRANS_SLICE_LOWTRANS_SLICE_HIGHTRANS_SLICE_MAXTRANS_STAR_XTRANS_REF_XTRANS_TYPEASSERT_XTRANS_TYPEASSERT_TYPETRANS_UNARY_XTRANS_COMPOSITE_TYPETRANS_COMPOSITE_KEYTRANS_COMPOSITE_VALUETRANS_FUNCLIT_TYPETRANS_FUNCLIT_BODYTRANS_FIELDTYPE_TYPETRANS_FIELDTYPE_TAGTRANS_ARRAYTYPE_LENTRANS_ARRAYTYPE_ELTTRANS_SLICETYPE_ELTTRANS_INTERFACETYPE_METHODTRANS_CHANTYPE_VALUETRANS_FUNCTYPE_PARAMTRANS_FUNCTYPE_RESULTTRANS_MAPTYPE_KEYTRANS_MAPTYPE_VALUETRANS_STRUCTTYPE_FIELDTRANS_MAYBENATIVETYPE_TYPETRANS_ASSIGN_LHSTRANS_ASSIGN_RHSTRANS_BLOCK_BODYTRANS_DECL_BODYTRANS_DEFER_CALLTRANS_EXPR_XTRANS_FOR_INITTRANS_FOR_CONDTRANS_FOR_POSTTRANS_FOR_BODYTRANS_GO_CALLTRANS_IF_INITTRANS_IF_CONDTRANS_IF_BODYTRANS_IF_ELSETRANS_IF_CASE_BODYTRANS_INCDEC_XTRANS_RANGE_XTRANS_RANGE_KEYTRANS_RANGE_VALUETRANS_RANGE_BODYTRANS_RETURN_RESULTTRANS_PANIC_EXCEPTIONTRANS_SELECT_CASETRANS_SELECTCASE_COMMTRANS_SELECTCASE_BODYTRANS_SEND_CHANTRANS_SEND_VALUETRANS_SWITCH_INITTRANS_SWITCH_XTRANS_SWITCH_CASETRANS_SWITCHCASE_CASETRANS_SWITCHCASE_BODYTRANS_FUNC_RECVTRANS_FUNC_TYPETRANS_FUNC_BODYTRANS_IMPORT_PATHTRANS_CONST_TYPETRANS_CONST_VALUETRANS_VAR_TYPETRANS_VAR_VALUETRANS_TYPE_TYPETRANS_FILE_BODY"

var _TransField_index = [...]uint16{0, 10, 27, 45, 60, 74, 87, 104, 120, 133, 148, 164, 179, 191, 202, 220, 241, 254, 274, 293, 314, 332, 350, 370, 389, 408, 427, 446, 472, 492, 512, 533, 550, 569, 591, 617, 633, 649, 665, 680, 696, 708, 722, 736, 750, 764, 777, 790, 803, 816, 829, 847, 861, 874, 889, 906, 922, 941, 962, 979, 1000, 1021, 1036, 1052, 1069, 1083, 1100, 1121, 1142, 1157, 1172, 1187, 1204, 1220, 1237, 1251, 1266, 1281, 1296}

func (i TransField) String() string {
	if i >= TransField(len(_TransField_index)-1) {
		return "TransField(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TransField_name[_TransField_index[i]:_TransField_index[i+1]]
}
