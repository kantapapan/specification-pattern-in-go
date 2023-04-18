## specification-pattern-in-go

このコードは、Go言語でSpecificationパターンを実装する例です。Specificationパターンは、ビジネスルールをカプセル化し、組み合わせ可能な方法で表現するデザインパターンです。以下にコードの各部分の説明を示します。

### 基本構造体

- `Invoice`構造体: 請求書を表す構造体です。`Day`、`Notice`、および`IsSent`といったプロパティを持っています。

### Specificationインターフェース

- `Specification`インターフェース: すべての仕様に共通のメソッドを定義します。これには、`IsSatisfiedBy`、`And`、`Or`、`Not`、および`Relate`メソッドが含まれます。

### 基本仕様

- `BaseSpecification`構造体: すべての具体的な仕様が継承する基本的な仕様です。`Specification`インターフェースのメソッドをデフォルトで実装しています。

### 複合仕様

- `AndSpecification`構造体: 2つの仕様が両方とも満たされている場合に満たされる仕様を表します。
- `OrSpecification`構造体: 2つの仕様のいずれかが満たされている場合に満たされる仕様を表します。
- `NotSpecification`構造体: 与えられた仕様が満たされていない場合に満たされる仕様を表します。

### 具体的な仕様

- `OverDueSpecification`構造体: 請求日から30日以上経過した請求書を表す仕様です。
- `NoticeSentSpecification`構造体: 通知が3回以上送信された請求書を表す仕様です。
- `InCollectionSpecification`構造体: まだ回収されていない請求書を表す仕様です。

これらの仕様は、複雑なビジネスルールを効率的に表現し、組み合わせることができます。例えば、`OverDueSpecification`と`NoticeSentSpecification`を組み合わせて、期限切れで通知が3回以上送信された請求書を特定することができます。

## 使用例

Specificationパターンを使用して、複雑なビジネスルールを簡単に表現できます。以下に、実装した仕様を使用して複数の条件を組み合わせる例を示します。

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
この例では、OverDueSpecification、NoticeSentSpecification、およびInCollectionSpecificationを組み合わせて、期限切れで通知が送信され、まだ回収されていない請求書を特定します。
このように、Specificationパターンを使用することで、ビジネスルールを効率的に表現し、組み合わせることができます。



