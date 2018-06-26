package server

import (
	"log"
	"time"
	"net/http"
	"wuyuyun-datastore-v0.2/api/handler"
	"os"
	"os/signal"
	"flag"
	"context"
	"github.com/gorilla/mux"
	"fmt"
)

func StartServer(){
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	// Add your routes as needed
	// 开启web服务
	srv := &http.Server{
		Handler:       Middleware( Router()),
		Addr:         "127.0.0.1:8004",
		// 设置超时避免Slowloris攻击。
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// 在goroutine中运行我们的服务器，以便它不会阻塞.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("服务器状态：", err)
		}
	}()
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

/**
	路由中间件，到达匹配路由指定的执行函数前调用
 */
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("✔路由中间件输出：%s:%s%s \n", r.Method, r.Host, r.URL)
		// Call the next handler, which can be another middleware in the chain, or the final handler.

		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

/**
路由配置
 */
func Router() http.Handler {
	// 创建一个路由实例
	mx := mux.NewRouter()

	// 匹配结尾斜杠的行为
	mx.StrictSlash(true)
	// 或者将定义路由的代码分离，方便维护，如下：
	// 注册自定义匹配函数的一种新途径, 可以获取错误的请求状态信息，如"405 Method Not Allowed"、"404 Page Not Found"
	mx.MatcherFunc( func(req *http.Request, rm *mux.RouteMatch) bool {
		// 如果请求的方法不匹配（405 Method Not Allowed）或者路由不匹配（404 Page Not Found）
		if err := rm.MatchErr; err != nil {
			//panic(fmt.Sprintf("%q\n", rm.MatchErr.Error()))
			log.Printf("%q\n", rm.MatchErr.Error())
		}
		return req.ProtoMajor == 0
	})
	//注册自己的路由
	handler.CustomHandler(mx)
	return mx
}

