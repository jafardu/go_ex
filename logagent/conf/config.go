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
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type tailLogConf struct {
	FileName string `ini:"filename"`
}

type AppConf struct {
	kafkaConf
	tailLogConf
}
