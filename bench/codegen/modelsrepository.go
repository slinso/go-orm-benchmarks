package codegen

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ModelsRepository

// GENERATED BY CODEGEN.
/* ModelsRepository interface definition. */
type ModelsRepository interface {
	Create(data *Models) error
	Update(data *Models) error
	UpdatePartial(data *ModelsPartial) error
	Delete(data *Models) error
	OneByID(id int) (*Models, error)

	// ^^ END OF GENERATED BY CODEGEN. ^^
}
