package helpers

func NewResourceCollection[T any, U any](
	items []T,
	newResource func(item *T) U,
) []U {
	resources := make([]U, 0, len(items))

	for _, item := range items {
		resources = append(resources, newResource(&item))
	}

	return resources
}
