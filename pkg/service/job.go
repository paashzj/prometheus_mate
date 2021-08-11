package service

import (
	"encoding/json"
	"github.com/gogf/gf/os/gfile"
	"os"
	"path/filepath"
	"prometheus_mate/pkg/model"
	"prometheus_mate/pkg/path"
)

func AddJob(req model.CreateJobReq) (resp model.CreateJobResp, err error) {
	resp = model.CreateJobResp{}
	resp.Job = req.Job
	file, err := gfile.OpenWithFlag(filepath.FromSlash(path.PromJobs+"/job-"+req.Job+"-v1.json"), os.O_RDWR|os.O_CREATE)
	if err != nil {
		return
	}
	defer file.Close()
	bytes, err := json.Marshal(req)
	if err != nil {
		return
	}
	_, err = file.Write(bytes)
	if err != nil {
		return
	}
	return
}
