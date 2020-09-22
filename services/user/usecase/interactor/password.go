package interactor

import "github.com/takuya911/go_pf/services/user/errors"

func compareHashAndPass(truePass string, formPass string) error {
	// パスワードが正しいか確認している。ユーザー登録ができるようになってから検証
	// if err := bcrypt.CompareHashAndPassword([]byte(truePass), []byte(formPass)); err != nil {
	// 	return err
	// }
	if formPass != truePass {
		return errors.BadRequestError
	}
	return nil
}
