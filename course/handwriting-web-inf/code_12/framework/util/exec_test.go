/**
 * @Author: vincent
 * @Description:
 * @File:  exec_test
 * @Version: 1.0.0
 * @Date: 2021/10/29 13:46
 */

package util

import (
	"os"
	"testing"
)

func TestGetExecDirectory(t *testing.T) {
	path := GetExecDirectory()
	if path == "" {
		t.Error("path is empty")
	}
	t.Log("path is " + path)
}

func TestCheckProcessExist(t *testing.T) {
	pid := os.Getpid()
	if false == CheckProcessExist(pid) {
		t.Errorf("pid(%d) is existed, but not check out", pid)
	}
	t.Logf("pid(%d) is checked", pid)
}
