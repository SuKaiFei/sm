// 作者 sukaifei
// 日期 2019/4/8

// 文件控制器

package http

import (
	"errors"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/ecode"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/render"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path"
	"regexp"
	"strconv"
	"time"
)

var (
	ErrFileTypeNotImage = errors.New("文件不是图片类型")
)

type FileConfig struct {
	Path string `dsn:"path"`
}

type File struct {
	*FileConfig
}

func NewFile() (file *File) {
	var (
		fc struct {
			Conf *FileConfig
		}
	)
	if err := paladin.Get("file.toml").UnmarshalTOML(&fc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	file = new(File)
	file.FileConfig = fc.Conf
	return
}

func (file *File) Upload(bm *bm.Context) {
	_, uploadFile, err := bm.Request.FormFile("file")
	if err != nil {
		bm.JSON(nil, err)
		return
	}

	if len(regexp.MustCompile(`image/*`).FindAllString(uploadFile.Header.Get("Content-Type"), 1)) == 0 {
		bm.Render(ecode.RequestErr.Code(), render.JSON{
			Code:    ecode.RequestErr.Code(),
			Message: ErrFileTypeNotImage.Error(),
			Data:    nil,
		})
		return
	}

	now := time.Now()
	strTime := path.Join(strconv.Itoa(now.Year()),
		fmt.Sprintf("%0.2v", int(now.Month())),
		fmt.Sprintf("%0.2v", now.Day()))

	idPath := uuid.New().String()
	rootPath := path.Join(file.Path,
		strTime,
		idPath)

	err = os.MkdirAll(rootPath, os.ModePerm)
	if err != nil {
		bm.JSON(nil, err)
		return
	}
	err = SaveUploadedFile(uploadFile, path.Join(rootPath, uploadFile.Filename))
	if err != nil {
		bm.JSON(nil, err)
		return
	}

	bm.JSON(path.Join(strTime, idPath, uploadFile.Filename), err)
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	io.Copy(out, src)
	return nil
}
