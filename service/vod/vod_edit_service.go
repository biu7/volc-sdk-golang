// Code generated by protoc-gen-volcengine-sdk
// source: VodEditService
// DO NOT EDIT!

package vod

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/pkg/errors"

	"github.com/volcengine/volc-sdk-golang/service/base/models/base"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/response"
)

// SubmitDirectEditTaskAsync
/*
 * @param *request.VodSubmitDirectEditTaskAsyncRequest
 * @return *response.VodSubmitDirectEditTaskAsyncResponse, int, error
 */
func (p *Vod) SubmitDirectEditTaskAsync(req *request.VodSubmitDirectEditTaskAsyncRequest) (*response.VodSubmitDirectEditTaskAsyncResponse, int, error) {
	reqMap := make(map[string]interface{})
	req.ProtoReflect().Range(func(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		if !value.IsValid() {
			return true
		}
		if descriptor.Name() != "EditParam" {
			reqMap[string(descriptor.Name())] = value.Interface()
		} else {
			reqMap[string(descriptor.Name())] = json.RawMessage(value.Bytes())
		}
		return true
	})
	jsonData, _ := json.Marshal(reqMap)
	respBody, status, err := p.Json("SubmitDirectEditTaskAsync", url.Values{}, string(jsonData))

	output := &response.VodSubmitDirectEditTaskAsyncResponse{}
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	errUnmarshal := unmarshaler.Unmarshal(respBody, output)
	if err != nil || status != http.StatusOK {
		// if exist http err,check whether the respBody's type is defined struct,
		// if it is ,
		// return struct,
		// otherwise return nil body
		// if httpCode is not 200,check whether the respBody's type is defined struct,
		// if it is ,
		// use errorCode as err and return struct,
		// otherwise use respBody string as error and return
		if errUnmarshal != nil || len(output.GetResponseMetadata().GetError().GetCode()) == 0 {
			if err == nil {
				err = errors.New(string(respBody))
			}
			return nil, status, err
		} else {
			return output, status, errors.New(output.GetResponseMetadata().GetError().GetCode())
		}
	}
	return output, status, nil
}

// GetDirectEditResult
/*
 * @param *request.VodGetDirectEditResultRequest
 * @return *response.VodGetDirectEditResultResponse, int, error
 */
func (p *Vod) GetDirectEditResult(req *request.VodGetDirectEditResultRequest) (*response.VodGetDirectEditResultResponse, int, error) {
	marshaler := protojson.MarshalOptions{
		Multiline:       false,
		AllowPartial:    false,
		UseProtoNames:   true,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
	}
	jsonData := marshaler.Format(req)
	respBody, status, err := p.Json("GetDirectEditResult", url.Values{}, jsonData)
	resultMap := struct {
		ResponseMetadata *base.ResponseMetadata
		Result           []map[string]interface{}
	}{}
	err = json.Unmarshal(respBody, &resultMap)
	if err != nil {
		return nil, status, errors.New(string(respBody))
	}
	output := &response.VodGetDirectEditResultResponse{}
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	errUnmarshal := unmarshaler.Unmarshal(respBody, output)
	if err != nil || status != http.StatusOK {
		// if exist http err,check whether the respBody's type is defined struct,
		// if it is ,
		// return struct,
		// otherwise return nil body
		// if httpCode is not 200,check whether the respBody's type is defined struct,
		// if it is ,
		// use errorCode as err and return struct,
		// otherwise use respBody string as error and return
		if errUnmarshal != nil || len(output.GetResponseMetadata().GetError().GetCode()) == 0 {
			if err == nil {
				err = errors.New(string(respBody))
			}
			return nil, status, err
		} else {
			return output, status, errors.New(output.GetResponseMetadata().GetError().GetCode())
		}
	}
	return output, status, nil
}

// GetDirectEditProgress
/*
 * @param *request.VodGetDirectEditProgressRequest
 * @return *response.VodGetDirectEditProgressResponse, int, error
 */
func (p *Vod) GetDirectEditProgress(req *request.VodGetDirectEditProgressRequest) (*response.VodGetDirectEditProgressResponse, int, error) {
	query := url.Values{}
	form := url.Values{}
	marshaler := protojson.MarshalOptions{
		Multiline:       false,
		AllowPartial:    false,
		UseProtoNames:   true,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
	}
	jsonData := marshaler.Format(req)
	reqMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonData), &reqMap)
	if err != nil {
		return nil, 0, err
	}
	for k, v := range reqMap {
		var sv string
		switch ov := v.(type) {
		case string:
			sv = ov
		case int:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int8:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint8:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int16:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint16:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int32:
			sv = strconv.FormatInt(int64(ov), 10)
		case uint32:
			sv = strconv.FormatUint(uint64(ov), 10)
		case int64:
			sv = strconv.FormatInt(ov, 10)
		case uint64:
			sv = strconv.FormatUint(ov, 10)
		case bool:
			sv = strconv.FormatBool(ov)
		case float32:
			sv = strconv.FormatFloat(float64(ov), 'f', -1, 32)
		case float64:
			sv = strconv.FormatFloat(ov, 'f', -1, 64)
		case []byte:
			sv = string(ov)
		default:
			v2, e2 := json.Marshal(ov)
			if e2 != nil {
				return nil, 0, e2
			}
			sv = string(v2)
		}
		query.Set(k, sv)
		form.Add(k, sv)
	}

	respBody, status, err := p.Query("GetDirectEditProgress", query)
	resultMap := struct {
		ResponseMetadata *base.ResponseMetadata
		Result           int32
	}{}
	err = json.Unmarshal(respBody, &resultMap)
	if err != nil {
		return nil, status, errors.New(string(respBody))
	}
	result := struct {
		ResponseMetadata *base.ResponseMetadata
		Result           *struct {
			Result int32
		}
	}{}
	result.ResponseMetadata = resultMap.ResponseMetadata
	result.Result = &struct {
		Result int32
	}{Result: resultMap.Result}
	marshal, _ := json.Marshal(result)
	output := &response.VodGetDirectEditProgressResponse{}
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	errUnmarshal := unmarshaler.Unmarshal(marshal, output)
	if err != nil || status != http.StatusOK {
		// if exist http err,check whether the respBody's type is defined struct,
		// if it is ,
		// return struct,
		// otherwise return nil body
		// if httpCode is not 200,check whether the respBody's type is defined struct,
		// if it is ,
		// use errorCode as err and return struct,
		// otherwise use respBody string as error and return
		if errUnmarshal != nil || len(output.GetResponseMetadata().GetError().GetCode()) == 0 {
			if err == nil {
				err = errors.New(string(respBody))
			}
			return nil, status, err
		} else {
			return output, status, errors.New(output.GetResponseMetadata().GetError().GetCode())
		}
	}
	return output, status, nil
}
