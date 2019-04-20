package tests

import (
	"encoding/json"
	"testing"

	"gitlab.goiot.net/sde-base/front-api/utils"
)

func TestMap(t *testing.T) {
	b, err := utils.FileGetContents("data/jsonmap.txt")
	if err != nil {
		t.Error(err)
	}
	m := map[string]interface{}{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Error(err)
	}
	str := utils.JSONWithMap(m)
	utils.FilePutContents(str, "result/jsonmap.txt")
	t.Error(str)
}

func TestJSONSlice(t *testing.T) {
	b, err := utils.FileGetContents("data/jsonslice.txt")
	if err != nil {
		t.Error(err)
	}
	m := []interface{}{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Error(err)
	}
	str := utils.JSONWithSlice(m)
	utils.FilePutContents(str, "result/jsonslice.txt")
	t.Error(str)
}

func TestSignature(t *testing.T) {
	b, err := utils.FileGetContents("data/jsonmap.txt")
	if err != nil {
		t.Error(err)
	}
	str := utils.Sign(string(b), "0f90529eeccc1539b5cf6f0101a97ff2")
	utils.FilePutContents(str, "result/signature.txt")
	t.Error(str)
}
