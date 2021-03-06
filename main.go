package main

/*
#cgo LDFLAGS: -lTDBAPI -lstdc++
#include "include/TDBAPI.h"
#include "include/TDBAPIStruct.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func String2char(str string, des uintptr, sizeOf uintptr){
	bytes := []byte(str)
	for i:=0; i<len(bytes); i++{
		unit := (*C.char)(unsafe.Pointer(des))
		*unit = C.char(bytes[i])
		des += sizeOf
	}
}


func Char2byte(des uintptr, sizeOf uintptr, leng int)[256]byte{
	var bytes [256]byte
	for i:=0; i < leng; i++ {
		unit := (*C.char)(unsafe.Pointer(des))
		bytes[i] = byte(*unit)
		des += sizeOf
	}
	return bytes
}

func main(){
	var hTdb C.THANDLE = nil

	var settings C.OPEN_SETTINGS

	//================================================
	String2char("114.80.154.34",uintptr(unsafe.Pointer(&settings.szIP)),unsafe.Sizeof(settings.szIP[0]))
	String2char("6261",uintptr(unsafe.Pointer(&settings.szPort)),unsafe.Sizeof(settings.szPort[0]))
	String2char("TD3446699001",uintptr(unsafe.Pointer(&settings.szUser)),unsafe.Sizeof(settings.szUser[0]))
	String2char("43449360",uintptr(unsafe.Pointer(&settings.szPassword)),unsafe.Sizeof(settings.szPassword[0]))
	//================================================
	settings.nRetryCount = 15
	settings.nRetryGap = 1
	settings.nTimeOutVal = 1

	//proxy
/*	var proxy_setting C.TDB_PROXY_SETTING

	proxy_setting.nProxyType = C.TDB_PROXY_HTTP11
	//================================================
	string2char("10.100.3.42",uintptr(unsafe.Pointer(&proxy_setting.szProxyHostIp)),unsafe.Sizeof(proxy_setting.szProxyHostIp[0]))
	string2char("12345",uintptr(unsafe.Pointer(&proxy_setting.szProxyPort)),unsafe.Sizeof(proxy_setting.szProxyPort[0]))
	string2char("1",uintptr(unsafe.Pointer(&proxy_setting.szProxyUser)),unsafe.Sizeof(proxy_setting.szProxyUser[0]))
	string2char("1",uintptr(unsafe.Pointer(&proxy_setting.szProxyPwd)),unsafe.Sizeof(proxy_setting.szProxyPwd[0]))
	//================================================
	*/

	var LoginRes C.TDBDefine_ResLogin
	//TDB_OpenProxy
	//hTdb = C.TDB_OpenProxy(&settings, &proxy_setting, &LoginRes)

	hTdb = C.TDB_Open(&settings, &LoginRes)
	if hTdb == nil {
		fmt.Println("连接失败！")
		return
	}

	//TDB_GetCOdeInfo
	var pCode *C.TDBDefine_Code
	pCode = C.TDB_GetCodeInfo(hTdb, C.CString("000001.SZ"), C.CString("SZ-2-0"))
	fmt.Printf("交易所代码 chWindCode:%s \n", Char2byte(uintptr(unsafe.Pointer(&pCode.chCode)),unsafe.Sizeof(pCode.chCode[0]),len(pCode.chCode)))



	var pCount C.int = 0
	C.TDB_GetCodeTable(hTdb,C.CString("SZ"),&pCode,&pCount);
	tmpPtr := uintptr(unsafe.Pointer(pCode))
	sizeOf := unsafe.Sizeof(*pCode)
	if pCount!=0 && pCode!=nil{
		for i := 0; i < 2; i++{
		pC := (*C.TDBDefine_Code)(unsafe.Pointer(tmpPtr))
		fmt.Println("-------------code table ----------------------------");
		fmt.Printf("chWindCode:%s \n", pC.chCode);
		fmt.Printf("chWindCode:%s \n", pC.chMarket);
		fmt.Printf("chWindCode:%s \n", pC.chCNName);
		fmt.Printf("chWindCode:%s \n", pC.chENName);
		fmt.Printf("chWindCode:%s \n", pC.nType);
		tmpPtr += sizeOf
		}
	}
	GetKData(hTdb, "600715.SH", "SH-2-0", 20151126, 20151126, C.CYC_MINUTE, 0, 0, 1);//autocomplete k-minute
	GetTickData(hTdb, "000001.sz", "SZ-2-0", 20150910);//tick
	GetTransaction(hTdb, "000001.sz", "SZ-2-0", 20150910);//Transaction
	GetOrder(hTdb, "000001.sz", "SZ-2-0", 20150910);//Order
	GetOrderQueue(hTdb, "000001.sz", "SZ-2-0", 20150910);//OrderQueue
	UseEZFFormula(hTdb);//test for formula

}
