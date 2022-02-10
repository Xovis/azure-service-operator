// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

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

func Test_EventSubscriptions_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscriptions_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionsSPECARM, EventSubscriptionsSPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionsSPECARM runs a test to see if a specific instance of EventSubscriptions_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionsSPECARM(subject EventSubscriptions_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscriptions_SPECARM
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

// Generator of EventSubscriptions_SPECARM instances for property testing - lazily instantiated by
//EventSubscriptionsSPECARMGenerator()
var eventSubscriptionsSPECARMGenerator gopter.Gen

// EventSubscriptionsSPECARMGenerator returns a generator of EventSubscriptions_SPECARM instances for property testing.
// We first initialize eventSubscriptionsSPECARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventSubscriptionsSPECARMGenerator() gopter.Gen {
	if eventSubscriptionsSPECARMGenerator != nil {
		return eventSubscriptionsSPECARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionsSPECARM(generators)
	eventSubscriptionsSPECARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptions_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionsSPECARM(generators)
	AddRelatedPropertyGeneratorsForEventSubscriptionsSPECARM(generators)
	eventSubscriptionsSPECARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptions_SPECARM{}), generators)

	return eventSubscriptionsSPECARMGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionsSPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionsSPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForEventSubscriptionsSPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventSubscriptionsSPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(EventSubscriptionPropertiesSpecARMGenerator())
}

