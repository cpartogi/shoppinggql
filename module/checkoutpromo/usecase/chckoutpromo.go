package usecase

import (
	"checkoutpromo/graph/model"
	repository "checkoutpromo/module/checkoutpromo/repo"
	"errors"
)

func ProductList() ([]*model.Product, error) {
	productList, err := repository.ProductList()

	if err != nil {
		return nil, err
	}
	return productList, nil
}

func CustomerList() ([]*model.Customer, error) {
	customerList, err := repository.CustomerList()

	if err != nil {
		return nil, err
	}
	return customerList, nil
}

func CartList(customerId string) ([]*model.Cart, error) {

	if customerId == "" {
		return nil, errors.New("customer_id must be filled")
	}

	cartList, err := repository.CartList(customerId)

	if err != nil {
		return nil, err
	}
	return cartList, nil
}

func OrderList(customerId string) ([]*model.Order, error) {

	if customerId == "" {
		return nil, errors.New("customer_id must be filled")
	}

	orderList, err := repository.OrderList(customerId)

	if err != nil {
		return nil, err
	}
	return orderList, nil
}

func OrderDetail(orderNum string) ([]*model.OrderDetail, error) {

	if orderNum == "" {
		return nil, errors.New("order_num must be filled")
	}

	orderDetail, err := repository.OrderDetail(orderNum)

	if err != nil {
		return nil, err
	}
	return orderDetail, nil
}

func AddCart(customerID string, productID string, qty int) (res *model.ResponseData, err error) {

	//parameter validation
	if customerID == "" {
		return nil, errors.New("customer id must be filled")
	}

	if productID == "" {
		return nil, errors.New("product id must be filled")
	}

	if qty == 0 {
		return nil, errors.New("qty must be more than 0")
	}

	// cek quantity on cart
	cart_qty, err := repository.CheckCartStock(customerID, productID)

	if err != nil {
		return nil, err
	}

	var qty_update int
	if cart_qty != 0 {
		qty_update = qty + cart_qty
	} else {
		qty_update = qty
	}

	product_stock, err := repository.CheckStock(productID)

	if err != nil {
		return nil, err
	}

	var addCart *model.ResponseData
	if qty < product_stock {

		// update product stock quantity
		err = repository.ReduceStock(productID, qty)
		if err != nil {
			return nil, err
		}

		if cart_qty != 0 {
			// update shopping cart
			addCart, err = repository.UpdateCart(customerID, productID, qty_update)

			if err != nil {
				return nil, err
			}

		} else {
			// add to shopping cart
			addCart, err = repository.AddCart(customerID, productID, qty)

			if err != nil {
				return nil, err
			}
		}
	} else {
		return nil, errors.New("product stock is not available")
	}

	return addCart, nil
}

func DeleteCart(customerID string, productID string) (res *model.ResponseData, err error) {
	//parameter validation
	if customerID == "" {
		return nil, errors.New("customer id must be filled")
	}

	if productID == "" {
		return nil, errors.New("product id must be filled")
	}

	// cek quantity on cart
	cart_qty, err := repository.CheckCartStock(customerID, productID)

	if err != nil {
		return nil, err
	}

	// update product stock quantity
	err = repository.AddStock(productID, cart_qty)
	if err != nil {
		return nil, err
	}

	// remove from shopping cart
	delCart, err := repository.DeleteCart(customerID, productID)

	if err != nil {
		return nil, err
	}

	return delCart, nil

}

func Checkout(customerID string) (orderNumber string, err error) {
	//parameter validation
	if customerID == "" {
		return orderNumber, errors.New("customer id must be filled")
	}

	orderNum, err := repository.Checkout(customerID)

	if err != nil {
		return orderNumber, err
	}

	return orderNum, nil
}
