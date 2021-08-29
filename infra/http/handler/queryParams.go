package handler

func getQueryParamOne(queryParams map[string][]string, paramName string) string {
	if len(queryParams) > 0 {
		if param, ok := queryParams[paramName]; ok {
			return param[0]
		}
	}
	return ""
}
