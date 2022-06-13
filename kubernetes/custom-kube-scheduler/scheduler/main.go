/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/13 14:19
 */
package main

import (
	"github.com/julienschmidt/httprouter"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	"log"
	"net/http"
)

// filter 根据扩展程序定义的预选规则来过滤节点
func filter(args schedulerapi.ExtenderArgs) *schedulerapi.ExtenderFilterResult {
	//
}
func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.POST("/filter", Filter)
	router.POST("/prioritize", Prioritize)

	log.Fatal(http.ListenAndServe(":8888", router))
	//fmt.Println(router)

}
