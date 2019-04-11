package main

import "github.com/girlvr/seed/model"

func searchVideo(s string) *model.Video {
	video := &model.Video{}
	if b, err := model.FindVideo(s, video); err != nil || !b {
		return nil
	}
	return video
}
