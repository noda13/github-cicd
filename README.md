# github-cicd

### cache key に、ファイルキャッシュを指定

key: test-${{ runner.os }}-${{ runner.arch }}-${{ hashFiles('**/package-lock.json') }}
