package fluent

import (
	"fmt"
	"strings"
	"time"

	"github.com/summit-fi/wordsdk-go/fluent/cldr"
	"github.com/summit-fi/wordsdk-go/fluent/parser/ast"
)

// Bundle represents a collection of messages and terms collected from one or many resources.
// It provides the main API to format messages.
type Bundle struct {
	locales   []cldr.Language
	messages  Map[string, *ast.Message]
	terms     Map[string, *ast.Term]
	functions map[string]Function
}

// NewBundle creates a new empty bundle
func NewBundle(primaryLocale cldr.Language, fallbackLocales ...cldr.Language) *Bundle {
	locales := make([]cldr.Language, 0, len(fallbackLocales)+1)
	locales = append(locales, primaryLocale)
	for _, fallback := range fallbackLocales {
		locales = append(locales, fallback)
	}

	return &Bundle{
		locales:  locales,
		messages: NewMap[string, *ast.Message](),
		terms:    NewMap[string, *ast.Term](),
	}
}

// AddResource adds a Resource to the Bundle.
// If a message or term was already defined by another resource, an error is raised and the entry is skipped.
func (bundle *Bundle) AddResource(resource *Resource) (errs []error) {
	for _, message := range resource.messages {
		id := message.ID.Name

		if _, ok := bundle.messages.Exist(id); ok {
			errs = append(errs, fmt.Errorf("message '%s' is already defined", id))
			continue
		}
		bundle.messages.Set(id, message)
	}
	for _, term := range resource.terms {
		id := term.ID.Name
		if bundle.terms.Get(id) != nil {
			errs = append(errs, fmt.Errorf("term '%s' is already defined", id))
			continue
		}
		bundle.terms.Set(id, term)
	}
	return
}

// AddResourceOverriding adds a Resource to the Bundle.
// If a message or term was already defined by another resource, the already existing one gets overridden.
func (bundle *Bundle) AddResourceOverriding(resource *Resource) {
	for _, message := range resource.messages {
		bundle.messages.Set(message.ID.Name, message)
	}
	for _, term := range resource.terms {
		bundle.terms.Set(term.ID.Name, term)
	}
}

func (bundle *Bundle) RetrieveMessages() map[string]string {
	if !bundle.messages.IsInitialized() {
		return nil
	}
	result := make(map[string]string, len(bundle.messages.RetrieveAll()))
	//for key, message := range bundle.messages.RetrieveAll() {
	//	if message == nil {
	//		continue
	//	}
	//	//result[key] = message.Attributes
	//}
	return result
}

// RegisterFunction registers a function with the given name.
func (bundle *Bundle) RegisterFunction(name string, function Function) {
	if bundle.functions == nil {
		bundle.functions = make(map[string]Function)
	}
	bundle.functions[strings.ToUpper(name)] = function
}

func (bundle *Bundle) PrimaryLocale() cldr.Language {
	if len(bundle.locales) > 0 {
		return bundle.locales[0]
	}
	return cldr.LanguageEnUa
}

// A FormatContext holds variables and functions to pass them to Bundle.FormatMessage
type FormatContext struct {
	variables map[string]Value
	functions map[string]Function
}

// WithVariable creates a FormatContext with a single variable
func WithVariable(key string, value interface{}) *FormatContext {
	resolved := resolveValue(value)
	if resolved == nil {
		return &FormatContext{
			variables: nil,
			functions: nil,
		}
	}
	return &FormatContext{
		variables: map[string]Value{strings.TrimSpace(key): resolved},
		functions: nil,
	}
}

// WithVariables creates a FormatContext with multiple variables
func WithVariables(variables map[string]interface{}) *FormatContext {
	cleaned := make(map[string]Value, len(variables))
	for name, variable := range variables {
		resolved := resolveValue(variable)
		if resolved == nil {
			continue
		}
		cleaned[strings.TrimSpace(name)] = resolved
	}
	return &FormatContext{
		variables: cleaned,
		functions: nil,
	}
}

