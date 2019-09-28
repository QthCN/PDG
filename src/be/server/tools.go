package server

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ResMsgS struct {
	Msg string `json:"msg"`
}

func ParseResMsg(r string) string {
	data := ResMsgS{}
	reqContent := []byte(r)
	err := json.Unmarshal(reqContent, &data)
	if err != nil {
		log.WithFields(log.Fields{
			"target_str": r,
			"err":        err.Error(),
		}).Error("内容解析失败")
		return "JSON格式错误"
	}
	return data.Msg
}

func ResSuccessMsg(res http.ResponseWriter, code int, msg string) {
	resMsg_ := ResMsgS{Msg: msg}
	result := ""
	b, err := json.Marshal(resMsg_)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("生成JSON串失败")
		result = `{"msg": "内部错误"}`
		code = 500

	} else {
		result = string(b)
	}
	res.WriteHeader(code)
	res.Write([]byte(result))
}

func ResMsg(res http.ResponseWriter, code int, msg string) {
	resMsg_ := ResMsgS{Msg: msg}
	result := ""
	if code != 200 {
		b, err := json.Marshal(resMsg_)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err.Error(),
			}).Error("生成JSON串失败")
			result = `{"msg": "内部错误"}`
			code = 500

		} else {
			result = string(b)
		}
	} else {
		result = msg
	}

	res.WriteHeader(code)
	res.Write([]byte(result))
}

func ResInvalidRequestBody(res http.ResponseWriter) {
	ResMsg(res, 400, "请求报文格式错误")
}

func ResInvalidToken(res http.ResponseWriter) {
	ResMsg(res, 400, "认证Token或获取Token失败")
}

func ResNotAuth(res http.ResponseWriter) {
	ResMsg(res, 400, "未授权")
}

func ResInternalError(res http.ResponseWriter) {
	ResMsg(res, 500, "内部错误")
}

func CheckReqBody(res http.ResponseWriter, req *http.Request, body interface{}) error {
	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.WithFields(log.Fields{}).Error("请求报文解析失败")
		return err
	}
	err = json.Unmarshal(reqContent, body)
	if err != nil {
		log.WithFields(log.Fields{
			"body": string(reqContent),
			"err":  err.Error(),
		}).Error("请求报文解析失败")
		return err
	}
	return nil
}

func ParseJsonStr(str string, body interface{}) error {
	data := []byte(str)
	err := json.Unmarshal(data, body)
	if err != nil {
		log.WithFields(log.Fields{
			"body": str,
			"err":  err.Error(),
		}).Error("报文解析失败")
		return err
	} else {
		return nil
	}
}

/*
*
* 返回n位数的随机字符串
* From: http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
*
 */

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
