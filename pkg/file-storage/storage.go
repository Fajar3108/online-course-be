package file_storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

var storageBasePath = "./storage"

func Store(file *multipart.FileHeader, dir string, isPublic bool) (path string, err error) {
	currentStoragePath := storageBasePath

	if isPublic {
		currentStoragePath += "/public/"
	} else {
		currentStoragePath += "/private/"
	}

	safeFilename := filepath.Base(file.Filename)
	fileName := fmt.Sprintf("%s-%s", strconv.Itoa(int(time.Now().Unix())), safeFilename)
	storagePath := fmt.Sprintf("%s%s", currentStoragePath, dir)

	if _, err = os.Stat(storagePath); os.IsNotExist(err) {
		if err = os.MkdirAll(storagePath, 0755); err != nil {
			return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	fullPath := fmt.Sprintf("%s/%s", storagePath, fileName)

	src, err := file.Open()
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return fmt.Sprintf("%s/%s", dir, fileName), nil
}

func Remove(path string, isPublic bool) (err error) {
	currentStoragePath := storageBasePath

	if isPublic {
		currentStoragePath += "/public/"
	} else {
		currentStoragePath += "/private/"
	}

	filePath := fmt.Sprintf("%s%s", currentStoragePath, path)

	if err = os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return fiber.NewError(fiber.StatusNotFound, "File not found")
		}

		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
