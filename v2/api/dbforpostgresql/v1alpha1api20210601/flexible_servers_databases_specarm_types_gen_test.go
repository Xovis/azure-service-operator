// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

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

func Test_FlexibleServersDatabases_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersDatabases_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersDatabasesSPECARM, FlexibleServersDatabasesSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersDatabasesSPECARM runs a test to see if a specific instance of FlexibleServersDatabases_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersDatabasesSPECARM(subject FlexibleServersDatabases_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersDatabases_SPECARM
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

// Generator of FlexibleServersDatabases_SPECARM instances for property testing - lazily instantiated by
//FlexibleServersDatabasesSPECARMGenerator()
var flexibleServersDatabasesSPECARMGenerator gopter.Gen

// FlexibleServersDatabasesSPECARMGenerator returns a generator of FlexibleServersDatabases_SPECARM instances for property testing.
// We first initialize flexibleServersDatabasesSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersDatabasesSPECARMGenerator() gopter.Gen {
	if flexibleServersDatabasesSPECARMGenerator != nil {
		return flexibleServersDatabasesSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersDatabasesSPECARM(generators)
	flexibleServersDatabasesSPECARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabases_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersDatabasesSPECARM(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersDatabasesSPECARM(generators)
	flexibleServersDatabasesSPECARMGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabases_SPECARM{}), generators)

	return flexibleServersDatabasesSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersDatabasesSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersDatabasesSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForFlexibleServersDatabasesSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersDatabasesSPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabasePropertiesSpecARMGenerator())
}

func Test_DatabaseProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabasePropertiesSpecARM, DatabasePropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabasePropertiesSpecARM runs a test to see if a specific instance of DatabaseProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabasePropertiesSpecARM(subject DatabaseProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_SpecARM
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

// Generator of DatabaseProperties_SpecARM instances for property testing - lazily instantiated by
//DatabasePropertiesSpecARMGenerator()
var databasePropertiesSpecARMGenerator gopter.Gen

// DatabasePropertiesSpecARMGenerator returns a generator of DatabaseProperties_SpecARM instances for property testing.
func DatabasePropertiesSpecARMGenerator() gopter.Gen {
	if databasePropertiesSpecARMGenerator != nil {
		return databasePropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabasePropertiesSpecARM(generators)
	databasePropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_SpecARM{}), generators)

	return databasePropertiesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabasePropertiesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabasePropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
}