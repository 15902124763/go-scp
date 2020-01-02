package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yarm/go-scp/connect"
	"github.com/yarm/go-scp/file"
	"os"
	"strings"
)

func main() {

	//fmt.Println("命令行的参数有", len(os.Args))
	//// 遍历 os.Args 切片，就可以得到所有的命令行输入参数值
	//for i, v := range os.Args {
	//	fmt.Printf("args[%v]=%v\n", i, v)
	//}
	//scpCopy("D:/temp/copy/rrr.txt", "/usr/local/src/")

	str := "&"
	for index := range str {
		fmt.Println( str[index])
	}

	// 解析入参
	inputArgs, err := getOsArgs()
	if err != nil{
		fmt.Println(err)
		return
	}

	file.ScpSsh("./temp/out.txt", "/usr/local/src/", inputArgs.Conn)
}

func getOsArgs() (inputArgs InputArgs, err error){
	args := os.Args
	fmt.Println("命令行的参数有", len(args))
	// 遍历 os.Args 切片，就可以得到所有的命令行输入参数值
	for i, v := range args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}

	len := len(args)

	if(len < 3){
		help()
		return inputArgs, errors.New("error:输入参数格式错误")
	}

	if args[1] == "-R"{
		fmt.Println("ok")
		inputArgs.LocalFilePath = args[2]
		inputArgs = remoteArgs(args[3], inputArgs)
		fmt.Println(inputArgs)
	}

	fmt.Println(inputArgs)

	fmt.Println("input password:")
	Stdin := os.Stdin
	input := bufio.NewScanner(Stdin)
	input.Scan()
	inputPassWord := input.Text()

	inputArgs.Conn.Password = inputPassWord

	return inputArgs, nil
}

func remoteArgs(s string, args InputArgs) (inputArgs InputArgs){
	userName := ""
	arrLeft := strings.Split(s, "@")
	userName = arrLeft[0]
	inputArgs.Conn.User = userName
	fmt.Println("userName:", userName)

	fmt.Println(arrLeft[1])
	arrRight := strings.Split(arrLeft[1], ":")
	inputArgs.Conn.Host = arrRight[0]
	inputArgs.RemoteDir = arrRight[1]
	return inputArgs
}

// 帮助文档
func help()  {
	fmt.Println("\nExample:\n")
	fmt.Println("scp -R /usr/local/upload.txt root@youId:/opt/src/","\n")
	fmt.Println("scp:输入命令","\n")
	fmt.Println("-R:指定的文件夹，不填默认是当前路径","\n")
	fmt.Println("root@youId:用户名和ip","\n")
	fmt.Println("/opt/src/:远程服务器文件夹路径","\n")
}

type InputArgs struct {
	Conn connect.Conn
	LocalFilePath string
	RemoteDir string
}