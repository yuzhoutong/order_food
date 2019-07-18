package beego

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/utils"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

var (
	myskip                   = make(map[string]bool, 10)
	MyGlobalControllerRouter = make(map[string]ControllerComments)
	mygenInfoList            map[string][]ControllerComments
)

func MyInclude(cList ...ControllerInterface) *App {
	BeeApp.Handlers.MyInclude(cList...)
	return BeeApp
}
func MyAutoRouter(c ControllerInterface) *App {
	MyInclude(c)
	BeeApp.Handlers.MyAddAuto(c)
	return BeeApp
}

func (p *ControllerRegister) MyInclude(cList ...ControllerInterface) {
	if BConfig.RunMode == DEV {
		for _, c := range cList {
			reflectVal := reflect.ValueOf(c)
			t := reflect.Indirect(reflectVal).Type()
			gopath := os.Getenv("GOPATH")
			if gopath == "" {
				panic("you are in dev mode. So please set gopath")
			}
			pkgpath := ""

			wgopath := filepath.SplitList(gopath)
			for _, wg := range wgopath {
				wg, _ = filepath.EvalSymlinks(filepath.Join(wg, "src", t.PkgPath()))
				if utils.FileExists(wg) {
					pkgpath = wg
					break
				}
			}
			if pkgpath != "" {
				if _, ok := myskip[pkgpath]; !ok {
					// Info(pkgpath)
					myskip[pkgpath] = true
					MyparserPkg(pkgpath, t.PkgPath())
				}
			}
		}
	}
}

func (p *ControllerRegister) MyAddAuto(c ControllerInterface) {
	p.MyAddAutoPrefix("/", c)
}

func (p *ControllerRegister) MyAddAutoPrefix(prefix string, c ControllerInterface) {
	reflectVal := reflect.ValueOf(c)
	rt := reflectVal.Type()
	ct := reflect.Indirect(reflectVal).Type()
	key := ct.PkgPath() + ":" + ct.Name()
	controllerName := strings.TrimSuffix(ct.Name(), "Controller")
	for i := 0; i < rt.NumMethod(); i++ {
		if !utils.InSlice(rt.Method(i).Name, exceptMethod) {
			fn := rt.Method(i).Name
			ck := key + "/" + fn
			if comm, ok := MyGlobalControllerRouter[ck]; ok {
				p.Add("/"+strings.ToLower(controllerName)+comm.Router, c, strings.Join(comm.AllowHTTPMethods, ",")+":"+comm.Method)
				//如果使用了注解路由，不添加默认路由
				continue
			}
			//转为大写
			upfn := strings.ToUpper(fn)
			if _, ok := HTTPMETHOD[upfn]; ok {
				//restful 如果是get ，post之类的 只添加get，post之类的请求方式
				route := &controllerInfo{}
				route.routerType = routerTypeBeego
				route.methods = map[string]string{upfn: fn}
				route.controllerType = ct
				pattern := path.Join(prefix, strings.ToLower(controllerName), strings.ToLower(rt.Method(i).Name), "*")
				route.pattern = pattern
				p.addToRouter(upfn, pattern, route)
				route2 := &controllerInfo{}
				route2.routerType = routerTypeBeego
				route2.methods = map[string]string{upfn: fn}
				route2.controllerType = ct
				pattern2 := strings.ToLower(controllerName)
				route2.pattern = pattern2
				p.addToRouter(upfn, pattern2, route2)
			} else {
				route := &controllerInfo{}
				route.routerType = routerTypeBeego
				route.methods = map[string]string{"*": rt.Method(i).Name}
				route.controllerType = ct
				pattern := path.Join(prefix, strings.ToLower(controllerName), strings.ToLower(rt.Method(i).Name), "*")
				patternFixInit := path.Join(prefix, controllerName, rt.Method(i).Name)
				route.pattern = pattern
				for _, m := range HTTPMETHOD {
					p.addToRouter(m, pattern, route)
					p.addToRouter(m, patternFixInit, route)
				}
			}
		}
	}
}

