package service

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

func CallApi(url string, method string, reqObj any, respObj any) error {
	logrus.Info(fmt.Sprintf("--Start %s %s--", "baseService.CallApi ", url))
	reqJson, err := json.Marshal(reqObj)
	payload := strings.NewReader(string(reqJson))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		logrus.Error(err)
		return err
	} else {
		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			logrus.Error(err)
			return err
		} else {
			defer func() {
				err = res.Body.Close()
			}()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				logrus.Error(err)
				return err
			} else {
				err = json.Unmarshal(body, &respObj)
				if err != nil {
					logrus.Error(err)
					return err
				}
			}
		}

	}
	logrus.Info(fmt.Sprintf("--Finish %s %s--", "baseService.CallApi ", url))
	return err
}
