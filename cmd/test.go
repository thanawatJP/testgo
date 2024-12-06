package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// สร้าง Gin Engine
	router := gin.Default()

	// จำกัดขนาดไฟล์ที่อัปโหลด (50MB)
	router.MaxMultipartMemory = 50 << 20 // 50MB

	// เส้นทางสำหรับอัปโหลดไฟล์
	router.POST("/upload", func(c *gin.Context) {
		// รับไฟล์จากฟอร์ม
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": "Unable to read file"})
			return
		}
		defer file.Close()

		// สร้างโฟลเดอร์สำหรับเก็บไฟล์ (ถ้ายังไม่มี)
		if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
			os.Mkdir("./uploads", os.ModePerm)
		}

		// ระบุชื่อไฟล์ปลายทาง
		dst := "./uploads/" + header.Filename

		// สร้างไฟล์ใหม่และเขียนข้อมูลลงไป
		out, err := os.Create(dst)
		if err != nil {
			c.JSON(500, gin.H{"error": "Unable to save file"})
			return
		}
		defer out.Close()

		if _, err := io.Copy(out, file); err != nil {
			c.JSON(500, gin.H{"error": "Unable to save file"})
			return
		}

		// ส่งข้อความตอบกลับ
		c.JSON(200, gin.H{
			"message": "File uploaded successfully",
			"file":    header.Filename,
		})
	})

	// รันเซิร์ฟเวอร์ที่พอร์ต 8080
	router.Run(":8080")
}
