package util

import (
	"bytes"
	"github.com/gogf/gf/os/glog"
	"io/ioutil"
	"regexp"
	"strings"
)

func ProcessExists(processName string) (bool, error) {
	result := false
	fileInfos, err := ioutil.ReadDir("/proc")
	if err != nil {
		return false, err
	}
	for _, info := range fileInfos {
		name := info.Name()
		matched, err := regexp.MatchString("[0-9]+", name)
		if err != nil {
			return false, err
		}
		if !matched {
			continue
		}
		cmdLine, err := parseCmdLine("/proc/" + info.Name() + "/cmdline")
		if err != nil {
			glog.Error("read cmd line failed ", err)
			// the proc may end during the loop, ignore it
			continue
		}
		if strings.Contains(cmdLine, processName) {
			result = true
		}
	}
	return result, err
}

func parseCmdLine(path string) (string, error) {
	cmdData, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	if len(cmdData) < 1 {
		return "", nil
	}

	split := strings.Split(string(bytes.TrimRight(cmdData, string("\x00"))), string(byte(0)))
	return strings.Join(split, " "), nil
}
