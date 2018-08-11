package models

import (
	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"testing"
)

func Test(t *testing.T) {
	id := "123"

	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Add", func() {

		g.BeforeEach(func() {
			RemoveByID(id)
		})

		g.AfterEach(func() {
			RemoveByID(id)
		})

		g.It("Should add asset to db successfully", func() {
			asset := NewAsset(id, "initialized")
			err := asset.Add()
			Expect(err).To(BeNil())

			savedAsset, _ := GetAssetByID(id)
			Expect(savedAsset.ID).To(Equal(id))
			Expect(savedAsset.Status).To(Equal("initialized"))
		})

		g.It("Should throw error for empty asset", func() {
			asset := &Asset{}
			err := asset.Add()
			Expect(err.Error()).To(Equal("id of the asset should not be empty"))
		})
	})

	g.Describe("Update", func() {

		g.Before(func() {
			RemoveByID(id)
		})

		g.After(func() {
			RemoveByID(id)
		})

		g.It("Should insert asset details if asset is not present", func() {
			asset := NewAsset(id, "uploaded")
			changeInfo, err := asset.Update()
			Expect(err).To(BeNil())
			Expect(changeInfo.UpsertedId).To(Equal(id))

			savedAsset, _ := GetAssetByID(id)
			Expect(savedAsset.ID).To(Equal(id))
			Expect(savedAsset.Status).To(Equal("uploaded"))
		})

		g.It("Should update asset details if asset is present in the db", func() {
			asset := NewAsset(id, "test-status")
			changeInfo, err := asset.Update()
			Expect(err).To(BeNil())
			Expect(changeInfo.Updated).To(Equal(1))

			updatedAsset, _ := GetAssetByID(id)
			Expect(updatedAsset.ID).To(Equal(id))
			Expect(updatedAsset.Status).To(Equal("test-status"))
		})
	})

	g.Describe("GetAssetByID", func() {

		g.Before(func() {
			asset := NewAsset(id, "initiated")
			asset.Add()
		})

		g.After(func() {
			RemoveByID(id)
		})

		g.It("Should get asset details for given id", func() {
			asset, err := GetAssetByID(id)
			Expect(err).To(BeNil())
			Expect(asset.ID).To(Equal(id))
			Expect(asset.Status).To(Equal("initiated"))
		})

		g.It("Should throw error while getting non existing id", func() {
			_, err := GetAssetByID("test")
			Expect(err.Error()).To(Equal("not found"))
		})
	})
}
