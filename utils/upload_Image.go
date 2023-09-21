package utils

import (
	"crypto/md5"
	"encoding/hex"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"restaurant/errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nfnt/resize"
)

type UploadedImage struct {
	Filename string
	MD5Hash  string
}

func UploadAndResizeImageHandler(c *fiber.Ctx, pathName string) (*UploadedImage, error) {
	// Retrieve the uploaded file from the request
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Failed to retrieve uploaded file:", err)
		return nil, errors.NewNotFoundError("ไม่สามารถเรียกไฟล์ที่อัปโหลด")
	}

	// Check if the file is an image
	err = isSupportedImageFormat(c) 
	if err != nil {
		return nil, errors.NewNotFoundError("รูปแบบภาพไม่ถูกต้อง")
	}

	// Create a new file to store the uploaded image
	uploadedFile, err := file.Open()
	if err != nil {
		log.Println("Failed to open uploaded file:", err)
		return nil, errors.NewNotFoundError("ไม่สามารถเปิดไฟล์ที่อัปโหลด")
	}
	defer uploadedFile.Close()

	// Create a new file to store the resized image
	resizedFilename := pathName + "resized_" + file.Filename

	// Check if the file with the same MD5 hash already exists
	duplicateFilename := checkDuplicateImage("resized_"+file.Filename, pathName)
	if duplicateFilename == nil {
		return nil, errors.NewNotFoundError("พบภาพที่ซ้ำกัน")
	}

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(resizedFilename), os.ModePerm); err != nil {
		log.Println("Failed to create directory:", err)
		return nil, errors.NewNotFoundError("ไม่สามารถสร้างไดเรกทอรีได้")
	}

	resizedFile, err := os.Create(resizedFilename)
	if err != nil {
		log.Println("Failed to create resized image file:", err)
		return nil, errors.NewNotFoundError("สร้างไฟล์ภาพที่ปรับขนาดไม่สำเร็จ")
	}
	defer resizedFile.Close()

	// Resize the image to a desired width and height using the resize library
	image, _, err := image.Decode(uploadedFile)
	if err != nil {
		log.Println("Failed to decode image:", err)
		return nil, errors.NewNotFoundError("ถอดรหัสรูปภาพไม่สำเร็จ")
	}
	resizedImage := resize.Resize(1000, 0, image, resize.Lanczos3)

	// Encode the resized image as JPEG and write it to the file
	err = encodeImage(resizedFile, resizedImage, file.Filename)
	if err != nil {
		log.Println("Failed to encode resized image:", err)
		return nil, errors.NewNotFoundError("เข้ารหัสภาพที่ปรับขนาดไม่สำเร็จ")
	}

	// Calculate the MD5 hash of the uploaded file
	uploadedFile.Seek(0, 0)
	hash := md5.New()
	if _, err := io.Copy(hash, uploadedFile); err != nil {
		log.Println("Failed to calculate MD5 hash:", err)
		return nil, errors.NewNotFoundError("ไม่สามารถคำนวณแฮช MD5")
	}
	md5Hash := hex.EncodeToString(hash.Sum(nil))

	// Create an UploadedImage struct with the filename and MD5 hash
	uploadedImage := &UploadedImage{
		Filename: resizedFilename,
		MD5Hash:  md5Hash,
	}

	return uploadedImage, nil
}

func isSupportedImageFormat(c *fiber.Ctx) error {
	allowedExtensions := []string{".png", ".jpg", ".jpeg"}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Failed to retrieve uploaded file:", err)
		return err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return nil
		}
	}

	return err
}



func checkDuplicateImage(filename string, path string) error {
	filepath := filepath.Join(path, filename)
	_, err := os.Stat(filepath)
	return err
}
func encodeImage(file io.Writer, img image.Image, filename string) error {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".png":
		return png.Encode(file, img)
	case ".jpg", ".jpeg":
		return jpeg.Encode(file, img, nil)
	default:
		return errors.NewNotFoundError("Unsupported image format: " + ext)
	}
}


func DeleteImage(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
