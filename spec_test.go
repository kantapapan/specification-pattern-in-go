package main

import (
	"testing"
)

// このテストコードでは、
// OverDueSpecification、
// NoticeSentSpecification、
// InCollectionSpecification、
// AndSpecification、
// OrSpecification、およびNotSpecificationの各仕様が期待通りに動作することを検証しています。
// 
// 実行方法
// go test -v
func TestSpecifications(t *testing.T) {
	overdue := NewOverDueSpecification()
	noticeSent := NewNoticeSentSpecification()
	inCollection := NewInCollectionSpecification()

	invoices := []Invoice{
		{Day: 35, Notice: 4, IsSent: false},
		{Day: 20, Notice: 2, IsSent: true},
		{Day: 40, Notice: 1, IsSent: true},
		{Day: 15, Notice: 5, IsSent: false},
	}

	// Test OverDueSpecification
	if !overdue.IsSatisfiedBy(invoices[0]) {
		t.Error("Expected invoice 0 to be overdue")
	}
	if overdue.IsSatisfiedBy(invoices[1]) {
		t.Error("Expected invoice 1 not to be overdue")
	}

	// Test NoticeSentSpecification
	if !noticeSent.IsSatisfiedBy(invoices[0]) {
		t.Error("Expected invoice 0 to have notice sent")
	}
	if noticeSent.IsSatisfiedBy(invoices[1]) {
		t.Error("Expected invoice 1 not to have notice sent")
	}

	// Test InCollectionSpecification
	if !inCollection.IsSatisfiedBy(invoices[0]) {
		t.Error("Expected invoice 0 to be in collection")
	}
	if inCollection.IsSatisfiedBy(invoices[1]) {
		t.Error("Expected invoice 1 not to be in collection")
	}

	// Test AndSpecification
	critical := overdue.And(noticeSent).And(inCollection)
	if !critical.IsSatisfiedBy(invoices[0]) {
		t.Error("Expected invoice 0 to be critical")
	}
	if critical.IsSatisfiedBy(invoices[1]) {
		t.Error("Expected invoice 1 not to be critical")
	}

	// Test OrSpecification
	overdueOrNoticeSent := overdue.Or(noticeSent)
	if !overdueOrNoticeSent.IsSatisfiedBy(invoices[0]) {
		t.Error("Expected invoice 0 to be overdue or have notice sent")
	}
	if !overdueOrNoticeSent.IsSatisfiedBy(invoices[3]) {
		t.Error("Expected invoice 3 to be overdue or have notice sent")
	}

	// Test NotSpecification
	notOverdue := overdue.Not()
	if notOverdue.IsSatisfiedBy(invoices[0]) {
		t.Error("Expected invoice 0 not to be not overdue")
	}
	if !notOverdue.IsSatisfiedBy(invoices[1]) {
		t.Error("Expected invoice 1 to be not overdue")
	}
}

