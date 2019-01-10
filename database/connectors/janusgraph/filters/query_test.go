package filters

import (
	"testing"
	"time"

	"github.com/creativesoftwarefdn/weaviate/database/connectors/janusgraph/state"
	"github.com/creativesoftwarefdn/weaviate/database/schema"
	cf "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/common_filters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_EmptyFilters(t *testing.T) {
	result, err := New(nil, nil).String()

	require.Nil(t, err, "no error should have occurred")
	assert.Equal(t, "", result, "should return an empty string")
}

func Test_SingleProperties(t *testing.T) {
	t.Run("with propertyType Int", func(t *testing.T) {
		t.Run("with various operators and valid values", func(t *testing.T) {
			tests := testCases{
				{"'City.population == 10000'", cf.OperatorEqual, `.has("population", eq(10000))`},
				{"'City.population != 10000'", cf.OperatorNotEqual, `.has("population", neq(10000))`},
				{"'City.population < 10000'", cf.OperatorLessThan, `.has("population", lt(10000))`},
				{"'City.population <= 10000'", cf.OperatorLessThanEqual, `.has("population", lte(10000))`},
				{"'City.population > 10000'", cf.OperatorGreaterThan, `.has("population", gt(10000))`},
				{"'City.population >= 10000'", cf.OperatorGreaterThanEqual, `.has("population", gte(10000))`},
			}

			tests.AssertFilter(t, "population", int64(10000), schema.DataTypeInt)
		})

		t.Run("an invalid value", func(t *testing.T) {
			tests := testCases{{"should fail with wrong type", cf.OperatorEqual, ""}}

			// Note the mismatch between the specified type (arg4) and the actual type (arg3)
			tests.AssertFilterErrors(t, "population", "200", schema.DataTypeInt)
		})
	})

	t.Run("with propertyType Number (float)", func(t *testing.T) {
		t.Run("with various operators and valid values", func(t *testing.T) {
			tests := testCases{
				{"'City.energyConsumption == 953.280000'", cf.OperatorEqual, `.has("energyConsumption", eq(953.280000))`},
				{"'City.energyConsumption != 953.280000'", cf.OperatorNotEqual, `.has("energyConsumption", neq(953.280000))`},
				{"'City.energyConsumption < 953.280000'", cf.OperatorLessThan, `.has("energyConsumption", lt(953.280000))`},
				{"'City.energyConsumption <= 953.280000'", cf.OperatorLessThanEqual, `.has("energyConsumption", lte(953.280000))`},
				{"'City.energyConsumption > 953.280000'", cf.OperatorGreaterThan, `.has("energyConsumption", gt(953.280000))`},
				{"'City.energyConsumption >= 953.280000'", cf.OperatorGreaterThanEqual, `.has("energyConsumption", gte(953.280000))`},
			}

			tests.AssertFilter(t, "energyConsumption", float64(953.28), schema.DataTypeNumber)
		})

		t.Run("an invalid value", func(t *testing.T) {
			tests := testCases{{"should fail with wrong type", cf.OperatorEqual, ""}}

			// Note the mismatch between the specified type (arg4) and the actual type (arg3)
			tests.AssertFilterErrors(t, "energyConsumption", "200", schema.DataTypeNumber)
		})
	})

	t.Run("with propertyType date (time.Time)", func(t *testing.T) {
		t.Run("with various operators and valid values", func(t *testing.T) {
			dateString := "2017-08-17T12:47:00+02:00"
			dateTime, err := time.Parse(time.RFC3339, dateString)
			require.Nil(t, err)

			tests := testCases{
				{`City.foundedWhen == "2017-08-17T12:47:00+02:00"`, cf.OperatorEqual,
					`.has("foundedWhen", eq("2017-08-17T12:47:00+02:00"))`},
				{`City.foundedWhen != "2017-08-17T12:47:00+02:00"`, cf.OperatorNotEqual,
					`.has("foundedWhen", neq("2017-08-17T12:47:00+02:00"))`},
				{`City.foundedWhen < "2017-08-17T12:47:00+02:00"`, cf.OperatorLessThan,
					`.has("foundedWhen", lt("2017-08-17T12:47:00+02:00"))`},
				{`City.foundedWhen <= "2017-08-17T12:47:00+02:00"`, cf.OperatorLessThanEqual,
					`.has("foundedWhen", lte("2017-08-17T12:47:00+02:00"))`},
				{`City.foundedWhen > "2017-08-17T12:47:00+02:00"`, cf.OperatorGreaterThan,
					`.has("foundedWhen", gt("2017-08-17T12:47:00+02:00"))`},
				{`City.foundedWhen >= "2017-08-17T12:47:00+02:00"`, cf.OperatorGreaterThanEqual,
					`.has("foundedWhen", gte("2017-08-17T12:47:00+02:00"))`},
			}

			tests.AssertFilter(t, "foundedWhen", dateTime, schema.DataTypeDate)
		})

		t.Run("an invalid value", func(t *testing.T) {
			tests := testCases{{"should fail with wrong type", cf.OperatorEqual, ""}}

			// Note the mismatch between the specified type (arg4) and the actual type (arg3)
			tests.AssertFilterErrors(t, "foundedWhen", "200", schema.DataTypeDate)
		})
	})

	t.Run("with propertyType string", func(t *testing.T) {
		t.Run("with various operators and valid values", func(t *testing.T) {
			tests := testCases{
				{`'City.name == "Berlin"'`, cf.OperatorEqual, `.has("name", eq("Berlin"))`},
				{`'City.name != "Berlin"'`, cf.OperatorNotEqual, `.has("name", neq("Berlin"))`},
			}

			tests.AssertFilter(t, "name", "Berlin", schema.DataTypeString)
		})

		t.Run("with an operator that does not make sense for this type", func(t *testing.T) {
			tests := testCases{
				{`City.name < "Berlin"`, cf.OperatorLessThan, ""},
				{`City.name <= "Berlin"`, cf.OperatorLessThanEqual, ""},
				{`City.name > "Berlin"`, cf.OperatorGreaterThan, ""},
				{`City.name >= "Berlin"`, cf.OperatorGreaterThanEqual, ""},
			}

			tests.AssertFilterErrors(t, "name", "Berlin", schema.DataTypeString)
		})

		t.Run("an invalid value", func(t *testing.T) {
			tests := testCases{{"should fail with wrong type", cf.OperatorEqual, ""}}

			// Note the mismatch between the specified type (arg4) and the actual type (arg3)
			tests.AssertFilterErrors(t, "name", int(200), schema.DataTypeString)
		})
	})

	t.Run("with propertyType bool", func(t *testing.T) {
		t.Run("with various operators and valid values", func(t *testing.T) {
			tests := testCases{
				{`'City.isCapital == true'`, cf.OperatorEqual, `.has("isCapital", eq(true))`},
				{`'City.isCapital != true'`, cf.OperatorNotEqual, `.has("isCapital", neq(true))`},
			}

			tests.AssertFilter(t, "isCapital", true, schema.DataTypeBoolean)
		})

		t.Run("with an operator that does not make sense for this type", func(t *testing.T) {
			tests := testCases{
				{`City.isCapital < true`, cf.OperatorLessThan, ""},
				{`City.isCapital <= true`, cf.OperatorLessThanEqual, ""},
				{`City.isCapital > true`, cf.OperatorGreaterThan, ""},
				{`City.isCapital >= true`, cf.OperatorGreaterThanEqual, ""},
			}

			tests.AssertFilterErrors(t, "isCapital", true, schema.DataTypeBoolean)
		})

		t.Run("an invalid value", func(t *testing.T) {
			tests := testCases{{"should fail with wrong type", cf.OperatorEqual, ""}}

			// Note the mismatch between the specified type (arg4) and the actual type (arg3)
			tests.AssertFilterErrors(t, "isCapital", int(200), schema.DataTypeBoolean)
		})
	})
}

