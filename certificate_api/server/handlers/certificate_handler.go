package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/certificate_api/server/models"
)

var (
	certificates []models.Certificate
	mutex        sync.Mutex
)

func sendJSONResponse(c *gin.Context, data interface{}, statusCode int) {
	c.JSON(statusCode, data)
}

func GetCertificateByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendJSONResponse(c, gin.H{"error": "Invalid certificate ID"}, http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for _, cert := range certificates {
		if cert.ID == id {
			sendJSONResponse(c, cert, http.StatusOK)
			return
		}
	}
	sendJSONResponse(c, gin.H{"error": "Certificate not found"}, http.StatusNotFound)
}

func CreateCertificate(c *gin.Context) {
	var cert models.Certificate
	if err := c.ShouldBindJSON(&cert); err != nil {
		sendJSONResponse(c, gin.H{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	cert.ID = len(certificates) + 1
	certificates = append(certificates, cert)

	sendJSONResponse(c, cert, http.StatusCreated)
}

func GetAllCertificates(c *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()
	sendJSONResponse(c, certificates, http.StatusOK)
}

func UpdateCertificate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendJSONResponse(c, gin.H{"error": "Invalid certificate ID"}, http.StatusBadRequest)
		return
	}

	var updatedCert models.Certificate
	if err := c.ShouldBindJSON(&updatedCert); err != nil {
		sendJSONResponse(c, gin.H{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i, cert := range certificates {
		if cert.ID == id {
			updatedCert.ID = id
			certificates[i] = updatedCert
			sendJSONResponse(c, updatedCert, http.StatusOK)
			return
		}
	}
	sendJSONResponse(c, gin.H{"error": "Certificate not found"}, http.StatusNotFound)
}

func UploadCertificateData(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		sendJSONResponse(c, gin.H{"error": "File upload failed"}, http.StatusBadRequest)
		return
	}

	savePath := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		sendJSONResponse(c, gin.H{"error": "Failed to save file"}, http.StatusInternalServerError)
		return
	}

	certData, err := readCSVToCertificates(savePath)
	if err != nil {
		sendJSONResponse(c, gin.H{"error": "Failed to read CSV file"}, http.StatusInternalServerError)
		return
	}

	mutex.Lock()
	certificates = append(certificates, certData...)
	mutex.Unlock()

	sendJSONResponse(c, gin.H{"message": "Certificates added successfully"}, http.StatusOK)
}

func readCSVToCertificates(filePath string) ([]models.Certificate, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("CSV file must have at least one data row")
	}

	var certs []models.Certificate
	for i, row := range records[1:] {
		if len(row) < 6 {
			continue
		}
		cert := models.Certificate{
			ID:         len(certificates) + i + 1,
			Name:       row[0],
			Course:     row[1],
			IssuedTo:   row[2],
			IssueDate:  row[3],
			ExpiryDate: row[4],
			Issuer:     row[5],
			Content:    fmt.Sprintf("Certificate of Completion awarded to %s for successfully completing %s, issued by %s", row[2], row[1], row[5]),
		}
		certs = append(certs, cert)
	}

	return certs, nil
}
