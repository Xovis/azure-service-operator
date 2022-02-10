// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20181130

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

func Test_UserAssignedIdentities_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentities_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentitiesSPECARM, UserAssignedIdentitiesSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentitiesSPECARM runs a test to see if a specific instance of UserAssignedIdentities_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentitiesSPECARM(subject UserAssignedIdentities_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentities_SPECARM
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

// Generator of UserAssignedIdentities_SPECARM instances for property testing - lazily instantiated by
//UserAssignedIdentitiesSPECARMGenerator()
var userAssignedIdentitiesSPECARMGenerator gopter.Gen

// UserAssignedIdentitiesSPECARMGenerator returns a generator of UserAssignedIdentities_SPECARM instances for property testing.
func UserAssignedIdentitiesSPECARMGenerator() gopter.Gen {
	if userAssignedIdentitiesSPECARMGenerator != nil {
		return userAssignedIdentitiesSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSPECARM(generators)
	userAssignedIdentitiesSPECARMGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentities_SPECARM{}), generators)

	return userAssignedIdentitiesSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}