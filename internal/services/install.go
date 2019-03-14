package services

import "github.com/betterde/ects/config"

type InstallService interface {
	WriteConfigFile(conf config.Config) (bool, error)
}

type installService struct {
	
} 

func NewInstallService() InstallService {
	return &installService{
		
	}
}

// 将配置写入配置文件
func (s *installService) WriteConfigFile(conf config.Config) (bool, error) {
	return true, nil
}