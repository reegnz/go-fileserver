build:
	goreleaser release --snapshot --skip-publish --rm-dist

clean:
	rm -rf dist
