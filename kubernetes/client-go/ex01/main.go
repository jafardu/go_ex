/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/13 8:24
 */
package main

import "k8s.io/client-go/tools/clientcmd"

func main() {
	clientcmd.BuildConfigFromFlags("", *restclient.Config)

}
