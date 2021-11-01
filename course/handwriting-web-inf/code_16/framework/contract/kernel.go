/**
 * @Author: vincent
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2021/10/29 20:30
 */

package contract

import "go-examples/course/handwriting-web-inf/code_16/framework/gin"

const KernelKey = "hade:kernel"

// Kernel 接口提供框架核心的结构
type Kernel interface {
	// HttpEngine http.Handler结构，作为net/http框架使用, 实际上是gin.Engine
	HttpEngine() *gin.Engine
}
