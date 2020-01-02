package file

import (
	"github.com/pkg/sftp"
	"log"
	"os"
	"path"
	"yarm.yang/scp/connect"
)

func ScpSsh(localFilePath, remoteDir string, conn connect.Conn) error {
	var (
		client *sftp.Client
		err        error
	)

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


	var remoteFileName = path.Base(localFilePath)
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

