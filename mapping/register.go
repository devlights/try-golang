package mapping

type (
	// Register は、各パッケージ毎のサンプルを登録するためのインターフェースです.
	Register interface {
		// Regist -- 指定されたマッピングに自身のサンプル情報を登録します.
		Regist(m ExampleMapping)
	}
)
