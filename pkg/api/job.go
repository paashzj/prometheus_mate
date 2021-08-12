package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"net/http"
	"prometheus_mate/pkg/model"
	"prometheus_mate/pkg/prom"
	"prometheus_mate/pkg/service"
)

func AddJob(r *ghttp.Request) {
	var req model.CreateJobReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	}
	resp, err := service.AddJob(req)
	if err != nil {
		glog.Info("err occurred ", err)
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	}
	prom.ReloadChannel <- struct{}{}
	r.Response.WriteHeader(http.StatusCreated)
	err = r.Response.WriteJsonExit(&resp)
	if err != nil {
		return
	}
}

func DelJob(r *ghttp.Request) {
	job := gconv.String(r.Get("job"))
	if job == "" {
		glog.Info("job name is empty, do nothing")
		r.Response.WriteStatusExit(http.StatusBadRequest, "job name empty")
	}
	err := service.DelJob(job)
	if err != nil {
		glog.Error("err occurred ", err)
		r.Response.WriteStatusExit(http.StatusInternalServerError)
	}
	prom.ReloadChannel <- struct{}{}
	r.Response.WriteStatusExit(http.StatusNoContent)
}
