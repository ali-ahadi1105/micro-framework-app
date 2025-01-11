package quokka

<<<<<<< HEAD
=======
import (
	"fmt"

	"github.com/joho/godotenv"
)

>>>>>>> 5b5318a (Initial an project structure with micro framework)
const version = "1.0.0"

type Quokka struct {
	AppName string
	Debug   bool
	Version string
}

func (quokka *Quokka) New(rootPath string) error {
	initConfigs := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "middlewares", "logs", "public", "data", "migrations", "views", "tmp"},
	}

	err := quokka.Init(initConfigs)

	if err != nil {
		return err
	}

<<<<<<< HEAD
=======
	err = quokka.checkDotenvFile(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath)
	if err != nil {
		return err
	}

>>>>>>> 5b5318a (Initial an project structure with micro framework)
	return nil
}

func (quokka *Quokka) Init(path initPaths) error {
	root := path.rootPath
	for _, path := range path.folderNames {
		err := quokka.createDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}
<<<<<<< HEAD
=======

func (quokka *Quokka) checkDotenvFile(path string) error {
	err := quokka.createFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}
>>>>>>> 5b5318a (Initial an project structure with micro framework)
