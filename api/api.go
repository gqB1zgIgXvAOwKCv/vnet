package api

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rc452860/vnet/model"
	"github.com/rc452860/vnet/utils/langx"
	"github.com/rc452860/vnet/utils/stringx"
	"github.com/tidwall/gjson"
	"gopkg.in/resty.v1"
	"net/http"
	"strconv"
	"time"
)

func init() {
	resty.SetTimeout(3 * time.Second)
}

// implement for vnet api get request
func get(url string,header map[string]string)(result string,err error){
	r, err := resty.R().SetHeaders(header).Get(url)
	if err != nil {
		return "", errors.Wrap(err, "get request error")
	}
	if r.StatusCode() != http.StatusOK{
		return "",errors.New(fmt.Sprintf("get request status: %d body: %s",r.StatusCode(),string(r.Body())))
	}
	responseJson := stringx.BUnicodeToUtf8(r.Body())
	value := gjson.Get(responseJson, "data").String()
	if value == ""{
		return "",errors.New("get data not found: " + responseJson)
	}
	return value,nil
}

func post(url,param string,header map[string]string)(result string,err error){
	header["Content-Type"] = "application/json"
	r,err := resty.R().SetHeaders(header).SetBody(param).Post(url)
	if err != nil {
		return "", errors.Wrap(err, "get request error")
	}
	if r.StatusCode() != http.StatusOK{
		return "",errors.New(fmt.Sprintf("get request status: %d body: %s",r.StatusCode(),string(r.Body())))
	}
	responseJson := stringx.BUnicodeToUtf8(r.Body())
	return responseJson,nil
}


/*------------------------------ code below is webapi implement ------------------------------*/

// GetNodeInfo
func GetNodeInfo(nodeID int, key string) model.NodeInfo {
	value,err := get(fmt.Sprintf("http://localhost/api/node/%s",strconv.Itoa(nodeID)),map[string]string{
		"key":       key,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	})
	if err != nil{
		panic(err)
	}
	result := model.NodeInfo{}
	err = json.Unmarshal([]byte(value), &result)
	if err != nil {
		panic(err)
	}
	return result
}

// GetUserList
func GetUserList(nodeID int, key string) []model.UserInfo {
	value,err := get(fmt.Sprintf("http://localhost/api/userList/%s", strconv.Itoa(nodeID)),map[string]string{
		"key":       key,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	})
	if err != nil{
		panic(err)
	}
	result := []model.UserInfo{}
	err = json.Unmarshal([]byte(value), &result)
	if err != nil {
		panic(err)
	}
	return result
}

func PostAllUserTraffic(allUserTraffic []model.UserTraffic,nodeID int, key string) {
	value, err := post(fmt.Sprintf("http://localhost/api/userTraffic/%s", strconv.Itoa(nodeID)),
	string(langx.Must(func() (interface{}, error) {
		return json.Marshal(allUserTraffic)
	}).([]byte)),
	map[string]string{
		"key":       key,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	})

	if err != nil{
		panic(err)
	}
	if gjson.Get(value,"status").String() != "success" {
		panic(stringx.UnicodeToUtf8(gjson.Get(value,"message").String()))
	}
}

func PostNodeOnline(nodeOnline []model.NodeOnline,nodeID int, key string){
	value, err := post(fmt.Sprintf("http://localhost/api/nodeOnline/%s", strconv.Itoa(nodeID)),
		string(langx.Must(func() (interface{}, error) {
			return json.Marshal(nodeOnline)
		}).([]byte)),
		map[string]string{
			"key":       key,
			"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		})

	if err != nil{
		panic(err)
	}
	if gjson.Get(value,"status").String() != "success" {
		panic(stringx.UnicodeToUtf8(gjson.Get(value,"message").String()))
	}
}

func PostNodeStatus(status model.NodeStatus,nodeID int, key string){
	value, err := post(fmt.Sprintf("http://localhost/api/nodeStatus/%s", strconv.Itoa(nodeID)),
		string(langx.Must(func() (interface{}, error) {
			return json.Marshal(status)
		}).([]byte)),
		map[string]string{
			"key":       key,
			"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		})

	if err != nil{
		panic(err)
	}
	if gjson.Get(value,"status").String() != "success" {
		panic(stringx.UnicodeToUtf8(gjson.Get(value,"message").String()))
	}
}