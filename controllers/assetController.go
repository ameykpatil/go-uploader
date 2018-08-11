package controllers

import (
	"github.com/ameykpatil/go-uploader/models"
	"github.com/ameykpatil/go-uploader/utils/s3"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//AssetBody is a struct for binding body of update request
type AssetBody struct {
	Status string `form:"status" json:"status,omitempty" binding:"required"`
}

//AddAsset add an asset to db & returns signed url for uploading
func AddAsset(c *gin.Context) {
	assetID := models.GetNewAssetID()
	signedURL, err := s3.SignedURLForPut(assetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	asset := models.NewAsset(assetID, "initiated")
	asset.Add()

	c.JSON(http.StatusOK, gin.H{"upload_url": signedURL, "id": assetID})
}

//UpdateAsset updates an asset in a db
func UpdateAsset(c *gin.Context) {
	assetID := c.Param("assetId")

	var assetBody AssetBody
	if err := c.Bind(&assetBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	asset := models.NewAsset(assetID, assetBody.Status)

	_, err := asset.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, asset)
}

//GetAsset returns an asset along with signed url for downloading
func GetAsset(c *gin.Context) {

	timeoutString := c.DefaultQuery("timeout", "60")
	timeout, err := strconv.ParseInt(timeoutString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	assetID := c.Param("assetId")
	asset, err := models.GetAssetByID(assetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if asset.Status != "uploaded" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "asset is not completely uploaded yet"})
		return
	}

	signedURL, err := s3.SignedURLForGet(assetID, timeout)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"download_url": signedURL})
}
