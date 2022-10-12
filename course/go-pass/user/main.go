package main

import (
	"flag"
	"fmt"
	"path/filepath"
	//"git.imooc.com/coding-535/common"
	"git.imooc.com/coding-447/common"
	"user/domain/repository"

	//"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/opentracing/opentracing-go"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	service2 "user/domain/service"
	"user/handler"
	//hystrix2 "user/plugin/hystrix"
	"strconv"
	user "user/proto/user"
)

var (
    //服务地址
	hostIp = "192.168.0.108"
    //服务地址
    serviceHost = hostIp
    //服务端口
	servicePort = "8081"
     
	//注册中心配置
	consulHost  = hostIp
	consulPort int64 = 8500
	//链路追踪
	tracerHost = hostIp
	tracerPort = 6831
	//熔断端口，每个服务不能重复 
	//hystrixPort = 9092
	//监控端口，每个服务不能重复
	prometheusPort = 9192
)

func main() {
    //需要本地启动，mysql，consul中间件服务
	//1.注册中心
	consul:=consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost+":"+strconv.FormatInt(consulPort,10),
		}
	})
	//2.配置中心，存放经常变动的变量
	consulConfig,err := common.GetConsulConfig(consulHost,consulPort,"/micro/config")
	if err !=nil {
		common.Error(err)
	}
	//3.使用配置中心连接 mysql
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	//初始化数据库
	db,err := gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@("+mysqlInfo.Host+":3306)/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
        //命令行输出下，方便查看错误
		fmt.Println(err)
		common.Fatal(err)
	}
	defer db.Close()
	//禁止复表
	db.SingularTable(true)

	//4.添加链路追踪
	t,io,err := common.NewTracer("go.micro.service.user",tracerHost+":"+strconv.Itoa(tracerPort))
	if err !=nil {
		common.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//添加熔断器，作为客户端需要启用
	//hystrixStreamHandler := hystrix.NewStreamHandler()
	//hystrixStreamHandler.Start()

	//添加日志中心
	//1）需要程序日志打入到日志文件中
	//2）在程序中添加filebeat.yml 文件
	//3) 启动filebeat，启动命令 ./filebeat -e -c filebeat.yml
	fmt.Println("日志统一记录在根目录 micro.log 文件中，请点击查看日志！")

	//启动监听程序
	//go func() {
	//	//http://192.168.0.112:9092/turbine/turbine.stream
	//	//看板访问地址 http://127.0.0.1:9002/hystrix，url后面一定要带 /hystrix
	//	err = http.ListenAndServe(net.JoinHostPort("0.0.0.0",strconv.Itoa(hystrixPort)),hystrixStreamHandler)
	//	if err !=nil {
	//		common.Error(err)
	//	}
	//}()

	//5.添加监控
	common.PrometheusBoot(prometheusPort)


	//下载kubectl：https://kubernetes.io/docs/tasks/tools/#tabset-2
	//macos：
	// 1.curl -LO "https://dl.k8s.io/release/v1.21.0/bin/darwin/amd64/kubectl"
	// 2.chmod +x ./kubectl
	// 3.sudo mv ./kubectl /usr/local/bin/kubectl
	//   sudo chown root: /usr/local/bin/kubectl
	// 5.kubectl version --client
	// 6.集群模式下直接拷贝服务端~/.kube/config 文件到本机 ~/.kube/confg 中
	//   注意：- config中的域名要能解析正确
	//        - 生产环境可以创建另一个证书
	// 7.kubectl get ns 查看是否正常
	//
	//6.创建k8s连接
	//在集群外面连接
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		common.Fatal(err.Error())
	}

	//在集群中外的配置
	//config, err := rest.InClusterConfig()
	//if err != nil {
	//	panic(err.Error())
	//}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		common.Fatal(err.Error())
	}

	//7.创建服务
	service := micro.NewService(
		//自定义服务地址，且必须写在其它参数前面
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise =serviceHost+":"+servicePort
		})),
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		//指定服务端口
		micro.Address(":"+servicePort),
		//添加注册中心
		micro.Registry(consul),
		//添加链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//只作为客户端的时候起作用，如果存在调用别人的情况，原则上不去主动调用
		//micro.WrapClient(hystrix2.NewClientHystrixWrapper()),
		//添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(1000)),
	)
 
	service.Init()

	//只能执行一遍
	//err = repository.NewUserRepository(db).InitTable()
	//if err != nil {
	//	common.Fatal(err)
	//}

	// 注册句柄，可以快速操作已开发的服务
	userDataService:=service2.NewUserDataService(repository.NewUserRepository(db),clientset)
	user.RegisterUserHandler(service.Server(), &handler.UserHandler{ UserDataService:userDataService})

	// 启动服务
	if err := service.Run(); err != nil {
        //输出启动失败信息
		common.Fatal(err)
	}
}

