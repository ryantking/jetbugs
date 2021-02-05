# GoLand Coverage Window Bug

This repo demonstrates a bug where scopes don't work with multiple content
roots.

## GoLand Version

```
GoLand 2020.3.2
Build #GO-203.7148.71, built on January 26, 2021
Licensed to Solo.io, Inc. / Ryan King
You have a perpetual fallback license for this version.
Subscription is active until December 2, 2021.
Runtime version: 11.0.9.1+11-b1145.77 x86_64
VM: OpenJDK 64-Bit Server VM by JetBrains s.r.o.
macOS 10.16
GC: ParNew, ConcurrentMarkSweep
Memory: 3987M
Cores: 16
Registry: ide.completion.variant.limit=500, suggest.all.run.configurations.from.context=true
Non-Bundled Plugins: IdeaVIM, Key Promoter X, com.arcticicestudio.nord.jetbrains, com.intellij.tasks, com.vincentp.gruvbox-theme, mobi.hsz.idea.gitignore, name.kropp.intellij.makefile, com.intellij.kubernetes, org.toml.lang, ro.florinpatan.gopher, intellij.prettierJS, izhangzhihao.rainbow.brackets, PythonCore, idea.plugin.protoeditor, ru.adelf.idea.dotenv, org.rust.lang
```

## Bug Description

The primary issue that I can deduce from experimentation is that all content
roots are viewed by GoLand as one "module." I have found this to manifest
clearly in the following three use cases:

1. When setting scope, there is no way to select a content root as a scope, only
   the main content root will be visible, which includes all files.

2. When building a custom scope, all files in the top level of each content root
   behave as though they're all part of one directory. This leads to strings
   such as `file[module-name]:pkg//*` selecting files in all `pkg` directories
   across all content roots.

The issue boils down to when building the file tree, all content roots are
listed and merged, which makes doing operations like find/replace and refactors
a bad user experience when working with on a large project with many
repositories added as content roots, especially since Go projects share a lot of
conventional folder and file names.

## Reproduction

1. Clone this repository
2. Open the `dev` folder in GoLand, the other content roots should be set
   properly.
3. Start a project wide search and type `Config` as the search term. You will
   see that two of the content roots have a separate `Config` struct with a
   library using them.
4. Click on the "Module" tab. Note that only one module is available, which
   encompasses all files across all content roots.
5. Switch to the "Scope" tab, click on the "..." button to setup some custom
scopes, and create a new local one.
6. Now let's look at different patterns and their affect on the scope:
    - `file[JetBugs]:*`: All files in all content roots are selected.
    - `file[JetBugs]:pkg//*`: All files in both content roots' `pkg` directories
        are selected.
    - `file[JetBugs]:math//*`: Only files in `myapp`'s top level `math`
        directory are selected.
    - `file[JetBugs]:httputil//*`: Only files in `mylib`'s top level `httputil`
        directory are selected.
    - `file[myapp]:*`: No files are selected.
    - `file[mylib]:*`: No files are selected.
