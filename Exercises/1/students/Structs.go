package student

import (
	"fmt"
	"time"
)

type Student struct {
	Name string
	ID   int
	Age  int
	DOB  time.Time
}

func (s *Student) String() string {
	formattedDOB := s.DOB.Format("2006-01-02")
	return fmt.Sprintf("Student: %-18s Age: %-6d Born: %s", s.Name, s.Age, formattedDOB)
}
