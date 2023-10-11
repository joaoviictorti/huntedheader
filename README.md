# huntedheader

Hunteheader check for missing security headers in an HTTP/HTTPS response

<p align="left">
	<a href="https://go.dev"><img src="https://img.shields.io/badge/made%20with-go-blue"></a>
	<a href="#"><img src="https://img.shields.io/badge/platform-osx%2Flinux%2Fwindows-blueviolet"></a>
	<a href="https://github.com/joaoviictorti/hunterheader/releases"><img src="https://img.shields.io/github/release/joaoviictorti/hunterheader?color=blue"></a>
</p>

![hunteheader](/img/hunterheader.png)

Huntedheader is a tool that aims to check whether security headers exist in an HTTP/HTTPS response from a mentioned url. I designed `huntedheader` and maintained a consistently passive model to make it useful for penetration testers.

- [Installation](#installation)
- [Features](#features)
- [Usage](#usage)
- [Details](#details)
- [Running huntedheader](#running-huntedheader)


---


# Features

 - Check whether security headers exist in an HTTP/HTTPS response.

# Usage

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
