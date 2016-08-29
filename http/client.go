/*
* (C) 2001-2015 gtlService Inc.
*
* gutils source code
* version: 1.0.0
* author: bobliu0909@gmail.com
* datetime: 2015-10-14
*
 */

package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
)

var DefaultClient *HttpClient = NewHttpClient(nil, 0)

type HttpClient struct {
	client *http.Client
	retry  int
}

func NewHttpClient(client *http.Client, retry int) *HttpClient {

	if client == nil {
		client = http.DefaultClient
	}

	return &HttpClient{
		client: client,
		retry:  retry,
	}
}

func (c *HttpClient) PostJson(url string, request interface{}) (io.ReadCloser, error) {

	buffer := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buffer).Encode(request); err != nil {
		return nil, err
	}

	httpResp, err := http.Post(url, "application/json", bytes.NewReader(buffer.Bytes()))
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	return httpResp.Body, nil
}

func (c *HttpClient) GetJson(url string) (io.ReadCloser, error) {

	//预留方法，后期完善
	return nil, nil
}

func (c *HttpClient) PostFile(url string, buf []byte) error {

	//预留方法，后期完善
	return nil
}

func (c *HttpClient) GetFile(fd *os.File, url string, resumable bool) error {

RETRY:
	c.retry = c.retry - 1
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ErrHttpNewRequest
	}

	if resumable { //续传
		st, err := fd.Stat()
		if err != nil {
			return ErrFileStatException
		}
		seek := st.Size()
		fd.Seek(seek, 0)
		req.Header.Add("Range", "bytes="+strconv.FormatInt(seek, 10)+"-")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		if c.retry > 0 { //网络失败重试
			goto RETRY
		}
		return ErrHttpRequestFailed
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		if _, err := io.Copy(fd, resp.Body); err != nil {
			return ErrHttpIOCopyFailed
		}
		return nil
	}
	return errors.New("http response error(" + strconv.Itoa(resp.StatusCode) + ").")
}
