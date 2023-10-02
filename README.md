<img width=100% src="https://capsule-render.vercel.app/api?type=waving&color=0000FF&height=120&section=header"/>

![hunteheader](/img/hunterheader.png)

<p align="center">
	<a href="https://www.python.org/"><img src="https://img.shields.io/badge/made%20with-go-blue"></a>
	<a href="#"><img src="https://img.shields.io/badge/platform-osx%2Flinux%2Fwindows-blueviolet"></a>
	<a href="https://github.com/joaoviictorti/hunterheader/releases"><img src="https://img.shields.io/github/release/joaoviictorti/hunterheader?color=blue"></a>
</p>

<h4 align="center">Hunteheader check for missing security headers in an HTTP/HTTPS response</h4>


<p align="center">
  <a href="#características">Features</a> •
  <a href="#instalação">Installation</a> •
  <a href="#how-to-use"> How to use</a> •
  <a href="#details">Details</a> •
  <a href="#running-huntedheader">Running huntedheader</a>  
</p>

---


Huntedheader is a tool that aims to check whether security headers exist in an HTTP/HTTPS response from a mentioned url.

I designed `huntedheader` and maintained a consistently passive model to make it useful for penetration testers.

# Features

 - Check whether security headers exist in an HTTP/HTTPS response.

# How to use

```sh
huntedheader -u https://teste.com
```

# Details

![hunteheader](/img/help.png)


# Installation

huntedheader requires **golang** installed, to use:

```sh
go install -v github.com/joaoviictorti/huntedheader/cmd/hunterheader@latest
```

# Running huntedheader

![hunteheader](/img/exec.png)


<img width=100% src="https://capsule-render.vercel.app/api?type=waving&color=0000FF&height=120&section=footer"/>
