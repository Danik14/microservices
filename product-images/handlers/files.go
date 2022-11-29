package handlers

import (
	"fmt"
	"net/http"

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
	fmt.Println("name:", fileName)

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, "./imagestore/"+fileName)
	if err != nil {
		f.log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", fileName))
}
