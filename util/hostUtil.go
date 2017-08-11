// hostUtil
package util

const (
	HTTP_LOCAL_HOST  string = "http://localhost:8081"
	HTTP_REMOTE_HOST string = "http://gzzhizhuo.com:8081"

	HTTP_APPLICATION string = "application/json;charset=utf-8"
)

var (
	local_test string
)

func GetFactoryUrl() string {
	if local_test == "" {
		local_test = Param("local_test")
		PrintLog("GetFactoryUrl, local_test", local_test)
	}
	if local_test == "0" {
		return HTTP_LOCAL_HOST + "/factory"
	} else {
		return HTTP_REMOTE_HOST + "/factory"
	}
}

func GetHistoryUrl() string {
	if local_test == "" {
		local_test = Param("local_test")
		PrintLog("GetHistoryUrl, local_test", local_test)
	}
	if local_test == "0" {
		return HTTP_LOCAL_HOST + "/history"
	} else {
		return HTTP_REMOTE_HOST + "/history"
	}
}

func GetForceUrl() string {
	if local_test == "" {
		local_test = Param("local_test")
		PrintLog("GetForceUrl, local_test", local_test)
	}
	if local_test == "0" {
		return HTTP_LOCAL_HOST + "/force"
	} else {
		return HTTP_REMOTE_HOST + "/force"
	}
}

func GetOrderUrl() string {
	if local_test == "" {
		local_test = Param("local_test")
		PrintLog("GetForceUrl, local_test", local_test)
	}
	if local_test == "0" {
		return HTTP_LOCAL_HOST + "/order"
	} else {
		return HTTP_REMOTE_HOST + "/order"
	}
}

func GetFinishInfoUrl() string {
	if local_test == "" {
		local_test = Param("local_test")
		PrintLog("GetForceUrl, local_test", local_test)
	}
	if local_test == "0" {
		return HTTP_LOCAL_HOST + "/finish_info"
	} else {
		return HTTP_REMOTE_HOST + "/finish_info"
	}
}
