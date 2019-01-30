package main

import (
	"net/http"
	"log"
	"fmt"
	"html/template"
	"os"
)

func main() {
	//定义3个路由
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/loginsuccess", handleLoginSuccess)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))//启动服务器

}

//登录成功界面
func handleLoginSuccess(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"登录成功")//输出登录成功
}

//登录处理
func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")//获取实体值
	password := r.PostFormValue("password")//获取实体值
	if username == "admin" && password == "123456"{//登录成功
		http.Redirect(w,r,"/loginsuccess",http.StatusFound)//重定向到登录成功界面
	}else{//登录失败
		t,err := template.ParseFiles("src/WebServer/html/login.html")
		//错误处理
		if err != nil{
			fmt.Println("找不到解析的文件")
			os.Exit(0)
		}
		t = template.Must(t,err)
		err = t.Execute(w,"用户名或密码错误")//解析模板
		//错误处理
		if err != nil{
			fmt.Println("解析错误")
			os.Exit(0)
		}
		http.Redirect(w,r,"/",http.StatusFound)//重定向到错误界面
	}
}

//首页处理
func handleIndex(w http.ResponseWriter, r *http.Request) {
	t,err := template.ParseFiles("src/WebServer/html/login.html")
	//错误处理
	if err != nil{
		fmt.Println("找不到解析的文件")
		os.Exit(0)
	}
	t = template.Must(t,err)
	err = t.Execute(w,"")//解析模板
	//错误处理
	if err != nil{
		fmt.Println("解析错误")
		os.Exit(0)
	}
}




























