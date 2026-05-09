package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	//http://127.0.0.1:8000/add?a=1&b=2
	//1.CallId的问题 每一个uri 都可以对应一个方法
	//2.数据传输协议 URL 参数传输协议
	//3.网络传输协议 使用HTTP解决

	// =========================
	// HTTP 服务核心概念
	// =========================
	// 1. Go 使用 net/http 标准库实现 Web 服务
	// 2. 每个 URL Path（如 /add）会绑定一个处理函数（Handler）
	// 3. Handler 本质：函数签名固定
	//    func(w http.ResponseWriter, r *http.Request)
	//    - w：用于构造响应（写回客户端）
	//    - r：封装客户端请求（参数、Header、Body等）

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {

		// =========================
		// 请求解析（Request Parsing）
		// =========================

		// ParseForm 会解析 URL 中的 query 参数（?a=1&b=2）
		// 并填充到 r.Form（map[string][]string）
		_ = r.ParseForm()

		// 调试用：查看访问路径
		fmt.Println("path:", r.URL.Path)

		// =========================
		// 参数获取（重点）
		// =========================

		// r.Form["a"] 返回 []string（因为一个 key 可以有多个值）
		// 直接取 [0] 存在风险：如果参数不存在会 panic
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])

		// 更安全写法（推荐）：
		// aStr := r.Form.Get("a") // 内部自动取第一个值
		// bStr := r.Form.Get("b")
		// a, _ := strconv.Atoi(aStr)
		// b, _ := strconv.Atoi(bStr)

		// =========================
		// 响应构造（Response）
		// =========================

		// 设置响应类型为 JSON（非常关键）
		w.Header().Set("Content-Type", "application/json")

		// Go 中常用 map / struct -> JSON
		// json.Marshal 会返回 []byte（序列化后的数据）
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})

		// 将数据写入 Response Body，返回给客户端
		_, _ = w.Write(jData)
	})

	// =========================
	// 启动 HTTP 服务（必须）
	// =========================
	// :8000 表示监听本机 8000 端口
	// 第二个参数 nil 表示使用默认路由器 DefaultServeMux
	// 这个调用会阻塞（一直运行，直到程序退出）
	fmt.Println("server running at http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
