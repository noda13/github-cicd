# GitHub Actions

## ワークフローファイル

```sh
$ mkdir .github/workflows
$ touch .github/workflows/hello.yml
```

## GitHub Actionsの構成要素

- workflow
  1ファイルに１つ
- event
  トリガー
- job
  実行単位
- runner
  実行環境（Ubuntu, Windows, macOS）
- step
  処理の最小単位

## GitHub Actionsの実行

```sh
$ git add .github/workflows/hello.yml
$ git commit -m "Add hello.yml"
$ git push
```

実行ログの確認

```sh
$ git run watch
$ git run view
```

## workflowの手動実行

イベントに`workflow-dispatch`を指定する

- github cliから起動
```sh
$ gh workflow run manual.yml -f greeting=goodbye
```

## アクションの探し方

[GitHub Marketplace](https://github.com/marketplace)

`Creator verified by GitHub`のものが比較的安全性が高い

## GitHub Actionsの課金モデル

- 使用時間
  - 実行時間 x ランナーごとの乗率
- ランナーの乗率
  - Ubuntu  1
  - Windows 2
  - macOS 10
