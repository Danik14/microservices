package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Danik14/microservices/product-images/utils"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
)

// Files is a handler for reading and writing files
type Files struct {
	log hclog.Logger
	// store files.Storage
}

// NewFiles creates a new File handler
func NewFiles(l hclog.Logger) *Files {
	return &Files{log: l}
}

// ^[0-9]*$
func (f *Files) SaveFile(c *gin.Context) {
	id := c.Param("id")
	fileName := c.Param("filename")

	idCheck, fileNameCheck := utils.CheckRegEx(id, fileName)

	if !idCheck {
		f.log.Error("Id parameter did not match regex pattern")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "Id must only consist of numbers"})
		return
	}

	if !fileNameCheck {
		f.log.Error("filename parameter did not match regex pattern")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "File's name must start with character"})
		return
	}

	// Single file
	file, err := c.FormFile("file")
	if err != nil {
		f.log.Error("Error with file forming")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	extension := filepath.Ext(file.Filename)

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, "./imagestore/"+fileName+extension)
	if err != nil {
		f.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", fileName+extension))
}

const DOWNLOADS_PATH = "imagestore/"

func (f *Files) ServeFile(c *gin.Context) {
	fmt.Println(12234242342)
	fileName := c.Param("filename")
	if fileName == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error": "Empty Filename"})
		return
	}
	targetPath := filepath.Join(DOWNLOADS_PATH, fileName)
	fmt.Println(targetPath)
	//This ckeck is for example, I not sure is it can prevent all possible filename attacks
	//will be much better if real filename will not come from user side. I not even tryed this code
	if !strings.HasPrefix(filepath.Clean(targetPath), DOWNLOADS_PATH) {
		c.String(403, "Look like you attacking me")
		return
	}
	//Seems this headers needed for some browsers
	//(for example without this headers Chrome will download files as txt)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}
