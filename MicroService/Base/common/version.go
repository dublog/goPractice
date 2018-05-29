package common



var (
	VERSION string = "0.0.1"
	COMPILE_TIME string = "2018-05-28"
	ApplicationName string
	ApplicationDir string
)


func GetVersionInfo() string{
	return "\r\n=====================================\r\n"+
		"Copyright (C) xxx \r\n"+
		"Application : " + ApplicationName +"\r\n"+
		"Version : "  + VERSION + "\r\n" +
		"Complie Time : " + COMPILE_TIME  + "\r\n" +
		   "=====================================\r\n"
}