/**
 * @Author: vincent
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/11/22 17:14
 */

package ssh

import (
	"context"
	"fmt"
	"go-examples/course/handwriting-web-inf/code_28/framework"
	"go-examples/course/handwriting-web-inf/code_28/framework/contract"
	"sync"

	"golang.org/x/crypto/ssh"
)

// HadeSSH 代表hade框架的ssh实现
type HadeSSH struct {
	container framework.Container    // 服务容器
	clients   map[string]*ssh.Client // key为uniqKey，value为ssh.Client(连接池)

	lock *sync.RWMutex
}

// NewHadeSSH 代表实例化Client
func NewHadeSSH(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	clients := make(map[string]*ssh.Client)
	lock := &sync.RWMutex{}

	return &HadeSSH{
		container: container,
		clients:   clients,
		lock:      lock,
	}, nil
}

// GetClient 获取Client实例
func (app *HadeSSH) GetClient(option ...contract.SSHOption) (*ssh.Client, error) {
	logService := app.container.MustMake(contract.LogKey).(contract.Log)
	// 读取默认配置
	config := GetBaseConfig(app.container)
	fmt.Printf("config is %+v\n", config)
	// option对opt进行修改
	for _, opt := range option {
		if err := opt(app.container, config); err != nil {
			logService.Error(context.Background(), "ssh opt error", map[string]interface{}{
				"err": err,
			})
			return nil, err
		}
		fmt.Printf("*** config is %+v\n", config)
	}

	// key
	key := config.UniqKey()

	// 判断是否已经实例化了
	app.lock.RLock()
	if db, ok := app.clients[key]; ok {
		// 已经实例化，马上解锁
		app.lock.RUnlock()
		return db, nil
	}
	app.lock.RUnlock()

	// 没有实例化，那么就要进行实例化操作
	app.lock.Lock()
	defer app.lock.Unlock()

	// 实例化
	addr := config.Host + ":" + config.Port
	client, err := ssh.Dial(config.Network, addr, config.ClientConfig)
	if err != nil {
		logService.Error(context.Background(), "ssh dial error", map[string]interface{}{
			"err":  err,
			"addr": addr,
		})
	}

	// 挂载到map中，结束配置
	app.clients[key] = client

	return client, nil
}
