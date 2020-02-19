package advanced

import (
	"github.com/devlights/try-golang/advanced/async"
	"github.com/devlights/try-golang/advanced/closure"
	"github.com/devlights/try-golang/advanced/crypto"
	"github.com/devlights/try-golang/advanced/generate"
	"github.com/devlights/try-golang/advanced/reflection"
	"github.com/devlights/try-golang/advanced/sets"
	"github.com/devlights/try-golang/advanced/xdgspec"
	"github.com/devlights/try-golang/interfaces"
)

type (
	advancedExampleRegister struct{}
)

// NewRegister は、advanced パッケージ用の lib.Register を返します.
func NewRegister() interfaces.Register {
	r := new(advancedExampleRegister)
	return r
}

func (r *advancedExampleRegister) Regist(m interfaces.SampleMapping) {
	m["async01"] = async.Async01
	m["async_producer_consumer"] = async.ProducerConsumer
	m["closure01"] = closure.Closure01
	m["crypto_bcrypt_password_hash"] = crypto.BcryptPasswordHash
	m["generate_generic_stack"] = generate.UseGenericStack
	m["generate_generic_queue"] = generate.UseGenericQueue
	m["reflection01"] = reflection.Reflection01
	m["set01"] = sets.Set01
	m["set02"] = sets.Set02
	m["set03"] = sets.Set03
	m["set04"] = sets.Set04
	m["set05"] = sets.Set05
	m["xdg_base_directory"] = xdgspec.XdgBaseDirectory
	m["xdg_user_directory"] = xdgspec.XdgUserDirectory
	m["xdg_file_operation"] = xdgspec.XdgFileOperation
}
