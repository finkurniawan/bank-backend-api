package model

import "testing"

func TestCustomer_TableName(t *testing.T) {
	customer := &Customer{}
	expectedTableName := "customers"
	if customer.TableName() != expectedTableName {
		t.Errorf("Expected TableName to be %s, got %s", expectedTableName, customer.TableName())
	}
}

func TestMerchant_TableName(t *testing.T) {
	merchant := &Merchant{}
	expectedTableName := "merchants"
	if merchant.TableName() != expectedTableName {
		t.Errorf("Expected TableName to be %s, got %s", expectedTableName, merchant.TableName())
	}
}

func TestPayment_TableName(t *testing.T) {
	payment := &Payment{}
	expectedTableName := "payments"
	if payment.TableName() != expectedTableName {
		t.Errorf("Expected TableName to be %s, got %s", expectedTableName, payment.TableName())
	}
}
