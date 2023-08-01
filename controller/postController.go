package controller

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/shinohara320/travel-agent/database"
	"github.com/shinohara320/travel-agent/models"
	"github.com/shinohara320/travel-agent/util"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parse body")
	}

	// Handle uploaded image file
	if file, err := c.FormFile("image"); err == nil {
		src, err := file.Open()
		if err != nil {
			c.Status(500)
			return c.JSON(fiber.Map{
				"message": "Unable to open the uploaded file",
			})
		}
		defer src.Close()

		// Read the file into a []byte
		imageBytes, err := io.ReadAll(src)
		if err != nil {
			c.Status(500)
			return c.JSON(fiber.Map{
				"message": "Unable to read the uploaded file",
			})
		}

		blogpost.Image = imageBytes
	}

	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Congratulations, your post is online",
	})
}

// ... Rest of the code ...

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 6
	offset := (page - 1) * limit
	var total int64
	var getblog []models.Blog
	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&getblog)
	database.DB.Model(&models.Blog{}).Count(&total)
	lastPage := int(math.Ceil(float64(total) / float64(limit)))

	return c.JSON(fiber.Map{
		"data": getblog,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage,
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}
	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}

	// Separate handling for the image data
	fileHeader, err := c.FormFile("image")
	if err == nil {
		file, err := fileHeader.Open()
		if err == nil {
			defer file.Close()

			// Read the image data into a byte slice
			var imageData bytes.Buffer
			_, err = io.Copy(&imageData, file)
			if err == nil {
				// Update the Image field in the blog struct
				blog.Image = imageData.Bytes()
			} else {
				// Handle the error if reading image data fails
				fmt.Println("Error reading image data:", err)
			}
		} else {
			// Handle the error if opening the image file fails
			fmt.Println("Error opening image file:", err)
		}
	}

	// Update the data in the database, including the Image field
	database.DB.Save(&blog)

	return c.JSON(fiber.Map{
		"message": "post has been updated successfully",
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	Id, err := util.ParseJwt(cookie)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id = ?", Id).Preload("User").Find(&blog)

	return c.JSON(fiber.Map{
		"userId": Id,
		"blogs":  blog,
		"token":  cookie,
	})
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}
	deleteQuery := database.DB.Delete(&blog)
	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Opps!, Record not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "post deleted successfully",
	})

}

func GetLatestBlogs(c *fiber.Ctx) error {
	var latestBlogs []models.Blog
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 6
	offset := (page - 1) * limit
	var total int64
	database.DB.Model(&models.Blog{}).Count(&total)
	lastPage := int(math.Ceil(float64(total) / float64(limit)))

	err := database.DB.Order("id desc").Offset(offset).Limit(limit).Find(&latestBlogs).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil daftar berita terbaru",
		})
	}

	return c.JSON(fiber.Map{
		"data": latestBlogs,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": lastPage,
		},
	})
}
