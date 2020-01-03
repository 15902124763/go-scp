package file

import (
	"github.com/pkg/sftp"
	"github.com/yarm/go-scp/connect"
	"log"
	"os"
	"path"
	"strings"
)

func ScpSsh(localFilePath, remoteDir string, conn connect.Conn) error {
	var (
		client *sftp.Client
		err    error
	)

	// 连接
	client, err = connect.SftpSsh(conn)

	if err != nil {
		log.Println("scpCopy:", err)
		return err
	}

	defer client.Close()
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		log.Println("scpCopy:", err)
		return err
	}
	defer srcFile.Close()

	var remoteFileName = Base4Windows(localFilePath)
	dstFile, err := client.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		log.Println("scpCopy:", err)
		return err
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}

		dstFile.Write(buf[0:n])
	}
	return nil
}

func Base4Windows(dir string) string {
	if strings.Contains(dir, "/") {

		return path.Base(dir)
	} else {
		return base(dir)
	}
}

func base(path string) string {
	if path == "" {
		return "."
	}
	// Strip trailing slashes.
	for len(path) > 0 && path[len(path)-1] == '\\' {
		path = path[0 : len(path)-1]
	}
	// Find the last element
	if i := strings.LastIndex(path, "\\"); i >= 0 {
		path = path[i+1:]
	}
	// If empty now, it had only slashes.
	if path == "" {
		return "\\"
	}
	return path
}
