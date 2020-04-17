package v2

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// TKE implements TKE service
type TKE struct {
	SecretID  string `json:"secretID"`
	SecretKey string `json:"secretKey"`
}

// New returns a new TKE service
func New(secretID, secretKey string) *TKE {
	return &TKE{
		SecretID:  secretID,
		SecretKey: secretKey,
	}
}

// UpdateImage updates image for given TKE cluster
func (t *TKE) UpdateImage(image string, namespace string, svcName string, region string, clusterID string) error {
	method := "GET"
	params := make(map[string]string)
	params["SecretId"] = t.SecretID
	params["Action"] = "ModifyClusterServiceImage"
	params["Region"] = region

	nowUnix := time.Now().Unix()
	nonce := rand.Intn(100000)
	params["Timestamp"] = strconv.FormatInt(nowUnix, 10)
	params["Nonce"] = strconv.Itoa(nonce)
	params["SignatureMethod"] = "HmacSHA256"
	params["clusterId"] = clusterID
	params["serviceName"] = svcName
	params["image"] = image
	params["namespace"] = namespace

	params["Signature"] = sign(t.SecretKey, method, params)

	req, err := request(params, method)
	if err != nil {
		return err
	}

	b, err := doRequest(req)
	if err != nil {
		return err
	}

	type codeResp struct {
		Code int    `json:"code"`
		Msg  string `json:"message"`
	}

	var r codeResp
	err = json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	if r.Code == 0 {
		return nil
	}
	return fmt.Errorf("Code: %d Message: %s", r.Code, r.Msg)
}

// Service implements k8s service.
type Service struct {
	Name            string `json:"serviceName"`
	Desc            string `json:"serviceDesc"`
	ExternalIP      string `json:"externalIp"`
	CreatedAt       string `json:"createdAt"`
	CurrentReplicas int    `json:"currentReplicas"`
	DesiredReplicas int    `json:"desiredReplicas"`
	LBID            string `json:"lbId"`
	LBStatus        string `json:"lbStatus"`
	Status          string `json:"status"`
	IP              string `json:"serviceIp"`
	Namespace       string `json:"namespace"`
	RegionID        int    `json:"regionId"`
	SubnetID        string `json:"subnetId"`
	UnHubID         string `json:"unHubId"`
}

type svcResp struct {
	Svc Service `json:"service"`
}

// GetServiceInfo returns given service status.
func (t *TKE) GetServiceInfo(namespace, svcName, region, clusterID string) (Service, error) {
	var svc Service
	method := "GET"
	params := make(map[string]string)
	params["SecretId"] = t.SecretID
	params["Action"] = "DescribeClusterServiceInfo"
	params["Region"] = region

	nowUnix := time.Now().Unix()
	nonce := rand.Intn(100000)
	params["Timestamp"] = strconv.FormatInt(nowUnix, 10)
	params["Nonce"] = strconv.Itoa(nonce)
	params["SignatureMethod"] = "HmacSHA256"
	params["clusterId"] = clusterID
	params["serviceName"] = svcName
	params["namespace"] = namespace

	params["Signature"] = sign(t.SecretKey, method, params)

	req, err := request(params, method)
	if err != nil {
		return svc, err
	}

	b, err := doRequest(req)
	if err != nil {
		return svc, err
	}

	type codeResp struct {
		Code int     `json:"code"`
		Msg  string  `json:"message"`
		Data svcResp `json:"data"`
	}

	var r codeResp
	err = json.Unmarshal(b, &r)
	if err != nil {
		return svc, err
	}
	if r.Code == 0 {
		return r.Data.Svc, nil
	}
	return svc, fmt.Errorf("Code: %d Message: %s", r.Code, r.Msg)
}
