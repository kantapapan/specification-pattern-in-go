package spec_test

import (
	"go-specification/spec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Spec", func() {

	Context("SUCCESS：Testing Specification", func() {
		It(" Example of use...", func() {

			overDue := spec.NewOverDueSpecification()
			noticeSent := spec.NewNoticeSentSpecification()
			inCollection := spec.NewInCollectionSpecification()

			sendToCollection := overDue.And(noticeSent).And(inCollection.Not())

			invoice := spec.Invoice{
				Day:    31,    // >= 30
				Notice: 4,     // >= 3
				IsSent: false, // false
			}

			// true!
			result := sendToCollection.IsSatisfiedBy(invoice)

			actual := result
			expected := true
			Expect(actual).To(Equal(expected))
		})
	})

})
