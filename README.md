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

### コンテナイメージの命名規則とビルド

ghcr.io/<名前空間>/<イメージ名>:<イメージタグ>

- 個人アカウント名を環境変数にセット（container Registryへのログインなどで利用）
gh auth login

export GHCR_USER=$(gh config get -h github.com user)

- ビルド
docker build -t ghcr.io/${GHCR_USER}/example:latest docker/example/

### Container Registryへのログインとプッシュ

- GitHub CLIのクレデンシャルでGitHub Packagesへのアクセス権限追加
gh auth refresh --scopes write:packages

- login
gh auth token | docker login ghcr.io -u ${GHCR_USER} --password-stdin

- コンテナイメージをプッシュ
docker push ghcr.io/${GHCR_USER}/example:latest

- コンテナイメージのプル
docker pull ghcr.io/${GHCR_USER}/example:latest

### パッケージの自動リンクとパーミッションの継承

- 自動リンク
docker build -t ghcr.io/${GHCR_USER}/auto-link:latest \
--label "org.opencontainers.image.source=https://github.om/${GHCR_USER}/github-cicd
" \
docker/example/

docker push ghcr.io/${GHCR_USER}/auto-link:latest
