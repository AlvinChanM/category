package main

import (
	"category/common"
	"category/domain/model"
	"category/domain/repository"
	service2 "category/domain/service"
	"category/handler"
	user "category/proto/category"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("localhost", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options){
		options.Addrs = []string{
			"localhost:8500",
			}
	})

	//服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		//这里设置地址和需要暴露的端口
		micro.Address("localhost:8082"),
		//添加consul作为注册中心
		micro.Registry(consulRegistry),
	)

	// 获取配置中心的mysql配置, 路径中不带前缀
	mysqlInfo := common.GetMysqlFormConsul(consulConfig, "mysql")


	//初始化服务
	srv.Init()

	//初始化数据库
	model.InitMysql(*mysqlInfo)
	defer model.DB.Close()

	// 初始化CategoryHandler
	rp :=repository.NewCategoryRepository(model.DB)
	rp.InitTable()
	categoryDataService := service2.NewCategoryDataService(rp)
	//注册Handler
	err = user.RegisterCategoryHandler(srv.Server(),&handler.Category{categoryDataService})

	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

