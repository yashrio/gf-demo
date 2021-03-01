package green

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var DefaultClient = new(defaultClient)

type defaultClient struct {
	Profile profile
}

type profile struct {
	AccessKeyId     string
	AccessKeySecret string
}

func init() {
	fmt.Println("初始化client")
	fmt.Println(g.Cfg().GetString("aliyun.accessKeyId"))
	DefaultClient.Profile = profile{
		AccessKeyId:     g.Cfg().GetString("aliyun.accessKeyId"),
		AccessKeySecret: g.Cfg().GetString("aliyun.accessKeySecret"),
	}
}

func (defaultClient *defaultClient) GetResponse(path string, clinetInfo ClinetInfo, bizData BizData) string {
	clientInfoJson, _ := json.Marshal(clinetInfo)
	bizDataJson, _ := json.Marshal(bizData)

	client := &http.Client{}
	req, err := http.NewRequest(method, host+path+"?clientInfo="+url.QueryEscape(string(clientInfoJson)), strings.NewReader(string(bizDataJson)))

	if err != nil {
		// handle error
		return ErrorResult(err)
	} else {
		addRequestHeader(string(bizDataJson), req, string(clientInfoJson), path, defaultClient.Profile.AccessKeyId, defaultClient.Profile.AccessKeySecret)

		response, _ := client.Do(req)

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			// handle error
			return ErrorResult(err)
		} else {
			return string(body)
		}
	}
}

func (defaultClient *defaultClient) GetScanResult(bizData BizData) *ScanResponse {
	scanResponse := new(ScanResponse)
	clinetInfo := ClinetInfo{
		Ip: "127.0.0.1",
	}
	response := defaultClient.GetResponse("/green/image/scan", clinetInfo, bizData)
	if j, err := gjson.DecodeToJson(response); err == nil {
		j.Struct(scanResponse)
	}
	return scanResponse
}

type IAliYunClient interface {
	GetResponse(path string, clinetInfo ClinetInfo, bizData BizData) string
	GetScanResult(bizData BizData) *ScanResponse
}
