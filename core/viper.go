package core

import (
	"blog_server/core/internal"
	"blog_server/global"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Viper 初始化
// 优先级： 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	var config string

	// 获取配置
	if len(path) == 0 {
		// 增加命令函输入 -c 增加文件名名称
		flag.StringVar(&config, "c", "", "choose config file.")
		// 一定要运行这个 要不不会保存config
		flag.Parse()
		if config == "" { // 判断命令函参数是不是为空
			// 判断 internal.configEnv 常量是否为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用的是gin模式的%s环境名称，config路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用的是gin模式的%s环境名称，config路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用的是gin模式的%s环境名称，config路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else {
				config = configEnv
				fmt.Printf("您正在使用的是%s环境变量，config路径为%s\n", internal.ConfigEnv, config)

			}
		} else {
			fmt.Printf("您正在使用命令行-c参数传递值，config的路径为%s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper() 传递的值，config的路径为%s\n", config)
	}

	v := viper.New()
	// 设置文件名
	v.SetConfigFile(config)
	// 设置文件类型
	v.SetConfigType("yaml")
	// 读取配置 看看格式是否正确
	err := v.ReadInConfig()
	// 打印报错信息
	if err != nil {
		panic(fmt.Errorf("Fatal error config file:%s\n", err))
	}
	// 查看配置
	v.WatchConfig()

	// 监听器 如果环境修改
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.BLOG_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.BLOG_CONFIG); err != nil {
		fmt.Println(err)
	}

	// 获取当前绝对路径
	global.BLOG_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	fmt.Println(global.BLOG_CONFIG.AutoCode)
	return v
}
