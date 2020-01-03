package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/qianlnk/pgbar"
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

	// 解析入参
	inputArgs, err := getOsArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 默认端口
	inputArgs.Conn.Port = 22
	//marshal, err := json.Marshal(inputArgs)
	//fmt.Print(string(marshal))

	// 进度条
	pgb := pgbar.New("upload file")
	b := pgb.NewBar("upload bar", 1)

	b.SetSpeedSection(10, 100)
	for {
		b.Add()
		// 执行上传
		file.ScpSsh(inputArgs.LocalFilePath, inputArgs.RemoteDir, inputArgs.Conn)
		return
	}
}

func getOsArgs() (inputArgs InputArgs, err error) {
	var result InputArgs
	args := os.Args
	//fmt.Println("命令行的参数有", len(args))
	//// 遍历 os.Args 切片，就可以得到所有的命令行输入参数值
	//for i, v := range args {
	//	fmt.Printf("args[%v]=%v\n", i, v)
	//}

	if len(args) < 3 {
		help()
		return result, errors.New("error:输入参数格式错误")
	}

	if args[1] == "-R" {
		result.LocalFilePath = args[2]
		result, err = remoteArgs(args[3], result)
		if err != nil {
			help()
			return result, errors.New("error:远程服务入参错误")
		}
	} else {
		localFilePath := ""
		localFilePath = args[1]
		if !strings.Contains(localFilePath, "./") {
			localFilePath = "./" + localFilePath
		}

		temp, err := remoteArgs(args[2], result)

		if err != nil {
			help()
			return result, errors.New("error:远程服务入参错误")
		}
		result.LocalFilePath = localFilePath
		result.Conn.Host = temp.Conn.Host
		result.Conn.User = temp.Conn.User
		result.Conn.Password = temp.Conn.Password
		result.RemoteDir = temp.RemoteDir
	}

	fmt.Println("input password:")
	Stdin := os.Stdin
	input := bufio.NewScanner(Stdin)
	input.Scan()
	result.Conn.Password = input.Text()

	return result, nil
}

func remoteArgs(s string, args InputArgs) (inputArgs InputArgs, err error) {

	if !strings.Contains(s, "@") || !strings.Contains(s, ":") {
		help()
		return args, errors.New("error:入参错误")
	}

	arrLeft := strings.Split(s, "@")
	args.Conn.User = arrLeft[0]

	arrRight := strings.Split(arrLeft[1], ":")

	if len(arrRight) < 2 {
		return args, errors.New("error:入参错误")
	}
	args.Conn.Host = arrRight[0]
	args.RemoteDir = arrRight[1]
	return args, nil
}

// 帮助文档
func help() {
	fmt.Println("\nExample:\n")
	fmt.Println("scp -R /usr/local/upload.txt root@youId:/opt/src/", "\n")
	fmt.Println("scp:输入命令", "\n")
	fmt.Println("-R:指定的文件夹，不填默认是当前路径", "\n")
	fmt.Println("root@youId:用户名和ip", "\n")
	fmt.Println("/opt/src/:远程服务器文件夹路径", "\n")
}

type InputArgs struct {
	Conn          connect.Conn
	LocalFilePath string
	RemoteDir     string
}
