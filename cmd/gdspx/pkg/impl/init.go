package impl

import (
	"go/build"
	"os"
	"os/exec"
	"path"
	"runtime"

	_ "embed"

	"github.com/realdream-ai/gdspx/cmd/gdspx/pkg/util"
)

func downloadPack(dstDir, tagName, postfix string) error {
	urlHeader := "https://github.com/JiepengTan/godot/releases/download/"
	fileName := tagName + postfix
	url := urlHeader + tagName + "/" + fileName
	// download pc
	err := util.DownloadFile(url, path.Join(dstDir, fileName))
	if err != nil {
		return err
	}
	// download web
	fileName = tagName + "_web.zip"
	url = urlHeader + tagName + "/" + fileName
	err = util.DownloadFile(url, path.Join(dstDir, fileName))
	if err != nil {
		return err
	}
	// download webpack
	fileName = tagName + "_webpack.zip"
	url = urlHeader + tagName + "/" + fileName
	err = util.DownloadFile(url, path.Join(dstDir, fileName))
	if err != nil {
		return err
	}
	return err
}

func CheckAndGetAppPath(tag, version string) (string, string, error) {
	binPostfix := ""
	if runtime.GOOS == "windows" {
		binPostfix = "_win.exe"
	} else if runtime.GOOS == "darwin" {
		binPostfix = "_darwin"
	} else if runtime.GOOS == "linux" {
		binPostfix = "_linux"
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	tagName := tag + version
	dstFileName := tagName + binPostfix
	gdx, err := exec.LookPath(dstFileName)
	if err == nil {
		if _, err := exec.Command(gdx, "--version").CombinedOutput(); err == nil {
			return binPostfix, gdx, nil
		}
	}

	dstDir := gopath + "/bin"
	CmdPath := path.Join(dstDir, dstFileName)
	info, err := os.Stat(CmdPath)
	if os.IsNotExist(err) {
		println("Downloading pack...")
		err := downloadPack(dstDir, tagName, binPostfix)
		if err != nil {
			print("downloadPack error:" + err.Error())
			return binPostfix, CmdPath, err
		}
		if err := os.Chmod(CmdPath, 0755); err != nil {
			return binPostfix, CmdPath, err
		}
	} else if err != nil {
		return binPostfix, "", err
	} else {
		if info.Mode()&0111 == 0 {
			if err := os.Chmod(CmdPath, 0755); err != nil {
				return binPostfix, CmdPath, err
			}
		}
	}
	return binPostfix, CmdPath, nil
}
