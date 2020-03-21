package handler

import (
    "net/http"
)

// HTTPInterceptor : http 请求拦截器
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            r.ParseForm()
            username := r.Form.Get("username")
            token := r.Form.Get("token")

            // 验证登录 token 是否有效
            if len(username) < 3 || !IsTokenValid(token) {
                // w.WriteHeader(http.StatusForbidden)
                // token 校验失败则跳转到登录页面
                http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
                return
            }
            h(w, r)
        })
}
