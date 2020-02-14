//Package passwordbox provides a textbox that hides its input with dots.
package passwordbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seeds/textbox"
)

func init() {
	var js, _ = Asset("argon2.js")

	seed.Embed("/argon2.js", js)

	seed.Embed("/argon2-browser.js", []byte(Lib))

	var wasm, _ = Asset("argon2.wasm")

	seed.Embed("/argon2.wasm", wasm)
}

//Seed is a passwordbox.
type Seed struct {
	textbox.Seed
}

//New passwordbox.
func New() Seed {
	var PasswordBox = textbox.New(textbox.Options{
		Discard: true,
	})

	PasswordBox.SetAttributes("type='password'")
	PasswordBox.Require("/argon2-browser.js")

	return Seed{PasswordBox}
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var PasswordBox = New()
	parent.Root().Add(PasswordBox)
	return PasswordBox
}

//Ctx is the script context to this seed.
type Ctx struct {
	textbox.Ctx
	with script.Args
}

//Ctx returns the script context to this seed.
func (passwordbox Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{passwordbox.Seed.Ctx(q), nil}
}

//Hash returns a securely hashed passkey promise from the password box.
//The hostname is used as a salt, so do not change the hostname or users will be unable to login.
func (passwordbox Ctx) Hash() script.Promise {
	var q = passwordbox.Q
	var value = passwordbox.Ctx.Value()
	return q.Value(`argon2.hash({
		// required
		pass: ` + value.LanguageType().Raw() + `,
		salt: window.location.hostname,

		time: 1,
		mem: 1024*64,
		hashLen: 32,
		parallelism: 1,
		type: argon2.ArgonType.Argon2di,
		distPath: '' // asm.js script location, without trailing slash
	}).then(function(result) {
		return result.hashHex;
	})`).Promise()
}

//Value returns the password box's value securely hashed.
func (passwordbox Ctx) Value() script.String {
	return passwordbox.Hash().Wait().String()
}

//With returns a context with the provided args used for HashAndGo calls.
func (passwordbox Ctx) With(args script.Args) Ctx {
	passwordbox.with = args
	return passwordbox
}
