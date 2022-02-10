// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARM, DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARM runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARM(subject DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM instances for property testing -
//lazily instantiated by DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARMGenerator()
var databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARMGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARMGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARMGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARMGenerator != nil {
		return databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARM(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARM(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPECARM{}), generators)

	return databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = SqlUserDefinedFunctionCreateUpdatePropertiesSpecARMGenerator()
}

func Test_SqlUserDefinedFunctionCreateUpdateProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionCreateUpdateProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdatePropertiesSpecARM, SqlUserDefinedFunctionCreateUpdatePropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdatePropertiesSpecARM runs a test to see if a specific instance of SqlUserDefinedFunctionCreateUpdateProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdatePropertiesSpecARM(subject SqlUserDefinedFunctionCreateUpdateProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionCreateUpdateProperties_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SqlUserDefinedFunctionCreateUpdateProperties_SpecARM instances for property testing - lazily
//instantiated by SqlUserDefinedFunctionCreateUpdatePropertiesSpecARMGenerator()
var sqlUserDefinedFunctionCreateUpdatePropertiesSpecARMGenerator gopter.Gen

// SqlUserDefinedFunctionCreateUpdatePropertiesSpecARMGenerator returns a generator of SqlUserDefinedFunctionCreateUpdateProperties_SpecARM instances for property testing.
func SqlUserDefinedFunctionCreateUpdatePropertiesSpecARMGenerator() gopter.Gen {
	if sqlUserDefinedFunctionCreateUpdatePropertiesSpecARMGenerator != nil {
		return sqlUserDefinedFunctionCreateUpdatePropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdatePropertiesSpecARM(generators)
	sqlUserDefinedFunctionCreateUpdatePropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionCreateUpdateProperties_SpecARM{}), generators)

	return sqlUserDefinedFunctionCreateUpdatePropertiesSpecARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdatePropertiesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdatePropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsSpecARMGenerator())
	gens["Resource"] = SqlUserDefinedFunctionResourceSpecARMGenerator()
}

func Test_SqlUserDefinedFunctionResource_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionResource_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionResourceSpecARM, SqlUserDefinedFunctionResourceSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionResourceSpecARM runs a test to see if a specific instance of SqlUserDefinedFunctionResource_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionResourceSpecARM(subject SqlUserDefinedFunctionResource_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionResource_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SqlUserDefinedFunctionResource_SpecARM instances for property testing - lazily instantiated by
//SqlUserDefinedFunctionResourceSpecARMGenerator()
var sqlUserDefinedFunctionResourceSpecARMGenerator gopter.Gen

// SqlUserDefinedFunctionResourceSpecARMGenerator returns a generator of SqlUserDefinedFunctionResource_SpecARM instances for property testing.
func SqlUserDefinedFunctionResourceSpecARMGenerator() gopter.Gen {
	if sqlUserDefinedFunctionResourceSpecARMGenerator != nil {
		return sqlUserDefinedFunctionResourceSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResourceSpecARM(generators)
	sqlUserDefinedFunctionResourceSpecARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionResource_SpecARM{}), generators)

	return sqlUserDefinedFunctionResourceSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResourceSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResourceSpecARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.AlphaString()
}