# GoLand Vendor Bug

This repo is meant to demonstrate a bug that occurs when vendoring a repo into a
different repo such that both are content roots in a GoLand project.

## GoLand Version

```
GoLand 2020.2.1
Build #GO-202.6397.124, built on August 4, 2020
Licensed to XXXX
Subscription is active until May 23, 2021
Runtime version: 11.0.7+10-b944.20 x86_64
VM: OpenJDK 64-Bit Server VM by JetBrains s.r.o.
macOS 10.15.6
GC: ParNew, ConcurrentMarkSweep
Memory: 1979M
Cores: 12
Registry: ide.completion.variant.limit=500, debugger.watches.in.variables=false, suggest.all.run.configurations.from.context=true, ide.balloon.shadow.size=0, ideFeaturesTrainer.welcomeScreen.tutorialsTree=TRUE
Non-Bundled Plugins: BashSupport, IdeaVIM, Key Promoter X, com.alayouni.ansiHighlight, com.arcticicestudio.nord.jetbrains, com.intellij.ideolog, com.jetbrains.plugins.ini4idea, org.jetbrains.plugins.sass, ideanginx9, name.kropp.intellij.makefile, net.seesharpsoft.intellij.plugins.csv, org.jetbrains.plugins.hocon, mobi.hsz.idea.gitignore, com.intellij.kubernetes, org.zalando.intellij.swagger, org.toml.lang, NodeJS, com.intellij.plugins.html.instantEditing, intellij.prettierJS, AngularJS, com.dmarcotte.handlebars, izhangzhihao.rainbow.brackets, PythonCore, ru.adelf.idea.dotenv, org.rust.lang
```

## Bug Description

The common repo configuration that I use for developing a system with several
microservices is as follows:

- An API gateway repo that has the API as well as all the common packages such
    as models, auth, errors, etc.
- A repo per additional microservice
- A "dev" repo that contains the GoLand project, dev documentation, common
    resources such as HTTP requests, docker compose files, etc. that has the
    other repos as content roots.

The common packages of the API repo are vendored into the other microservice
repos.

When trying to add new tests to one of the common packages in the API repository,
GoLand seems to place priority on the vendored packages in the other
repositories when running a test from the sidebar. The following error was
outputted:

```
/usr/local/opt/go/libexec/bin/go tool test2json -t /private/var/folders/ft/pyybg1qd53d7sgz0f6sfydlw0000gp/T/___TestSquare2_in_github_com_ryantking_jetbugs_mylib -test.v -test.run ^TestSquare2$
Process finished with exit code 0
testing: warning: no tests to run
PASS
```

Then, eventually the error switched to:

```
test2json: fork/exec /private/var/folders/ft/pyybg1qd53d7sgz0f6sfydlw0000gp/T/___TestSquare_in_github_com_ryantking_jetbugs_mylib: exec format error
```

## Confirmation

Both errors were resolved by deleting the vendor folder then returned again
after rerunning `go mod vendor` from the microservice repository.

Another way I tested was by replacing the `go` binary used by GoLand with a bash
script that outputted the go environment, and I noticed `GOMOD` was set
incorrectly to the vendored module:

```
GOMOD="/Users/rking/Projects/jetbugs/myapp/vendor/github.com/ryantking/jetbugs/mylib/go.mod"
```

What is strange is that it was only set incorrectly during the `go test -c`
setup call, but set correctly during the `go tool test2json` call.

## Reproduction

1. Clone this repository
2. Open the `dev` folder in GoLand, the other content roots should be set
   properly.
3. Try to run `TestSquare` in `mylib_test.go`, it should fail with one of the
   two above errors.
4. Delete the vendor folder in `myapp` then run the test again, it should pass

