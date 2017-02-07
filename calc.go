package stringUtils

func AddNum(iArgs ... int) int {
	result := 0

	for _,argsVal := range iArgs {
		result += argsVal
	}

	return result
}