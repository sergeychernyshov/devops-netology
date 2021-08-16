#добавил сертификат на github
# корневым каталогом считаю каталог в котором находится файл .gitignor
**/.terraform/* - игнорируем во всех вложенных папках все файлы внутри папок .terraform

*.tfstate	- исключаем все файлы с раширением .tfstate во всех каталогах 

*.tfstate.* - исключаем все файлы, в названии которых присутствует .tfstate., 
во всех вложенных каталогах начиная от корневого и всех вложенных

crash.log - исключаем файл crash.log в корневом каталоге

*.tfvars - исключаем все файлы с раширением .tfvars во всех каталогах начиная от корневого и всех вложенных

override.tf	- исключаем файл override.tf в корневом каталоге

override.tf.json - исключаем файл override.tf.json в корневом каталоге

*_override.tf	- исключаем все файлы, которые заканчиваются на _override.tf во всех каталогах, 
начиная от корневого и всех вложенных

*_override.tf.json - исключаем все файлы, которые заканчиваются на _override.tf.json  во всех каталогах, 
начиная от корневого и всех вложенных

.terraformrc - исключаем файл  .terraformrc в корневом каталоге

terraform.rc - исключаем файл  terraform.rc в корневом каталоге 

Домашнее задание к занятию «2.4. Инструменты Git»

1) git show --format="%H" -s aefea

	aefead2207ef7e2aa5dc81a34aedf0cad4c32545
2) git show -s 85024d3

commit 85024d3100126de36331c6982bfaac02cdab9e76 (tag: v0.12.23)

3) git rev-list --parents -n 2 b8d720


b8d720f8340221f2146e4e4870bf2ee0bc48f2d5 56cd7859e05c36c06b56d013b55a252d0bb7e158 9ea88f22fc6269854151c571162c5bcf958bee2b


b8d720f8340221f2146e4e4870bf2ee0bc48f2d5 - текущий коммит


56cd7859e05c36c06b56d013b55a252d0bb7e158 9ea88f22fc6269854151c571162c5bcf958bee2b - родительские коммиты 

4) git rev-list v0.12.23..v0.12.24 --oneline 

33ff1c03b v0.12.24

b14b74c49 [Website] vmc provider links

3f235065b Update CHANGELOG.md

6ae64e247 registry: Fix panic when server is unreachable

5c619ca1b website: Remove links to the getting started guide's old location

06275647e Update CHANGELOG.md

d5f9411f5 command: Fix bug when using terraform login on Windows

4b6d06cc5 Update CHANGELOG.md

dd01a3507 Update CHANGELOG.md

225466bc3 Cleanup after v0.12.23 release

5) git log -S "func providersource("  --oneline --regexp-ignore-case

8c928e835 main: Consult local directories as potential mirrors of providers

6) git log -S "func globalPluginDirs"  --oneline --regexp-ignore-case

plugins.go

git log -L :globalPluginDirs:plugins.go --pretty=format:"%H" --no-patch

78b12205587fe839f10d946ea3fdc06719decb05

52dbf94834cb970b510f2fba853a5b49ad9b1a46

41ab0aef7a0fe030e84018973a64135b11abcd70

66ebff90cdfaa6938f26f908c7ebad8d547fea17

8364383c359a6b738a436d1b7745ccdce178df47

7) git log -S "synchronizedWriters"  --oneline --regexp-ignore-case

bdfea50cc remove unused
fd4f7eb0b remove prefixed io
5ac311e2a main: synchronize writes to VT100-faker on Windows

git show 5ac311e2a

git show --format="%H %an %ae %cn %ce" -s 5ac311e2a
5ac311e2a91e381e2f52234668b49ba670aa0fe5 Martin Atkins mart@degeneration.co.uk Martin Atkins mart@degeneration.co.uk