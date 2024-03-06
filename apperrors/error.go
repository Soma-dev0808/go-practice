package apperrors

type MyAppError struct {
	ErrCode			// レスポンスとログに表示するエラーコード
	Message string  // レスポンスに表示するエラーメッセージ
	Err     error `json:"-"`   // エラーチェーンのための内部エラー
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// エラーチェーンのための内部エラー
// Unwrapを呼び出すことで入れ子になっているエラーを取得できるように実装
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
