package processor

import "github.com/pkg/errors"

type Processor struct {
	Validator    Validator
	SchemaLoader SchemaLoader
	Parser       Parser
}

type Validator interface {
	ValidateData(data, schema []byte) error
	ValidateDocument(doc, schema []byte) error
}

type SchemaLoader interface {
	Load(url string) (schema []byte, extension string, err error)
}

type Parser interface {
	ParseSlots(data, schema []byte) (index, value []byte, err error)
}

var (
	errParserNotDefined    = errors.New("parser is not defined")
	errLoaderNotDefined    = errors.New("loader is not defined")
	errValidatorNotDefined = errors.New("validator is not defined")
)

// Opt returns configuration options for processor suite
type Opt func(opts *Processor)

// WithValidator return new options
func WithValidator(s Validator) Opt {
	return func(opts *Processor) {
		opts.Validator = s
	}
}

// WithSchemaLoader return new options
func WithSchemaLoader(s SchemaLoader) Opt {
	return func(opts *Processor) {
		opts.SchemaLoader = s
	}
}

// WithParser return new options
func WithParser(s Parser) Opt {
	return func(opts *Processor) {
		opts.Parser = s
	}
}

// InitProcessorOptions initializes processor with options.
func InitProcessorOptions(processor *Processor, opts ...Opt) *Processor {
	for _, opt := range opts {
		opt(processor)
	}
	return processor
}

// Load will load a schema by given url.
func (s *Processor) Load(url string) (schema []byte, extension string, err error) {
	if s.SchemaLoader == nil {
		return nil, "", errLoaderNotDefined
	}
	return s.SchemaLoader.Load(url)
}

// ParseSlots will serialize input data to index and value fields.
func (s *Processor) ParseSlots(data []byte, schema []byte) (index []byte, value []byte, err error) {
	if s.Parser == nil {
		return nil, nil, errParserNotDefined
	}
	return s.Parser.ParseSlots(data, schema)
}

// ValidateData will validate a claim content by given schema.
func (s *Processor) ValidateData(data, schema []byte) error {
	if s.Validator == nil {
		return errValidatorNotDefined
	}
	return s.Validator.ValidateData(data, schema)
}

// ValidateDocument will validate a document content by given schema.
func (s *Processor) ValidateDocument(data, schema []byte) error {
	if s.Validator == nil {
		return errValidatorNotDefined
	}
	return s.Validator.ValidateDocument(data, schema)
}
