package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/adityaupadhyay0/certificate_api/server/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	certRoutes := r.Group("/certificates")
	{
		certRoutes.POST("/import", handlers.UploadCertificateData)
		certRoutes.GET("/:id", handlers.GetCertificateByID)
		certRoutes.POST("/", handlers.CreateCertificate)
		certRoutes.GET("/", handlers.GetAllCertificates)
		certRoutes.PUT("/:id", handlers.UpdateCertificate)
	}
}
