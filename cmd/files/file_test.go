package files

import (
	"fmt"
	"github.com/apulisai/sdk/go-utils/utils"
	"github.com/go-labs/internal/logging"
	"path/filepath"
	"testing"
)

func TestName(t *testing.T) {
	rf := NewRf("../../build", "/")
	err := rf.Fill()

	fmt.Print(rf, err)
}

func TestBase(t *testing.T) {
	//path := "./a/v/log.png"
	p := filepath.Join("", "code")
	fmt.Print(p)
}

func TestFiles(t *testing.T) {
	files := []string{
		"code",
		"code/train/train.py",
		"code/train/train1.py",
		"code/train/train2.py",
		"model",
		"model/run/model.py",
		"model/run/model1.py",
		"model/run/model2.py",
	}
	root := utils.BuildFileObject(files)
	logging.Debug().Interface("root", root).Send()
	root.Child[0].Degrade("code")
	logging.Debug().Interface("root", root).Send()
}
