# GoLand Coverage Window Bug

This repo demonstrates a bug where the coverage suite window only works for tests in the same content root as the `.idea` project folder. Coverage suits for tests in other content roots are not found

## GoLand Version

```
GoLand 2020.2.2
Build #GO-202.6948.9, built on August 11, 2020
Licensed to XXXX
Subscription is active until May 23, 2021
Runtime version: 11.0.8+10-b944.31 x86_64
VM: OpenJDK 64-Bit Server VM by JetBrains s.r.o.
macOS 10.15.6
GC: ParNew, ConcurrentMarkSweep
Memory: 1979M
Cores: 12
Registry: ide.completion.variant.limit=500, debugger.watches.in.variables=false, suggest.all.run.configurations.from.context=true, ide.balloon.shadow.size=0, ideFeaturesTrainer.welcomeScreen.tutorialsTree=TRUE
Non-Bundled Plugins: BashSupport, IdeaVIM, Key Promoter X, com.alayouni.ansiHighlight, com.arcticicestudio.nord.jetbrains, com.intellij.ideolog, com.jetbrains.plugins.ini4idea, org.jetbrains.plugins.sass, com.paperetto.dash, ideanginx9, name.kropp.intellij.makefile, net.seesharpsoft.intellij.plugins.csv, org.jetbrains.plugins.hocon, mobi.hsz.idea.gitignore, com.intellij.kubernetes, org.zalando.intellij.swagger, org.toml.lang, NodeJS, com.intellij.plugins.html.instantEditing, intellij.prettierJS, AngularJS, com.dmarcotte.handlebars, izhangzhihao.rainbow.brackets, PythonCore, ru.adelf.idea.dotenv, org.rust.lang
```

## Bug Description

Test suites that exist outside of the main content root do not show up in the test coverage window as one would expect. It seems like the test coverage window only picks up coverage profile files from the "main" content root (where the `.idea` folder is) so any test run with coverage in a different content root do not show up in the coverage window.

What does work is:

- Coverage gutter indicator
- Coverage indication in the file tree sidebar

## Reproduction

1. Clone this repository
2. Open the `dev` folder in GoLand, the other content roots should be set
   properly.
3. Run the `Test in the same folder as .idea` test configuration with coverage, click "Replace with Suite" if prompted.
4. Note that the window looks as it should.
5. Run the `Test in other content root` test configuration, click "Replace with Suite" if prompted.
6. Note that the test coverage is not displayed in the window.
