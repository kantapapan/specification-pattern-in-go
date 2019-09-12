package spec

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
	return false
}

// And ...
func (s *BaseSpecification) And(spec Specification) Specification {
	a := &AndSpecification{
		s.Specification, spec,
	}
	a.Relate(a)
	return a
}

// Or ...
func (s *BaseSpecification) Or(spec Specification) Specification {
	a := &OrSpecification{
		s.Specification, spec,
	}
	a.Relate(a)
	return a
}

// Not ...
func (s *BaseSpecification) Not() Specification {
	a := &NotSpecification{
		s.Specification,
	}
	a.Relate(a)
	return a
}

// Relate to specification
func (s *BaseSpecification) Relate(spec Specification) {
	s.Specification = spec
}

/////

// AndSpecification ...
type AndSpecification struct {
	Specification
	compare Specification
}

// IsSatisfiedBy ...
func (s *AndSpecification) IsSatisfiedBy(elm Invoice) bool {
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
	return s.Specification.IsSatisfiedBy(elm) || s.compare.IsSatisfiedBy(elm)
}

/////

// NotSpecification ...
type NotSpecification struct {
	Specification
}

// IsSatisfiedBy ...
func (s *NotSpecification) IsSatisfiedBy(elm Invoice) bool {
	return s.Specification.IsSatisfiedBy(elm)
}

/////

// OverDueSpecification ...
type OverDueSpecification struct {
	Specification
}

// IsSatisfiedBy ...
func (s *OverDueSpecification) IsSatisfiedBy(elm Invoice) bool {
	return elm.Day >= 30
}

// NewOverDueSpecification ...
func NewOverDueSpecification() Specification {
	a := &OverDueSpecification{&BaseSpecification{}}
	a.Relate(a)
	return a
}

// NoticeSentSpecification ...
type NoticeSentSpecification struct {
	Specification
}

// IsSatisfiedBy ...
func (s *NoticeSentSpecification) IsSatisfiedBy(elm Invoice) bool {
	return elm.Notice >= 3
}

// NewNoticeSentSpecification ...
func NewNoticeSentSpecification() Specification {
	a := &NoticeSentSpecification{&BaseSpecification{}}
	a.Relate(a)
	return a
}

// InCollectionSpecification ...
type InCollectionSpecification struct {
	Specification
}

// IsSatisfiedBy ...
func (s *InCollectionSpecification) IsSatisfiedBy(elm Invoice) bool {
	return !elm.IsSent
}

// NewInCollectionSpecification ...
func NewInCollectionSpecification() Specification {
	a := &InCollectionSpecification{&BaseSpecification{}}
	a.Relate(a)
	return a
}
