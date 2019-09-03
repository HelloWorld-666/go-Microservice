package common

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net"
	"strings"
)

func FormatConversion(x interface{}) string {
	jsonRsp, err := json.Marshal(x)
	if err != nil {
		return err.Error()
	}
	strRsp := string(jsonRsp)
	return strRsp
}

func FormatBool(value bool) string {
	if value {
		return "1"
	}
	return "0"
}

func VerifySnKey(snKey, sn, salt string) bool {
	return strings.ToUpper(GetMd5String(strings.ToUpper(GetMd5String(sn))+salt)) == snKey
}

func GetMd5String(text string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(text))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func GetMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "error"
	}
	var mac string
	for _, inter := range interfaces {
		mac = inter.HardwareAddr.String()
		if len(mac) > 0 { // mac can be empty
			break
		}
	}
	return mac
}

func GetPages(total, pageSize int32) (pages int32) {
	if pageSize <= 0 {
		pages = 1
		return
	} else {
		pages = total / pageSize
		if total%pageSize > 0 {
			pages++
		}
		return
	}
}
