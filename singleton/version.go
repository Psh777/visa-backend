package singleton

var version string
var node string
var uptime int64

func SetData(v, n string) {
	version = v
	node = n
}

func GetData() (string, string){
	return version, node
}

func GetUptime() int64{
	return uptime
}

func SetUptime(i int64) {
	uptime = i
}