func Test_SinglePropertiesWithMappedNames(t *testing.T) {
	tests := testCases{
		{"'City.population == 10000'", cf.OperatorEqual, `.has("prop_20", eq(10000))`},
		{"'City.population != 10000'", cf.OperatorNotEqual, `.has("prop_20", neq(10000))`},
		{"'City.population < 10000'", cf.OperatorLessThan, `.has("prop_20", lt(10000))`},
		{"'City.population <= 10000'", cf.OperatorLessThanEqual, `.has("prop_20", lte(10000))`},
		{"'City.population > 10000'", cf.OperatorGreaterThan, `.has("prop_20", gt(10000))`},
		{"'City.population >= 10000'", cf.OperatorGreaterThanEqual, `.has("prop_20", gte(10000))`},
	}

	tests.AssertFilterWithNameSource(t, "population", int64(10000), schema.DataTypeInt, &fakeNameSource{})
}

func Test_InvalidOperator(t *testing.T) {
	filter := buildFilter("population", "200", cf.Operator(27), schema.DataTypeInt)

	_, err := New(filter, nil).String()

	assert.NotNil(t, err, "it should error due to the wrong type")
}

func Test_MultipleConditions(t *testing.T) {
	t.Run("with operator and", func(t *testing.T) {
		filter := &cf.LocalFilter{
			Root: &cf.Clause{
				Operator: cf.OperatorAnd,
				Operands: []cf.Clause{
					cf.Clause{
						Operator: cf.OperatorGreaterThan,
						On: &cf.Path{
							Class:    schema.ClassName("City"),
							Property: schema.PropertyName("population"),
						},
						Value: &cf.Value{
							Value: int64(70000),
							Type:  schema.DataTypeInt,
						},
					},
					cf.Clause{
						Operator: cf.OperatorNotEqual,
						On: &cf.Path{
							Class:    schema.ClassName("City"),
							Property: schema.PropertyName("name"),
						},
						Value: &cf.Value{
							Value: "Rotterdam",
							Type:  schema.DataTypeString,
						},
					},
				},
			},
		}
		expectedResult := `.and(has("population", gt(70000)), has("name", neq("Rotterdam")))`

		result, err := New(filter, nil).String()

		require.Nil(t, err, "should not error")
		assert.Equal(t, expectedResult, result, "should match the gremlin query")
	})

}

