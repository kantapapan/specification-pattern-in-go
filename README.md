# Specification Pattern in Go

This repository demonstrates an implementation of the **Specification Pattern** in Go. The Specification Pattern encapsulates business rules in a reusable and composable way, making it easier to express complex logic. This README is written to be accessible to both Japanese and international engineers.

---

## Overview

### Core Structures

- **`Invoice` struct**: Represents an invoice with the following properties:
  - `Day`: Number of days since the invoice was issued.
  - `Notice`: Number of notices sent.
  - `IsSent`: Whether the invoice has been collected (`true` if collected).

### Specification Interface

The `Specification` interface defines the following methods:

- `IsSatisfiedBy(Invoice) bool`: Checks if the specification is satisfied by a given invoice.
- `And(Specification) Specification`: Combines two specifications with an AND condition.
- `Or(Specification) Specification`: Combines two specifications with an OR condition.
- `Not() Specification`: Negates a specification.
- `Relate(Specification)`: Relates specifications internally (optional for advanced use).

### Base Specification

- **`BaseSpecification` struct**: A base implementation of the `Specification` interface. All concrete specifications inherit from this base.

### Composite Specifications

- **`AndSpecification`**: Satisfied when both specifications are met.
- **`OrSpecification`**: Satisfied when at least one specification is met.
- **`NotSpecification`**: Satisfied when the given specification is not met.

### Concrete Specifications

- **`OverDueSpecification`**: Satisfied when the invoice is overdue (e.g., more than 30 days).
- **`NoticeSentSpecification`**: Satisfied when the invoice has received at least 3 notices.
- **`InCollectionSpecification`**: Satisfied when the invoice has not yet been collected.

These specifications can be combined to express complex business rules.

---

## Example Usage

The following example demonstrates how to use the Specification Pattern to filter invoices based on multiple conditions:

```go
package main

import (
	"fmt"
	"your_repository/spec"
)

func main() {
	invoices := []spec.Invoice{
		{Day: 35, Notice: 4, IsSent: false},
		{Day: 20, Notice: 2, IsSent: true},
		{Day: 40, Notice: 1, IsSent: true},
		{Day: 15, Notice: 5, IsSent: false},
	}

	overdue := spec.NewOverDueSpecification()
	noticeSent := spec.NewNoticeSentSpecification()
	inCollection := spec.NewInCollectionSpecification()

	criticalInvoices := overdue.And(noticeSent).And(inCollection)

	for _, invoice := range invoices {
		if criticalInvoices.IsSatisfiedBy(invoice) {
			fmt.Printf("Critical invoice found: %+v\n", invoice)
		}
	}
}
```

### Explanation

In this example:
- `OverDueSpecification` checks if the invoice is overdue.
- `NoticeSentSpecification` checks if at least 3 notices have been sent.
- `InCollectionSpecification` checks if the invoice has not been collected.

The combined specification (`criticalInvoices`) identifies invoices that are overdue, have received notices, and are still uncollected.

---

## Why Use the Specification Pattern?

- **Reusability**: Specifications can be reused across different parts of the application.
- **Composability**: Combine specifications using logical operators (AND, OR, NOT).
- **Readability**: Business rules are expressed in a clear and declarative manner.

---

## Running Tests

To ensure the implementation works as expected, run the tests using the following command:

```bash
go test ./...
```

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.



