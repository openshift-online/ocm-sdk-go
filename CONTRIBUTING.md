# Contributing to the OCM SDK

### Validating model updates

The easiest way to validate changes to the ocm-api-model is to generate the sdk based on the updated model.
Ensure ocm-sdk-go is cloned locally alongside your cloned ocm-api-model directory where changes are made.

In ocm-sdk-go, run the following to generate the sdk using the local ocm-api-model. Replace
`/path/to/ocm-api-model` with the path to your cloned ocm-api-model repository where changes have been made.
Replace `HEAD` with the commit SHA in ocm-api-model to build against if necessary.

```shell
make clean
make model_version=HEAD model_url=/path/to/ocm-api-model generate
```

Review the output for errors. If none, the changes are at least syntactically proper.

### Style notes

Comments can be added to types with the `//` notation. Comments should be added to each type's attribute where
possible.

Hard tabs are required rather than spaces.

## Releasing a new OCM API Model version

To use any updates to the [ocm-api-model](https://github.com/openshift-online/ocm-api-model), the version
must be incremented for consumption in ocm-sdk-go generation. The version is defined by the latest git tag.

Once all changes to the OCM API Model have been defined and reviewed the client types for the model need to be generated via `make update` target
in the `ocm-api-model` project.

Once all changes to the OCM API Model have been committed to the main branch you will need to update these changes in the SDK.

## Updating the OCM SDK

The OCM SDK can be generated simply by running the following after all changes have been made:

```shell
./hack/update-model.sh
make update
```

The `./hack/update-model.sh` script will ensure the `ocm-api-model` modules are all up to date with the latest version across the OCM-SDK project.
To verify that they are all in-sync one can use the `./hack/verify-model-version.sh` script.

One can add an optional commit SHA or version to the `./update-model.sh <vX.Y.Z>` script to update the go modules to a specific version.

Whenever an update is made, ensure that the corresponding example in [examples](examples) is also updated where
necessary. It is *highly recommended* that new endpoints have a new example created.

## Releasing a new OCM SDK Version

Releasing a new version requires submitting an MR for review/merge with an update to the `Version` constant in
[version.go](version.go). Additionally, update the [CHANGES.md](CHANGES.md) file to include the new version and
describe all changes included.

Below is an example CHANGES.md update:

```
== 0.1.39 Oct 7 2019

- Update to model 0.0.9:
  - Add `type` attribute to the `ResourceQuota` type.
  - Add `config_managed` attribute to the `RoleBinding` type.
```

Submit an MR for review/merge with the CHANGES.md and version.go update.

Finally, create and submit a new tag with the new version following the below example:

```shell
git checkout main
git pull
git tag -a -m 'Release 0.1.39' v0.1.39
git push origin v0.1.39
```

Note that a repository administrator may need to push the tag to the repository due to access restrictions.
