package storage

import (
	"fmt"

	"github.com/IaK3lwin/ac/internal/config"
	"github.com/IaK3lwin/ac/internal/domain/actions"
	"github.com/IaK3lwin/ac/internal/domain/syserror"

	"gopkg.in/yaml.v3"
	"os"
)

type YmlStorage struct {
}

func (ys YmlStorage) GetActions(workspaceName string) (map[string]actions.Actions, error) {
	actionsactionsList := map[string]actions.Actions{}

	pathWs := config.Config().GetPathWs()

	err := createDirIfNotExist(pathWs)

	if err != nil {
		return actionsactionsList, err
	}

	entry, err := os.ReadDir(pathWs)

	if err != nil {

		return actionsactionsList, syserror.StorageErrorHandle(err.Error(), pathWs, "infra/storage/ymlStorage")
	}

	for _, file := range entry {

		data, err := os.ReadFile(pathWs + "/" + file.Name())

		if err != nil {
			return actionsactionsList, syserror.StorageErrorReadFile(err.Error(), file.Name(), "infra/storage/ymlStorage/getActions")
		}

		var typeAction actions.ActionsBaseType

		err = yaml.Unmarshal(data, &typeAction)
		if err != nil {
			errorFormat := syserror.StorageErrorHandle(fmt.Sprintf("[Error file configuration .yml]\n %v", err.Error()),
				file.Name(), "infra/storage/yml/getAction")
			return actionsactionsList, errorFormat
		}

		actionCurrent, err := parseAction[typeAction.Type](data)

		if err != nil {
			return actionsactionsList, err
		}

		if len(actionCurrent.GetAlias()) == 0 {
			return actionsactionsList, syserror.StorageErrorHandle("Action dont have alias!: ", file.Name(), "infra/storage/yaml/getAlias")
		}

		actionsactionsList[actionCurrent.GetAlias()] = actionCurrent

	}

	return actionsactionsList, nil
}

func createDirIfNotExist(dirPath string) error {

	_, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		//folder dont exist

		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return syserror.StorageDirNotExist(dirPath, "domain/storage/GetActions", err.Error())
		}

		fmt.Println("file create with complete: ", dirPath)
	}

	return nil
}

func (ys YmlStorage) workspaceExist() bool {
	return false
}

func (ys YmlStorage) GetWorkspaceList() []string {
	wslist := []string{"teste", "code"}
	return wslist // remove, its dont clean code!!
}

func (ys YmlStorage) UpdateWsCurrent(curentWs string) error {
	return nil
}