func resolveValue(value interface{}) Value {
	if strVal, ok := value.(string); ok {
		return String(strVal)
	}
	if strVal, ok := value.(*StringValue); ok {
		return strVal
	}
	if boolVal, ok := value.(bool); ok {
		if boolVal {
			return String("true")
		}
		return String("false")
	}
	if float32Val, ok := value.(float32); ok {
		return NumberLiteral(float32Val)
	}
	if float64Val, ok := value.(float64); ok {
		return NumberLiteral(float32(float64Val))
	}
	if uintVal, ok := value.(uint); ok {
		return NumberLiteral(float32(uintVal))
	}
	if uint8Val, ok := value.(uint8); ok {
		return NumberLiteral(float32(uint8Val))
	}
	if uint16Val, ok := value.(uint16); ok {
		return NumberLiteral(float32(uint16Val))
	}
	if uint32val, ok := value.(uint32); ok {
		return NumberLiteral(float32(uint32val))
	}
	if uint64val, ok := value.(uint64); ok {
		return NumberLiteral(float32(uint64val))
	}
	if intVal, ok := value.(int); ok {
		return NumberLiteral(float32(intVal))
	}
	if int8Val, ok := value.(int8); ok {
		return NumberLiteral(float32(int8Val))
	}
	if int16Val, ok := value.(int16); ok {
		return NumberLiteral(float32(int16Val))
	}
	if int32val, ok := value.(int32); ok {
		return NumberLiteral(float32(int32val))
	}
	if int64val, ok := value.(int64); ok {
		return NumberLiteral(float32(int64val))
	}
	if numVal, ok := value.(*NumberValue); ok {
		return numVal
	}
	if timeVal, ok := value.(time.Time); ok {
		return &DateTimeValue{Value: timeVal}
	}
	return nil
}

// WithFunction creates a FormatContext with a single function
func WithFunction(key string, function Function) *FormatContext {
	return &FormatContext{
		variables: nil,
		functions: map[string]Function{strings.TrimSpace(strings.ToUpper(key)): function},
	}
}

// WithFunctions creates a FormatContext with multiple functions
func WithFunctions(functions map[string]Function) *FormatContext {
	cleaned := make(map[string]Function, len(functions))
	for name, function := range functions {
		cleaned[strings.TrimSpace(strings.ToUpper(name))] = function
	}
	return &FormatContext{
		variables: nil,
		functions: cleaned,
	}
}

// TODO: Builtin functions (NUMBER, DATETIME)
func assembleContexts(options ...*FormatContext) (map[string]Value, map[string]Function) {
	variables := make(map[string]Value)
	functions := make(map[string]Function)
	for _, option := range options {
		if option.variables != nil {
			for key, variable := range option.variables {
				variables[key] = variable
			}
		}
		if option.functions != nil {
			for key, function := range option.functions {
				functions[key] = function
			}
		}
	}

	functions["NUMBER"] = NumberFunc
	functions["DATETIME"] = DateTimeFunc
	functions["CLDRDATETIME"] = CLDRDateTimeFunc

	return variables, functions
}

// FormatMessage formats the message with the given key.
// To pass variables or functions, pass contexts created using WithVariable, WithVariables, WithFunction or WithFunctions.
// Besides the formatted message, this method returns the errors the resolver stumbled upon during resolving specific values
// and an optional error if there is no message with the given key.
// If the resolver returns errors it does not automatically mean that the whole message could not be resolved.
// It may be just incomplete.
func (bundle *Bundle) FormatMessage(key string, contexts ...*FormatContext) (string, []error, error) {
	if bundle.messages.Get(key) == nil {
		return "", nil, fmt.Errorf("message '%s' does not exist", key)
	}

	msg := bundle.messages.Get(key)
	variables, functions := assembleContexts(contexts...)

	// Add the bundle's functions to the resolver's functions
	for name, value := range bundle.functions {
		functions[name] = value
	}

	res := &resolver{
		bundle:          bundle,
		primaryLanguage: bundle.locales[0],
		params:          nil,
		variables:       variables,
		functions:       functions,
		errors:          []error{},
	}
	result := res.resolvePattern(msg.Value).String()
	if strings.TrimSpace(result) == "" || result == " " {
		result = key
	}
	return result, res.errors, nil
}

// Checks whether the bundle contains a message with the given key.
func (bundle *Bundle) HasMessage(key string) bool {
	return bundle.messages.Get(key) != nil
}
