/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/13 10:02
 */
package conf

/*

[kafka]
address=127.0.0.1:9092
topic=web_log

[taillog]
filename="./log.txt"
*/

type kafkaConf struct {
	Address     string `ini:"address"`
	Topic       string `ini:"topic"`
	ChanMaxSize int    `ini:"chanmaxsize"`
}

type TailLogConf struct {
	FileName string `ini:"filename"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"key"`
}
type AppConf struct {
	kafkaConf
	TailLogConf
	EtcdConf
}
