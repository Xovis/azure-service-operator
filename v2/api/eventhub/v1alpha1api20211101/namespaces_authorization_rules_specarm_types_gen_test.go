// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

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

func Test_NamespacesAuthorizationRules_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRules_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRulesSPECARM, NamespacesAuthorizationRulesSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRulesSPECARM runs a test to see if a specific instance of NamespacesAuthorizationRules_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRulesSPECARM(subject NamespacesAuthorizationRules_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRules_SPECARM
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

// Generator of NamespacesAuthorizationRules_SPECARM instances for property testing - lazily instantiated by
//NamespacesAuthorizationRulesSPECARMGenerator()
var namespacesAuthorizationRulesSPECARMGenerator gopter.Gen

// NamespacesAuthorizationRulesSPECARMGenerator returns a generator of NamespacesAuthorizationRules_SPECARM instances for property testing.
// We first initialize namespacesAuthorizationRulesSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesAuthorizationRulesSPECARMGenerator() gopter.Gen {
	if namespacesAuthorizationRulesSPECARMGenerator != nil {
		return namespacesAuthorizationRulesSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSPECARM(generators)
	namespacesAuthorizationRulesSPECARMGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRules_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSPECARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRulesSPECARM(generators)
	namespacesAuthorizationRulesSPECARMGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRules_SPECARM{}), generators)

	return namespacesAuthorizationRulesSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRulesSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRulesSPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NamespacesAuthorizationRulesSPECPropertiesARMGenerator())
}

func Test_NamespacesAuthorizationRules_SPEC_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRules_SPEC_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRulesSPECPropertiesARM, NamespacesAuthorizationRulesSPECPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRulesSPECPropertiesARM runs a test to see if a specific instance of NamespacesAuthorizationRules_SPEC_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRulesSPECPropertiesARM(subject NamespacesAuthorizationRules_SPEC_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRules_SPEC_PropertiesARM
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

// Generator of NamespacesAuthorizationRules_SPEC_PropertiesARM instances for property testing - lazily instantiated by
//NamespacesAuthorizationRulesSPECPropertiesARMGenerator()
var namespacesAuthorizationRulesSPECPropertiesARMGenerator gopter.Gen

// NamespacesAuthorizationRulesSPECPropertiesARMGenerator returns a generator of NamespacesAuthorizationRules_SPEC_PropertiesARM instances for property testing.
func NamespacesAuthorizationRulesSPECPropertiesARMGenerator() gopter.Gen {
	if namespacesAuthorizationRulesSPECPropertiesARMGenerator != nil {
		return namespacesAuthorizationRulesSPECPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSPECPropertiesARM(generators)
	namespacesAuthorizationRulesSPECPropertiesARMGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRules_SPEC_PropertiesARM{}), generators)

	return namespacesAuthorizationRulesSPECPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSPECPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSPECPropertiesARM(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(NamespacesAuthorizationRulesSPECPropertiesRightsListen, NamespacesAuthorizationRulesSPECPropertiesRightsManage, NamespacesAuthorizationRulesSPECPropertiesRightsSend))
}