package spec_test

import (
	"fmt"
	"go-specification/spec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Spec", func() {

	Context("SUCCESS：Testing Specification", func() {
		It(" Example of use...", func() {

			fmt.Println("# === NewOverDue")
			overDue := spec.NewOverDueSpecification()
			fmt.Println("# === NewNotice")
			noticeSent := spec.NewNoticeSentSpecification()
			fmt.Println("# === NewInCollect")
			inCollection := spec.NewInCollectionSpecification()

			fmt.Println("### ========BuileRule========")
			sendToCollection := overDue.And(noticeSent).And(inCollection.Not())

			invoice := spec.Invoice{
				Day:    31,    // >= 30
				Notice: 4,     // >= 3
				IsSent: false, // false
			}

			// true!
			fmt.Println("=======================")
			fmt.Printf("%#v\n", sendToCollection)
			fmt.Println("=======================")
			result := sendToCollection.IsSatisfiedBy(invoice)

			actual := result
			expected := true
			Expect(actual).To(Equal(expected))
		})
	})

})
