package svc

import (
	"activity/answer/api/internal/config"
	"activity/answer/api/internal/types"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	SIGN_VERSION = "2021-04-15"
)

func Signature(method string, path string, query map[string][]string, body []byte, secretKey string) (string, error) {
	ss := make([]string, 0)
	hasSecretId := false
	hasTimeStamp := false
	hasNonce := false
	hasVersion := false
	// hasSignature := false
	for k, v := range query {
		for i := 0; i < len(v); i++ {
			if k == "Signature" {
				// hasSignature = true
				continue
			} else if k == "SecretId" {
				hasSecretId = true
			} else if k == "Timestamp" {
				hasTimeStamp = true
			} else if k == "Nonce" {
				hasNonce = true
			} else if k == "Version" {
				hasVersion = true
			}
			ss = append(ss, fmt.Sprintf("%s=%s", k, v[i]))
		}
	}
	if !hasSecretId {
		return "", errors.New("miss param SecretId")
	}
	if !hasTimeStamp {
		return "", errors.New("miss param Timestamp")
	}
	if !hasNonce {
		return "", errors.New("miss param Nonce")
	}
	if !hasVersion {
		return "", errors.New("miss param Version")
	}
	sort.Strings(ss)
	sss := strings.Join(ss, "&")
	str := ""
	if len(body) > 0 {
		bodyMD5 := fmt.Sprintf("%x", md5.Sum(body))
		str = fmt.Sprintf("%s%s?%s%s", method, path, sss, bodyMD5)
	} else {
		str = fmt.Sprintf("%s%s?%s", method, path, sss)
	}
	key := []byte(secretKey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(str))
	enc := mac.Sum(nil)
	res := base64.StdEncoding.EncodeToString(enc)
	log.Printf("DEBUG,signature,str:%s,key:%s,sign:%s", str, secretKey, res)
	return res, nil
}

//⽣成query的⽅法：
func MakeQuery(method string, path string, secretKey string, query map[string][]string, body []byte) (string, error) {

	sign, err := Signature(method, path, query, body, secretKey)
	if err != nil {
		return "", err
	}
	query["Signature"] = []string{sign}
	arr := []string{}
	for k, v := range query {
		for i := 0; i < len(v); i++ {
			arr = append(arr, fmt.Sprintf("%s=%s", k, url.QueryEscape(v[i])))
		}
	}
	return strings.Join(arr, "&"), nil
}

//SendGameEmail 发送邮件
func SendGameEmail(apiConfig config.GameApiConfig, params types.SendMailParams) (int, error) {
	timestamp := int(time.Now().Unix())
	nonce := strconv.FormatInt(time.Now().Unix(), 10)
	query := make(map[string][]string, 0)
	query["SecretId"] = []string{apiConfig.SecretId}
	query["Timestamp"] = []string{strconv.Itoa(timestamp)}
	query["Nonce"] = []string{nonce}
	query["Version"] = []string{SIGN_VERSION}
	body, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}
	logx.Info("send mail body ", string(body))
	urlPath := "/api/egmt/v1/gmt_send_mail"
	queryString, err := MakeQuery("POST", urlPath, apiConfig.SecretKey, query, body)
	if err != nil {
		return 0, err
	}
	gameApiUrl := apiConfig.GameUrl + urlPath + "?" + queryString
	logx.Info("send mail url ", gameApiUrl)
	payload := strings.NewReader(string(body))
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, gameApiUrl, payload)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	bodyResp, err := ioutil.ReadAll(res.Body)
	logx.Info("send mail resq ", string(bodyResp))
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, errors.New("http code error")
	}
	var result types.SendMailResp
	if err := json.Unmarshal(bodyResp, &result); err != nil {
		return 0, err
	}
	if result.Result != 0 {
		return 0, errors.New(result.Msg)
	}
	if len(result.Data) < 1 {
		return 0, errors.New("send mail data len error")
	}
	for _, v := range result.Data {
		if v.Ret != 0 {
			return 0, errors.New("send mail response error")
		}
	}
	logx.Info("send mail resp ", string(bodyResp))
	return 0, nil
}

func SendCode(sendCodeConfig config.SendCodeConfig, request types.SendCodeParams) (int, error) {
	values := url.Values{}
	values.Add("type", request.Type)
	values.Add("uid", request.Uid)
	values.Add("code", request.Code)
	signStr := values.Encode() + sendCodeConfig.ApiSecret
	sign := fmt.Sprintf("%x", md5.Sum([]byte(signStr)))
	values.Add("sign", sign)
	body := values.Encode()
	logx.Info("send code body ", body)
	payload := strings.NewReader(body)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, sendCodeConfig.GameUrl, payload)
	if err != nil {
		return 0, err
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	bodyResp, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return 0, errors.New("http code error")
	}
	logx.Info("send code resq ", string(bodyResp))
	if err != nil {
		return 0, err
	}
	var sendCodeData types.SendCodeRes
	if err := json.Unmarshal(bodyResp, &sendCodeData); err != nil {
		return 0, err
	}
	
	if sendCodeData.Result != "success" {
		return 0, errors.New("response result error")
	}

	return 0, nil
}
