package routes

import (
	"awesomeProject/database"
	"awesomeProject/database/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func GetOneBlogHandler(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog
	result := database.DB.First(&blog, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}

func CreateBlogHandler(c *gin.Context) {
	const MaxFileSize = 50 * 1024 * 1024 // 50MB
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxFileSize)

	// ตรวจสอบ Content-Type
	contentType := c.ContentType()
	var blog models.Blog

	if contentType == "application/json" {
		// ถ้าเป็น JSON ให้ bind JSON
		if err := c.ShouldBindJSON(&blog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload", "details": err.Error()})
			return
		}
	} else if contentType == "multipart/form-data" {
		// ถ้าเป็น multipart ให้ bind form และจัดการไฟล์
		if err := c.ShouldBind(&blog); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data", "details": err.Error()})
			return
		}

		// จัดการไฟล์
		file, err := c.FormFile("image")
		if err == nil {
			uploadDir := "./uploads"
			if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
				return
			}

			// สร้างชื่อไฟล์
			timestamp := time.Now().Unix()
			ext := filepath.Ext(file.Filename)
			filename := fmt.Sprintf("%d%s", timestamp, ext)
			filePath := filepath.Join(uploadDir, filename)

			// บันทึกไฟล์
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file", "details": err.Error()})
				return
			}

			// เก็บ path ใน blog
			blog.ImageURL = filepath.ToSlash(filePath)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported Content-Type"})
		return
	}

	// บันทึก blog ลงในฐานข้อมูล
	result := database.DB.Create(&blog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog", "details": result.Error.Error()})
		return
	}

	// ส่งข้อมูลกลับ
	c.JSON(http.StatusOK, blog)
}

func UpdateBlogHandler(c *gin.Context) {

}

func UploadFileHandler(c *gin.Context) {
	const MaxFileSize = 1000 * 1024 * 1024 // 50MB
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxFileSize)

	// ตรวจสอบว่าเป็น multipart
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded", "details": err.Error()})
		return
	}

	// สร้างไดเรกทอรีสำหรับการอัปโหลดไฟล์
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// สร้างชื่อไฟล์
	timestamp := time.Now().Unix()
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", timestamp, ext)
	filePath := filepath.Join(uploadDir, filename)

	// บันทึกไฟล์ไปยังเซิร์ฟเวอร์
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file", "details": err.Error()})
		return
	}

	// ส่ง URL ของไฟล์กลับ
	c.JSON(http.StatusOK, gin.H{"image_url": filepath.ToSlash(filePath)})
}
