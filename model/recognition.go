package model

import (
	"github.com/glvd/seed/model"
)

// RecognitionStatus ...
type RecognitionStatus int

// StatusReady ...
const StatusReady RecognitionStatus = 0

// StatusFinished ...
const StatusFinished RecognitionStatus = iota

// StatusFailed ...
const StatusFailed RecognitionStatus = iota

// Recognition ...
type Recognition struct {
	model.Model `xorm:"extends"`
	VID         string
	FilePath    string
	Status      RecognitionStatus
	Checksum    string
	Bangumi     string
	Role        []string
}

func init() {
	model.RegisterTable(Recognition{})
}

// UpdateFromVideo ...
func (r *Recognition) UpdateFromVideo() (e error) {
	video, e := model.FindVideo(nil, r.Bangumi)
	if e != nil {
		return e
	}
	r.Role = video.Role
	r.VID = video.ID
	return r.AddOrUpdate()
}

// AddOrUpdate ...
func (r *Recognition) AddOrUpdate() (e error) {
	var tmp Recognition
	b, _ := model.DB().Where("checksum = ?", r.Checksum).Get(&tmp)
	if r.ID != "" || b {
		r.Version++
		_, e := model.DB().Update(r)
		if e != nil {
			return e
		}
		return nil
	}
	_, e = model.DB().InsertOne(r)
	return
}
