package options

import (
	"encoding/json"
	"fmt"
	pixiuConfig "github.com/caoyingjunz/pixiulib/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"rbac-demo/cmd/app"
	"rbac-demo/cmd/app/config"
	"rbac-demo/cmd/app/router"
	"rbac-demo/pkg/db"
)

const (
	defaultConfigFile = "/etc/rbac-demo/config.yaml"
)

type Options struct {
	ConfigFile      string
	ComponentConfig config.Config
	DB              *gorm.DB

	Factory db.ShareDaoFactory

	ginEngine *gin.Engine
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Complete() error {
	configFile := o.ConfigFile
	if len(configFile) == 0 {
		configFile = os.Getenv("ConfigFile")
	}
	if len(configFile) == 0 {
		configFile = defaultConfigFile
	}

	// 解析 yaml 文件
	c := pixiuConfig.New()
	c.SetConfigFile(configFile)
	c.SetConfigType("yaml")
	if err := c.Binding(&o.ComponentConfig); err != nil {
		return err
	}

	// 注册依赖组件
	if err := o.register(); err != nil {
		return err
	}
	return nil
}

func (o *Options) register() error {
	if err := o.registerMysqlClient(); err != nil {
		return err
	}

	o.registerGinClient()

	return nil
}

func (o *Options) registerMysqlClient() error {
	data, err := json.Marshal(&o.ComponentConfig.Mysql)
	if err != nil {
		return err
	}
	DB, err := db.Connection(data)
	if err != nil {
		return err
	}

	o.DB = DB

	o.Factory = db.NewDaoFactory(o.DB)

	app.NewGlobal(o.Factory)

	return nil
}

func (o *Options) registerGinClient() {
	o.ginEngine = gin.Default()
	router.RegistryRoutes(o.ginEngine)

}

func (o *Options) Run() {
	o.ginEngine.Run(fmt.Sprintf(":%d", o.ComponentConfig.Default.Listen))
}
