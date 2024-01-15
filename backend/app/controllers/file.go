package controllers

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type File struct {
	Dirs       []string `json:"dirs"`
	Files      []string `json:"files"`
	CurrentDir []string `json:"current_path"`
	BasePath   string   `json:"base_path"`
}

type ResponseFileBody struct {
	Item  string   `json:"item"`
	Items []string `json:"items"`
	Old   string   `json:"old"`
	New   string   `json:"new"`
	News  []string `json:"news"`
}

func GetFiles(c *fiber.Ctx) error {
	file := File{}
	return DirsFilesList(c, file)
}

func DirsFilesList(c *fiber.Ctx, file File) error {

	file.BasePath = os.Getenv("BASE_PATH") + "/"
	currPath := filepath.Join(file.BasePath, filepath.Join(file.CurrentDir...))
	listItems, err := os.ReadDir(currPath)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	file.Dirs = make([]string, 0, len(listItems))
	file.Files = make([]string, 0, len(listItems))

	for _, item := range listItems {
		if item.IsDir() {
			file.Dirs = append(file.Dirs, item.Name())
		} else {
			file.Files = append(file.Files, item.Name())
		}
	}
	return c.Status(200).JSON(file)
}

func PostFiles(c *fiber.Ctx) error {
	file, resp := File{}, ResponseFileBody{}
	err := c.BodyParser(&resp)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	file.BasePath = os.Getenv("BASE_PATH") + "/"

	switch c.Params("action") {
	case "open":
		currPath := filepath.Join(file.BasePath, filepath.Join(file.CurrentDir...), resp.Item)
		info, err := os.Stat(currPath)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		if info.IsDir() {
			file.CurrentDir = append(file.CurrentDir, resp.Item)
		}

	case "parent":
		currPath := filepath.Join(file.BasePath, filepath.Join(file.CurrentDir...))
		info, err := os.Stat(currPath)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		if info.IsDir() {
			file.CurrentDir = append(file.CurrentDir, resp.Item)
		}

	case "download":
		currPath := filepath.Join(file.BasePath, filepath.Join(file.CurrentDir...), resp.Item)
		info, err := os.Stat(currPath)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		if !info.IsDir() {
			return c.Download(currPath)
		}

	case "create":
		currPath := filepath.Join(file.BasePath, filepath.Join(file.CurrentDir...), "New Folder")
		err = os.MkdirAll(currPath, 0755)
		if err != nil {
			return c.Status(500).JSON(err)
		}

	case "copy":
		for _, item := range resp.News {
			currPath := filepath.Join(file.BasePath, resp.Old, item)
			newPath := filepath.Join(file.BasePath, filepath.Join(file.CurrentDir...), item)
			info, err := os.Stat(currPath)
			if err != nil {
				return c.Status(500).JSON(err)
			}
			if !info.IsDir() {
				err = copyFile(currPath, newPath)
				if err != nil {
					return c.Status(500).JSON(err)
				}
			} else {
				err = dirCopyMove(currPath, newPath, "copy")
				if err != nil {
					return c.Status(500).JSON(err)
				}
			}
		}

	case "cut":
		for _, item := range resp.News {
			currPath := filepath.Join(file.BasePath, resp.Old, item)
			newPath := filepath.Join(file.BasePath, filepath.Join(file.CurrentDir...), item)
			info, err := os.Stat(currPath)
			if err != nil {
				return c.Status(500).JSON(err)
			}
			if !info.IsDir() {
				err = copyFile(currPath, newPath)
				if err != nil {
					return c.Status(500).JSON(err)
				}
				err = os.Remove(currPath)
				if err != nil {
					return c.Status(500).JSON(err)
				}
			} else {
				err = dirCopyMove(currPath, newPath, "move")
				if err != nil {
					return c.Status(500).JSON(err)
				}
			}
		}

	case "rename":

		newPath := filepath.Join(file.BasePath, filepath.Join(file.CurrentDir...), resp.New)
		oldPath := filepath.Join(file.BasePath, filepath.Join(file.CurrentDir...), resp.Old)
		err = os.Rename(oldPath, newPath)
		if err != nil {
			return c.Status(500).JSON(err)
		}

	case "delete":
		for _, item := range resp.Items {
			currPath := filepath.Join(file.BasePath, filepath.Join(file.CurrentDir...), item)
			info, err := os.Stat(currPath)
			if err != nil {
				return c.Status(500).JSON(err)
			}
			if !info.IsDir() {
				err = os.Remove(currPath)
				if err != nil {
					return c.Status(500).JSON(err)
				}
			} else {
				err = os.RemoveAll(currPath)
				if err != nil {
					return c.Status(500).JSON(err)
				}
			}
		}
	}
	return DirsFilesList(c, file)
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = destinationFile.Sync()
	if err != nil {
		return err
	}
	return nil
}
func dirCopyMove(src string, dest string, action string) error {
	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		return err
	}

	srcDir, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, f := range srcDir {
		if f.IsDir() {
			err = dirCopyMove(filepath.Join(src, f.Name()), filepath.Join(dest, f.Name()), action)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(filepath.Join(src, f.Name()), filepath.Join(dest, f.Name()))
			if err != nil {
				return err
			}
		}
	}
	if action == "move" {
		if err := os.RemoveAll(src); err != nil {
			return err
		}
	}
	return nil
}
