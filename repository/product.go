package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/AndiGanesha/gamified/application"
	"github.com/AndiGanesha/gamified/model"
)

// define interface
type IProductRepository interface {
	GetProducts() ([]model.Product, error)
	BuyProduct(buyerID int, productID string, quantityToBuy int) error
	GetSales(userId int) ([]model.SalesTransaction, error)
}

// define a scallable struct if needed in the future
type ProductRepository struct {
	DB *sql.DB
}

// create new product func
func NewProductRepository(app *application.App) IProductRepository {
	return &ProductRepository{
		DB: app.DB,
	}
}

func (r *ProductRepository) GetProducts() (products []model.Product, err error) {
	// Use Query to fetch multiple rows, not QueryRow.
	rows, err := r.DB.Query("SELECT id, name, quantity FROM product")
	if err != nil {
		log.Printf("Error querying products: %v", err)
		return nil, err
	}
	// Defer closing the rows to ensure the connection is released.
	defer rows.Close()

	// Loop through each row in the result set.
	for rows.Next() {
		var p model.Product

		// Scan the values from the current row into the fields of the Product struct.
		if err := rows.Scan(&p.Id, &p.Name, &p.Quantity); err != nil {
			log.Printf("Error scanning product row: %v", err)
			// Decide if you want to continue or stop. For now, we'll stop and return the error.
			return nil, err
		}
		products = append(products, p)
	}

	// After the loop, check for any errors that occurred during iteration.
	if err = rows.Err(); err != nil {
		log.Printf("Error iterating product rows: %v", err)
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) BuyProduct(buyerID int, productID string, quantityToBuy int) error {
	// A transaction is a sequence of operations performed as a single logical unit of work.
	tx, err := r.DB.Begin()
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return err
	}
	defer tx.Rollback()

	var currentQuantity int
	query := "SELECT quantity FROM product WHERE id = ? FOR UPDATE"
	err = tx.QueryRow(query, productID).Scan(&currentQuantity)

	// Handle cases where the product might not exist.
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("product with id %s not found", productID)
		}
		log.Printf("Failed to query product for update: %v", err)
		return err
	}

	// Check if there is enough quantity in stock.
	if currentQuantity < quantityToBuy {
		return errors.New("insufficient stock available")
	}

	// If stock is sufficient, update the product's quantity.
	newQuantity := currentQuantity - quantityToBuy
	updateQuery := "UPDATE product SET quantity = ? WHERE id = ?"
	_, err = tx.Exec(updateQuery, newQuantity, productID)
	if err != nil {
		log.Printf("Failed to update product quantity: %v", err)
		return err
	}

	// Record the transaction in the 'sales_transaction' table.
	// This creates a permanent record of the sale.
	insertQuery := "INSERT INTO sales_transaction (user_id, product_id, quantity) VALUES (?, ?, ?)"
	_, err = tx.Exec(insertQuery, buyerID, productID, quantityToBuy)
	if err != nil {
		log.Printf("Failed to insert into sales_transaction: %v", err)
		return err
	}

	// Add 10 experience points to the user for the successful transaction.
	// This is part of the gamification requirement and is included in the transaction.
	updateUserXPQuery := "UPDATE user SET experience = experience + 10 WHERE id = ?"
	_, err = tx.Exec(updateUserXPQuery, buyerID)
	if err != nil {
		log.Printf("Failed to update user experience: %v", err)
		return err
	}

	// all previous steps were successful, commit the transaction.
	// This makes all the changes permanent in the database.
	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return err
	}

	log.Printf("Successfully processed purchase for product %s by user %d.\n", productID, buyerID)
	return nil
}

func (r *ProductRepository) GetSales(userId int) ([]model.SalesTransaction, error) {
	// Initialize an empty slice to hold the sales transactions.
	var transactions []model.SalesTransaction

	// The SQL query to select all sales for a given user.
	query := "SELECT id, user_id, product_id, quantity FROM sales_transaction WHERE user_id = ?"

	// Use Query to fetch multiple rows.
	rows, err := r.DB.Query(query, userId)
	if err != nil {
		log.Printf("Error querying sales for user %d: %v", userId, err)
		return nil, err // Return nil slice and the error.
	}
	// Defer closing the rows to ensure the database connection is released.
	defer rows.Close()

	// Loop through each row in the result set.
	for rows.Next() {
		// Create a new SalesTransaction struct for each row.
		var t model.SalesTransaction

		// Scan the values from the current row into the fields of the struct.
		// The order of fields in Scan must match the order in the SELECT statement.
		if err := rows.Scan(&t.BuyerId, &t.ProductId, &t.Quantity); err != nil {
			log.Printf("Error scanning sales transaction row: %v", err)
			// Decide if you want to continue or stop. For now, we'll stop and return the error.
			return nil, err
		}

		// Append the successfully scanned transaction to the slice.
		transactions = append(transactions, t)
	}

	// After the loop, check for any errors that occurred during iteration.
	if err = rows.Err(); err != nil {
		log.Printf("Error iterating sales transaction rows: %v", err)
		return nil, err
	}

	// If the loop completes without error, return the slice of transactions.
	return transactions, nil
}
