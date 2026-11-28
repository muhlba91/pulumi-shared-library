# Changelog

## [2.0.0](https://github.com/muhlba91/pulumi-shared-library/compare/v1.0.0...v2.0.0) (2026-11-28)


### ⚠ BREAKING CHANGES

* gitlab actions secret.Create returns (pulumi.Output, error)
* github actions secret.Create returns an error and fails on a failed creation

### Features

* add iam condition and role attachment ([ae63ee6](https://github.com/muhlba91/pulumi-shared-library/commit/ae63ee6c9e5f39aad6dfdccc4ad5d33b82ee1810))
* return error from gitlab secret create ([46be2ca](https://github.com/muhlba91/pulumi-shared-library/commit/46be2caf9c41e9c47917bcac2dbd6477cfc48cd9))


### Bug Fixes

* align role attachment docs and naming ([5ef4b74](https://github.com/muhlba91/pulumi-shared-library/commit/5ef4b74a2b5cdfc2c97e69f8488bcd40e997febd))
* **ci:** add mise lock sidecars ([4764f08](https://github.com/muhlba91/pulumi-shared-library/commit/4764f088f0cefcad2f77cdbdb27ab9588043f29a))


### Miscellaneous Chores

* align code and docs to the style guide ([a0ab0e0](https://github.com/muhlba91/pulumi-shared-library/commit/a0ab0e04d7f6c76b785172bfee4dd77d67b9efa2))
* **ci:** add actionlint job ([c201cdf](https://github.com/muhlba91/pulumi-shared-library/commit/c201cdf763a0217b7c0d67e571011a849f826c8c))
* **deps:** update github/codeql-action action to v4.39.5 ([1a73fb4](https://github.com/muhlba91/pulumi-shared-library/commit/1a73fb4a472d9fe719b6ef1f9383e37cba55e7b2))
* **deps:** update module github.com/pulumi/pulumi-kubernetes/sdk/v4 to v4.35.0 ([8293486](https://github.com/muhlba91/pulumi-shared-library/commit/8293486f6101e80699ea18b30c5c6ed362499a66))
* **deps:** update module github.com/pulumi/pulumi-pulumiservice/sdk to v1.10.0 ([90f3eda](https://github.com/muhlba91/pulumi-shared-library/commit/90f3eda6363ec21235bf12b9d88f06940816ef27))
* **deps:** update module github.com/pulumi/pulumi-pulumiservice/sdk to v1.10.1 ([f6173b5](https://github.com/muhlba91/pulumi-shared-library/commit/f6173b5fa29cbd81282233c4f889bc5bdea49e3c))
* **deps:** update module github.com/pulumi/pulumi-pulumiservice/sdk to v1.9.0 ([447f911](https://github.com/muhlba91/pulumi-shared-library/commit/447f91130882e714b523241a049dc9e8a9686e79))
* **deps:** update module github.com/pulumi/pulumi/sdk/v3 to v3.278.0 ([050e990](https://github.com/muhlba91/pulumi-shared-library/commit/050e9906384126298c6078b9f583a96cac30c88f))
* **deps:** update pre-commit hook gitleaks/gitleaks to v8.30.1 ([d0b1ce9](https://github.com/muhlba91/pulumi-shared-library/commit/d0b1ce99a44f691dd636be09a744d092e7d8f00e))
* **release:** set bootstrap sha ([304805b](https://github.com/muhlba91/pulumi-shared-library/commit/304805b7b3385b2fdbe6fd3cf4e947b3d5b7e579))
* update pre-commit config ([a3f667f](https://github.com/muhlba91/pulumi-shared-library/commit/a3f667fb3536885fc74415e00e07a5b516aa2bdf))
