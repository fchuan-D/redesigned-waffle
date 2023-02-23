package test

import (
	"soft-pro/conf"
	"soft-pro/dao"
	"soft-pro/service"
	"testing"
)

func TestCheckLoginUser(t *testing.T) {

}

func TestCheckRegisterUser(t *testing.T) {

}

func TestEarthDistance(t *testing.T) {

}

func TestGetPhoneCode(t *testing.T) {
	conf.InitConfig(".")
	dao.Init()

	service.SendPhoneCode("17882380261")
}

func TestGetUserByID(t *testing.T) {

}

func TestGetUserByPhone(t *testing.T) {

}

func TestInsertUser(t *testing.T) {

}

func TestSaveBufferToRd(t *testing.T) {

}
