# Changelog

All notable changes to this project will be documented in this file.

## [0.2.16](https://github.com/inference-gateway/documentation-agent/compare/v0.2.15...v0.2.16) (2025-09-18)

### 🐛 Bug Fixes

* Update Docker build command to conditionally tag images with 'latest' ([ec33aad](https://github.com/inference-gateway/documentation-agent/commit/ec33aad69518e3befe40c54a5dbbc190e9f458f5))
* Update Docker build command to use 'index:' prefix for image annotations ([8263c38](https://github.com/inference-gateway/documentation-agent/commit/8263c38bce0bdb3bd14737974188ff06c9863f51))

## [0.2.15](https://github.com/inference-gateway/documentation-agent/compare/v0.2.14...v0.2.15) (2025-09-18)

### 🐛 Bug Fixes

* Update Docker build command to use annotations instead of labels for image metadata ([ca803d0](https://github.com/inference-gateway/documentation-agent/commit/ca803d0c3e3ae15645d9dc42bb973a7523290bcb))

## [0.2.14](https://github.com/inference-gateway/documentation-agent/compare/v0.2.13...v0.2.14) (2025-09-18)

### ♻️ Improvements

* Update Docker build command to use a variable for image description ([3de82c6](https://github.com/inference-gateway/documentation-agent/commit/3de82c6f56f51e74c0f494bf0e94da34e576ef3c))

### 🐛 Bug Fixes

* Correct Docker build command to use consistent variable syntax for image source label ([ee79757](https://github.com/inference-gateway/documentation-agent/commit/ee7975765370d8a2e12bd0f71c072866ed6a9c20))

## [0.2.13](https://github.com/inference-gateway/documentation-agent/compare/v0.2.12...v0.2.13) (2025-09-18)

### 🐛 Bug Fixes

* Update Docker build command to use lowercase repository name for image source label ([5c33637](https://github.com/inference-gateway/documentation-agent/commit/5c33637ee55320e3460105560b128f1215639d6b))

## [0.2.12](https://github.com/inference-gateway/documentation-agent/compare/v0.2.11...v0.2.12) (2025-09-18)

### 🐛 Bug Fixes

* Improve Docker build command with OCI image labels and conditional logic for release channels ([890f2c4](https://github.com/inference-gateway/documentation-agent/commit/890f2c45726cb39f984fa75c3b47a682994d2a06))

## [0.2.11](https://github.com/inference-gateway/documentation-agent/compare/v0.2.10...v0.2.11) (2025-09-18)

### ♻️ Improvements

* Add OCI image labels for container metadata in Dockerfile ([352a3eb](https://github.com/inference-gateway/documentation-agent/commit/352a3eb99ebfcb4b62a98e364646efc228a41116))
* Update agent description to accurately reflect functionality across documentation ([32cbdce](https://github.com/inference-gateway/documentation-agent/commit/32cbdced372429c3ccd4a8f38d86a474cd223bb6))

### 🐛 Bug Fixes

* Update ADL CLI version to 0.20.13 in generated files ([e59cfb8](https://github.com/inference-gateway/documentation-agent/commit/e59cfb8c087cd85e357891b733f2d0ddb0bd4579))

## [0.2.10](https://github.com/inference-gateway/documentation-agent/compare/v0.2.9...v0.2.10) (2025-09-18)

### 🐛 Bug Fixes

* integrate container publishing into semantic-release workflow ([9d7b3c7](https://github.com/inference-gateway/documentation-agent/commit/9d7b3c71fd48c0580aadc28ff40def1a269ae493))

## [0.2.9](https://github.com/inference-gateway/documentation-agent/compare/v0.2.8...v0.2.9) (2025-09-17)

### 🐛 Bug Fixes

* Correct Docker image tag formatting in publish command ([c0d8a6a](https://github.com/inference-gateway/documentation-agent/commit/c0d8a6ad1b3847defb1e66e155819dd4e678b510))

## [0.2.8](https://github.com/inference-gateway/documentation-agent/compare/v0.2.7...v0.2.8) (2025-09-17)

### ♻️ Improvements

* Improve CI/CD pipeline with Docker image build and push steps ([78677ac](https://github.com/inference-gateway/documentation-agent/commit/78677ac5f808307b877de580592ea073eb44b7e6))

### 🐛 Bug Fixes

* Add installation step for ADL CLI in CI/CD workflow ([e84279c](https://github.com/inference-gateway/documentation-agent/commit/e84279cf34804a994a9875e7d6bbf8a129682090))

### 🔧 Miscellaneous

* Run task genrate - update ADL CLI version to 0.20.12 and increment agent version to 0.2.7 ([2b0213a](https://github.com/inference-gateway/documentation-agent/commit/2b0213a915042b904704f37bfd265eb94c786419))

## [0.2.7](https://github.com/inference-gateway/documentation-agent/compare/v0.2.6...v0.2.7) (2025-09-17)

### ♻️ Improvements

* Make codegen less noisy ([a2062a7](https://github.com/inference-gateway/documentation-agent/commit/a2062a737b1a1979dd7f0c1bc823b0ab3fb1af68))

### 🐛 Bug Fixes

* Update Go version to 1.24.5 and increment agent version to 0.2.6 ([6741726](https://github.com/inference-gateway/documentation-agent/commit/674172668e2fd5e4a3b8f8a32e4c3db54ccc85de))

## [0.2.6](https://github.com/inference-gateway/documentation-agent/compare/v0.2.5...v0.2.6) (2025-09-17)

### ♻️ Improvements

* Add @semantic-release/exec plugin for version update in agent.yaml ([4c99456](https://github.com/inference-gateway/documentation-agent/commit/4c99456e937c13b2c2bc518f4d8d8c2c98997994))

### 🎨 Miscellaneous

* Format tags in resolve_library_id skill for consistency ([35f05f9](https://github.com/inference-gateway/documentation-agent/commit/35f05f917361a37d3367c7ba91ddbbc295ddbcbf))
* Format YAML arrays for default input/output modes and library tags ([e6b6609](https://github.com/inference-gateway/documentation-agent/commit/e6b6609bc462666ae0d673e1f8029ff44f262822))

## [0.2.5](https://github.com/inference-gateway/documentation-agent/compare/v0.2.4...v0.2.5) (2025-09-17)

### ♻️ Improvements

* Update generated files to reflect ADL CLI v0.20.7 and bump ADK version to v0.11.0 ([9f25c0c](https://github.com/inference-gateway/documentation-agent/commit/9f25c0ccb8cc125adee61f574070845b7bf97170))

## [0.2.4](https://github.com/inference-gateway/documentation-agent/compare/v0.2.3...v0.2.4) (2025-09-16)

### ♻️ Improvements

* Update generated files to reflect ADL CLI v0.20.3 and enhance logging configuration ([8919123](https://github.com/inference-gateway/documentation-agent/commit/891912331441051f80996d86153740ff97b06cd6))

## [0.2.3](https://github.com/inference-gateway/documentation-agent/compare/v0.2.2...v0.2.3) (2025-09-16)

### ♻️ Improvements

* Bump ADK version to 0.10.1 ([dc2b3ce](https://github.com/inference-gateway/documentation-agent/commit/dc2b3cee736a65dc2b59dc3c14e9410119b54cf1))

## [0.2.2](https://github.com/inference-gateway/documentation-agent/compare/v0.2.1...v0.2.2) (2025-09-12)

### ♻️ Improvements

* Re-generate files ([4f8a4cf](https://github.com/inference-gateway/documentation-agent/commit/4f8a4cf69dbe67039271a65138cab71d82ff8be8))
* Re-run generate - update deployment instructions and remove Kubernetes manifests ([b52416e](https://github.com/inference-gateway/documentation-agent/commit/b52416e5c46f2bc52afb6e32973e630daf84bf75))
* Remove PRD.md as it is no longer needed ([97335cc](https://github.com/inference-gateway/documentation-agent/commit/97335cca887c86efcc15fd4cd5d9587e555162d1))
* Update documentation and configuration files; add agent capabilities and deployment instructions ([3bfe662](https://github.com/inference-gateway/documentation-agent/commit/3bfe6622f8085c86d8bf4a59b6500412278c2b54))
* Update generated files to reflect ADL CLI v0.19.12 and clean up documentation ([62a1144](https://github.com/inference-gateway/documentation-agent/commit/62a1144f6629e1747e365b7b8f814fc0a9e8b1af))

## [0.2.1](https://github.com/inference-gateway/documentation-agent/compare/v0.2.0...v0.2.1) (2025-09-03)

### 🐛 Bug Fixes

* Correct output variable syntax in release job ([cf947a2](https://github.com/inference-gateway/documentation-agent/commit/cf947a2ba42a6d6a2d13ba34b5c3a461abbf5449))

## [0.2.0](https://github.com/inference-gateway/documentation-agent/compare/v0.1.0...v0.2.0) (2025-09-03)

### ✨ Features

* Improve CD workflow with GitHub App token management and user ID configuration ([2db29c9](https://github.com/inference-gateway/documentation-agent/commit/2db29c9672e35501b37f028b3da2ec6ce90db4ac))

### 🐛 Bug Fixes

* Update GITHUB_TOKEN to use GitHub App token for release creation ([52ea5b6](https://github.com/inference-gateway/documentation-agent/commit/52ea5b6af7152ce9128132a10975de78c8ccdd9a))

### 🔧 Miscellaneous

* Remove document that was previously relevant ([19a6003](https://github.com/inference-gateway/documentation-agent/commit/19a60038c9ee44237b50cb32d581b2c8d657e410))

### 🔨 Miscellaneous

* Add claude-code package configuration to manifest files ([1a3d4ee](https://github.com/inference-gateway/documentation-agent/commit/1a3d4eea5b6b7aff9a4ca1f7fa1a29b5fc76c11b))
