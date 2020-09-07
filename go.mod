module startGo

go 1.14

require (
	github.com/imdario/mergo v0.3.11 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/utils v0.0.0-20200821003339-5e75c0163111 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20200904131630-697df40f2d58
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200904051630-d8e5c2b33a59
)
