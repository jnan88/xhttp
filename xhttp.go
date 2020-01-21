package main

import (
	"flag"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	//dir, _ := os.Executable()
	//exPath := filepath.Dir(dir)
	cmd := exec.Command("pwd")
	buf, _ := cmd.Output()
	exPath := strings.Replace(string(buf), "\n", "", -1)
	path := flag.String("d", exPath, "监听目录绝对路径,xhttp -d /www/dist")
	port := flag.String("p", "9000", "监听端口,xhttp -p 9001 ")
	// 【必须调用】从 arguments 中解析注册的 flag
	flag.Parse()
	println("---------------xhttp v1.0---------")
	println("运行路径: " + *path)
	println("访问路径: http://localhost:" + *port)
	println("----------------------------------")
	http.Handle("/", http.FileServer(http.Dir(*path)))
	http.ListenAndServe(":"+*port, nil)
}
