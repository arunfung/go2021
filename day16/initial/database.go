package initial
import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func initDatabase() {
	user,_ := beego.AppConfig.String("mysql_user")
	pass,_ := beego.AppConfig.String("mysql_pass")
	host,_ := beego.AppConfig.String("mysql_host")
	port,_ := beego.AppConfig.Int("mysql_port")
	dbname,_ := beego.AppConfig.String("mysql_dbname")
	// "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4"
	fmt.Println("MySQL连接配置 ： ", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", user,pass,host,port,dbname))
	// 定义连接配置
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", user,pass,host,port,dbname))

	orm.RegisterDriver("mysql", orm.DRMySQL)
}
