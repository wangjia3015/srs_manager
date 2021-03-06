package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	URL_STREAMS_PATH   = "api/v1/streams"
	URL_CLIENTS_PATH   = "api/v1/clients"
	URL_SUMMARIES_PATH = "api/v1/summaries"

	HTTP_GET    = "GET"
	HTTP_PUT    = "PUT"
	HTTP_DELETE = "DELETE"
)

func sendRequest(method, url string) (respBody []byte, err error) {
	var (
		req *http.Request
	)
	client := http.DefaultClient
	if req, err = http.NewRequest(method, url, nil); err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.ContentLength != 0 {
		defer resp.Body.Close()
		respBody, _ = ioutil.ReadAll(resp.Body)
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unavali code :%v response :%v", resp.StatusCode, string(respBody))
	}

	return
}

type KbpsInfo struct {
	Recv30s int `json:"recv_30s"`
	Send30s int `json:"send_30s"`
}

type Publisher struct {
	Active bool `json:active` // 是否工作
	CID    int  `json:cid`    // publisher ID
}

type Stream struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	VHost     int       `json:"vhost"`
	AppName   string    `json:"app"`
	LiveMs    int64     `json:"live_ms"`
	ClientNum int       `json:"clients"`
	SendBytes int64     `json:"send_bytes"`
	RecvBytes int64     `json:"recv_bytes"`
	Kbps      KbpsInfo  `json:"kbps"`
	Publish   Publisher `json:"publish"`
}

type RspStream struct {
	Code     int      `json:"code"`
	ServerID int      `json:"server"`
	Streams  []Stream `json:"streams"`
}

func GetStreams(host string) (stream RspStream, err error) {
	var body []byte
	url := fmt.Sprintf("http://%s/%s", host, URL_STREAMS_PATH)
	if body, err = sendRequest(HTTP_GET, url); err != nil {
		return
	}
	err = json.Unmarshal(body, &stream)

	return
}

type RspBase struct {
	Code int `json:"code"`
}

func KickOffClient(host string, clientID int) (rsp RspBase, err error) {
	var body []byte
	url := fmt.Sprintf("http://%s/%s/%d", host, URL_CLIENTS_PATH, clientID)
	if body, err = sendRequest(HTTP_DELETE, url); err == nil {
		return
	}
	err = json.Unmarshal(body, &rsp)
	return
}

type SelfInfo struct {
	Version    string  `json:"version"`
	PID        int64   `json:"pid"`
	PPID       int64   `json:"ppid"`
	Argv       string  `json:"argv"`
	Cwd        string  `json:"cwd"`
	Mem        int64   `json:"mem_kbyte"`
	MemPercent float64 `json:"mem_percent"`
	CPUPercent float64 `json:"cpu_percent"`
	SrsUptime  float64 `json:"srs_uptime"`
}

type SystemInfo struct {
	CPUPercent      float64 `json:"cpu_percent"`
	DiskReadKBps    int64   `json:"disk_read_KBps"`
	DiskWriteKBps   int64   `json:"disk_write_KBps"`
	DiskBusyPercent float64 `json:"disk_busy_percent"`
	MemRam          int64   `json:"mem_ram_kbyte"`
	MemRamPercent   float64 `json:"mem_ram_percent"`
	MemSwap         int64   `json:"mem_swap_kbyte"`
	MemSwapPercent  float64 `json:"mem_swap_percent"`
	CPUNum          int     `json:"cpus"`
	CPUOnline       int     `json:"cpus_online"`
	Uptime          float64 `json:"uptime"`
	IldeTime        float64 `json:"ilde_time"`
	Load1m          float64 `json:"load_1m"`
	Load5m          float64 `json:"load_5m"`
	Load15m         float64 `json:"load_15m"`
	NetSampleTime   int64   `json:"net_sample_time"`
	NetRecv         int64   `json:"net_recv_bytes"`
	NetSend         int64   `json:"net_send_bytes"`
	NetRecvi        int64   `json:"net_recvi_bytes"`
	NetSendi        int64   `json:"net_sendi_bytes"`
	SrsSampleTime   int64   `json:"srs_sample_time"`
	SrsRecv         int64   `json:"srs_recv_bytes"`
	SrsSend         int64   `json:"srs_send_bytes"`
	ConnSys         int     `json:"conn_sys"`
	ConnSysET       int     `json:"conn_sys_et"`
	COnnSysTW       int     `json:"conn_sys_tw"`
	ConnSysUdp      int     `json:"conn_sys_udp"`
	ConnSrs         int     `json:"conn_srs"`
}

type SummaryData struct {
	Self       SelfInfo   `json:"self"`
	Sys        SystemInfo `json:"system"`
	UpdateTime int64
}

type RspSummary struct {
	RspBase
	Data SummaryData `json:"data"`
}

func GetSummaries(host string) (info *RspSummary, err error) {
	var body []byte
	url := fmt.Sprintf("http://%s/%s", host, URL_SUMMARIES_PATH)
	if body, err = sendRequest(HTTP_GET, url); err != nil {
		return
	}
	err = json.Unmarshal(body, &info)

	return
}
