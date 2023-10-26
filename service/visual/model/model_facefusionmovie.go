package model

import "github.com/volcengine/volc-sdk-golang/base"

// 请求参数

type FaceFusionMovieLogoInfo struct {
	AddLogo  bool    `json:"add_logo"`
	Position int     `json:"position"`
	Language int     `json:"language"`
	Opacity  float64 `json:"opacity"`
}

type FaceFusionMovieSubmitTaskRequest struct {
	ReqKey             string                   `json:"req_key"`
	ImageUrl           string                   `json:"image_url"`
	VideoUrl           string                   `json:"video_url"`
	EnableFaceBeautify bool                     `json:"enable_face_beautify"`
	RefImgUrl          string                   `json:"ref_img_url"`
	SourceSimilarity   string                   `json:"source_similarity"`
	LogoInfo           *FaceFusionMovieLogoInfo `json:"logo_info"`
}

type FaceFusionMovieGetResultRequest struct {
	ReqKey string `json:"req_key"`
	TaskId string `json:"task_id"`
}

type FaceFusionMovieRequest struct {
	ReqKey             string   `json:"req_key"`
	BinaryDataBase64   []string `json:"binary_data_base64"`
	ImageUrl           string   `json:"image_url"`
	VideoUrl           string   `json:"video_url"`
	EnableFaceBeautify bool     `json:"enable_face_beautify"`
	RefImgUrl          string   `json:"ref_img_url"`
}

// 响应参数

type FaceFusionMovieSubmitTaskData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	TaskId           string   `json:"task_id"`
}

type FaceFusionMovieSubmitTaskResult struct {
	ResponseMetadata *base.ResponseMetadata         `json:",omitempty"`
	RequestId        string                         `json:"request_id"`
	Code             int                            `json:"code"`
	Message          string                         `json:"message"`
	Status           int                            `json:"status"`
	TimeElapsed      string                         `json:"time_elapsed"`
	Data             *FaceFusionMovieSubmitTaskData `json:"data"`
}

type FaceFusionMovieGetResultData struct {
	AlgorithmBaseResp *struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"algorithm_base_resp"`
	BinaryDataBase64 []string `json:"binary_data_base64"`
	RespData         string   `json:"resp_data"`
	Status           string   `json:"status"`
}

type FaceFusionMovieGetResultResult struct {
	ResponseMetadata *base.ResponseMetadata        `json:",omitempty"`
	RequestId        string                        `json:"request_id"`
	Code             int                           `json:"code"`
	Message          string                        `json:"message"`
	Status           int                           `json:"status"`
	TimeElapsed      string                        `json:"time_elapsed"`
	Data             *FaceFusionMovieGetResultData `json:"data"`
}

type FaceFusionMovieResult struct {
	ResponseMetadata *base.ResponseMetadata        `json:",omitempty"`
	RequestId        string                        `json:"request_id"`
	Code             int                           `json:"code"`
	Message          string                        `json:"message"`
	Status           int                           `json:"status"`
	TimeElapsed      string                        `json:"time_elapsed"`
	Data             *FaceFusionMovieGetResultData `json:"data"`
}
