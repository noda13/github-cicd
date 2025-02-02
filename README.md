# github-cicd

### cache key に、ファイルキャッシュを指定

key: test-${{ runner.os }}-${{ runner.arch }}-${{ hashFiles('**/package-lock.json') }}

### GitHub CLIによるリリースノートの作成

gh release create v0.1.0 --title "v0.1.0" --notes "Wonderful Text"

### GitHub CLIによるアセットのアップロード

gh release upload v0.1.0 example.txt

- リリースノート作成と同時実行も可能
gh release create v0.1.0 --title "v0.1.0" --notes "Wonderful Text" example.txt
