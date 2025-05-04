web应用

处理请求
模板
中间件
存储数据
HTTPS, HTTP2
测试
部署

创建Web Server
http.Server这是一个struct
Addr
Handler, 为nil就是DefaultServeMux
ListenAndServe()

handler是一个接口
handler定义了ServerHTTP()

http request -> DefaultServerMux -(一对多)->(Handler1, Handler2, ...)
http request -> myHandler

http.Handle("/", myHandler)
http.HandleFunc("/", func())

func NotFoundHandler() Handler
返回404
func RedirectHandler(url string, code int) Handler
跳转到url, code是跳转码
func StripPrefix(prefix string, h Handler) Handler
去掉请求url中的prefix前缀，将请求发送给h这个Handler
类似中间件，这种行为可以叫做 修饰，其本身叫做 修饰器
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
在指定时间内，将请求发送给h并运行，超时返回msg给请求
也相当于是 修饰器
func FileServer(root FileSystem) Handler
使用基于root的文件系统响应请求
type FileSystem interface {
    Open(name string) (File, error)
}
使用时用的是操作系统的文件系统
type Dir string
func (d Dir) Open(name string) (File, error)  // 不需要是自定义的struct也可以定义方法
