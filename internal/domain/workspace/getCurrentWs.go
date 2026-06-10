package workspace

import (
	"fmt"
	"io"
	"os"

	"github.com/IaK3lwin/ac/internal/config"
	"github.com/IaK3lwin/ac/internal/domain/syserror"
)

var pathToFolderCurrentWs = config.Config().GetPathCurrentWs()

func GetNameWsCurrent() (string, error) {

	filePath := fmt.Sprintf("%v/%v", pathToFolderCurrentWs, config.Config().GetFileNameCurrentWs())

	fmt.Printf("directory exist? -> %v\n", directoryCurrentWsExist())

	if !directoryCurrentWsExist() {

		err := os.MkdirAll(pathToFolderCurrentWs, 0o777)

		if err != nil {
			return "", err
		}

		err = createFile(filePath)

		if err != nil {
			return "", err
		}

		fmt.Println("file created! with path: ", filePath)
	}

	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func directoryCurrentWsExist() bool {
	fmt.Println("directory current: ", pathToFolderCurrentWs)

	info, err := os.Stat(pathToFolderCurrentWs)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()

}

func createFile(nameFile string) error {

	_, err := os.Create(nameFile)

	if err != nil {
		return syserror.FactoryFileError(err.Error(), "domain/workspace/getWsCurrent")
	}

	return nil
}
