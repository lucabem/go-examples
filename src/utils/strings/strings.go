package StringsUtils

func IsRowEmpty(row []string) bool {
	isEmpty := true
	for _, cell := range row {
		if cell != "" {
			isEmpty = !isEmpty
			break
		}
	}

	return isEmpty
}

func TransformPath(inputPath string, transformations []func(string) string) string {
	transformedPath := inputPath
	for _, transformation := range transformations {
		transformedPath = transformation(transformedPath)
	}

	return transformedPath
}
