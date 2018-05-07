# Schema

## Version v0.16.2

The Schema repository holds the protocol buffer message schemas used to exchange data between systems.
The commands below will create Protobuf classes for use in iOS, Python, Golang, and JVM (Clojure, Java, Android) runtimes.
Javascript will use the `actions.json` and `events.json` since the contents of these files are already native to the language.

## Building Schema Artifacts
- sh and Bash (**version 4+**) are required to build (available on
most Linux distributions by default)
- on OS X
  - `brew install bash` to get v4.
  - make sure scripts run with bash v4 -- can likely be accomplished by changing script headers to be `#!/usr/local/bin/bash`
- Cigwin may work on Windows, but it has not been
tested

### Commits merged to `master` **must** increment the version in `script/version.json`.

### Also run golang.sh to update the go source as it's not auto-generated and deployed by Jenkins

---
The schema definitions are contained in `actions.json` and `events.json`. For artifacts that need/use a version number,
the version is in `script/version.json`. Any PRs that make changes to the schema (or add/delete schema definition files)
should always bump the version in `script/version.json`. The helper scripts (documented below) will convert the schema
JSON files into appropriate source for the target language/platform and compile the artifact, versioning them where
applicable. The bulk of the conversion logic is contained and documented in `script/index.sh`.

Currently, there is a Jenkins job that will automatically deploy the Java and Python artifacts, located here:

https://ciapi.boundlessgeo.io/job/boundless-schema-deploy

Any time a commit is pushed to `master`, that job will pick up the changes and deploy new artifacts.

### Java
This will put a schema jar in your local maven repo

```
sh script/java.sh
```

After running the above command, if you want to publish the artifact,

```
cd java
mvn deploy
```
> Note: your ~/.m2/settings.xml must be setup with authentication
> credentials required to publish to the Boundless maven repo


### Objective-C
This will build to a folder called `objc_build` in the root dir

```
sh script/objc.sh
```

### Python
This will build a distribution in the `python/dist` directory

```
bash script/python.sh
```

### Golang
Places go source in `pkg/schema`

```
sh script/golang.sh
```