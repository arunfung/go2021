package controller

var (
	views       map[string][][3]string
	controllers map[string]interface{}
)

func init() {
	views = make(map[string][][3]string, 0)
	controllers = make(map[string]interface{}, 0)

	initView()
	addController()
}

func initView() {
	views["login_view"] = [][3]string{
		0: {
			0: "auto",
			1: "Login",
			2: "登入系统",
		},
		1: {
			0: "auto",
			1: "Register",
			2: "注册用户",
		},
	}
	views["index_view"] = [][3]string{
		0: {
			0: "index",
			1: "Index",
			2: "进入首页",
		},
		1: {
			0: "user",
			1: "List",
			2: "展示用户信息",
		},
	}
}

func addController() {
	// addIndexController()
	// addUserController()
	// addAutoController()

	controllers["index"] = &IndexController{}
	controllers["user"] = &UserController{}
	controllers["auto"] = &AutoController{}
}

// [][3]string{
// 0: {
// 	0: "auto",
// 	1: "Login",
// 	2: "登入系统",
// },
// 1: {
// 	0: "auto",
// 	1: "Register",
// 	2: "注册用户",
// },
// 格式化输出方法和描述
func toModelFormat(methods [][3]string) ([]string, []string) {
	var method = make([]string, len(methods))
	var decs = make([]string, len(methods))

	for k, v := range methods {
		method[k] = v[0] + "::" + v[1]
		decs[k] = v[2]
	}

	return method, decs
}