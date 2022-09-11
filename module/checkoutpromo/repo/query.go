package repository

import (
	"checkoutpromo/graph/model"
	"checkoutpromo/internal/db"
	"log"
)

func ProductList() ([]*model.Product, error) {
	stmt, err := db.Db.Prepare(`SELECT product_id, product_sku, product_name, product_price, product_qty FROM products ORDER BY product_name `)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	var productList []*model.Product

	for rows.Next() {
		var product model.Product
		rows.Scan(&product.ProductID, &product.Sku, &product.ProductName, &product.ProductPrice, &product.ProductQty)

		productList = append(productList, &product)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	defer rows.Close()

	return productList, err
}

func CustomerList() ([]*model.Customer, error) {
	stmt, err := db.Db.Prepare(`SELECT customer_id, customer_name, customer_email FROM customers ORDER BY customer_name `)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	var customerList []*model.Customer

	for rows.Next() {
		var customer model.Customer
		rows.Scan(&customer.CustomerID, &customer.CustomerName, &customer.CustomerEmail)

		customerList = append(customerList, &customer)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	defer rows.Close()

	return customerList, err
}

func CartList(customerId string) ([]*model.Cart, error) {
	stmt, err := db.Db.Prepare(`SELECT sc.cart_id , sc.customer_id, sc.product_id , p.product_sku , p.product_name , p.product_price, sc.qty , sc.qty*p.product_price as total_price, sc.created_at  FROM shopping_carts sc
	left join products p on p.product_id = sc.product_id WHERE sc.customer_id =?  `)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(customerId)

	if err != nil {
		return nil, err
	}

	var cartList []*model.Cart

	for rows.Next() {
		var cart model.Cart
		rows.Scan(&cart.CartID, &cart.CustomerID, &cart.ProductID, &cart.Sku, &cart.ProductName, &cart.UnitPrice, &cart.Quantity, &cart.TotalPrice, &cart.CreatedAt)

		cartList = append(cartList, &cart)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	defer rows.Close()

	return cartList, err
}

func OrderList(customerId string) ([]*model.Order, error) {
	stmt, err := db.Db.Prepare(`SELECT order_num, customer_id FROM orders WHERE customer_id =?  `)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(customerId)

	if err != nil {
		return nil, err
	}

	var orderList []*model.Order

	for rows.Next() {
		var order model.Order
		rows.Scan(&order.OrderNum, &order.CustomerID)

		orderList = append(orderList, &order)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	defer rows.Close()

	return orderList, err
}

func OrderDetail(orderNum string) ([]*model.OrderDetail, error) {
	stmt, err := db.Db.Prepare(`SELECT o.order_id,o.order_num, o.customer_id , o.product_id , p.product_sku , p.product_name , p.product_price , o.qty , o.total_price  FROM orders o 
	LEFT JOIN products p on p.product_id = o.product_id where o.order_num = ?`)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(orderNum)

	if err != nil {
		return nil, err
	}

	var orderDetail []*model.OrderDetail

	for rows.Next() {
		var order model.OrderDetail
		rows.Scan(&order.OrderID, &order.OrderNum, &order.CustomerID, &order.ProductID, &order.Sku, &order.ProductName, &order.UnitPrice, &order.Quantity, &order.TotalPrice)

		orderDetail = append(orderDetail, &order)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	defer rows.Close()

	return orderDetail, err
}

func CheckStock(productID string) (qty int, err error) {
	stmt, err := db.Db.Prepare(`SELECT product_qty FROM products WHERE product_id= ?`)

	if err != nil {
		return 0, err
	}

	row, err := stmt.Query(productID)

	if row.Next() {
		err = row.Scan(&qty)

		if err != nil {
			return 0, err
		}
	}

	if err != nil {
		return 0, err
	}

	return qty, err
}

func CheckCartStock(customerID, productID string) (qty int, err error) {
	stmt, err := db.Db.Prepare(`SELECT qty FROM shopping_carts WHERE customer_id= ? AND product_id = ?`)

	if err != nil {
		return 0, err
	}

	row, err := stmt.Query(customerID, productID)

	if row.Next() {
		err = row.Scan(&qty)

		if err != nil {
			return 0, err
		}
	}

	if err != nil {
		return 0, err
	}

	return qty, err
}
