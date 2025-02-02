# github-cicd

### cache key に、ファイルキャッシュを指定

key: test-${{ runner.os }}-${{ runner.arch }}-${{ hashFiles('**/package-lock.json') }}

### GitHub CLIによるリリースノートの作成

gh release create v0.1.0 --title "v0.1.0" --notes "Wonderful Text"

### GitHub CLIによるアセットのアップロード

gh release upload v0.1.0 example.txt

- リリースノート作成と同時実行も可能
gh release create v0.1.0 --title "v0.1.0" --notes "Wonderful Text" example.txt

### プルリクエストへのラベル付与

- １つ目のプルリクエスト：ラベルは付与せず、リリースノートから除外
- enhancementラベルを付与し、リリースノートへ記載

```
git switch -c add-release-setting main
git add .github/release.yml
git commit -m "Add release setting"
git push origin add-release-setting
gh pr create --fill-first
gh pr merge --merge
```

```
git switch -c enhancement-pr main
git commit -m "Add super useful features" --allow-empty
git push origin enhancement-pr
gh pr create --fill-first --label "enhancement"
gh pr merge --merge
```

gh release create v0.2.0 --title "v0.2.0" --generate-notes

### タグを作成してプッシュ

git tag v0.3.0
git push origin v0.3.0