func MyparserPkg(pkgRealpath, pkgpath string) error {
	rep := strings.NewReplacer("/", "_", ".", "_")
	commentFilename = "My" + commentPrefix + rep.Replace(pkgpath) + ".go"
	if !compareFile(pkgRealpath) {
		Info(pkgRealpath + " no changed")
		return nil
	}
	mygenInfoList = make(map[string][]ControllerComments)
	fileSet := token.NewFileSet()
	astPkgs, err := parser.ParseDir(fileSet, pkgRealpath, func(info os.FileInfo) bool {
		name := info.Name()
		return !info.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
	}, parser.ParseComments)

	if err != nil {
		return err
	}
	for _, pkg := range astPkgs {
		for _, fl := range pkg.Files {
			for _, d := range fl.Decls {
				switch specDecl := d.(type) {
				case *ast.FuncDecl:
					if specDecl.Recv != nil {
						exp, ok := specDecl.Recv.List[0].Type.(*ast.StarExpr) // Check that the type is correct first beforing throwing to parser
						if ok {
							MyparserComments(specDecl.Doc, specDecl.Name.String(), fmt.Sprint(exp.X), pkgpath)
						}
					}
				}
			}
		}
	}
	MygenRouterCode()
	// fmt.Println("pkgRealpath", pkgRealpath)
	savetoFile(pkgRealpath)
	return nil
}

func MyparserComments(comments *ast.CommentGroup, funcName, controllerName, pkgpath string) error {
	if comments != nil && comments.List != nil {
		for _, c := range comments.List {
			t := strings.TrimSpace(strings.TrimLeft(c.Text, "//"))
			if strings.HasPrefix(t, "@router") {
				elements := strings.TrimLeft(t, "@router ")
				e1 := strings.SplitN(elements, " ", 2)
				if len(e1) < 1 {
					return errors.New("you should has router infomation")
				}
				key := pkgpath + ":" + controllerName + "/" + funcName
				cc := ControllerComments{}
				cc.Method = funcName
				cc.Router = e1[0]
				if len(e1) == 2 && e1[1] != "" {
					e1 = strings.SplitN(e1[1], " ", 2)
					if len(e1) >= 1 {
						cc.AllowHTTPMethods = strings.Split(strings.Trim(e1[0], "[]"), ",")
					} else {
						cc.AllowHTTPMethods = append(cc.AllowHTTPMethods, "get")
					}
				} else {
					cc.AllowHTTPMethods = append(cc.AllowHTTPMethods, "get")
				}
				if len(e1) == 2 && e1[1] != "" {
					keyval := strings.Split(strings.Trim(e1[1], "[]"), " ")
					for _, kv := range keyval {
						kk := strings.Split(kv, ":")
						cc.Params = append(cc.Params, map[string]string{strings.Join(kk[:len(kk)-1], ":"): kk[len(kk)-1]})
					}
				}
				mygenInfoList[key] = append(mygenInfoList[key], cc)
			}
		}
	}
	return nil
}

func MygenRouterCode() {
	os.Mkdir(path.Join(AppPath, "routers"), 0755)
	var (
		globalinfo string
		sortKey    []string
	)
	for k := range mygenInfoList {
		sortKey = append(sortKey, k)
	}
	sort.Strings(sortKey)
	for _, k := range sortKey {
		cList := mygenInfoList[k]
		for _, c := range cList {
			allmethod := "nil"
			if len(c.AllowHTTPMethods) > 0 {
				allmethod = "[]string{"
				for _, m := range c.AllowHTTPMethods {
					allmethod += `"` + m + `",`
				}
				allmethod = strings.TrimRight(allmethod, ",") + "}"
			}

			globalinfo = globalinfo + `
	beego.MyGlobalControllerRouter["` + k + `"] =beego.ControllerComments{
			"` + strings.TrimSpace(c.Method) + `",
			` + "`" + c.Router + "`" + `,
			` + allmethod + `,
			nil}
`
		}
	}
	if globalinfo != "" {
		f, err := os.Create(path.Join(AppPath, "routers", commentFilename))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString(strings.Replace(globalRouterTemplate, "{{.globalinfo}}", globalinfo, -1))
	} else {
		//不使用注解路由时 删除生成的路由文件
		if utils.FileExists(path.Join(AppPath, "routers", commentFilename)) {
			os.Remove(path.Join(AppPath, "routers", commentFilename))
		}
	}
}
