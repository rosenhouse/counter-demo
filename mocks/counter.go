package mocks

type Counter struct {
	CountLinesCall struct {
		CallCount int
		Receives  struct {
			PackageName string
		}
		Returns struct {
			Count int
			Error error
		}
	}
}

func (c *Counter) CountLines(importPath string) (int, error) {
	c.CountLinesCall.CallCount++
	c.CountLinesCall.Receives.PackageName = importPath
	return c.CountLinesCall.Returns.Count, c.CountLinesCall.Returns.Error
}
