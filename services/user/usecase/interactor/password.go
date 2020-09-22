package interactor

func compareHashAndPass(truePass string, formPass string) error {
	// パスワードが正しいか確認している。ユーザー登録ができるようになってから検証
	// if err := bcrypt.CompareHashAndPassword([]byte(truePass), []byte(formPass)); err != nil {
	// 	return err
	// }
	return nil
}
