# これは何？

[gitignore.io](https://www.toptal.com/developers/gitignore/) のAPIを利用してプログラム言語の ```.gitignore``` ファイルを作成するツールです。

windows用とlinux用のバイナリを用意しています。

- Windows
  - bin/windows/amd64/gitignore.exe
- Linux
  - bin/linux/amd64/gitignore

## 使い方

```sh
$ ./gitignore.exe help
gitignore - .gitignoreを生成するツール

  Usage:
    gitignore [lang|list]

  Subcommands: 
    lang   指定した言語で .gitignore を出力
    list   利用可能言語リストを出力

  Flags: 
       --version   Displays the program version string.
    -h --help      Displays help with available flag, subcommand, and positional value parameters.
```

サブコマンドには ```lang``` と ```list``` があります。

例として java の ```.gitignore``` を出力したい場合は

```sh
$ ./gitignore.exe lang java > .gitignore
```

Cの場合は

```sh
$ ./gitignore.exe lang c > .gitignore
```

C++の場合は

```sh
$ ./gitignore.exe lang c++ > .gitignore
```

Goの場合は

```sh
$ ./gitignore.exe lang go > .gitignore
```

Pythonの場合は

```sh
$ ./gitignore.exe lang python > .gitignore
```

.NETの場合は

```sh
$ ./gitignore.exe lang visualstudio > .gitignore
```

対応している言語のリストを見る場合は ```list``` を使います。

```sh
$ ./gitignore.exe list
1c,1c-bitrix,a-frame,actionscript,ada
adobe,advancedinstaller,adventuregamestudio,agda,al
alteraquartusii,altium,amplify,android,androidstudio
angular,anjuta,ansible,ansibletower,apachecordova
apachehadoop,appbuilder,appceleratortitanium,appcode,appcode+all
appcode+iml,appengine,aptanastudio,arcanist,archive
archives,archlinuxpackages,asdf,aspnetcore,assembler
astro,ate,atmelstudio,ats,audio
autohotkey,automationstudio,autotools,autotools+strict,awr
azurefunctions,azurite,backup,ballerina,basercms
basic,batch,bazaar,bazel,bitrise
bitrix,bittorrent,blackbox,blender,bloop
bluej,bookdown,bower,bricxcc,buck
c,c++,cake,cakephp,cakephp2
cakephp3,calabash,carthage,certificates,ceylon
cfwheels,chefcookbook,chocolatey,circuitpython,clean
clion,clion+all,clion+iml,clojure,cloud9
cmake,cocoapods,cocos2dx,cocoscreator,codeblocks
codecomposerstudio,codeigniter,codeio,codekit,codesniffer
coffeescript,commonlisp,compodoc,composer,compressed
compressedarchive,compression,conan,concrete5,coq
cordova,craftcms,crashlytics,crbasic,crossbar
crystal,cs-cart,csharp,cuda,cvs
cypressio,d,dart,darteditor,data
database,datarecovery,dbeaver,dbt,defold
delphi,deno,dframe,diff,direnv
diskimage,django,dm,docfx,docpress
docusaurus,docz,dotenv,dotfilessh,dotnetcore
dotsettings,doxygen,dreamweaver,dropbox,drupal
drupal7,drupal8,e2studio,eagle,easybook
eclipse,eiffelstudio,elasticbeanstalk,elisp,elixir
elm,emacs,ember,ensime,episerver
erlang,espresso,executable,exercism,expressionengine
extjs,fancy,fastlane,finale,firebase
fish,flashbuilder,flask,flatpak,flex
flexbuilder,floobits,flutter,font,fontforge
forcedotcom,forgegradle,fortran,freecad,freepascal
fsharp,fuelphp,fusetools,games,gatsby
gcov,genero4gl,geth,ggts,gis
git,gitbook,go,godot,goland
goland+all,goland+iml,goodsync,gpg,gradle
grails,greenfoot,groovy,grunt,gwt
haskell,helm,hexo,hol,homeassistant
homebrew,hsp,hugo,hyperledgercomposer,iar
iar_ewarm,iarembeddedworkbench,idapro,idris,igorpro
images,infer,inforcms,inforcrm,intellij
intellij+all,intellij+iml,ionic3,jabref,janet
java,jboss,jboss-4-2-3-ga,jboss-6-x,jboss4
jboss6,jdeveloper,jekyll,jenv,jetbrains
jetbrains+all,jetbrains+iml,jgiven,jigsaw,jmeter
joe,joomla,jsonnet,jspm,julia
jupyternotebooks,justcode,kaldi,kate,kdevelop4
kdiff3,keil,kentico,keystonejs,kicad
kirby2,kirby3,kobalt,kohana,komodoedit
konyvisualizer,kotlin,labview,labviewnxg,lamp
laravel,latex,lazarus,leiningen,lemonstand
less,liberosoc,librarian-chef,libreoffice,lighthouseci
lilypond,linux,lithium,localstack,logtalk
lsspice,ltspice,lua,lyx,macos
magento,magento1,magento2,magic-xpa,matlab
maven,mavensmate,mdbook,mean,mercurial
mercury,meson,metals,metalsmith,metaprogrammingsystem
meteor,meteorjs,microsoftoffice,mikroc,mill
moban,modelsim,modx,momentics,monodevelop
mplabx,mule,nanoc,nativescript,ncrunch
nesc,netbeans,nette,nextjs,nikola
nim,ninja,node,nodechakratimetraveldebug,nohup
notepadpp,nova,now,nuxtjs,nwjs
objective-c,obsidian,ocaml,octave,octobercms
opa,opencart,opencv,openfoam,openframeworks
openframeworks+visualstudio,oracleforms,orcad,osx,otto
oxideshop,oxygenxmleditor,packer,pants,particle
patch,pawn,perl,perl6,ph7cms
phalcon,phoenix,php-cs-fixer,phpcodesniffer,phpstorm
phpstorm+all,phpstorm+iml,phpunit,pico8,pimcore
pimcore4,pimcore5,pinegrow,platformio,playframework
plone,polymer,powershell,premake-gmake,prepros
prestashop,processing,progressabl,psoccreator,pulumi
pulumi+stacks,puppet,puppet-librarian,purebasic,purescript
putty,pvs,pycharm,pycharm+all,pycharm+iml
pydev,python,pythonvanilla,qml,qooxdoo
qt,qtcreator,r,racket,rails
react,reactnative,reasonml,red,redcar
redis,remix,remix+arc,remix+cloudflarepages,remix+cloudflareworkers
remix+netlify,remix+vercel,renpy,replit,retool
rhodesrhomobile,rider,robotframework,root,ros
ros2,ruby,rubymine,rubymine+all,rubymine+iml
rust,rust-analyzer,salesforce,salesforcedx,sam
sam+config,sas,sass,sbt,scala
scheme,scons,scrivener,sdcc,seamgen
senchatouch,serverless,shopware,silverstripe,sketchup
slickedit,smalltalk,snap,snapcraft,snyk
solidity,soliditytruffle,sonar,sonarqube,sourcepawn
spark,specflow,splunk,spreadsheet,ssh
standardml,stata,stdlib,stella,stellar
storybookjs,strapi,stylus,sublimetext,sugarcrm
svelte,svn,swift,swiftpackagemanager,swiftpm
symfony,symphonycms,synology,synopsysvcs,tags
tarmainstallmate,terraform,terragrunt,test,testcomplete
testinfra,tex,text,textmate,textpattern
theos-tweak,thinkphp,tla+,toit,tortoisegit
tower,turbo,turbogears2,twincat3,tye
typings,typo3,typo3-composer,umbraco,unity
unrealengine,vaadin,vagrant,valgrind,vapor
vcpkg,venv,vercel,vertx,video
vim,virtualenv,virtuoso,visualbasic,visualstudio
visualstudiocode,vivado,vlab,vrealizeorchestrator,vs
vue,vuejs,vvvv,waf,wakanda
web,webmethods,webstorm,webstorm+all,webstorm+iml
werckercli,windows,wintersmith,wordpress,wyam
xamarinstudio,xcode,xcodeinjection,xilinx,xilinxise
xilinxvivado,xill,xmake,xojo,xtext
y86,yalc,yarn,yeoman,yii
yii2,zendframework,zephir,zig,zsh
zukencr8000
```

## 参考情報

- https://zenn.dev/kiri_i/articles/gitignore_io
- https://docs.gitignore.io/use/api
