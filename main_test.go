package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/ameykpatil/go-uploader/models"
	. "github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	if flag.Lookup("test.v") != nil {
		gin.SetMode(gin.ReleaseMode)
	}
}

func performRequest(router *gin.Engine, method, url string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func Test(t *testing.T) {
	//id := "123"

	router := getMainEngine()
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("GET /ping", func() {

		g.It("Should return pong successfully", func() {
			w := performRequest(router, "GET", "/ping", nil)
			Expect(w.Code).To(Equal(200))
			Expect(w.Body.String()).To(Equal("pong"))
		})

		g.It("Should throw error for non existing API", func() {
			w := performRequest(router, "GET", "/pong", nil)
			Expect(w.Code).To(Equal(404))
		})
	})

	g.Describe("POST /asset", func() {

		g.It("Should add asset successfully & return upload url", func() {
			w := performRequest(router, "POST", "/asset", nil)
			Expect(w.Code).To(Equal(200))
			var body map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &body)
			Expect(err).To(BeNil())
			Expect(body["id"]).ToNot(Equal(""))
			Expect(body["upload_url"]).ToNot(Equal(""))

			assetID := body["id"]
			savedAsset, _ := models.GetAssetByID(assetID)
			Expect(savedAsset.ID).To(Equal(assetID))
			Expect(savedAsset.Status).To(Equal("initiated"))

			models.RemoveByID(assetID)
		})
	})

	g.Describe("PUT /asset/:assetId", func() {

		id := "put-asset"

		g.After(func() {
			models.RemoveByID(id)
		})

		g.It("Should insert asset if not present already", func() {
			reqBody := bytes.NewBuffer([]byte("{\"status\":\"initiated\"}"))
			w := performRequest(router, "PUT", "/asset/"+id, reqBody)
			Expect(w.Code).To(Equal(200))
			var body map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &body)
			Expect(err).To(BeNil())
			Expect(body["id"]).To(Equal(id))
			Expect(body["status"]).To(Equal("initiated"))

			updatedAsset, _ := models.GetAssetByID(id)
			Expect(updatedAsset.Status).To(Equal("initiated"))
		})

		g.It("Should update asset successfully", func() {
			reqBody := bytes.NewBuffer([]byte("{\"status\":\"uploaded\"}"))
			w := performRequest(router, "PUT", "/asset/"+id, reqBody)
			Expect(w.Code).To(Equal(200))
			var body map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &body)
			Expect(err).To(BeNil())
			Expect(body["id"]).To(Equal(id))
			Expect(body["status"]).To(Equal("uploaded"))

			updatedAsset, _ := models.GetAssetByID(id)
			Expect(updatedAsset.Status).To(Equal("uploaded"))
		})
	})

	g.Describe("GET /asset/:assetId", func() {

		id := "123"

		g.After(func() {
			models.RemoveByID(id)
		})

		g.It("Should return error for non-existing asset", func() {
			w := performRequest(router, "GET", "/asset/"+id, nil)
			Expect(w.Code).To(Equal(400))
			var body map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &body)
			Expect(err).To(BeNil())
			Expect(body["message"]).To(Equal("not found"))
		})

		g.It("Should return error if the status of the asset is not uploaded", func() {
			w := performRequest(router, "GET", "/asset/1533972640563", nil)
			Expect(w.Code).To(Equal(400))
			var body map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &body)
			Expect(err).To(BeNil())
			Expect(body["message"]).To(Equal("asset is not completely uploaded yet"))
		})

		g.It("Should return download url for an uploaded asset", func() {
			w := performRequest(router, "GET", "/asset/1533912247399", nil)
			Expect(w.Code).To(Equal(200))
			var body map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &body)
			Expect(err).To(BeNil())
			Expect(body["download_url"]).ToNot(BeEmpty())
			Expect(body["download_url"]).To(ContainSubstring("1533912247399"))
		})

		g.It("Should return download url for an uploaded asset with given timeout", func() {
			w := performRequest(router, "GET", "/asset/1533912247399?timeout=100", nil)
			Expect(w.Code).To(Equal(200))
			var body map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &body)
			Expect(err).To(BeNil())
			Expect(body["download_url"]).ToNot(BeEmpty())
			Expect(body["download_url"]).To(ContainSubstring("1533912247399"))
			Expect(body["download_url"]).To(ContainSubstring("X-Amz-Expires=100"))
		})

		g.It("Should return download url with default timeout if not given", func() {
			w := performRequest(router, "GET", "/asset/1533912247399", nil)
			Expect(w.Code).To(Equal(200))
			var body map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &body)
			Expect(err).To(BeNil())
			Expect(body["download_url"]).ToNot(BeEmpty())
			Expect(body["download_url"]).To(ContainSubstring("1533912247399"))
			Expect(body["download_url"]).To(ContainSubstring("X-Amz-Expires=60"))
		})
	})
}
