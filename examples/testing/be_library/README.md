# beとは？

[nalgeon/be](https://github.com/nalgeon/be)は、シンプルでとても使いやすいテスト用アサートライブラリです。

作者のブログ記事は以下です。

- Expressive tests without testify/assert
  - https://antonz.org/do-not-testify/

関数は３つしか存在しない。

- Equal
- Err
- True

それぞれがジェネリック対応となっているため、どのような型でもアサート出来るようになっている。
