package mocks

type Counter struct {
	CountCall struct {
		Receives struct {
			Package string
		}
		Returns struct {
			Lines int
			Error error
		}
	}
}

func (c *Counter) Count(pkgPath string) (int, error) {
	c.CountCall.Receives.Package = pkgPath
	return c.CountCall.Returns.Lines, c.CountCall.Returns.Error
}