func buildFilter(propName string, value interface{}, operator cf.Operator, schemaType schema.DataType) *cf.LocalFilter {
	return &cf.LocalFilter{
		Root: &cf.Clause{
			Operator: operator,
			On: &cf.Path{
				Class:    schema.ClassName("City"),
				Property: schema.PropertyName(propName),
			},
			Value: &cf.Value{
				Value: value,
				Type:  schemaType,
			},
		},
	}
}

type testCase struct {
	name           string
	operator       cf.Operator
	expectedResult string
}

type testCases []testCase

func (tests testCases) AssertFilter(t *testing.T, propName string, propValue interface{}, propType schema.DataType) {
	tests.AssertFilterWithNameSource(t, propName, propValue, propType, nil)
}

func (tests testCases) AssertFilterWithNameSource(t *testing.T, propName string, propValue interface{}, propType schema.DataType, nameSource nameSource) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filter := buildFilter(propName, propValue, test.operator, propType)

			result, err := New(filter, nameSource).String()

			require.Nil(t, err, "no error should have occurred")
			assert.Equal(t, test.expectedResult, result, "should form the right query")
		})
	}
}

func (tests testCases) AssertFilterErrors(t *testing.T, propName string, propValue interface{}, propType schema.DataType) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filter := buildFilter(propName, propValue, test.operator, propType)

			_, err := New(filter, nil).String()

			assert.NotNil(t, err, "should error")
		})
	}
}

type fakeNameSource struct{}

func (f *fakeNameSource) GetMappedPropertyName(className schema.ClassName,
	propName schema.PropertyName) state.MappedPropertyName {
	return state.MappedPropertyName("prop_20")
}