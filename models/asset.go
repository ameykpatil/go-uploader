package models

import (
	"errors"
	"github.com/ameykpatil/go-uploader/constants"
	"github.com/ameykpatil/go-uploader/utils/helper"
	"github.com/ameykpatil/go-uploader/utils/logger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

var collection *mgo.Collection

func init() {
	session, err := mgo.Dial(constants.Env.MongoURL)
	if err != nil {
		panic(err)
	}
	collection = session.DB(constants.Env.MongoDBName).C(constants.Env.MongoCollection)
}

//Asset is a struct for Asset details
type Asset struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Status    string    `bson:"status,omitempty" json:"status,omitempty"`
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

//NewAsset is a constructor for Asset
func NewAsset(id string, status string) *Asset {
	return &Asset{
		ID:        id,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

//Add adds corresponding asset to db
func (asset Asset) Add() error {
	if asset.ID == "" {
		return errors.New("id of the asset should not be empty")
	}
	err := collection.Insert(&asset)
	if err != nil {
		logger.Err(asset.ID, "error saving asset to db", err)
		return err
	}
	return nil
}

//Update updates a status of an asset in a db
func (asset Asset) Update() (*mgo.ChangeInfo, error) {
	selector := bson.M{"_id": asset.ID}
	update := bson.M{"$set": bson.M{"status": asset.Status, "updatedAt": time.Now()}}
	changeInfo, err := collection.Upsert(selector, update)
	if err != nil {
		logger.Err(asset.ID, "error saving asset to db", err)
		return nil, err
	}
	return changeInfo, nil
}

//GetAssetByID returns asset for a given id
func GetAssetByID(ID string) (*Asset, error) {
	asset := &Asset{}
	err := collection.Find(bson.M{"_id": ID}).One(asset)
	if err != nil {
		logger.Err(asset.ID, "error saving asset to db", err)
		return nil, err
	}
	return asset, nil
}

//GetNewAssetID returns new asset id based on timestamp
func GetNewAssetID() string {
	currentTime := helper.GetCurrentTime()
	assetID := strconv.FormatInt(currentTime, 10)
	return assetID
}

//RemoveByID removes asset by ID
func RemoveByID(ID string) error {
	return collection.RemoveId(ID)
}
