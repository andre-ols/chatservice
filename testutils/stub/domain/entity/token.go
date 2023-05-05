package stub

type TokenCounterStub struct {
}

func (tcs *TokenCounterStub) CountTokens(model string, text string) (int, error) {
	return 42, nil
}
