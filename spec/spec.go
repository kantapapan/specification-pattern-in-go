package spec

import "fmt"

// Invoice ...
type Invoice struct {
	Day    int
	Notice int
	IsSent bool
}

/////

// Specification ...
type Specification interface {
	IsSatisfiedBy(Invoice) bool
	And(Specification) Specification
	Or(Specification) Specification
	Not() Specification
	Relate(Specification)
}

/////

// BaseSpecification ...
type BaseSpecification struct {
	Specification
}

// IsSatisfiedBy ...
func (s *BaseSpecification) IsSatisfiedBy(elm Invoice) bool {
	fmt.Println("Base -> IsSatisfiedBy()")
	return false
}

// And ...
func (s *BaseSpecification) And(spec Specification) Specification {
	fmt.Println("BaseSpec -> And()")
	a := &AndSpecification{
		s.Specification, spec,
	}
	fmt.Println("Create AndSpec(Base.Specification, compare=param Specification)  --> Relate ")
	a.Relate(a)
	return a
}

// Or ...
func (s *BaseSpecification) Or(spec Specification) Specification {
	fmt.Println("BaseSpec -> Or()")
	a := &OrSpecification{
		s.Specification, spec,
	}
	fmt.Println("Create OrSpec(Base.Specification, param Specification)  --> Relate ")
	a.Relate(a)
	return a
}

// Not ...
func (s *BaseSpecification) Not() Specification {
	fmt.Println("BaseSpec -> Not")
	a := &NotSpecification{
		s.Specification,
	}
	fmt.Println("Create NotSpec(Base.Specification)  --> Relate ")
	a.Relate(a)
	return a
}

// Relate to specification
func (s *BaseSpecification) Relate(spec Specification) {
	fmt.Println("BaseSpec -> Relate ---- (")
	s.Specification = spec
	fmt.Println("BaseSpec.Specification = param spec Specification")
	fmt.Println("BaseSpec -> Relate ---- )")
}

/////

// AndSpecification ...
type AndSpecification struct {
	Specification
	compare Specification
}

// IsSatisfiedBy ...
func (s *AndSpecification) IsSatisfiedBy(elm Invoice) bool {
	fmt.Println("AndSpec -> IsSatisfiedBy")
	fmt.Println("AndSpec.Specification.IsSatisfiedBy && AndSpec.compare.IsSatisfiedBy")
	return s.Specification.IsSatisfiedBy(elm) && s.compare.IsSatisfiedBy(elm)
}

/////

// OrSpecification ...
type OrSpecification struct {
	Specification
	compare Specification
}

// IsSatisfiedBy ...
func (s *OrSpecification) IsSatisfiedBy(elm Invoice) bool {
	fmt.Println("OrSpec.IsSatisfiedBy")
	fmt.Println("OrSpec.Specification.IsSatisfiedBy && AndSpec.compare.IsSatisfiedBy")
	return s.Specification.IsSatisfiedBy(elm) || s.compare.IsSatisfiedBy(elm)
}

/////

// NotSpecification ...
type NotSpecification struct {
	Specification
}

// IsSatisfiedBy ...
func (s *NotSpecification) IsSatisfiedBy(elm Invoice) bool {
	fmt.Println("NotSpec.IsSatisfiedBy")
	fmt.Println("NotSpec.Specification.IsSatisfiedBy")
	return s.Specification.IsSatisfiedBy(elm)
}

/////

// OverDueSpecification ...
type OverDueSpecification struct {
	Specification
}

// IsSatisfiedBy ...
func (s *OverDueSpecification) IsSatisfiedBy(elm Invoice) bool {
	fmt.Println("*OverDueSpec.IsSatisfiedBy")
	return elm.Day >= 30
}

// NewOverDueSpecification ...
func NewOverDueSpecification() Specification {
	fmt.Println("*OverDueSpec.New Start")
	a := &OverDueSpecification{&BaseSpecification{}}
	fmt.Println("Create *OverDueSpec(OverDueSpec.Specification=&BaseSpec)  --> Relate ")
	a.Relate(a)
	return a
}

// NoticeSentSpecification ...
type NoticeSentSpecification struct {
	Specification
}

// IsSatisfiedBy ...
func (s *NoticeSentSpecification) IsSatisfiedBy(elm Invoice) bool {
	fmt.Println("*NoticeSentSpec.IsSatisfiedBy")
	return elm.Notice >= 3
}

// NewNoticeSentSpecification ...
func NewNoticeSentSpecification() Specification {
	fmt.Println("*NoticeSentSpec.New Start")
	a := &NoticeSentSpecification{&BaseSpecification{}}
	fmt.Println("Create *NoticeSentSpec(NoticeSentSpec.Specification=&BaseSpec)  --> Relate ")
	a.Relate(a)
	return a
}

// InCollectionSpecification ...
type InCollectionSpecification struct {
	Specification
}

// IsSatisfiedBy ...
func (s *InCollectionSpecification) IsSatisfiedBy(elm Invoice) bool {
	fmt.Println("*InCollectionSpec.IsSatisfiedBy")
	return !elm.IsSent
}

// NewInCollectionSpecification ...
func NewInCollectionSpecification() Specification {
	fmt.Println("*InCollectionSpec.New Start")
	a := &InCollectionSpecification{&BaseSpecification{}}
	fmt.Println("Create *InCollectionSpec(InCollectionSpec.Specification=&BaseSpec)  --> Relate ")
	a.Relate(a)
	return a
}
