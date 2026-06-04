BOM（バイトオーダーマーク）は、Unicodeエンコーディングで符号化されたテキストファイルの先頭に付加される特殊なバイト列です。

バイナリ表現は以下のようになります：

- UTF-8 BOM:
  - バイナリ: 0xEF 0xBB 0xBF
- UTF-16 BE (Big Endian) BOM:
  - バイナリ: 0xFE 0xFF
- UTF-16 LE (Little Endian) BOM:
  - バイナリ: 0xFF 0xFE
- UTF-32 BE (Big Endian) BOM:
  - バイナリ: 0x00 0x00 0xFE 0xFF
- UTF-32 LE (Little Endian) BOM:
  - バイナリ: 0xFF 0xFE 0x00 0x00

BOMは主に以下の目的で使用されます：

- ファイルの文字エンコーディングを示す
- バイトオーダー（エンディアン）を示す（UTF-16やUTF-32の場合）

ただし、BOMの使用には注意が必要です。特にUTF-8の場合、BOMの使用は推奨されておらず、一部のシステムで問題を引き起こす可能性があります。**多くの場合、UTF-8エンコードされたファイルにはBOMを付加しないことが一般的です。**