func Test_EventSubscriptionProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscriptionProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionPropertiesSpecARM, EventSubscriptionPropertiesSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionPropertiesSpecARM runs a test to see if a specific instance of EventSubscriptionProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionPropertiesSpecARM(subject EventSubscriptionProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscriptionProperties_SpecARM
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

// Generator of EventSubscriptionProperties_SpecARM instances for property testing - lazily instantiated by
//EventSubscriptionPropertiesSpecARMGenerator()
var eventSubscriptionPropertiesSpecARMGenerator gopter.Gen

// EventSubscriptionPropertiesSpecARMGenerator returns a generator of EventSubscriptionProperties_SpecARM instances for property testing.
// We first initialize eventSubscriptionPropertiesSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventSubscriptionPropertiesSpecARMGenerator() gopter.Gen {
	if eventSubscriptionPropertiesSpecARMGenerator != nil {
		return eventSubscriptionPropertiesSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionPropertiesSpecARM(generators)
	eventSubscriptionPropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionProperties_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionPropertiesSpecARM(generators)
	AddRelatedPropertyGeneratorsForEventSubscriptionPropertiesSpecARM(generators)
	eventSubscriptionPropertiesSpecARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionProperties_SpecARM{}), generators)

	return eventSubscriptionPropertiesSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionPropertiesSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionPropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["EventDeliverySchema"] = gen.PtrOf(gen.OneConstOf(EventSubscriptionPropertiesSpecEventDeliverySchemaCloudEventSchemaV10, EventSubscriptionPropertiesSpecEventDeliverySchemaCustomInputSchema, EventSubscriptionPropertiesSpecEventDeliverySchemaEventGridSchema))
	gens["ExpirationTimeUtc"] = gen.PtrOf(gen.AlphaString())
	gens["Labels"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventSubscriptionPropertiesSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventSubscriptionPropertiesSpecARM(gens map[string]gopter.Gen) {
	gens["DeadLetterDestination"] = gen.PtrOf(DeadLetterDestinationSpecARMGenerator())
	gens["Destination"] = gen.PtrOf(EventSubscriptionDestinationSpecARMGenerator())
	gens["Filter"] = gen.PtrOf(EventSubscriptionFilterSpecARMGenerator())
	gens["RetryPolicy"] = gen.PtrOf(RetryPolicySpecARMGenerator())
}

func Test_DeadLetterDestination_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeadLetterDestination_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeadLetterDestinationSpecARM, DeadLetterDestinationSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeadLetterDestinationSpecARM runs a test to see if a specific instance of DeadLetterDestination_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDeadLetterDestinationSpecARM(subject DeadLetterDestination_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeadLetterDestination_SpecARM
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

// Generator of DeadLetterDestination_SpecARM instances for property testing - lazily instantiated by
//DeadLetterDestinationSpecARMGenerator()
var deadLetterDestinationSpecARMGenerator gopter.Gen

// DeadLetterDestinationSpecARMGenerator returns a generator of DeadLetterDestination_SpecARM instances for property testing.
func DeadLetterDestinationSpecARMGenerator() gopter.Gen {
	if deadLetterDestinationSpecARMGenerator != nil {
		return deadLetterDestinationSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeadLetterDestinationSpecARM(generators)
	deadLetterDestinationSpecARMGenerator = gen.Struct(reflect.TypeOf(DeadLetterDestination_SpecARM{}), generators)

	return deadLetterDestinationSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDeadLetterDestinationSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeadLetterDestinationSpecARM(gens map[string]gopter.Gen) {
	gens["EndpointType"] = gen.OneConstOf(DeadLetterDestinationSpecEndpointTypeStorageBlob)
}

func Test_EventSubscriptionDestination_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscriptionDestination_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionDestinationSpecARM, EventSubscriptionDestinationSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionDestinationSpecARM runs a test to see if a specific instance of EventSubscriptionDestination_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionDestinationSpecARM(subject EventSubscriptionDestination_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscriptionDestination_SpecARM
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

// Generator of EventSubscriptionDestination_SpecARM instances for property testing - lazily instantiated by
//EventSubscriptionDestinationSpecARMGenerator()
var eventSubscriptionDestinationSpecARMGenerator gopter.Gen

// EventSubscriptionDestinationSpecARMGenerator returns a generator of EventSubscriptionDestination_SpecARM instances for property testing.
func EventSubscriptionDestinationSpecARMGenerator() gopter.Gen {
	if eventSubscriptionDestinationSpecARMGenerator != nil {
		return eventSubscriptionDestinationSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionDestinationSpecARM(generators)
	eventSubscriptionDestinationSpecARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionDestination_SpecARM{}), generators)

	return eventSubscriptionDestinationSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionDestinationSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionDestinationSpecARM(gens map[string]gopter.Gen) {
	gens["EndpointType"] = gen.OneConstOf(
		EventSubscriptionDestinationSpecEndpointTypeAzureFunction,
		EventSubscriptionDestinationSpecEndpointTypeEventHub,
		EventSubscriptionDestinationSpecEndpointTypeHybridConnection,
		EventSubscriptionDestinationSpecEndpointTypeServiceBusQueue,
		EventSubscriptionDestinationSpecEndpointTypeServiceBusTopic,
		EventSubscriptionDestinationSpecEndpointTypeStorageQueue,
		EventSubscriptionDestinationSpecEndpointTypeWebHook)
}

func Test_EventSubscriptionFilter_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscriptionFilter_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionFilterSpecARM, EventSubscriptionFilterSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionFilterSpecARM runs a test to see if a specific instance of EventSubscriptionFilter_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionFilterSpecARM(subject EventSubscriptionFilter_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscriptionFilter_SpecARM
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

// Generator of EventSubscriptionFilter_SpecARM instances for property testing - lazily instantiated by
//EventSubscriptionFilterSpecARMGenerator()
var eventSubscriptionFilterSpecARMGenerator gopter.Gen

// EventSubscriptionFilterSpecARMGenerator returns a generator of EventSubscriptionFilter_SpecARM instances for property testing.
// We first initialize eventSubscriptionFilterSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventSubscriptionFilterSpecARMGenerator() gopter.Gen {
	if eventSubscriptionFilterSpecARMGenerator != nil {
		return eventSubscriptionFilterSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionFilterSpecARM(generators)
	eventSubscriptionFilterSpecARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionFilter_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionFilterSpecARM(generators)
	AddRelatedPropertyGeneratorsForEventSubscriptionFilterSpecARM(generators)
	eventSubscriptionFilterSpecARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionFilter_SpecARM{}), generators)

	return eventSubscriptionFilterSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionFilterSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionFilterSpecARM(gens map[string]gopter.Gen) {
	gens["IncludedEventTypes"] = gen.SliceOf(gen.AlphaString())
	gens["IsSubjectCaseSensitive"] = gen.PtrOf(gen.Bool())
	gens["SubjectBeginsWith"] = gen.PtrOf(gen.AlphaString())
	gens["SubjectEndsWith"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventSubscriptionFilterSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventSubscriptionFilterSpecARM(gens map[string]gopter.Gen) {
	gens["AdvancedFilters"] = gen.SliceOf(AdvancedFilterSpecARMGenerator())
}

func Test_RetryPolicy_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RetryPolicy_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRetryPolicySpecARM, RetryPolicySpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRetryPolicySpecARM runs a test to see if a specific instance of RetryPolicy_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRetryPolicySpecARM(subject RetryPolicy_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RetryPolicy_SpecARM
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

// Generator of RetryPolicy_SpecARM instances for property testing - lazily instantiated by RetryPolicySpecARMGenerator()
var retryPolicySpecARMGenerator gopter.Gen

// RetryPolicySpecARMGenerator returns a generator of RetryPolicy_SpecARM instances for property testing.
func RetryPolicySpecARMGenerator() gopter.Gen {
	if retryPolicySpecARMGenerator != nil {
		return retryPolicySpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRetryPolicySpecARM(generators)
	retryPolicySpecARMGenerator = gen.Struct(reflect.TypeOf(RetryPolicy_SpecARM{}), generators)

	return retryPolicySpecARMGenerator
}

// AddIndependentPropertyGeneratorsForRetryPolicySpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRetryPolicySpecARM(gens map[string]gopter.Gen) {
	gens["EventTimeToLiveInMinutes"] = gen.PtrOf(gen.Int())
	gens["MaxDeliveryAttempts"] = gen.PtrOf(gen.Int())
}

func Test_AdvancedFilter_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AdvancedFilter_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAdvancedFilterSpecARM, AdvancedFilterSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAdvancedFilterSpecARM runs a test to see if a specific instance of AdvancedFilter_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAdvancedFilterSpecARM(subject AdvancedFilter_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AdvancedFilter_SpecARM
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

// Generator of AdvancedFilter_SpecARM instances for property testing - lazily instantiated by
//AdvancedFilterSpecARMGenerator()
var advancedFilterSpecARMGenerator gopter.Gen

// AdvancedFilterSpecARMGenerator returns a generator of AdvancedFilter_SpecARM instances for property testing.
func AdvancedFilterSpecARMGenerator() gopter.Gen {
	if advancedFilterSpecARMGenerator != nil {
		return advancedFilterSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAdvancedFilterSpecARM(generators)
	advancedFilterSpecARMGenerator = gen.Struct(reflect.TypeOf(AdvancedFilter_SpecARM{}), generators)

	return advancedFilterSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForAdvancedFilterSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAdvancedFilterSpecARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(gen.AlphaString())
	gens["OperatorType"] = gen.OneConstOf(
		AdvancedFilterSpecOperatorTypeBoolEquals,
		AdvancedFilterSpecOperatorTypeNumberGreaterThan,
		AdvancedFilterSpecOperatorTypeNumberGreaterThanOrEquals,
		AdvancedFilterSpecOperatorTypeNumberIn,
		AdvancedFilterSpecOperatorTypeNumberLessThan,
		AdvancedFilterSpecOperatorTypeNumberLessThanOrEquals,
		AdvancedFilterSpecOperatorTypeNumberNotIn,
		AdvancedFilterSpecOperatorTypeStringBeginsWith,
		AdvancedFilterSpecOperatorTypeStringContains,
		AdvancedFilterSpecOperatorTypeStringEndsWith,
		AdvancedFilterSpecOperatorTypeStringIn,
		AdvancedFilterSpecOperatorTypeStringNotIn)
}