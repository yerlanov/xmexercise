package postgresql

import "github.com/Masterminds/squirrel"

func BuildFilter(builder squirrel.SelectBuilder, filter map[string]string) squirrel.SelectBuilder {
	return builder.Where(getConditions(filter))
}

func getConditions(filter map[string]string) squirrel.Sqlizer {
	var conditions []squirrel.Sqlizer

	for k, v := range filter {
		condition := squirrel.Eq{k: v}
		conditions = append(conditions, condition)
	}

	return and(conditions)
}

func and(conditions []squirrel.Sqlizer) squirrel.Sqlizer {
	result := squirrel.And{}
	for _, condition := range conditions {
		result = append(result, condition)
	}
	return result
}
