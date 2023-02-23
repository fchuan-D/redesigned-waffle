package service

import (
	"errors"
	"soft-pro/dao"
	"soft-pro/entity"
	"soft-pro/resp"
	"strconv"
)

func GetOrderByID(orderID any) (entity.Order, error) {
	order := dao.GetOrderByID(orderID)
	if order.ID == 0 {
		return order, errors.New(resp.NotFindMsg)
	}
	return order, nil
}

func GetOrdersByID(userID any) ([]entity.Order, error) {
	orders := dao.GetOrdersByUserID(userID)
	if orders[0].ID == 0 {
		return orders, errors.New(resp.NotFindMsg)
	}
	return orders, nil
}

func CreateOrder(order entity.Order) error {
	err := dao.CreateOrder(order)
	if err != nil {
		return errors.New(resp.NotOkMsg)
	}
	return nil
}
func DeleteOrder(orderID any) error {
	err := dao.DeleteOrder(orderID)
	if err != nil {
		return errors.New(resp.NotOkMsg)
	}
	return nil
}

func PayOrder(orderID any, UserID any) error {
	exist := dao.GetOrderByID(orderID)
	if exist.Status == 1 {
		return errors.New(resp.OrderAbortMsg)
	}

	err := dao.PayOrder(orderID)
	if err != nil {
		return errors.New(resp.NotOkMsg)
	}

	order := dao.GetOrderByID(orderID)
	user := dao.GetUserByID(UserID)
	amount, err := strconv.ParseFloat(order.Amount, 64)
	if err != nil {
		return errors.New(resp.NotOkMsg)
	}

	if user.Bal < amount {
		return errors.New(resp.NotEnoughMsg)
	}
	user.Bal -= amount

	err = dao.UpdateBal(user.ID, user.Bal)
	if err != nil {
		return errors.New(resp.NotOkMsg)
	}
	return nil
}

func AbortOrder(orderID any) error {
	err := dao.AbortOrder(orderID)
	if err != nil {
		return errors.New(resp.NotOkMsg)
	}
	return nil
}
