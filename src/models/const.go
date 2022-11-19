package models

import "fmt"

type DefaultFileAndFolderList struct {
	List map[string]FileInfo
}

type DefaultFileAndFolderListInterface interface {
	SetGetDefaultFileAndFolderList(projectName string)
	GetDefaultFileAndFolderList() map[string]FileInfo
}

func NewDefaultFileAndFolderList() DefaultFileAndFolderListInterface {
	return &DefaultFileAndFolderList{
		List: make(map[string]FileInfo),
	}
}

func (obj *DefaultFileAndFolderList) SetGetDefaultFileAndFolderList(projectName string) {
	obj.List = map[string]FileInfo{
		fmt.Sprintf("%s/app", projectName): {
			IsDir:   true,
			IsExist: false,
		},
		fmt.Sprintf("%s/assets", projectName): {
			IsDir:   true,
			IsExist: false,
		},
		fmt.Sprintf("%s/configs", projectName): {
			IsDir:   true,
			IsExist: false,
		},
		fmt.Sprintf("%s/modules", projectName): {
			IsDir:   true,
			IsExist: false,
		},
		fmt.Sprintf("%s/package", projectName): {
			IsDir:   true,
			IsExist: false,
		},
		fmt.Sprintf("%s/tests", projectName): {
			IsDir:   true,
			IsExist: false,
		},
		fmt.Sprintf("%s/.env.dev", projectName): {
			IsDir:   false,
			IsExist: false,
		},
		fmt.Sprintf("%s/.env.test", projectName): {
			IsDir:   false,
			IsExist: false,
		},
		fmt.Sprintf("%s/.env.prod", projectName): {
			IsDir:   false,
			IsExist: false,
		},
		fmt.Sprintf("%s/.gitignore", projectName): {
			IsDir:   false,
			IsExist: false,
		},
	}
}

func (obj *DefaultFileAndFolderList) GetDefaultFileAndFolderList() map[string]FileInfo {
	return obj.List
}
